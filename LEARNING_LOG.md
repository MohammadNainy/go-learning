# Learning Log

## Stage 1 — Getting Started  (Practice1)
- Installed and configured the Go environment  
- Verified setup using `go version`  
- Created base folders: `Basic/` and `Projects/`  
- Initialized module using `go mod init`  
- Wrote the first Go program (`Hello, World!`) using `fmt.Println`  
- Created documentation files (`README.md`, `LEARNING_LOG.md`, `.gitignore`)  

---

## Stage 2 — Variables & Time Formatting  (Practice2)
- Practiced variable declaration with `var` and shorthand `:=`  
- Collected user input using `fmt.Scanln`  
- Used `fmt.Printf` for formatted string output  
- Retrieved and formatted current time with `time.Now().Format()`  
- Added inline code comments following Go documentation conventions  

---

## Stage 3 — Conditions & Loops  (Practice3)
- Learned to control program flow using `if` and `else` statements  
- Implemented repetitive input with an infinite `for` loop  
- Used `break` to terminate user input when a negative number was entered  
- Calculated the total sum of all positive inputs  
- Enhanced user interaction and program clarity through clear console output  

---

## Stage 4 — Functions & Slices  (Practice4)
- Used `slice` to dynamically store user input values  
- Implemented `append()` for adding items to slices  
- Defined custom functions to calculate average and maximum values  
- Utilized the new `slices` package (`slices.Max()`) for simplicity  
- Practiced handling empty input cases to prevent runtime errors  
- Improved modular design and readability through function separation  

---

## Stage 5 — File Handling  (Practice5)
- Learned how to create, open, and write to files using the `os` package  
- Used `os.OpenFile()` with flags `O_APPEND`, `O_CREATE`, and `O_WRONLY` to safely append data  
- Implemented `defer f.Close()` to ensure proper resource management  
- Handled new or empty files by checking file size before writing headers  
- Utilized `bufio.Scanner` for full-line user input (supporting spaces in notes)  
- Implemented `break` and `continue` logic for better user control during input  
- Used `os.ReadFile()` to read and display all stored notes after writing  
- Practiced safe error handling with a helper function `check(err)`  
- Added detailed code comments to document program flow and structure

---

## Stage 6 — Structs & Methods  (Practice6)
- Defined a custom data type `Student` using `struct` to represent complex data  
- Implemented methods for displaying information and checking logical conditions  
- Practiced using **value receivers** for read-only operations (`Display`, `IsAdult`)  
- Used **pointer receivers** for modifying data (`UpdateGrade`)  
- Learned to separate logic and output for cleaner, modular code  
- Improved understanding of data encapsulation and method organization in Go  
- Followed Go naming conventions for exported types and fields  

---

## Stage 7 — Slices of Structs  (Practice7)
- Created a slice of `Student` structs to represent a list of students  
- Implemented `AddStudent` to dynamically add new items to the list  
- Added `FindByName` and `RemoveByName` for basic search and delete operations  
- Practiced slice manipulation using `append()` and index slicing  
- Implemented `MaxGrade` and `AverageGrade` for data aggregation and analysis  
- Added empty-list checks and improved function readability  
- Strengthened understanding of slice iteration, data filtering, and Go conventions  

---

## Stage 8 — JSON Persistence  (Practice8)
- Combined file handling and structs to create persistent data storage  
- Used the `encoding/json` package to serialize and deserialize data  
- Added JSON tags to `Student` fields for cleaner, lowercase keys  
- Implemented `SaveStudents()` using `json.MarshalIndent()` for readable output  
- Implemented `LoadStudents()` with proper error handling and empty-file checks  
- Practiced returning `(value, error)` patterns instead of using `panic`  
- Ensured data persists between executions by saving and reloading from `students.json`  
- Improved program structure with a clear flow: **Load → Modify → Save**

---