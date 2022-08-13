package main

import (
	"fmt"
	"strings"
)

// func userDetails(){
// var firstName, lastName, email string
// 	var userTickets int

// 	fmt.Print("Please enter your first name: ")
// 	fmt.Scan(&firstName)

// 	fmt.Print("Please enter your last name: ")
// 	fmt.Scan(&lastName)

// 	fmt.Print("Enter your email address: ")
// 	fmt.Scan(&email)

// 	fmt.Print("How mamy tickets do you want to buy: ")
// 	fmt.Scan(&userTickets)
// 	fmt.Printf("\nThank you %v for booking %v tickets! You will recieve a confirmation in your email at %v\n", firstName, userTickets, email)
// }

func main() {
	conferenceName := "Otaku conference"
	const conferenceTicket = 50
	remainingTickets := 50

	fmt.Printf("Hello! Welcome to %v booking application!\n", conferenceName)
	fmt.Print("Get your tickets here to attend!")
	fmt.Printf("\nThere are %v tickets available\n\n", remainingTickets)
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

		if userTickets > conferenceTicket {
			fmt.Print("\nYou cant book more tickets than available at this time!\n")
			continue
		}

		if userTickets <= remainingTickets {
			//admin-panel
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v for booking %v ticket(s)! You will recieve a confirmation in your email at %v\n", firstName, userTickets, email)
			fmt.Printf("\nThere are %v ticket(s) available\n\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("\nThe booked users are %v\n", firstNames)
			if remainingTickets == 0 {
				//end program
				fmt.Println("\nThe conference is booked out. Come again next year!")
				break
			}

		} else {
			fmt.Printf("We have only %v ticket(s) remaining \n", remainingTickets)
		}
	}
}

// func exit(){
// 	break
// }

// func admin-panel(){
// 	fmt.Printf("\nThe booked users are %v\n", firstNames)
// }
