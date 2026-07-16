package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BMIResponse struct {
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	BMI    float64 `json:"bmi"`
	Status string  `json:"status"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "BMI Calculator API is Running")
}

func calculateBMI(w http.ResponseWriter, r *http.Request) {

	heightStr := r.URL.Query().Get("height")
	weightStr := r.URL.Query().Get("weight")

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		http.Error(w, "Invalid height", http.StatusBadRequest)
		return
	}

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		http.Error(w, "Invalid weight", http.StatusBadRequest)
		return
	}

	bmi := weight / (height * height)

	var status string

	switch {
	case bmi < 18.5:
		status = "Underweight"
	case bmi < 25:
		status = "Normal"
	case bmi < 30:
		status = "Overweight"
	default:
		status = "Obese"
	}

	response := BMIResponse{
		Height: height,
		Weight: weight,
		BMI:    bmi,
		Status: status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/bmi", calculateBMI)

	fmt.Println("BMI API running on port 8080")

	http.ListenAndServe(":8080", nil)
}