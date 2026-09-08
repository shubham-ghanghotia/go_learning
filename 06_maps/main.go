package main

import (
	"fmt"
)

// Lesson 6: Maps
// Maps are key-value pairs (like dictionaries in Python, objects in JavaScript)

func main() {
	// === DECLARING MAPS ===
	fmt.Println("=== DECLARING MAPS ===")

	// Method 1: Using var
	var person map[string]string
	fmt.Printf("Empty map: %v\n", person) // nil

	// Method 2: Using make
	countries := make(map[string]string)
	fmt.Printf("Map with make: %v\n", countries)

	// Method 3: Initialize with values
	person = map[string]string{
		"name":  "Alice",
		"email": "alice@example.com",
		"city":  "New York",
	}
	fmt.Printf("Person: %v\n", person)

	// Short declaration
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}
	fmt.Printf("Scores: %v\n", scores)

	// === ACCESSING VALUES ===
	fmt.Println("\n=== ACCESSING VALUES ===")

	fmt.Printf("Person's name: %s\n", person["name"])
	fmt.Printf("Alice's score: %d\n", scores["Alice"])

	// Accessing non-existent key returns zero value
	fmt.Printf("Non-existent key: '%s'\n", person["age"])

	// Check if key exists (ok idiom)
	value, exists := person["name"]
	fmt.Printf("Name exists: %v, Value: %s\n", exists, value)

	value2, exists2 := person["phone"]
	fmt.Printf("Phone exists: %v, Value: %s\n", exists2, value2)

	// === ADDING AND MODIFYING ===
	fmt.Println("\n=== ADDING AND MODIFYING ===")

	// Add new key-value
	person["age"] = "30"
	fmt.Printf("After adding age: %v\n", person)

	// Modify existing
	person["city"] = "Boston"
	fmt.Printf("After updating city: %v\n", person)

	// Add to scores
	scores["David"] = 88
	scores["Eve"] = 99
	fmt.Printf("After adding scores: %v\n", scores)

	// === DELETING ===
	fmt.Println("\n=== DELETING ===")

	fmt.Printf("Before delete: %v\n", scores)
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob: %v\n", scores)

	// Deleting non-existent key is safe
	delete(scores, "NonExistent")
	fmt.Println("No error deleting non-existent key")

	// === LENGTH ===
	fmt.Println("\n=== LENGTH ===")
	fmt.Printf("Number of people: %d\n", len(person))
	fmt.Printf("Number of scores: %d\n", len(scores))

	// === ITERATING ===
	fmt.Println("\n=== ITERATING ===")

	fmt.Println("\nPerson details:")
	for key, value := range person {
		fmt.Printf("  %s: %s\n", key, value)
	}

	fmt.Println("\nScores:")
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// Iterate only keys
	fmt.Println("\nOnly keys:")
	for key := range person {
		fmt.Printf("  %s\n", key)
	}

	// Iterate only values
	fmt.Println("\nOnly values:")
	for _, value := range person {
		fmt.Printf("  %s\n", value)
	}

	// === MAP OF SLICES ===
	fmt.Println("\n=== MAP OF SLICES ===")

	classes := map[string][]string{
		"Math": {"Alice", "Bob", "Carol"},
		"Science": {"David", "Eve"},
	}
	fmt.Printf("Classes: %v\n", classes)
	fmt.Printf("Math students: %v\n", classes["Math"])

	// === NESTED MAPS ===
	fmt.Println("\n=== NESTED MAPS ===")

	contacts := map[string]map[string]string{
		"Alice": {
			"phone": "555-1234",
			"email": "alice@example.com",
		},
		"Bob": {
			"phone": "555-5678",
			"email": "bob@example.com",
		},
	}
	fmt.Printf("Alice's phone: %s\n", contacts["Alice"]["phone"])
	fmt.Printf("Bob's email: %s\n", contacts["Bob"]["email"])

	// === MAP WITH DIFFERENT TYPES ===
	fmt.Println("\n=== MAP WITH DIFFERENT TYPES ===")

	// Map with interface{} (can store any type)
	data := map[string]interface{}{
		"name": "Alice",
		"age":  30,
		"score": 95.5,
		"active": true,
	}
	fmt.Printf("Mixed type map: %v\n", data)
}
