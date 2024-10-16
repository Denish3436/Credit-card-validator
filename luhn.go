package main

import (
	"strconv"
)

// luhnAlgorithm checks if a credit card number is valid using the Luhn algorithm.
func luhnAlgorithm(cardNumber string) bool {
	var sum int
	nDigits := len(cardNumber)
	parity := nDigits % 2

	
	for i := 0; i < nDigits; i++ {
		
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}

		if i%2 == parity {
			digit *= 2
		}

		
		if digit > 9 {
			digit -= 9
		}

		
		sum += digit
	}

	// If the total modulo 10 is 0, the number is valid
	return sum%10 == 0
}
