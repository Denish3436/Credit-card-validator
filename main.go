package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ValidationResult represents the structure of the JSON response.
type ValidationResult struct {
	CardNumber string `json:"card_number"`
	IsValid    bool   `json:"is_valid"`
}

func validateCreditCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var input map[string]string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	cardNumber, exists := input["card_number"]
	if !exists {
		http.Error(w, "Missing card_number field", http.StatusBadRequest)
		return
	}

	isValid := luhnAlgorithm(cardNumber)

	response := ValidationResult{
		CardNumber: cardNumber,
		IsValid:    isValid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/validate", validateCreditCardHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
