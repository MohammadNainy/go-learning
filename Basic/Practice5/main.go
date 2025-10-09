package main

import (
	"bufio" // For reading user input line by line
	"fmt"   // For formatted input/output operations (e.g., printing to console)
	"os"    // For interacting with the operating system, especially file operations
)

// Main function where the program execution starts.
func main() {
	// Call the function responsible for creating or reading a file.
	CreateOrReadFile()
}

// Main function for file management: creating, writing, and reading the file.
func CreateOrReadFile() {
	// Define a constant filename for the file to be used.
	const filename = "notes.txt"

	// 1. Open/Create the file:
	// os.OpenFile opens the file. If it doesn't exist, os.O_CREATE creates it.
	// os.O_APPEND ensures data is appended to the end of the file.
	// os.O_WRONLY opens the file for write-only operations.
	// 0644 sets file permissions (read/write for owner, read-only for group and others).
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Check for errors when opening/creating the file.
	check(err)
	// `defer` ensures f.Close() is called just before CreateOrReadFile exits,
	// closing the file to free system resources.
	defer f.Close()

	// 2. Write a header if the file is empty:
	// os.Stat returns file information (e.g., size).
	// If no error (file exists) and size is 0 bytes (file is empty).
	if info, err := os.Stat(filename); err == nil && info.Size() == 0 {
		// Write the header "Current Notes:" at the start of the file.
		_, err = f.WriteString("Current Notes:\n")
		// Check for errors when writing to the file.
		check(err)
	}

	// 3. Main loop to collect user notes:
	for {
		// Get a new note from the user.
		note := UserNote()

		// If the user enters "0", break the loop and exit.
		if note == "0" {
			break
		}
		// If the user enters an empty line, skip this iteration and continue.
		if note == "" {
			continue
		}

		// Write the note to the file with a "-" prefix and a newline.
		_, err = f.WriteString("- " + note + "\n")
		// Check for errors when writing.
		check(err)

		// Notify the user that the note has been saved.
		fmt.Println("Your note has been saved!")
	} // End of the `for` loop

	// 4. Read and print the final file content:
	// os.ReadFile reads the entire file content as bytes.
	data, err := os.ReadFile(filename)
	// Check for errors when reading the file.
	check(err)

	// Print the file content to the console by converting bytes to string.
	fmt.Println("--- File Content ---")
	fmt.Println(string(data))
	fmt.Println("--------------------")
}

// Helper function to handle errors.
func check(e error) {
	// If the error is not `nil` (i.e., an error occurred).
	if e != nil {
		// Stop the program and display the error message.
		panic(e)
	}
}

// Function to get note input from the user.
func UserNote() string {
	var note string
	// Display a prompt to the user.
	fmt.Print("Enter your note(0 for EXIT): ")
	// Create a new scanner to read from standard input (keyboard).
	scanner := bufio.NewScanner(os.Stdin)
	// Read the next line of input.
	scanner.Scan()
	// Store the read text in the `note` variable.
	note = scanner.Text()
	// Return the entered note.
	return note
}