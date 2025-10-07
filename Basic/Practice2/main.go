package main

import (
	"fmt"
	"time"
)

/*
	main runs the Go learning practice for taking user input

and displaying a formatted message with the current time.
*/
func main() {

	// Declare variables for user's name and age.
	var name string
	var age int

	// Ask for user's name
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	// Ask for user's age.
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	// Display personalized greeting.
	fmt.Printf("Hi %s. You are %v years old. \n", name, age)

	// Get current date and time.
	dt := time.Now()

	// Print formatted date and time (YYYY/MM/DD HH:MM:SS).
	fmt.Println("The time is: ", dt.Format("2006/01/02 15:04:05"))
}
