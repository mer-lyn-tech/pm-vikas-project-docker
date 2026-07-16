package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Input       float64 `json:"input"`
	Unit        string  `json:"unit"`
	Result      float64 `json:"result"`
	ConvertedTo string  `json:"converted_to"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Temperature Converter API is Running")
}

func convert(w http.ResponseWriter, r *http.Request) {

	valueStr := r.URL.Query().Get("value")
	unit := r.URL.Query().Get("unit")

	value, err := strconv.ParseFloat(valueStr, 64)

	if err != nil {
		http.Error(w, "Invalid temperature", http.StatusBadRequest)
		return
	}

	var response Response

	if unit == "c" {

		response.Input = value
		response.Unit = "Celsius"
		response.Result = (value * 9 / 5) + 32
		response.ConvertedTo = "Fahrenheit"

	} else if unit == "f" {

		response.Input = value
		response.Unit = "Fahrenheit"
		response.Result = (value - 32) * 5 / 9
		response.ConvertedTo = "Celsius"

	} else {

		http.Error(w, "Use unit=c or unit=f", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/convert", convert)

	fmt.Println("Temperature API running on port 8080")

	http.ListenAndServe(":8080", nil)
}