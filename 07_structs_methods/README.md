# Lesson 7: Structs and Methods

## What are Structs?

A **struct** is a collection of fields grouped together. Think of it as a blueprint for creating objects with multiple properties.

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

## Defining Structs

### Basic Definition

```go
type Student struct {
    ID       int
    Name     string
    Email    string
    GPA      float64
    IsActive bool
}
```

### Struct Instantiation

```go
// Method 1: Using field names (preferred)
student1 := Student{
    ID:       1,
    Name:     "Alice",
    Email:    "alice@example.com",
    GPA:      3.9,
    IsActive: true,
}

// Method 2: Using positional values
student2 := Student{2, "Bob", "bob@example.com", 3.7, true}

// Method 3: Zero value (fields have default values)
var student3 Student
```

## Accessing and Modifying Fields

```go
// Access
fmt.Println(student1.Name)  // Alice

// Modify
student1.GPA = 3.95
student1.City = "New York"
```

## Methods

A **method** is a function associated with a struct. It's like a function that belongs to a type.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

### Using Methods

```go
rect := Rectangle{Width: 5, Height: 10}
fmt.Println("Area:", rect.Area())           // 50
fmt.Println("Perimeter:", rect.Perimeter()) // 30
```

## Receivers: Value vs Pointer

### Value Receiver (Copy)

```go
func (r Rectangle) Scale(factor float64) {
    r.Width *= factor   // Only modifies the copy
    r.Height *= factor
}

rect := Rectangle{5, 10}
rect.Scale(2)
fmt.Println(rect.Width)  // Still 5 (not modified!)
```

### Pointer Receiver (Reference)

```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor    // Modifies the actual struct
    r.Height *= factor
}

rect := Rectangle{5, 10}
rect.Scale(2)
fmt.Println(rect.Width)  // Now 10 (modified!)
```

**Rule of thumb:**
- Use **value receiver** for small structs when you don't need to modify
- Use **pointer receiver** when you need to modify the struct

## Embedded Structs (Composition)

```go
type Address struct {
    Street string
    City   string
    ZIP    string
}

type Person struct {
    Name    string
    Age     int
    Address Address  // Nested struct
}

// Usage
person := Person{
    Name: "Alice",
    Age:  30,
    Address: Address{
        Street: "123 Main St",
        City:   "New York",
        ZIP:    "10001",
    },
}

fmt.Println(person.Address.City)  // New York
```

## Field Tags (Metadata)

Tags provide metadata about struct fields:

```go
type User struct {
    ID       int    `json:"id" db:"user_id"`
    Name     string `json:"name" db:"username"`
    Email    string `json:"email" db:"user_email"`
    Password string `json:"-" db:"password_hash"`  // Ignored by JSON
}
```

Common tags:
- `json:"fieldName"` - JSON marshaling
- `db:"columnName"` - Database mapping
- `validate:"required"` - Validation

## Public vs Private Fields

```go
type BankAccount struct {
    AccountHolder string  // Public (starts with uppercase)
    balance       float64 // Private (starts with lowercase)
}

// Public field can be accessed outside
account.AccountHolder = "Alice"

// Private field cannot be accessed outside
// account.balance = 1000  // ERROR!
```

## Running the Example

```bash
cd 07_structs_methods
go run main.go
```

## 🎯 Exercises

### Exercise 1: Create a Car struct
Define a Car struct with fields: Make, Model, Year, Price
Create a method to calculate depreciation (10% per year).

### Exercise 2: Employee management
Create an Employee struct with Name, Salary, Department.
Add a method that gives a 10% raise.
Add a method that displays employee info.

### Exercise 3: Nested structs
Create a Company struct containing multiple Employee structs.
Display company information and employee details.

## Common Mistakes

1. **Forgetting pointer receiver for modifications**
   ```go
   func (p Person) SetAge(age int) {
       p.Age = age  // Doesn't work!
   }
   ```
   Should be: `func (p *Person) SetAge(age int)`

2. **Private fields (lowercase) inaccessible outside package**
   ```go
   struct User {
       id int  // Can't access from another package!
       ID int  // Use this instead
   }
   ```

3. **Not using pointers when needed**
   Always use pointers for methods that modify the receiver.

## Next Steps

→ [08 - Interfaces](../08_interfaces/)

## Resources

- [Go Structs Documentation](https://golang.org/ref/spec#Struct_types)
- [Go Methods Documentation](https://golang.org/ref/spec#Method_declarations)
- [Effective Go - Methods](https://golang.org/doc/effective_go#methods)
