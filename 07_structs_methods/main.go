package main

import "fmt"

// Basic struct definition
type Person struct {
	Name string
	Age  int
	City string
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Method with value receiver (doesn't modify)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Method with pointer receiver (can modify)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Embedded struct (composition)
type Address struct {
	Street string
	City   string
	ZIP    string
}

type Employee struct {
	Name    string
	Age     int
	Salary  float64
	Address Address
}

// Method on Employee
func (e Employee) DisplayInfo() {
	fmt.Printf("Employee: %s, Age: %d, Salary: $%.2f\n", e.Name, e.Age, e.Salary)
	fmt.Printf("Lives at: %s, %s %s\n", e.Address.Street, e.Address.City, e.Address.ZIP)
}

// Pointer receiver - modifies the struct
func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * (percent / 100)
}

// Struct with tags
type User struct {
	ID       int    `json:"id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

func main() {
	fmt.Println("=== STRUCTS AND METHODS ===")
	fmt.Println()

	// Example 1: Basic struct creation and access
	fmt.Println("1. Creating Structs:")
	person := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Printf("Name: %s, Age: %d, City: %s\n", person.Name, person.Age, person.City)

	// Modifying fields
	person.Age = 31
	fmt.Printf("After birthday: Age = %d\n\n", person.Age)

	// Example 2: Methods with value receiver
	fmt.Println("2. Methods (Value Receiver):")
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Printf("Rectangle: Width=%.1f, Height=%.1f\n", rect.Width, rect.Height)
	fmt.Printf("Area: %.1f\n", rect.Area())
	fmt.Printf("Perimeter: %.1f\n\n", rect.Perimeter())

	// Example 3: Methods with pointer receiver
	fmt.Println("3. Methods (Pointer Receiver):")
	fmt.Printf("Before Scale: Width=%.1f, Height=%.1f\n", rect.Width, rect.Height)
	rect.Scale(2)
	fmt.Printf("After Scale(2): Width=%.1f, Height=%.1f\n\n", rect.Width, rect.Height)

	// Example 4: Embedded structs
	fmt.Println("4. Embedded Structs (Composition):")
	emp := Employee{
		Name:   "Bob",
		Age:    28,
		Salary: 50000,
		Address: Address{
			Street: "456 Oak Ave",
			City:   "San Francisco",
			ZIP:    "94102",
		},
	}
	emp.DisplayInfo()
	fmt.Println()

	// Example 5: Pointer receiver that modifies
	fmt.Println("5. Pointer Receiver (Modifying):")
	fmt.Printf("Current Salary: $%.2f\n", emp.Salary)
	emp.GiveRaise(10)
	fmt.Printf("After 10%% raise: $%.2f\n\n", emp.Salary)

	// Example 6: Zero values
	fmt.Println("6. Zero Values:")
	var emptyPerson Person
	fmt.Printf("Empty Person: %+v\n", emptyPerson)
	fmt.Printf("Name: '%s', Age: %d\n\n", emptyPerson.Name, emptyPerson.Age)

	// Example 7: Struct with tags (just showing the structure)
	fmt.Println("7. Struct with Tags:")
	user := User{ID: 1, Username: "alice", Email: "alice@example.com"}
	fmt.Printf("User: %+v\n", user)
}
