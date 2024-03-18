package main

import "fmt"

func main() {
	confrerenceName := "Go Confrerence"
	const confrerenceTickets = 50
	confrerenceRemains := 50
	var bookings []string

	fmt.Printf("Welcome to our %v booking application\n", confrerenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrerenceTickets, confrerenceRemains)
	fmt.Printf("Get your tickets here to attend.\n")

	for {
		var firstName string
		var lastName string
		var userTickest int

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickest)

		if confrerenceRemains-userTickest < 0 {
			fmt.Println("Number of tickets is not enough, please check again!")
			break
		}

		fullname := firstName + " " + lastName

		// add fullname into bookgings array
		bookings = append(bookings, fullname)

		fmt.Printf("Your fullname is %v.\n", fullname)
		fmt.Printf("You have booked %v tickets.\n", userTickest)

		confrerenceRemains -= userTickest
		fmt.Printf("Current bookings slot: %v\n", bookings)
		fmt.Printf("Now, we have %v tickets remain.\n", confrerenceRemains)
	}
}
