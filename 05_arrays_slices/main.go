package main

import (
	"fmt"
)

// Lesson 5: Arrays and Slices
// Arrays: Fixed size, Slices: Dynamic size

func main() {
	// === ARRAYS ===
	fmt.Println("=== ARRAYS ===")
	fmt.Println("Arrays have fixed length")

	// Declare array with size
	var numbers [3]int
	fmt.Printf("Empty array: %v\n", numbers) // [0 0 0]

	// Initialize array
	var fruits [3]string = [3]string{"apple", "banana", "cherry"}
	fmt.Printf("Fruits: %v\n", fruits)

	// Short declaration
	colors := [4]string{"red", "green", "blue", "yellow"}
	fmt.Printf("Colors: %v\n", colors)

	// Array with ... (compiler counts elements)
	scores := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("Scores: %v (length: %d)\n", scores, len(scores))

	// Accessing elements
	fmt.Println("\n--- Accessing Array Elements ---")
	fmt.Printf("First fruit: %s\n", fruits[0])
	fmt.Printf("Last fruit: %s\n", fruits[len(fruits)-1])

	// Modifying elements
	fmt.Println("\n--- Modifying Array Elements ---")
	fruits[1] = "mango"
	fmt.Printf("Modified fruits: %v\n", fruits)

	// Iterating over arrays
	fmt.Println("\n--- Iterating Over Arrays ---")

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index %d: %s\n", i, fruits[i])
	}

	fmt.Println("\nUsing range:")
	for index, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", index, fruit)
	}

	// === SLICES ===
	fmt.Println("\n=== SLICES ===")
	fmt.Println("Slices are dynamic arrays")

	// Declare slice (no size)
	var emptySlice []int
	fmt.Printf("Empty slice: %v (length: %d, capacity: %d)\n", emptySlice, len(emptySlice), cap(emptySlice))

	// Initialize slice
	var numbers2 []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers: %v (length: %d, capacity: %d)\n", numbers2, len(numbers2), cap(numbers2))

	// Short declaration
	vegetables := []string{"carrot", "broccoli", "spinach"}
	fmt.Printf("Vegetables: %v\n", vegetables)

	// Using make to create slice
	mySlice := make([]int, 5)           // Length 5, capacity 5
	fmt.Printf("Slice with make: %v (len: %d, cap: %d)\n", mySlice, len(mySlice), cap(mySlice))

	mySlice2 := make([]int, 3, 10)      // Length 3, capacity 10
	fmt.Printf("Slice with capacity: %v (len: %d, cap: %d)\n", mySlice2, len(mySlice2), cap(mySlice2))

	// === SLICE OPERATIONS ===
	fmt.Println("\n--- Slice Operations ---")

	// Append
	fruits2 := []string{"apple", "banana"}
	fmt.Printf("Before append: %v\n", fruits2)
	fruits2 = append(fruits2, "orange")
	fmt.Printf("After append: %v\n", fruits2)

	// Append multiple
	fruits2 = append(fruits2, "grape", "kiwi")
	fmt.Printf("After adding multiple: %v\n", fruits2)

	// Append another slice
	morefruits := []string{"mango", "pineapple"}
	fruits2 = append(fruits2, morefruits...)
	fmt.Printf("After append slice: %v\n", fruits2)

	// === SLICING ===
	fmt.Println("\n--- Slicing Syntax ---")
	numbers3 := []int{10, 20, 30, 40, 50, 60}
	fmt.Printf("Original: %v\n", numbers3)

	fmt.Printf("numbers[1:3]: %v (elements at index 1, 2)\n", numbers3[1:3])
	fmt.Printf("numbers[:3]: %v (from start to index 2)\n", numbers3[:3])
	fmt.Printf("numbers[3:]: %v (from index 3 to end)\n", numbers3[3:])
	fmt.Printf("numbers[:]: %v (entire slice)\n", numbers3[:])

	// === COPYING SLICES ===
	fmt.Println("\n--- Copying Slices ---")

	original := []int{1, 2, 3}
	copyByReference := original  // Just a reference
	copyByValue := make([]int, len(original))
	copy(copyByValue, original)  // Creates true copy

	original[0] = 999
	fmt.Printf("Original: %v\n", original)           // [999 2 3]
	fmt.Printf("By reference: %v\n", copyByReference) // [999 2 3] (affected!)
	fmt.Printf("By value: %v\n", copyByValue)        // [1 2 3] (not affected)

	// === LENGTH vs CAPACITY ===
	fmt.Println("\n--- Length vs Capacity ---")

	mySlice3 := make([]int, 3, 10)
	fmt.Printf("Initial: %v (len: %d, cap: %d)\n", mySlice3, len(mySlice3), cap(mySlice3))

	// Length cannot exceed capacity
	mySlice3 = append(mySlice3, 1, 2, 3, 4, 5, 6, 7)  // 7 more elements
	fmt.Printf("After append: %v (len: %d, cap: %d)\n", mySlice3, len(mySlice3), cap(mySlice3))

	// === 2D SLICES ===
	fmt.Println("\n--- 2D Arrays/Slices ---")

	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("Matrix: %v\n", matrix)
	fmt.Printf("Element [0][1]: %d\n", matrix[0][1])

	// 2D slice
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Grid: %v\n", grid)
	fmt.Printf("Element [2][2]: %d\n", grid[2][2])
}
