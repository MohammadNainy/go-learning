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
