package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	//alternative syntax for variable definition
	soldTickets := 0

	fmt.Println(soldTickets)

	fmt.Printf("Type of variable conferenceName is %T\n", conferenceName)

	fmt.Println("Welocome to the", conferenceName)
	fmt.Println("We have",remainingTickets, "remaining")

	//using printf instead of Println
	fmt.Printf("Get your tickets for %s here\n", conferenceName)

	var userName string
	var userTickets int
	
	fmt.Println("Enter the username: ")
	fmt.Scan(&userName)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)
	fmt.Printf("user %v booked %v tickets\n", userName,userTickets)
	
}