package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Age        int    `json:"age"`
	Eligible   bool   `json:"eligible"`
	Message    string `json:"message"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Age Eligibility Checker API is Running")
}

func checkAge(w http.ResponseWriter, r *http.Request) {

	ageStr := r.URL.Query().Get("age")

	age, err := strconv.Atoi(ageStr)

	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	response := Response{
		Age: age,
	}

	if age >= 18 {
		response.Eligible = true
		response.Message = "Eligible to Vote"
	} else {
		response.Eligible = false
		response.Message = "Not Eligible to Vote"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/check", checkAge)

	fmt.Println("Age Eligibility Checker API running on port 8080")

	http.ListenAndServe(":8080", nil)
}