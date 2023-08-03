package main

import (
	"fmt"
	"go-booking-tutorial/helper"
	"strconv"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := collectUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all of our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("The conference is fully booked now! Come back next week.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name and last name needs to be at least 2 characters in length.")
			}
			if !isValidEmail {
				fmt.Println("Your email address doesn't contain an @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("You entered an incorrect amount of tickets.")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func collectUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for user name and amount of tickets they want to buy
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email sent to %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
}
