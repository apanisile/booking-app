package main

import (
	"fmt"
)

func adminPanel() {
	var adminSelect uint
	fmt.Println("What would you like to do?")
	fmt.Println("(1) List of registered users \n(2) Check remaining tickets \n(3) Top up tickets")
	fmt.Scan(&adminSelect)
	switch adminSelect {
	case 1:
		checkUsers()
		continueMenu()
	case 2:
		checkTicket()
	case 3:
		break
	default:
		fmt.Print("No valid selection")
	}
}

func checkTicket() {
	fmt.Printf("There are %v tickets available \n", remainingTickets)
	continueMenu()
}

func checkUsers() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	fmt.Printf("The booked users are %v\n", firstNames)
	return firstNames
}
