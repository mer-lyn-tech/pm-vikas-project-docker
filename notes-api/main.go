package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes []Note

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Notes API is Running")
}

// GET all notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// GET note by ID
func getNoteByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, note := range notes {

		if note.ID == id {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(note)
			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}

// POST
func addNote(w http.ResponseWriter, r *http.Request) {

	var note Note

	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	note.ID = len(notes) + 1

	notes = append(notes, note)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// PUT
func updateNote(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idStr)

	for i := range notes {

		if notes[i].ID == id {

			json.NewDecoder(r.Body).Decode(&notes[i])

			notes[i].ID = id

			json.NewEncoder(w).Encode(notes[i])

			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}

// DELETE
func deleteNote(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idStr)

	for i := range notes {

		if notes[i].ID == id {

			notes = append(notes[:i], notes[i+1:]...)

			fmt.Fprintln(w, "Note Deleted")

			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}

func notesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		id := r.URL.Query().Get("id")

		if id == "" {
			getNotes(w, r)
		} else {
			getNoteByID(w, r)
		}

	case http.MethodPost:
		addNote(w, r)

	case http.MethodPut:
		updateNote(w, r)

	case http.MethodDelete:
		deleteNote(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/", home)

	http.HandleFunc("/notes", notesHandler)

	fmt.Println("Notes API running on port 8080")

	http.ListenAndServe(":8080", nil)
}