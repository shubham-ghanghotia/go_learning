package main

import (
	"fmt"
)

// Lesson 4: Functions
// Functions are reusable blocks of code that perform specific tasks

// === BASIC FUNCTION ===
// Syntax: func functionName(parameter type) returnType { }

// Simple function with no parameters or return value
func greet() {
	fmt.Println("Hello! Welcome to Go functions.")
}

// Function with parameters
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a int, b int) int {
	return a + b
}

// Shorthand: Same type parameters
func multiply(x, y int) int {
	return x * y
}

// Function with multiple return values
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Error: Cannot divide by zero"
	}
	return a / b, "Success"
}

// Named return values
func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Returns area and perimeter
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function with default-like behavior
func power(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

// Recursive function (calls itself)
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Function that returns a function (higher-order function)
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	// === CALLING FUNCTIONS ===
	fmt.Println("=== BASIC FUNCTIONS ===")

	// No parameters, no return
	greet()

	// With parameters
	greetPerson("Alice")
	greetPerson("Bob")

	// With return value
	fmt.Println("\n=== FUNCTIONS WITH RETURN VALUES ===")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// Multiple parameters (shorthand)
	product := multiply(4, 7)
	fmt.Printf("4 × 7 = %d\n", product)

	// === MULTIPLE RETURN VALUES ===
	fmt.Println("\n=== MULTIPLE RETURN VALUES ===")

	quotient, status := divide(10, 2)
	fmt.Printf("10 ÷ 2 = %f, Status: %s\n", quotient, status)

	quotient2, status2 := divide(10, 0)
	fmt.Printf("10 ÷ 0 = %f, Status: %s\n", quotient2, status2)

	// Ignoring return values
	quotient3, _ := divide(15, 3) // _ ignores the status
	fmt.Printf("15 ÷ 3 = %f\n", quotient3)

	// === NAMED RETURN VALUES ===
	fmt.Println("\n=== NAMED RETURN VALUES ===")
	a, p := rectangle(5, 3)
	fmt.Printf("Rectangle (5×3) - Area: %f, Perimeter: %f\n", a, p)

	// === VARIADIC FUNCTIONS ===
	fmt.Println("\n=== VARIADIC FUNCTIONS ===")

	fmt.Printf("Sum of 1,2,3: %d\n", sum(1, 2, 3))
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("Sum of 10,20: %d\n", sum(10, 20))

	// Pass slice with ...
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of slice: %d\n", sum(numbers...))

	// === RECURSIVE FUNCTIONS ===
	fmt.Println("\n=== RECURSIVE FUNCTIONS ===")

	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Factorial of 3: %d\n", factorial(3))

	// === HIGHER-ORDER FUNCTIONS ===
	fmt.Println("\n=== HIGHER-ORDER FUNCTIONS ===")

	doubler := makeMultiplier(2)
	tripler := makeMultiplier(3)

	fmt.Printf("Double 5: %d\n", doubler(5))
	fmt.Printf("Triple 5: %d\n", tripler(5))

	// === ANONYMOUS FUNCTIONS ===
	fmt.Println("\n=== ANONYMOUS FUNCTIONS ===")

	func(name string) {
		fmt.Printf("Hello from anonymous function, %s!\n", name)
	}("Charlie")

	// Store anonymous function in variable
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 4: %d\n", square(4))

	// === DEFER ===
	fmt.Println("\n=== DEFER ===")

	fmt.Println("Start")
	defer fmt.Println("This prints last (deferred)")
	fmt.Println("Middle")
	fmt.Println("This prints before deferred")
}
