package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Result float64 `json:"result"`
}

func calculate(w http.ResponseWriter, r *http.Request) {

	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		http.Error(w, "Invalid value for a", http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		http.Error(w, "Invalid value for b", http.StatusBadRequest)
		return
	}

	var result float64

	switch op {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		}
		result = a / b
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: result})
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Calculator API is Running")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/calculate", calculate)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}


