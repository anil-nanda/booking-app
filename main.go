package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welocome to the", conferenceName)
	fmt.Println("We have",remainingTickets, "remaining")

	//using printf instead of Println
	fmt.Printf("Get your tickets for %s here\n", conferenceName)

	
}