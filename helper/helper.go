package helper

import (
	"strings"
)

// making the function name in caps makes in exportable and this function can be used in other packages 
func ValidateUserInputs(firstName string, lastName string, email string, userTickets int, city string, remainingTickets int) (bool, bool, bool,bool) {
	isValidName := len(firstName) >=2 && len(lastName) >=2

	isValidEmail := strings.Contains(email, "@")

	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	isValidCity := city == "Singapore" || city == "London"

	return isValidName, isValidEmail, isValidTickets, isValidCity
}