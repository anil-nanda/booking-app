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

	remainingTickets = remainingTickets - userTickets

	// define and array of size 50 with type string
	var bookings = [50]string{}

	// define array with elements
	var arr = [50]string{"a","b","c"}

	fmt.Printf("%v\n",arr)

	// add element to array
	bookings[0] = userName

	fmt.Printf("Whole array %v\n",bookings)
	fmt.Printf("Length of array is %v\n", len(bookings))
	fmt.Printf("Type of array is %T\n", bookings)
	fmt.Printf("First element of array is %v\n", bookings[0])


	//Slice is an array without size definition
	var booked = []string{}
	booked  = append(booked, userName)
	fmt.Printf("print slice %v\n", booked)
	
}