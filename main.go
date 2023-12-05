package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference 2024"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	fmt.Printf("Welcome to %v ticket booking application\n", conferenceName)
	fmt.Printf("We have %v and %v are still available. Hurry up!", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {
		var firstName string
		var lastName string
		var emailAddress string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address")
		fmt.Scan(&emailAddress)
		fmt.Println("Enter tickets number you want to book")
		fmt.Scan(&userTickets)

		fmt.Printf("Thank you %v %v, for bookking %v tickets.\nYou will recieve confirmation to %v soon.\n", firstName, lastName, userTickets, emailAddress)

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - int(userTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("Booking slice %v\n", firstNames)
		fmt.Printf("There are %v tickets still available to %v conference\n", remainingTickets, conferenceName)
	}

}
