package main

import (
	"fmt"
)

// Student represents a student with a name, age, and grade.
type Student struct {
	Name  string  // Name of the student
	Age   int     // Age of the student
	Grade float64 // Grade of the student
}

// AddStudent creates and returns a new Student instance with the given name, age, and grade.
func AddStudent(s string, i int, f float64) Student {
	return Student{s, i, f} // Returns a new Student struct
}

// FindByName searches for a student by name in the provided list and prints their details if found.
func FindByName(n string, list []Student) {
	for _, v := range list { // Iterate over the student list
		if v.Name == n { // Check if the student's name matches the search term
			fmt.Printf("Found: %s | Age: %v | Grade: %.2f\n", v.Name, v.Age, v.Grade)
			return // Exit the function after finding the student
		}
	}
	fmt.Println("Student not found!") // Print message if no student is found
}

// RemoveByName removes a student by name from the provided list and returns the updated list.
func RemoveByName(n string, list []Student) []Student {
	for i, v := range list { // Iterate over the student list with index
		if v.Name == n { // Check if the student's name matches
			// Remove the student by slicing the list
			list = append(list[:i], list[i+1:]...)
			return list // Return the updated list
		}
	}
	fmt.Println("Student not found!") // Print message if no student is found
	return list // Return the original list if no student is removed
}

// MaxGrade finds and prints the student with the highest grade in the list.
func MaxGrade(list []Student) {
	if len(list) == 0 { // Check if the list is empty
		fmt.Println("No students in the list.")
		return
	}

	max := list[0].Grade       // Initialize max grade with the first student's grade
	bestStudent := list[0].Name // Initialize best student with the first student's name

	for i := 1; i < len(list); i++ { // Iterate over the list starting from the second student
		if list[i].Grade > max { // Check if the current student's grade is higher
			max = list[i].Grade          // Update max grade
			bestStudent = list[i].Name   // Update best student
		}
	}

	// Print the name and grade of the student with the highest grade
	fmt.Printf("Best student is %s with %.2f Grade\n", bestStudent, max)
}

// AverageGrade calculates and prints the average grade of all students in the list.
func AverageGrade(list []Student) {
	n := len(list) // Get the number of students
	if n == 0 { // Check if the list is empty
		fmt.Println("Average: N/A (no students)")
		return
	}
	var Sum float64 // Variable to store the sum of grades
	for _, v := range list { // Iterate over the student list
		Sum += v.Grade // Add each student's grade to the sum
	}
	Avg := Sum / float64(len(list)) // Calculate the average grade
	fmt.Printf("Average: %.2f\n", Avg) // Print the average grade
}

func main() {
	// Initialize a slice of Student structs with sample data
	students := []Student{
		{"Ali", 24, 17.02},
		{"Mohammad", 25, 18.05},
		{"Soheil", 20, 18.00},
	}

	// Add a new student named "Reza" to the list
	NewStudent := AddStudent("Reza", 18, 12.35)
	students = append(students, NewStudent) // Append the new student to the list

	// Search for students by name and print their details
	FindByName("Ali", students)   // Should find and print Ali's details
	FindByName("Reza", students)  // Should find and print Reza's details

	// Remove a student by name and update the list
	students = RemoveByName("Reza", students) // Remove Reza from the list
	FindByName("Reza", students)              // Should print "Student not found!"

	// Find and print the student with the highest grade
	MaxGrade(students)

	// Calculate and print the average grade of all students
	AverageGrade(students)
}