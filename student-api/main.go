		package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}

var students []Student

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Student API is Running")
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func addStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	student.ID = len(students) + 1

	students = append(students, student)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		getStudents(w, r)

	case http.MethodPost:
		addStudent(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/", home)

	http.HandleFunc("/students", studentsHandler)

	fmt.Println("Student API running on port 8080")

	http.ListenAndServe(":8080", nil)
}