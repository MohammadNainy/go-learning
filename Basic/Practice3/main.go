package main

import "fmt"

func main() {

	sum := 0 // Holds the total sum of positive numbers

	for {
		var number int
		fmt.Print("Enter a number: ")
		fmt.Scanln(&number)
		if number >= 0 {
			sum += number // Add positive numbers to the sum
		} else {
			break // Stop the loop when a negative number is entered
		}
	}

	// Display the final result
	fmt.Printf("The sum of positive numbers is: %v \n", sum)

}
