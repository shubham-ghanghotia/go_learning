# Lesson 8: Interfaces

## What are Interfaces?

An **interface** is a collection of method signatures. It defines a contract that types must fulfill.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

## Defining Interfaces

```go
type Writer interface {
    Write(data string) error
}

type Speaker interface {
    Speak() string
    Listen() string
}
```

## Implementing Interfaces

A type implements an interface by defining all its methods:

```go
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return d.Name + " says: Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return c.Name + " says: Meow!"
}

// Both Dog and Cat implement the Speaker interface
type Speaker interface {
    Speak() string
}
```

## Using Interfaces

```go
func MakeAnimalSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

dog := Dog{Name: "Buddy"}
cat := Cat{Name: "Whiskers"}

MakeAnimalSpeak(dog)   // Buddy says: Woof!
MakeAnimalSpeak(cat)   // Whiskers says: Meow!
```

## Empty Interface

The `interface{}` accepts any type:

```go
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

PrintAnything("Hello")      // Works
PrintAnything(42)           // Works
PrintAnything(3.14)         // Works
PrintAnything([]int{1,2,3}) // Works
```

## Type Assertion

Convert interface value to concrete type:

```go
var i interface{} = "hello"

s := i.(string)  // Assert i is a string
fmt.Println(s)   // hello

// Safe assertion
val, ok := i.(int)  // ok is false if not an int
if ok {
    fmt.Println(val)
} else {
    fmt.Println("Not an int")
}
```

## Type Switches

```go
func CheckType(v interface{}) {
    switch v.(type) {
    case string:
        fmt.Println("It's a string")
    case int:
        fmt.Println("It's an int")
    case []int:
        fmt.Println("It's a slice of ints")
    default:
        fmt.Println("Unknown type")
    }
}
```

## Common Built-in Interfaces

### Reader
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Writer
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### Stringer
```go
type Stringer interface {
    String() string
}
```

## Running the Example

```bash
cd 08_interfaces
go run main.go
```

## 🎯 Exercises

### Exercise 1: Shape interface
Create an interface with Area() and Perimeter() methods.
Implement it with Circle and Square types.

### Exercise 2: Sort interface
Implement sort.Interface on a custom struct.

### Exercise 3: Reader/Writer
Create a custom reader/writer implementation.

## Next Steps

→ [09 - Error Handling](../09_error_handling/)

## Resources

- [Go Interfaces](https://golang.org/ref/spec#Interface_types)
- [Effective Go - Interfaces](https://golang.org/doc/effective_go#interfaces_and_types)
