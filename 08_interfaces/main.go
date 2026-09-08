package main

import (
	"fmt"
	"math"
)

// Interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Square struct
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Animal interface
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: Woof! Woof!"
}

func (d Dog) Move() string {
	return d.Name + " runs with four legs"
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return b.Name + " says: Tweet! Tweet!"
}

func (b Bird) Move() string {
	return b.Name + " flies in the sky"
}

func main() {
	fmt.Println("=== INTERFACES ===\n")

	// Example 1: Using interface with different types
	fmt.Println("1. Shape Interface:")
	shapes := []Shape{
		Circle{Radius: 5},
		Square{Side: 4},
		Rectangle{Width: 3, Height: 6},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}
	fmt.Println()

	// Example 2: Animal interface
	fmt.Println("2. Animal Interface:")
	animals := []Animal{
		Dog{Name: "Buddy"},
		Bird{Name: "Tweety"},
		Dog{Name: "Max"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
		fmt.Println()
	}

	// Example 3: Empty interface (interface{})
	fmt.Println("3. Empty Interface (Can hold any type):")
	var anything interface{} = "Hello"
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)

	anything = 42
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)

	anything = 3.14
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)

	anything = []int{1, 2, 3}
	fmt.Printf("Value: %v, Type: %T\n\n", anything, anything)

	// Example 4: Type assertion
	fmt.Println("4. Type Assertion:")
	var i interface{} = "hello"

	str, ok := i.(string)
	if ok {
		fmt.Printf("Successfully asserted to string: %s\n", str)
	}

	num, ok := i.(int)
	if !ok {
		fmt.Printf("Not an int (as expected), got type %T\n\n", i)
	}

	// Example 5: Type switch
	fmt.Println("5. Type Switch:")
	values := []interface{}{42, "hello", 3.14, true, []int{1, 2}}

	for _, v := range values {
		switch v.(type) {
		case int:
			fmt.Printf("%v is an int\n", v)
		case string:
			fmt.Printf("%v is a string\n", v)
		case float64:
			fmt.Printf("%v is a float64\n", v)
		case bool:
			fmt.Printf("%v is a bool\n", v)
		case []int:
			fmt.Printf("%v is a slice of ints\n", v)
		default:
			fmt.Printf("%v is unknown type\n", v)
		}
	}
	fmt.Println()

	// Example 6: Stringer interface
	fmt.Println("6. Stringer Interface:")
	dog := Dog{Name: "Buddy"}
	fmt.Println(dog)
}
