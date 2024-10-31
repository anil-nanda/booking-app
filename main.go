package main

import (
	"fmt"
	"booking-app/helper"
	"strconv"
	"time"
	"sync"
)

// package level variables, it doesn't support := syntax
const conferenceTickets = 50
var conferenceName = "Go conference"
var remainingTickets = 50
var bookings = make([]map[string]string, 0)
// var bookings = make([]userData, 0)

// define a struct with custom datatypes
//type userData struct {
//	firstName string
//	lastName string
//	email string
//	city string
//	numberOfTickets int
//}

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
				firstNames = append(firstNames, booking["firstName"])
				// when using struct
				// firstNames = append(firstNames, booking.firstName)
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
	// create a map for storing all the user details
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

	// when userData struct is used
	// var userData = userData {
		// firstName: firstName,
		//lastName: lastName,
		//email: email,
		//city: city,
		//numberOfTickets: numberOfTickets
	//}

	fmt.Printf("%v\n",userData)

	bookings  = append(bookings, userData)	
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	fmt.Printf("user %v %v booked %v tickets. confirmation mail send to %v\n", firstName, lastName,userTickets, email)
}

var wg = sync.WaitGroup{}
	
func main() {

	greetUsers()
	
	//for {

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
			
			// add the number of functions that are added as async
			wg.Add(1)

			// keyword go is used to specify to run the code on separate thread
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstName()

			fmt.Printf("First names of the booking are %v\n", firstNames)

			//fmt.Printf("Whole booking list %v\n",bookings)
			//fmt.Printf("Number of unique bookings is %v\n", len(bookings))

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Printf("All tickets are sold out")
				//break
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

		
	//}

	// wait to exit until the waitgroup count is 0
	wg.Wait()
	
}


func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}