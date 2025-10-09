package main

import "fmt"

// Student struct defines a structure to store student information
type student struct {
	Name  string  // Student's name
	Age   int     // Student's age
	Grade float64 // Student's grade
}

// Display method prints the student's information
func (s student) Display() {
	// Print the student's details in a formatted manner
	fmt.Printf("Student: %s | Age: %v | Grade: %.2f\n", s.Name, s.Age, s.Grade)
}

// IsAdult method checks if the student is an adult
func (s student) IsAdult() bool {
	// Returns true if the student's age is 18 or older
	return s.Age >= 18
}

// UpdateGrade method updates the student's grade
// Uses a pointer (*student) to modify the actual struct
func (s *student) UpdateGrade(g float64) {
	s.Grade = g // Update the grade
}

func main() {
	// Create a student instance named Ali
	Ali := student{"Ali", 20, 18.00}
	
	// Call the Display method to show Ali's initial information
	Ali.Display()
	
	// Check if Ali is an adult and print the result
	fmt.Println(Ali.Name, " is an adult: ", Ali.IsAdult())
	
	// Update Ali's grade to 12.5
	Ali.UpdateGrade(12.5)
	
	// Display Ali's updated information
	Ali.Display()
}