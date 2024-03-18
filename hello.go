package main

import "fmt"

func main() {
	var confrerenceName = "Go Confrerence"
	const confrerenceTickets = 50
	var confrerenceRemains = 50

	fmt.Printf("Welcome to our %v booking application\n", confrerenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrerenceTickets, confrerenceRemains)
	fmt.Printf("Get your tickets here to attend.\n")
}
