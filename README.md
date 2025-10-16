# Go Learning Journey

This repository documents my personal journey of learning the Go programming language —  
from basic syntax and simple programs to more structured exercises and small projects.

---

## 🧩 Learning Stages

### Stage 1 — Getting Started
Learned how to set up the Go environment, initialize a project, and run a simple “Hello World” program.

### Stage 2 — Variables & Time Formatting
Practiced declaring and using variables, reading user input, formatting strings, and displaying the current system time.

### Stage 3 — Conditions & Loops
Implemented logic using `if` statements and `for` loops to process multiple user inputs and control program flow.

### Stage 4 — Functions & Slices
Combined loops, conditions, and reusable functions.  
Used slices to dynamically store user inputs and calculate both average and maximum values.

### Stage 5 — File Handling
Learned how to create, open, and append data to text files using the `os` and `bufio` packages.  
Built a simple Notes App that saves user input to a file and displays all saved notes at the end.

### Stage 6 — Structs & Methods
Learned how to define custom data types using `struct` and attach methods to them.  
Built a simple Student Manager program that stores, displays, and updates student information using methods with both value and pointer receivers.

### Stage 7 — Slices of Structs
Managed a list of `Student` records using slices.  
Implemented functions to add, find, and remove students,  
as well as to calculate maximum and average grades.  
Practiced struct iteration, conditional logic, and slice manipulation in Go.

### Stage 8 — JSON Persistence
Implemented file-based data persistence using the `encoding/json` package.  
Learned how to convert Go structs and slices into JSON and save them to disk using `json.MarshalIndent` and `os.WriteFile`.  
Also practiced reading JSON data back into Go structs with `json.Unmarshal`, enabling data to persist between program runs.

### Project 1 — Price Processor

A modular Go program that reads price inputs, applies multiple tax rates, and outputs tax-included results through an I/O interface. The project demonstrates clean package organization, use of interfaces (`IOManager`) for flexible input/output handling, and a simple data-processing pipeline from reading user input to generating structured results.

---

## ⚙️ Tech Stack
- **Language:** Go (Golang)  
- **Environment:** Visual Studio Code  
- **Version Control:** Git & GitHub  

---
