package main

import "fmt"

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
					sendTicket(userTickets, firstName, email)
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
					sendTicket(userTickets, firstName, email)
					if remainingTickets == 0 {
						sendTicket(userTickets, firstName, email)
						fmt.Println("\nThe conference is booked out. Come again next year!")
					}
				} else {
					sendTicket(userTickets, firstName, email)
					fmt.Print("Bye!")
					abegWait()
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
