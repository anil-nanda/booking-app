package main

import (
	"fmt"
 	"strings"
	"booking-app/helper"
)

// package level variables, it doesn't support := syntax
const conferenceTickets = 50
var conferenceName = "Go conference"
var remainingTickets = 50
var bookings = []string{}

//no arguments are passed as the variables in function are refered from package level
func greetUsers() {
	fmt.Println("Welocome to the ", conferenceName)
	fmt.Println("We have",remainingTickets, "remaining")
}

// last parameter specifies the type of return value
func getFirstName() []string {
	firstNames := []string{}
			// _ is used to mention an unused variable as index is not reffered in the loop 
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			return firstNames
	}

func getUserInput() (string, string, string, string, int) {
	
	var firstName string
	var lastName string
	var email string
	var userTickets int
	var city string

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

	return firstName, lastName, email, city, userTickets
}

func bookTicket( firstName string, lastName string, email string, userTickets int) {
	bookings  = append(bookings, firstName + " "+ lastName)	
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	fmt.Printf("user %v %v booked %v tickets. confirmation mail send to %v\n", firstName, lastName,userTickets, email)
}
	
func main() {

	greetUsers()
	
	for {

		firstName, lastName, email, city, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets, isValidCity := helper.ValidateUserInputs(firstName, lastName, email, userTickets, city, remainingTickets)
		
		if isValidName && isValidEmail && isValidTickets && isValidCity {
			
			switch city {
				case "London", "Manchester":
					timeZone := "UTC"
					fmt.Println(timeZone)
				case "Singapore":
					timeZone := "SGT"
					fmt.Println(timeZone)
				default:
					timeZone := "GMT"
					fmt.Println(timeZone)
			}
			
			bookTicket( firstName, lastName, email, userTickets)
			//fmt.Printf("First element of slice is %v\n", bookings[0])
			//fmt.Printf("Type of slice is %T\n", bookings)

			firstNames := getFirstName()

			fmt.Printf("First names of the booking are %v\n", firstNames)

			//fmt.Printf("Whole booking list %v\n",bookings)
			//fmt.Printf("Number of unique bookings is %v\n", len(bookings))

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Printf("All tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Entered name is not valid\n")
			}
			if !isValidEmail {
				fmt.Printf("Email is not valid\n")
			}
			if !isValidCity {
				fmt.Printf("City is not valid\n")
			}
			if !isValidTickets {
				fmt.Printf("Tickets is not valid\n")
			}
		}

		
	}
	
}