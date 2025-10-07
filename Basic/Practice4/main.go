package main

import (
	"fmt"
	"slices"
)


// main runs the program that collects numbers and displays their average and maximum.
func main() {
	avg, max := calculateAverage()
	fmt.Printf("Average: %.2f \nMax: %.2f", avg, max)
}


// calculateAverage reads numbers from user input and returns the average and maximum.
func calculateAverage() (float64, float64) {
	numbers := []float64{}

	for {
		var number float64
		fmt.Print("Enter a number: ")
		fmt.Scanln(&number)

		if number >= 0 {
			numbers = append(numbers, number)
		} else {
			break
		}
	}

	if len(numbers) == 0 {
		fmt.Println("No numbers entered.")
		return 0, 0
	}

	max := slices.Max(numbers)
	avg := Average(numbers)

	return avg, max
}

func Average(numbers []float64) float64 {
	n := len(numbers)
	var sum float64

	for i := 0; i < n; i++ {
		sum += numbers[i]
	}

	avg := sum / float64(n)
	
	return avg
}
