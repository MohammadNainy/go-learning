package main

import (
	"encoding/json" // For JSON encoding and decoding
	"fmt"           // For formatted input/output operations (e.g., printing to console)
	"os"            // For interacting with the operating system, especially file operations
)

// Student struct defines a structure to store student information with JSON tags
type Student struct {
	Name  string  `json:"name"`  // Student's name, mapped to "name" in JSON
	Age   int     `json:"age"`   // Student's age, mapped to "age" in JSON
	Grade float64 `json:"grade"` // Student's grade, mapped to "grade" in JSON
}

// Constant defining the filename for storing student data
const filename = "students.json"

// LoadStudents reads and deserializes the student data from the JSON file
func LoadStudents() ([]Student, error) {
	// Read the content of the file
	data, err := os.ReadFile(filename)
	if err != nil {
		// If the file does not exist, return an empty slice with no error
		if os.IsNotExist(err) {
			return []Student{}, nil
		}
		// Return any other error encountered during file reading
		return nil, err
	}
	// If the file is empty, return an empty slice
	if len(data) == 0 {
		return []Student{}, nil
	}
	// Declare a slice to hold the deserialized student data
	var list []Student
	// Deserialize the JSON data into the list slice
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	// Return the deserialized student list with no error
	return list, nil
}

// SaveStudents serializes the student list to JSON and writes it to the file
func SaveStudents(s []Student) error {
	// Serialize the student slice to JSON with indentation for readability
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err // Return any error during serialization
	}
	// Write the JSON data to the file with permissions 0644
	return os.WriteFile(filename, data, 0644)
}

// AddStudent appends a new student to the provided list
func AddStudent(list []Student, name string, age int, grade float64) []Student {
	// Append a new Student struct to the list and return the updated list
	return append(list, Student{name, age, grade})
}

// PrintStudents displays the list of students with their details
func PrintStudents(list []Student) {
	// Check if the list is empty
	if len(list) == 0 {
		fmt.Println("No students")
		return
	}
	// Iterate over the student list and print details with index
	for i, info := range list {
		fmt.Printf("%d) %s | Age: %d | Grade: %.2f\n", i+1, info.Name, info.Age, info.Grade)
	}
}

// Main function where the program execution starts
func main() {
	// Load the student list from the JSON file
	students, err := LoadStudents()
	if err != nil {
		// Print error and exit if loading fails
		fmt.Println("load error:", err)
		return
	}
	// Print the loaded student list
	fmt.Println("Loaded from file:")
	PrintStudents(students)

	// Add two new students to the list
	students = AddStudent(students, "Reza", 32, 15.32)
	students = AddStudent(students, "Mohammad", 24, 18.01)

	// Save the updated student list to the JSON file
	if err := SaveStudents(students); err != nil {
		// Print error and exit if saving fails
		fmt.Println("save error:", err)
		return
	}

	// Reload the student list from the file to verify saving
	reloaded, err := LoadStudents()
	if err != nil {
		// Print error and exit if reloading fails
		fmt.Println("reload error:", err)
		return
	}
	// Print the reloaded student list
	fmt.Println("\nAfter save (reloaded):")
	PrintStudents(reloaded)
}