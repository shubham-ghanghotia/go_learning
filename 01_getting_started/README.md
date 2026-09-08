# Lesson 1: Getting Started with Go

## What is Go?

Go (also called Golang) is a modern programming language created by Google. It's designed to be:
- **Simple**: Easy syntax, quick to learn
- **Fast**: Compiles to machine code
- **Concurrent**: Built-in support for concurrent programming
- **Scalable**: Great for large systems and cloud applications

## Your First Program

### Understanding the Code

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}
```

### Breaking It Down

1. **`package main`** - Declares this file belongs to the "main" package
   - Every executable program needs a main package
   - There can only be ONE main() function per program

2. **`import ("fmt")`** - Imports the "fmt" package
   - fmt = formatted I/O (input/output)
   - Used for printing to console

3. **`func main()`** - The entry point of your program
   - Program execution starts here
   - Must be named exactly "main"

4. **`fmt.Println()`** - Prints text and adds a newline
   - Alternative: `fmt.Print()` (no newline)
   - Alternative: `fmt.Printf()` (formatted printing)

## Running Your First Program

### Method 1: Using `go run`
```bash
cd 01_getting_started
go run main.go
```
**Output:**
```
Hello, World!
Welcome to Go!
Hello, Gopher!
```

### Method 2: Compile then Run
```bash
# Compile to executable
go build -o hello main.go

# Run the executable
./hello
```

## Key Concepts

### Packages
- Every Go file belongs to a package
- The package name is declared at the top: `package main`
- Non-main packages are imported using `import`

### Functions
- Declared with `func` keyword
- Syntax: `func functionName(parameters) returnType { }`
- The main() function is special - it's where programs start

### Comments
- Single line: `// comment`
- Multi-line: `/* comment */`

## Common Printing Functions

```go
fmt.Print("Hello")           // No newline
fmt.Println("Hello")         // Adds newline
fmt.Printf("%s\n", "Hello") // Formatted printing
```

## 🎯 Exercises

### Exercise 1: Modify the output
Edit `main.go` to print your name instead of "Gopher".

### Exercise 2: Multiple prints
Add 3 more `fmt.Println()` calls to print different messages.

### Exercise 3: Formatted output
Use `fmt.Printf()` to print: "I am learning Go!"

## Next Steps

Once you're comfortable with this, move to:
→ [02 - Variables & Data Types](../02_variables_datatypes/)

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Package fmt](https://pkg.go.dev/fmt)
- [Go Code Style Guide](https://golang.org/doc/effective_go)
