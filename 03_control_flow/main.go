package main

import (
	"fmt"
)

// Lesson 3: Control Flow (if/else, for, switch)

func main() {
	// === IF STATEMENTS ===
	fmt.Println("=== IF STATEMENTS ===")

	age := 20

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// IF-ELSE
	fmt.Println("\n--- IF-ELSE ---")
	temperature := 15

	if temperature > 25 {
		fmt.Println("It's hot!")
	} else if temperature > 15 {
		fmt.Println("It's warm")
	} else {
		fmt.Println("It's cold")
	}

	// IF with variable initialization
	fmt.Println("\n--- IF with Variable Initialization ---")

	if score := 85; score >= 80 {
		fmt.Printf("Great score: %d\n", score)
	}

	// === FOR LOOPS ===
	fmt.Println("\n=== FOR LOOPS ===")

	// Traditional for loop
	fmt.Println("\n--- Traditional For Loop ---")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// While-like for loop (no init/increment)
	fmt.Println("\n--- While-like For Loop ---")
	count := 0
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// Infinite loop with break
	fmt.Println("\n--- For Loop with Break ---")
	for {
		fmt.Println("This will print once")
		break // Exit the loop
	}

	// For loop with continue
	fmt.Println("\n--- For Loop with Continue ---")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip this iteration
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Range loops (for slices/arrays - more in lesson 5)
	fmt.Println("\n--- Range Loop ---")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", index, fruit)
	}

	// === SWITCH STATEMENTS ===
	fmt.Println("\n=== SWITCH STATEMENTS ===")

	day := 3

	fmt.Println("\n--- Basic Switch ---")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	// Switch with multiple cases
	fmt.Println("\n--- Switch with Multiple Cases ---")
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B", "C":
		fmt.Println("Good job!")
	case "D":
		fmt.Println("Needs improvement")
	default:
		fmt.Println("Invalid grade")
	}

	// Switch with no expression (like if-else)
	fmt.Println("\n--- Switch with No Expression ---")
	marks := 75
	switch {
	case marks >= 80:
		fmt.Println("A grade")
	case marks >= 60:
		fmt.Println("B grade")
	case marks >= 40:
		fmt.Println("C grade")
	default:
		fmt.Println("Fail")
	}

	// === LOGICAL OPERATORS ===
	fmt.Println("\n=== LOGICAL OPERATORS ===")

	x := 10
	y := 20

	if x > 5 && y > 15 {
		fmt.Println("Both conditions are true (AND)")
	}

	if x > 5 || y < 15 {
		fmt.Println("At least one condition is true (OR)")
	}

	if !(x > 20) {
		fmt.Println("x is not greater than 20 (NOT)")
	}
}
