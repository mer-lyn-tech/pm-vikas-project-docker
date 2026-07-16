package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type SalaryResponse struct {
	Name        string  `json:"name"`
	BasicSalary float64 `json:"basic_salary"`
	HRA         float64 `json:"hra"`
	DA          float64 `json:"da"`
	Tax         float64 `json:"tax"`
	NetSalary   float64 `json:"net_salary"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Employee Salary Calculator API is Running")
}

func calculateSalary(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	basicStr := r.URL.Query().Get("basic")

	basic, err := strconv.ParseFloat(basicStr, 64)

	if err != nil {
		http.Error(w, "Invalid Salary", http.StatusBadRequest)
		return
	}

	hra := basic * 0.20
	da := basic * 0.10
	tax := basic * 0.05

	net := basic + hra + da - tax

	response := SalaryResponse{
		Name:        name,
		BasicSalary: basic,
		HRA:         hra,
		DA:          da,
		Tax:         tax,
		NetSalary:   net,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/", home)

	http.HandleFunc("/salary", calculateSalary)

	fmt.Println("Salary API running on port 8080")

	http.ListenAndServe(":8080", nil)

}