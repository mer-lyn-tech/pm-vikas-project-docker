package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Number int    `json:"number"`
	Result string `json:"result"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Even Odd Checker API is Running")
}

func checkEvenOdd(w http.ResponseWriter, r *http.Request) {

	numStr := r.URL.Query().Get("number")

	number, err := strconv.Atoi(numStr)

	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	response := Response{
		Number: number,
	}

	if number%2 == 0 {
		response.Result = "Even"
	} else {
		response.Result = "Odd"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/check", checkEvenOdd)

	fmt.Println("Even Odd Checker API running on port 8080")

	http.ListenAndServe(":8080", nil)
}