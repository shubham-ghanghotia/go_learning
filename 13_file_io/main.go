package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== FILE I/O ===\n")

	// Example 1: Write to file
	fmt.Println("1. Writing to File:")
	content := "Hello, File I/O!\nThis is line 2.\nThis is line 3.\n"
	filename := "example.txt"

	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Printf("File '%s' written successfully\n\n", filename)

	// Example 2: Read entire file
	fmt.Println("2. Reading Entire File:")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("File contents:\n%s\n", string(data))

	// Example 3: Read file line by line
	fmt.Println("3. Reading File Line by Line:")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	fmt.Println()

	// Example 4: Append to file
	fmt.Println("4. Appending to File:")
	appendContent := "\nAppended line!"
	err = ioutil.WriteFile(filename, append(data, []byte(appendContent)...), 0644)
	if err != nil {
		fmt.Println("Error appending:", err)
		return
	}
	fmt.Println("Content appended successfully\n")

	// Example 5: File information
	fmt.Println("5. File Information:")
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Printf("Filename: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Mode: %v\n", fileInfo.Mode())
	fmt.Printf("Is Directory: %v\n\n", fileInfo.IsDir())

	// Example 6: Check if file exists
	fmt.Println("6. Checking File Existence:")
	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("File '%s' exists\n", filename)
	} else if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", filename)
	}
	fmt.Println()

	// Example 7: Create directory
	fmt.Println("7. Creating Directory:")
	dirName := "mydir"
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Printf("Directory '%s' created successfully\n", dirName)
	}
	fmt.Println()

	// Example 8: List directory contents
	fmt.Println("8. Listing Directory Contents:")
	entries, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Current directory contents:")
	for i, entry := range entries {
		if i >= 5 { // Show first 5 entries
			fmt.Println("...")
			break
		}
		fileType := "file"
		if entry.IsDir() {
			fileType = "dir"
		}
		fmt.Printf("  %s (%s) - %d bytes\n", entry.Name(), fileType, entry.Size())
	}
	fmt.Println()

	// Example 9: Write formatted content
	fmt.Println("9. Writing Formatted Content:")
	csvContent := "Name,Age,City\nAlice,30,New York\nBob,25,San Francisco\n"
	csvFile := "data.csv"

	err = ioutil.WriteFile(csvFile, []byte(csvContent), 0644)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}

	data, _ = ioutil.ReadFile(csvFile)
	fmt.Printf("CSV content:\n%s\n", string(data))

	// Example 10: Cleanup
	fmt.Println("10. Cleanup:")
	os.Remove(filename)
	os.Remove(csvFile)
	os.Remove(dirName)
	fmt.Println("Cleanup completed - test files removed")
}
