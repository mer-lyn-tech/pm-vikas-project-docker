package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var todos []Todo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "To-Do API is Running")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {

	var todo Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	todo.ID = len(todos) + 1

	todos = append(todos, todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idStr)

	for i := range todos {

		if todos[i].ID == id {

			json.NewDecoder(r.Body).Decode(&todos[i])

			todos[i].ID = id

			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idStr)

	for i := range todos {

		if todos[i].ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			fmt.Fprintln(w, "Task Deleted")

			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		getTodos(w, r)

	case http.MethodPost:
		addTodo(w, r)

	case http.MethodPut:
		updateTodo(w, r)

	case http.MethodDelete:
		deleteTodo(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

	}

}

func main() {

	http.HandleFunc("/", home)

	http.HandleFunc("/todos", todoHandler)

	fmt.Println("To-Do API running on port 8080")

	http.ListenAndServe(":8080", nil)

}