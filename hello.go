package main

import "fmt"

func main() {
	confrerenceName := "Go Confrerence"
	const confrerenceTickets = 50
	confrerenceRemains := 50

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

		fmt.Printf("Your fullname is %v.\n", firstName+" "+lastName)
		fmt.Printf("You have booked %v tickets.\n", userTickest)

		confrerenceRemains -= userTickest
		fmt.Printf("Now, we have %v tickets remain.\n", confrerenceRemains)
	}
}
