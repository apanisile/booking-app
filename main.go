package main

import (
	"fmt"
	"strings"
)

var remainingTickets uint = 50

// var firstName, lastName, email string
var userTickets, extraTickets uint

const conferenceTicket uint = 50

var bookings = []string{}
var conferenceName = "Otaku conference"

func main() {
	greetUser()
}

// func exit(){
// 	break
// }

func greetUser() {
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
		bookTicket()
	case 2:
		adminPanel()
	case 3:
		break
	default:
		fmt.Print("No valid selection")
	}

}

func userDetails() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("\nPlease enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: \n")
	fmt.Scan(&email)

	fmt.Print("How mamy tickets do you want to buy: \n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket() {

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := userDetails()

		//Validation
		isValidEmail, isValidName, isValidTicketAmount := checkValidation(firstName, lastName, email, userTickets)

		if userTickets > conferenceTicket {
			fmt.Print("\nYou cant book more tickets than available at this time!\n")
		}

		if isValidEmail && isValidName && isValidTicketAmount {

			if userTickets <= remainingTickets {
				ticketConfirmation(firstName, lastName, email, userTickets)
				if remainingTickets == 0 {
					//end program
					fmt.Println("\nThe conference is booked out. Come again next year!")
					break
				}
				//More tickets function
				moreTickets(firstName, userTickets, email)
				//Ticket counter
			} else {
				var choice uint
				fmt.Printf("We have only %v ticket(s) remaining \n", remainingTickets)
				fmt.Print("Do you want to book the remaining tickets? \n")
				fmt.Print("(1) Yes \n(2) No \n")
				fmt.Scan(&choice)
				if choice == 1 {
					userTickets = remainingTickets
					fmt.Printf("\nThank you %v for booking %v ticket(s). You will receive a confirmation in your email at %v\n", firstName, userTickets, email)
					if remainingTickets == 0 {
						//end program
						fmt.Println("\nThe conference is booked out. Come again next year!")
					}
				} else {
					fmt.Print("Bye!")
					main()
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
		main()
	}

}

func checkValidation(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketAmount := userTickets > 0

	return isValidEmail, isValidName, isValidTicketAmount
}

func moreTickets(firstName string, userTickets uint, email string) {
	for remainingTickets > 0 {
		fmt.Print("\nDo you want to book another ticket(s)? \n")
		fmt.Print("(1) Yes \n(2) No \n")
		var choice uint
		fmt.Scan(&choice)
		if choice == 1 {

			fmt.Print("How many tickets do you want to buy: ")
			fmt.Scan(&extraTickets)
			if extraTickets <= remainingTickets {
				userTickets = extraTickets + userTickets
				fmt.Printf("\nThank you %v for booking %v extra ticket(s)! You will receive a confirmation in your email at %v\n", firstName, extraTickets, email)
				remainingTickets = remainingTickets - extraTickets
				fmt.Printf("There are %v ticket(s) available right now. \n", remainingTickets)
				continue
			} else {
				fmt.Print("\nYou cant book more tickets than available at this time!\n")
				fmt.Printf("There are %v ticket(s) available right now. \n", remainingTickets)
			}
			if remainingTickets == 0 {
				//end program
				fmt.Println("\nThe conference is booked out. Come again next year!")
				break
			}

		} else {
			fmt.Printf("\nTotal number of tickets booked by %v is : %v \n", firstName, userTickets)
			fmt.Println("Bye!")
			break
		}
	}
}

func ticketConfirmation(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	//remainingTickets = conferenceTicket - userTickets

	//use sql here
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\nThank you %v for booking %v ticket(s). You will receive a confirmation in your email at %v\n", firstName, userTickets, email)
	fmt.Printf("\nThere are %v ticket(s) left\n", remainingTickets)

}

func continueMenu() {
	var continueOption uint
	fmt.Println("Do you want to perform any other operation?")
	fmt.Scan(&continueOption)
	switch continueOption {
	case 1:
		main()
	case 2:
		break
	default:
		fmt.Print("No valid selection")
	}
}
