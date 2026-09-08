package main

import (
	"fmt"
)

// This is your first Go program!
// Let's break down what's happening:

// 1. package main - Every executable Go program must have a main package
// 2. import - We're importing the "fmt" package for formatted output
// 3. func main() - This is the entry point of our program

func main() {
	// fmt.Println prints a line of text and adds a newline
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go!")

	// You can also use Printf for formatted printing
	name := "Gopher"
	fmt.Printf("Hello, %s!\n", name)
}
