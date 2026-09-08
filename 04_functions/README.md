# Lesson 4: Functions

Functions are reusable blocks of code that perform specific tasks. They make your code more organized, readable, and maintainable.

## What is a Function?

A function is a block of code designed to perform a specific task.

```
┌─────────────────────┐
│   func add(a, b)    │  ← Function name and parameters
│   return a + b      │  ← Function body
│                     │
└─────────────────────┘
```

## Basic Function Syntax

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // Function body
    return value
}
```

## Types of Functions

### 1. No Parameters, No Return Value

```go
func greet() {
    fmt.Println("Hello!")
}

// Call it:
greet()  // Output: Hello!
```

### 2. With Parameters

```go
func greetPerson(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

greetPerson("Alice")  // Output: Hello, Alice!
```

**Multiple parameters:**

```go
func add(a int, b int) int {
    return a + b
}

// Shorthand (same type):
func add(a, b int) int {
    return a + b
}

result := add(5, 3)  // 8
```

### 3. With Return Value

```go
func add(a, b int) int {
    return a + b
}

result := add(5, 3)
fmt.Println(result)  // Output: 8
```

### 4. Multiple Return Values

Go allows functions to return multiple values:

```go
func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Error: Division by zero"
    }
    return a / b, "Success"
}

quotient, status := divide(10, 2)
fmt.Println(quotient, status)  // 5 Success
```

**Ignoring return values with `_`:**

```go
quotient, _ := divide(10, 2)  // Ignore the status message
```

### 5. Named Return Values

```go
func rectangle(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return  // No need to specify what to return
}

a, p := rectangle(5, 3)
fmt.Println(a, p)  // 15 16
```

## Advanced Function Concepts

### Variadic Functions (Variable Number of Arguments)

Accept any number of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

fmt.Println(sum(1, 2, 3))        // 6
fmt.Println(sum(1, 2, 3, 4, 5))  // 15

// Pass a slice:
nums := []int{1, 2, 3}
fmt.Println(sum(nums...))  // 6
```

### Recursive Functions

A function that calls itself:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5))  // 120
```

### Higher-Order Functions

Functions that return other functions:

```go
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

doubler := makeMultiplier(2)
fmt.Println(doubler(5))  // 10

tripler := makeMultiplier(3)
fmt.Println(tripler(5))  // 15
```

### Anonymous Functions

Functions without a name:

```go
// Execute immediately:
func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}("Alice")

// Store in a variable:
square := func(x int) int {
    return x * x
}
fmt.Println(square(4))  // 16
```

### Defer

Delay the execution of a statement until the function returns:

```go
func main() {
    fmt.Println("Start")
    defer fmt.Println("This prints last")
    fmt.Println("Middle")
}

// Output:
// Start
// Middle
// This prints last
```

**Use cases:**
- Cleanup operations
- Closing files
- Releasing resources

```go
func readFile() {
    file := openFile("data.txt")
    defer file.Close()  // Executes last
    // Work with file
}
```

## Parameter Passing

### By Value (Default)

Changes don't affect the original:

```go
func increment(x int) {
    x++  // Only changes the copy
}

num := 5
increment(num)
fmt.Println(num)  // Still 5
```

### By Pointer

Changes affect the original:

```go
func increment(x *int) {
    *x++  // Changes the actual value
}

num := 5
increment(&num)
fmt.Println(num)  // Now 6
```

*(You'll learn more about pointers in later lessons)*

## Function Best Practices

✅ **Good:**
```go
// Clear, descriptive names
func calculateArea(radius float64) float64 {
    return 3.14159 * radius * radius
}

// Single responsibility
func validateEmail(email string) bool {
    return strings.Contains(email, "@")
}
```

❌ **Bad:**
```go
// Vague name
func calc(x float64) float64 {
    return 3.14159 * x * x
}

// Too many responsibilities
func doStuff(data string) string {
    // Validates, transforms, saves...
}
```

## Running the Example

```bash
cd 04_functions
go run main.go
```

## 🎯 Exercises

### Exercise 1: Temperature Converter
Write a function that converts Celsius to Fahrenheit:
- Takes Celsius as parameter
- Returns Fahrenheit
- Formula: F = (C × 9/5) + 32

### Exercise 2: String Reversal
Write a recursive function that reverses a string.

### Exercise 3: Find Maximum
Write a variadic function that returns the maximum number from any number of integers.

### Exercise 4: Grade Validator
Write a function that:
- Takes marks (0-100)
- Returns (grade string, isPass bool)

## Common Mistakes

1. **Forgetting return type:**
   ```go
   // Wrong:
   func add(a, b int) { return a + b }
   
   // Right:
   func add(a, b int) int { return a + b }
   ```

2. **Type mismatch in return:**
   ```go
   // Wrong:
   func getValue() int {
       return "hello"  // String, not int
   }
   ```

3. **Not using all return values:**
   ```go
   // Warning (but works):
   divide(10, 2)  // Returns values but they're ignored
   
   // Better:
   quotient, _ := divide(10, 2)  // Explicitly ignore
   ```

## Next Steps

➜ [05 - Arrays & Slices](../05_arrays_slices/)

## Resources

- [Go Functions](https://golang.org/ref/spec#Function_declarations)
- [Effective Go - Functions](https://golang.org/doc/effective_go#functions)
- [Go by Example - Functions](https://gobyexample.com/functions)
