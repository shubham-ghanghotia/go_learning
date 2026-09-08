# Lesson 10: Packages & Imports

## What are Packages?

A **package** is a way to organize and reuse code. Every Go file belongs to a package.

```go
package main  // This file is part of the main package
```

## Package Structure

```
myproject/
├── main.go (package main)
├── math/
│   └── calculator.go (package math)
└── utils/
    └── helpers.go (package utils)
```

## Creating Packages

### Package Directory
```bash
mkdir math
```

### Package File (math/calculator.go)
```go
package math  // Same name as directory

// Exported function (starts with uppercase)
func Add(a, b int) int {
    return a + b
}

// Unexported function (starts with lowercase)
func privateHelper() {
    // Only accessible within this package
}
```

## Importing Packages

### Single Import
```go
import "fmt"
```

### Multiple Imports
```go
import (
    "fmt"
    "math"
    "os"
)
```

### Importing Custom Packages
```go
import (
    "fmt"
    "myproject/math"  // Full import path
)

func main() {
    result := math.Add(5, 3)  // Must use capitalized name (exported)
    fmt.Println(result)
}
```

## Exported vs Unexported

- **Exported**: Starts with **UPPERCASE** letter → Accessible outside package
- **Unexported**: Starts with **lowercase** letter → Only within package

```go
// calculator.go
func Add(a, b int) int { }       // Exported ✅
func subtract(a, b int) int { }  // Unexported ❌
```

## Import Aliases

```go
import (
    "fmt"
    m "myproject/math"      // Alias
    _ "database/sql/driver" // Blank import (side effects only)
)

m.Add(5, 3)  // Use with alias
```

## The `init()` Function

Runs automatically when package is imported:

```go
func init() {
    fmt.Println("Package initialized!")
}
```

## Public vs Private Identifiers

```go
type Calculator struct {
    lastResult int  // Private (lowercase)
    Name       string  // Public (uppercase)
}

func (c Calculator) Add(a, b int) int { }      // Public
func (c Calculator) internalCalc(a, b int) {} // Private
```

## Built-in Packages

### fmt - Formatted I/O
```go
import "fmt"
fmt.Println("Hello")
fmt.Printf("%d\n", 42)
```

### math - Mathematical functions
```go
import "math"
math.Sqrt(16)     // 4
math.Max(5, 3)    // 5
```

### strings - String operations
```go
import "strings"
strings.ToUpper("hello")      // HELLO
strings.Contains("hello", "ll") // true
```

### strconv - String conversion
```go
import "strconv"
num, _ := strconv.Atoi("42")  // Convert string to int
str := strconv.Itoa(42)        // Convert int to string
```

### os - Operating system functions
```go
import "os"
os.Getenv("PATH")
os.Exit(0)
```

### time - Time functions
```go
import "time"
time.Now()
time.Sleep(1 * time.Second)
```

## Package Documentation

```go
// Package calculator provides basic math operations
package calculator

// Add returns the sum of two numbers
func Add(a, b int) int {
    return a + b
}
```

## Running the Example

```bash
cd 10_packages_imports
go run main.go
```

## 🎯 Exercises

### Exercise 1: Create a math package
Create functions for Add, Subtract, Multiply, Divide.

### Exercise 2: Create a string utils package
Create functions for Reverse, IsPalindrome, CountVowels.

### Exercise 3: Use multiple packages
Write a program that imports and uses both packages.

## Common Built-in Packages

| Package | Purpose |
|---------|---------|
| fmt | Formatted I/O |
| math | Math functions |
| strings | String manipulation |
| strconv | String conversion |
| os | OS interaction |
| time | Date/time |
| io | Input/output |
| json | JSON parsing |
| net | Networking |
| crypto | Cryptography |

## Next Steps

→ [11 - Goroutines](../11_goroutines/)

## Resources

- [Go Packages Documentation](https://golang.org/ref/spec#Packages)
- [Standard Library Packages](https://pkg.go.dev/std)
- [How to Write Go Code](https://golang.org/doc/code)
