package main

import (
	"fmt"
	"strings"
)

// func userDetails(){

// }

func main() {
	conferenceName := "Otaku conference"
	const conferenceTicket = 50
	remainingTickets := 50

	fmt.Printf("Hello! Welcome to %v booking application!\n", conferenceName)
	fmt.Print("Get your tickets here to attend\n")
	bookings := []string{}

	for {

		//userDetails()
		var firstName, lastName, email string
		var userTickets int

		fmt.Print("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("How mamy tickets do you want to buy: ")
		fmt.Scan(&userTickets)
		fmt.Printf("\nThank you %v for booking %v tickets! You will recieve a confirmation in your email at %v\n", firstName, userTickets, email)

		//admin-panel
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("\nThere are %v tickets and %v tickets available\n", conferenceTicket, remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("\nThe booked users are %v\n", firstNames)
	}
}
