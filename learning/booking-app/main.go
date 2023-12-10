package main

import "fmt"

func main() {
	conferenceName := "Conference"
	const tickets uint = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to our booking application")
	fmt.Println("Get your tickets here")
	fmt.Println("Conference name is |", conferenceName, "|")
	fmt.Println("We have a total of", tickets, "tickets and have", remainingTickets, "tickets left")

	var bookings = [tickets]string{}

	var userName string
	var userTickets int

	fmt.Println("Please enter your name")
	fmt.Scanln(&userName)
	fmt.Println("Please enter the number of tickets you want to buy")
	fmt.Scanln(&userTickets)

}
