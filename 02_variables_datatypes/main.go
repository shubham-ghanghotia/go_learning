package main

import (
	"fmt"
)

// Lesson 2: Variables and Data Types

func main() {
	// === DECLARING VARIABLES ===

	// Method 1: Using var keyword (explicit type)
	var name string = "Alice"
	var age int = 25
	var height float64 = 5.8

	fmt.Println("=== Method 1: var keyword ===")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)

	// Method 2: Using var keyword (type inference)
	var city = "New York" // Go figures out it's a string
	var temperature = 72  // Go figures out it's an int

	fmt.Println("\n=== Method 2: var with type inference ===")
	fmt.Println("City:", city)
	fmt.Println("Temperature:", temperature)

	// Method 3: Using := (short declaration - PREFERRED in functions)
	country := "USA"
	population := 8000000
	isRaining := false

	fmt.Println("\n=== Method 3: short declaration (:=) ===")
	fmt.Println("Country:", country)
	fmt.Println("Population:", population)
	fmt.Println("Is Raining:", isRaining)

	// === DATA TYPES ===

	fmt.Println("\n=== DATA TYPES ===")

	// Integers
	var count int = 42
	var smallNum int8 = 127
	var largeNum int64 = 9223372036854775807
	unsignedNum := uint(100) // Only positive numbers

	fmt.Println("\n--- Integers ---")
	fmt.Printf("int: %v\n", count)
	fmt.Printf("int8: %v\n", smallNum)
	fmt.Printf("int64: %v\n", largeNum)
	fmt.Printf("uint: %v\n", unsignedNum)

	// Floating Point
	var pi float32 = 3.14
	var e float64 = 2.71828

	fmt.Println("\n--- Floating Point ---")
	fmt.Printf("float32: %v\n", pi)
	fmt.Printf("float64: %v\n", e)

	// Strings
	var message string = "Hello, Go!"
	var emoji string = "🎉"

	fmt.Println("\n--- Strings ---")
	fmt.Printf("message: %v\n", message)
	fmt.Printf("emoji: %v\n", emoji)

	// Booleans
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println("\n--- Booleans ---")
	fmt.Printf("isTrue: %v\n", isTrue)
	fmt.Printf("isFalse: %v\n", isFalse)

	// === CONSTANTS ===
	fmt.Println("\n=== CONSTANTS ===")

	const PI float64 = 3.14159
	const AUTHOR string = "Go Team"
	const MAX_USERS = 1000 // Type inferred as int

	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("AUTHOR: %v\n", AUTHOR)
	fmt.Printf("MAX_USERS: %v\n", MAX_USERS)

	// === TYPE CONVERSION ===
	fmt.Println("\n=== TYPE CONVERSION ===")

	var intValue int = 42
	var floatValue float64 = float64(intValue) // Convert int to float64
	var stringValue string = fmt.Sprintf("%d", intValue) // Convert int to string

	fmt.Printf("int to float: %v\n", floatValue)
	fmt.Printf("int to string: %v\n", stringValue)

	// === ZERO VALUES ===
	fmt.Println("\n=== ZERO VALUES (default values) ===")

	var defaultInt int
	var defaultString string
	var defaultBool bool
	var defaultFloat float64

	fmt.Printf("default int: %v\n", defaultInt)
	fmt.Printf("default string: '%v'\n", defaultString) // Empty string
	fmt.Printf("default bool: %v\n", defaultBool)       // false
	fmt.Printf("default float: %v\n", defaultFloat)     // 0.0
}
