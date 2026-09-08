package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Demonstrating package organization concepts

// Package-level variable
var packageVar = "I'm at package level"

// init() runs automatically when package is initialized
func init() {
	fmt.Println("Package initialized!")
	fmt.Println()
}

// Exported function (starts with uppercase)
func ExportedFunction() {
	fmt.Println("This function is exported")
}

// unexported function (starts with lowercase)
func unexportedFunction() {
	fmt.Println("This function is only accessible within the package")
}

// Struct with exported and unexported fields
type Calculator struct {
	LastResult float64 // Exported field
	history    string  // Unexported field
}

// Method on Calculator
func (c *Calculator) Add(a, b float64) float64 {
	c.LastResult = a + b
	c.history = fmt.Sprintf("Added %.2f + %.2f", a, b)
	return c.LastResult
}

func (c *Calculator) GetHistory() string {
	return c.history
}

func main() {
	fmt.Println("=== PACKAGES & IMPORTS ===\n")

	// Example 1: Using fmt package
	fmt.Println("1. Using fmt Package:")
	fmt.Println("Hello from fmt.Println")
	fmt.Printf("Formatted output: %d + %d = %d\n\n", 5, 3, 8)

	// Example 2: Using math package
	fmt.Println("2. Using math Package:")
	fmt.Printf("sqrt(16) = %.2f\n", math.Sqrt(16))
	fmt.Printf("Pi = %.4f\n", math.Pi)
	fmt.Printf("Max(10, 20) = %.0f\n", math.Max(10, 20))
	fmt.Printf("Pow(2, 3) = %.0f\n\n", math.Pow(2, 3))

	// Example 3: Using strings package
	fmt.Println("3. Using strings Package:")
	text := "hello world"
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("ToUpper: %s\n", strings.ToUpper(text))
	fmt.Printf("Title: %s\n", strings.Title(text))
	fmt.Printf("Contains 'world': %v\n", strings.Contains(text, "world"))
	fmt.Printf("Count 'l': %d\n\n", strings.Count(text, "l"))

	// Example 4: Using strconv package
	fmt.Println("4. Using strconv Package:")
	numStr := "42"
	num, _ := strconv.Atoi(numStr)
	fmt.Printf("String '%s' converted to int: %d\n", numStr, num)

	intToStr := strconv.Itoa(123)
	fmt.Printf("Int 123 converted to string: '%s'\n", intToStr)

	floatStr, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("String '3.14' converted to float: %.2f\n\n", floatStr)

	// Example 5: Using time package
	fmt.Println("5. Using time Package:")
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)
	fmt.Printf("Year: %d, Month: %v, Day: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n\n", now.Hour(), now.Minute(), now.Second())

	// Example 6: Custom package concepts (within main)
	fmt.Println("6. Custom Package Concepts (Exported/Unexported):")
	ExportedFunction()
	unexportedFunction()
	fmt.Println()

	// Example 7: Struct with exported/unexported fields
	fmt.Println("7. Exported/Unexported Struct Fields:")
	calc := Calculator{}
	result := calc.Add(10, 20)
	fmt.Printf("Calculation result: %.2f\n", result)
	fmt.Printf("Last result field: %.2f\n", calc.LastResult)
	fmt.Printf("History: %s\n", calc.GetHistory())
	fmt.Println()

	// Example 8: Package variable
	fmt.Println("8. Package-Level Variable:")
	fmt.Println(packageVar)
}
