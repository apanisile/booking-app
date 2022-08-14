package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var remainingTickets uint = 50

// var firstName, lastName, email string
var userTickets, extraTickets uint

const conferenceTicket uint = 50

var bookings = make([]map[string]string, 0)
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
		//fmt.Println("#######################################################")
		fmt.Println("please wait...")
		time.Sleep(4 * time.Second)
		bookTicket()
	case 2:
		time.Sleep(4 * time.Second)
		adminPanel()
	case 3:
		//exit
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
				sendTicket(userTickets, firstName, email)
				fmt.Println("\nThe conference is booked out. Come again next year!")
				break
			}

		} else {
			fmt.Printf("\nTotal number of tickets booked by %v is : %v \n", firstName, userTickets)
			sendTicket(userTickets, firstName, email)
			fmt.Print("Bye!\n")
			break
		}
	}
}

func ticketConfirmation(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	//remainingTickets = conferenceTicket - userTickets
	// Using map instead fo slice
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//use sql here
	bookings = append(bookings, userData)

	fmt.Printf("\nThank you %v for booking %v ticket(s). You will receive a confirmation in your email at %v\n", firstName, userTickets, email)
	fmt.Printf("\nThere are %v ticket(s) left\n", remainingTickets)

}

func continueMenu() {
	var continueOption uint
	fmt.Println("Do you want to perform any other operation?")
	fmt.Println("(1) Yes \n(2) No")
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

func sendTicket(userTickets uint, firstName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, firstName)
	fmt.Println("#######################################################")
	fmt.Printf("Sending ticket:\n %v to %v\n\n", ticket, email)
	fmt.Println("#######################################################")
}

// func abegWait(){
// 	x := 5
// 	for x > 0{
// 		fmt.Println("#")
// 	}
// }
