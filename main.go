package main

import (
	"fmt"
 	"strings"
)
	
func main() {
	const conferenceTickets = 50
	conferenceName := "Go conference"
	
	remainingTickets := 50

	bookings := []string{}
	
	fmt.Println("Welocome to the", conferenceName)
	fmt.Println("We have",remainingTickets, "remaining")

	for {
		
		var firstName string
		var lastName string
		var email string
		var userTickets int
		
		fmt.Println("Enter the first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter the last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter the email: ")
		fmt.Scan(&email)
		fmt.Println("Enter the city: ")
		fmt.Scan(&city)
		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >=2 && len(lastName) >=2

		isValidEmail := strings.Containts(email, "@")

		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		isValidCity := city == 'Singapore' || city == 'London'
		
		if isValidName && isValidEmail && isValidTickets && isValidCity {
			fmt.Printf("user %v %v booked %v tickets. confirmation mail send to %v\n", firstName, lastName,userTickets, email)
		
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("%v tickets remaining\n", remainingTickets)

			bookings  = append(bookings, firstName + " "+ lastName)
		
			fmt.Printf("First element of slice is %v\n", bookings[0])
			//fmt.Printf("Type of slice is %T\n", bookings)

			firstNames := []string{}
			// _ is used to mention an unused variable as index is not reffered in the loop 
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First names of the booking are %v\n", firstNames)
			fmt.Printf("Whole booking list %v\n",bookings)
			fmt.Printf("Number of unique bookings is %v\n", len(bookings))

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Printf("All tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Entered name is not valid")
			}
			if !isValidEmail {
				fmt.Printf("Email is not valid")
			}
			if !isValidCity {
				fmt.Printf("City is not valid")
			}
			if !isValidTickets {
				fmt.Printf("Tickets is not valid")
			}
		}

		
	}
	
}