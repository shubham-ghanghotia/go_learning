package main

import (
	"errors"
	"fmt"
	"strings"
)

// Example 1: Using errors.New()
var errDivisionByZero = errors.New("division by zero")
var errUserNotFound = errors.New("user not found")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errDivisionByZero
	}
	return a / b, nil
}

// Example 2: Using fmt.Errorf()
func Multiply(a, b float64) (float64, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("values must be positive: a=%.2f, b=%.2f", a, b)
	}
	return a * b, nil
}

// Example 3: Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in '%s': %s", e.Field, e.Message)
}

func ValidateEmail(email string) error {
	if email == "" {
		return ValidationError{
			Field:   "email",
			Message: "cannot be empty",
		}
	}
	if !strings.Contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Message: "must contain @",
		}
	}
	if !strings.Contains(email, ".") {
		return ValidationError{
			Field:   "email",
			Message: "must contain domain extension",
		}
	}
	return nil
}

// Example 4: Wrapping errors
type User struct {
	ID    int
	Name  string
	Email string
}

func FindUser(id int) (User, error) {
	if id < 1 {
		return User{}, fmt.Errorf("invalid user id: %w", errUserNotFound)
	}
	// Simulate database lookup
	if id == 1 {
		return User{ID: 1, Name: "Alice", Email: "alice@example.com"}, nil
	}
	return User{}, fmt.Errorf("user with id %d not found: %w", id, errUserNotFound)
}

// Example 5: Error handling in operations
func ProcessUser(id int, email string) error {
	// Validate email first
	if err := ValidateEmail(email); err != nil {
		return fmt.Errorf("email validation failed: %w", err)
	}

	// Find user
	user, err := FindUser(id)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	user.Email = email
	fmt.Printf("Updated user: %+v\n", user)
	return nil
}

// Example 6: Panic and Recover
func DangerousOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("Starting dangerous operation...")
	panic("Something went wrong!")
	fmt.Println("This never executes")
}

func SafeDivide(a, b float64) (result float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from divide panic: %v\n", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("cannot divide by zero")
	}
	return a / b
}

func main() {
	fmt.Println("=== ERROR HANDLING ===\n")

	// Example 1: Basic error handling with errors.New()
	fmt.Println("1. Basic Error Handling (errors.New):")
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// Example 2: fmt.Errorf()
	fmt.Println("2. Using fmt.Errorf():")
	result, err = Multiply(5, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("5 * 3 = %.2f\n", result)
	}

	result, err = Multiply(-5, 3)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// Example 3: Custom error types
	fmt.Println("3. Custom Error Types:")
	validEmails := []string{"alice@example.com", "bob@test.org", "", "invalid", "test@"}

	for _, email := range validEmails {
		err := ValidateEmail(email)
		if err != nil {
			fmt.Printf("Email '%s' - Error: %v\n", email, err)
		} else {
			fmt.Printf("Email '%s' is valid\n", email)
		}
	}
	fmt.Println()

	// Example 4: Error wrapping
	fmt.Println("4. Error Wrapping:")
	user, err := FindUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found: %+v\n", user)
	}

	user, err = FindUser(999)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// Example 5: Error handling in complex operations
	fmt.Println("5. Complex Error Handling:")
	err = ProcessUser(1, "alice.new@example.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = ProcessUser(1, "invalid-email")
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// Example 6: Panic and Recover
	fmt.Println("6. Panic and Recover:")
	DangerousOperation()
	fmt.Println("Program continues after panic\n")

	// Example 7: Safe divide with recovery
	fmt.Println("7. Safe Division (with recovery):")
	result = SafeDivide(10, 2)
	fmt.Printf("10 / 2 = %.2f\n", result)

	result = SafeDivide(10, 0)
	fmt.Printf("10 / 0 = %.2f\n", result)
}
