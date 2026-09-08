# Lesson 6: Maps

Maps are unordered collections of key-value pairs. Think of them like dictionaries in Python or objects in JavaScript.

## What is a Map?

A map is a data structure that maps keys to values:

```
Key → Value
"name" → "Alice"
"age" → 30
"email" → "alice@example.com"
```

## Declaring Maps

### Method 1: Using `var`

```go
var person map[string]string
fmt.Println(person)  // nil (uninitialized)
```

### Method 2: Using `make` (Preferred)

```go
countries := make(map[string]string)
fmt.Println(countries)  // map[]
```

### Method 3: Initialize with Values

```go
person := map[string]string{
    "name":  "Alice",
    "email": "alice@example.com",
    "city":  "New York",
}

scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}
```

## Map Syntax

```go
map[keyType]valueType
```

**Common combinations:**
- `map[string]string` - String keys, string values
- `map[string]int` - String keys, integer values
- `map[int]string` - Integer keys, string values
- `map[string][]int` - String keys, slice of integers
- `map[string]interface{}` - String keys, any type of value

## Accessing Values

```go
person := map[string]string{
    "name": "Alice",
    "email": "alice@example.com",
}

fmt.Println(person["name"])   // Alice
fmt.Println(person["email"])  // alice@example.com
```

### Handling Non-Existent Keys

Accessing a non-existent key returns the zero value:

```go
fmt.Println(person["age"])  // "" (empty string)
```

### The "ok" Idiom (Check if Key Exists)

```go
value, exists := person["name"]
if exists {
    fmt.Printf("Name: %s\n", value)
}

// Or combine:
if value, ok := person["name"]; ok {
    fmt.Printf("Name: %s\n", value)
} else {
    fmt.Println("Name not found")
}
```

## Adding and Modifying

```go
person := make(map[string]string)

// Add new key-value pair
person["name"] = "Alice"
person["email"] = "alice@example.com"

// Modify existing
person["name"] = "Bob"
```

## Deleting

```go
person := map[string]string{
    "name": "Alice",
    "email": "alice@example.com",
}

delete(person, "email")
fmt.Println(person)  // map[name:Alice]

// Deleting non-existent key is safe
delete(person, "age")  // No error
```

## Length

```go
scores := map[string]int{
    "Alice": 95,
    "Bob": 87,
}

fmt.Println(len(scores))  // 2
```

## Iterating Over Maps

### All key-value pairs

```go
person := map[string]string{
    "name": "Alice",
    "email": "alice@example.com",
}

for key, value := range person {
    fmt.Printf("%s: %s\n", key, value)
}

// Output (order is random):
// name: Alice
// email: alice@example.com
```

### Only keys

```go
for key := range person {
    fmt.Println(key)
}
```

### Only values

```go
for _, value := range person {
    fmt.Println(value)
}
```

**Note:** Map iteration order is random! Each iteration may have a different order.

## Map of Slices

```go
classes := map[string][]string{
    "Math": {"Alice", "Bob", "Carol"},
    "Science": {"David", "Eve"},
}

fmt.Println(classes["Math"])      // [Alice Bob Carol]
fmt.Println(len(classes["Math"])) // 3
```

## Nested Maps

```go
contacts := map[string]map[string]string{
    "Alice": {
        "phone": "555-1234",
        "email": "alice@example.com",
    },
    "Bob": {
        "phone": "555-5678",
        "email": "bob@example.com",
    },
}

fmt.Println(contacts["Alice"]["phone"])  // 555-1234
```

## Maps with interface{}

Store different types of values:

```go
data := map[string]interface{}{
    "name": "Alice",      // String
    "age": 30,            // Integer
    "score": 95.5,        // Float
    "active": true,       // Boolean
}

// Access and type assert
name := data["name"].(string)
age := data["age"].(int)
```

## Common Patterns

### Counting Occurrences

```go
func countWords(words []string) map[string]int {
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++
    }
    return counts
}

result := countWords([]string{"apple", "banana", "apple", "cherry", "apple"})
fmt.Println(result)  // map[apple:3 banana:1 cherry:1]
```

### Grouping

```go
type Student struct {
    Name string
    Grade string
}

students := []Student{...}

// Group by grade
gradeGroups := make(map[string][]string)
for _, student := range students {
    gradeGroups[student.Grade] = append(gradeGroups[student.Grade], student.Name)
}
```

### Lookup/Cache

```go
cache := make(map[int]string)

// Store expensive computation
cache[1] = "result1"

// Check if exists
if value, ok := cache[1]; ok {
    fmt.Println("Found in cache:", value)
}
```

## Running the Example

```bash
cd 06_maps
go run main.go
```

## 🎯 Exercises

### Exercise 1: Student Grades
Create a map that stores student names and their grades.
Write a function that returns the average grade.

### Exercise 2: Word Frequency
Write a function that takes a slice of words and returns a map showing how many times each word appears.

### Exercise 3: Reverse Lookup
Given a map `map[string]int`, write a function that creates a reverse map `map[int][]string` (value → keys).

### Exercise 4: Phone Book
Create a nested map to store contacts (name → phone/email).
Write functions to add, find, and delete contacts.

## Common Mistakes

1. **Using nil map:**
   ```go
   var person map[string]string  // nil
   person["name"] = "Alice"      // PANIC: assignment to entry in nil map
   
   // Correct:
   person := make(map[string]string)
   person["name"] = "Alice"
   ```

2. **Assuming order:**
   ```go
   // Maps are unordered!
   scores := map[string]int{"a": 1, "b": 2, "c": 3}
   for k := range scores {
       fmt.Println(k)  // Order is random each time
   }
   ```

3. **Forgetting ok check:**
   ```go
   value := data["key"]  // If key doesn't exist, value is zero value
   
   // Better:
   if value, ok := data["key"]; ok {
       // Use value
   }
   ```

4. **Type assertion errors:**
   ```go
   data := map[string]interface{}{"age": 30}
   age := data["age"].(string)  // PANIC! age is int, not string
   
   // Correct:
   if age, ok := data["age"].(int); ok {
       fmt.Println(age)
   }
   ```

## Next Steps

➜ [07 - Structs & Methods](../07_structs_methods/)

## Resources

- [Go Maps](https://golang.org/ref/spec#Map_types)
- [Effective Go - Maps](https://golang.org/doc/effective_go#maps)
- [Go by Example - Maps](https://gobyexample.com/maps)
