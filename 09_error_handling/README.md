# Lesson 9: Error Handling

## What are Errors?

In Go, errors are values. They implement the `error` interface:

```go
type error interface {
    Error() string
}
```

## The `error` Type

Every function that might fail should return an error as the last return value:

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## Handling Errors

### The Standard Pattern

```go
result, err := Divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

## Creating Custom Errors

### Using `errors.New()`

```go
import "errors"

var errUserNotFound = errors.New("user not found")

func GetUser(id int) (User, error) {
    if id < 1 {
        return User{}, errUserNotFound
    }
    return User{ID: id}, nil
}
```

### Using `fmt.Errorf()`

```go
import "fmt"

func OpenFile(filename string) (File, error) {
    if filename == "" {
        return File{}, fmt.Errorf("invalid filename: %s", filename)
    }
    return File{Name: filename}, nil
}
```

### Custom Error Types

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}

func ValidateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return ValidationError{
            Field:   "email",
            Message: "must contain @",
        }
    }
    return nil
}
```

## Error Wrapping

Use `fmt.Errorf()` with `%w` to wrap errors:

```go
file, err := os.Open("data.txt")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}

// Later: Check the wrapped error
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File not found")
}
```

## Panic and Recover

### Panic - Stop execution

```go
func BadFunction() {
    panic("Something went wrong!")
    fmt.Println("This never runs")
}

BadFunction()  // Program crashes here
```

### Recover - Catch panic

```go
func SafeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    
    panic("Error!")
    fmt.Println("This never runs")
}

SafeFunction()  // Program continues
```

## Best Practices

1. **Always check errors**
   ```go
   if err != nil {
       // Handle error
   }
   ```

2. **Don't ignore errors**
   ```go
   _ = file.Close()  // If you must ignore, be explicit
   ```

3. **Return errors early**
   ```go
   if err != nil {
       return err
   }
   ```

4. **Use custom error types for context**
   ```go
   type ConfigError struct {
       Key string
   }
   ```

## Running the Example

```bash
cd 09_error_handling
go run main.go
```

## 🎯 Exercises

### Exercise 1: Division function
Create a Divide function that returns error for division by zero.

### Exercise 2: Email validation
Create a ValidateEmail function with custom error.

### Exercise 3: File operations
Read a file and properly handle errors.

## Next Steps

→ [10 - Packages & Imports](../10_packages_imports/)

## Resources

- [Error Handling in Go](https://golang.org/doc/effective_go#errors)
- [errors Package](https://pkg.go.dev/errors)
- [fmt.Errorf](https://pkg.go.dev/fmt#Errorf)
