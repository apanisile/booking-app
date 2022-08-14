package main

import (
	"fmt"
	"strings"
)

var remainingTickets = 50

// var firstName, lastName, email string
var userTickets int

const conferenceTicket = 50

var bookings = []string{}
var conferenceName = "Otaku conference"

func main() {
	greetUser(userTickets, conferenceName, bookings, remainingTickets)
}

// func exit(){
// 	break
// }

func greetUser(userTickets int, conferenceName string, bookings []string, remainingTickets int) {
	var menuSelect uint
	//bookings := []string{}
	fmt.Println("Welcome to our conference")
	fmt.Printf("Hello! Welcome to %v booking application!\n", conferenceName)
	fmt.Println("Get your tickets here to attend!")

	fmt.Println("What would you like to do today?")
	fmt.Print("(1) Book ticket \n(2) Admin Panel \n(3) Exit\n")
	fmt.Scan(&menuSelect)

	switch menuSelect {
	case 1:
		bookTicket(userTickets, conferenceName, bookings, remainingTickets)
	case 2:
		//adminPanel(bookings)
	default:
		fmt.Print("No valid selection")
	}

}

func userDetails() (string, string, string, int) {
	var firstName, lastName, email string
	var userTickets int
	//var choice int
	//const conferenceTicket = 50

	fmt.Print("Please enter your first name: ")

	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How mamy tickets do you want to buy: ")
	fmt.Scan(&userTickets)

	//fmt.Printf("\nThank you %v for booking %v tickets! You will recieve a confirmation in your email at %v\n", firstName, userTickets, email)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, conferenceName string, bookings []string, remainingTickets int) {

	//fmt.Printf("\nThere are %v tickets available\n\n", remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := userDetails()

		//Validation
		isValidEmail, isValidName, isValidTicketAmount := checkValidation(firstName, lastName, email, userTickets, remainingTickets)

		if userTickets > conferenceTicket {
			fmt.Print("\nYou cant book more tickets than available at this time!\n")
		}

		if isValidEmail && isValidName && isValidTicketAmount {

			if userTickets <= remainingTickets {
				//admin-panel
				remainingTickets = remainingTickets - userTickets
				//remainingTickets = conferenceTicket - userTickets

				//use sql here
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Printf("\nThank you %v for booking %v ticket(s)! You will receive a confirmation in your email at %v\n", firstName, userTickets, email)
				fmt.Printf("\nThere are %v ticket(s) left\n", remainingTickets)

				//More tickets function
				// fmt.Print("Do you want to book another ticket(s)? \n")
				// fmt.Print("(1) Yes \n(2) No \n")
				// fmt.Scan(&choice)
				// if choice == 1 {

				// 	fmt.Print("How many tickets do you want to buy: ")
				// 	fmt.Scan(&extraTickets)

				// 	userTickets = extraTickets + userTickets
				// 	fmt.Printf("\nThank you %v for booking %v ticket(s)! You will receive a confirmation in your email at %v\n", firstName, userTickets, email)
				// } else {
				// 	fmt.Print("Bye!")
				// 	break
				// }

				//Ticket counter
				if remainingTickets == 0 {
					//end program
					fmt.Println("\nThe conference is booked out. Come again next year!")
					break
				}

			} else {
				var choice int
				fmt.Printf("We have only %v ticket(s) remaining \n", remainingTickets)
				fmt.Print("Do you want to book the remaining tickets? \n")
				fmt.Print("(1) Yes \n(2) No \n")

				fmt.Scan(&choice)
				if choice == 1 {
					userTickets = remainingTickets
					fmt.Printf("\nThank you %v for booking %v ticket(s)! You will receive a confirmation in your email at %v\n", firstName, userTickets, email)
					if remainingTickets == 0 {
						//end program
						fmt.Println("\nThe conference is booked out. Come again next year!")
					}
				} else {
					fmt.Print("Bye!")
					break
				}
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email address, try again")
			}
			if !isValidTicketAmount {
				fmt.Println("Number of tickets you enterd is invalid")
			}
			//fmt.Print("Your input data is invalid, try again\n\n")
		}
	}
	//fmt.Printf("\nThe booked users are %v\n", firstNames)

}

// func adminPanel(bookings []string) {
// 	firstNames := []string{}
// 	for _, booking := range bookings {
// 		var names = strings.Fields(booking)
// 		firstNames = append(firstNames, names[0])
// 	}

// 	fmt.Printf("The booked users are %v\n", firstNames)
// }

func checkValidation(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketAmount := userTickets > 0

	return isValidEmail, isValidName, isValidTicketAmount
}
