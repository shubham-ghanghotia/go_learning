# Lesson 2: Variables and Data Types

## What are Variables?

Variables are containers that store data values. Think of them as labeled boxes that hold information.

```
Name: Alice  ← Variable name
Value: 25    ← The data stored
```

## Declaring Variables

Go offers three ways to declare variables:

### Method 1: `var` with explicit type

```go
var name string = "Alice"
var age int = 25
var height float64 = 5.8
```

### Method 2: `var` with type inference

Go figures out the type from the value:

```go
var city = "New York"    // Go knows this is a string
var temperature = 72     // Go knows this is an int
```

### Method 3: `:=` (Short declaration - PREFERRED in functions)

This is the most common way inside functions:

```go
name := "Alice"     // String
age := 25          // Integer
height := 5.8      // Float
```

**Note:** `:=` only works inside functions, not at package level.

## Go Data Types

### 1. Integers

```go
var count int = 42           // Standard integer
var small int8 = 127         // 8-bit (-128 to 127)
var large int64 = 9223...    // 64-bit (very large numbers)
var positive uint = 100      // Unsigned (only positive)
```

| Type | Range | Use Case |
|------|-------|----------|
| `int` | -2^31 to 2^31 | Default choice |
| `int8` | -128 to 127 | Memory-constrained |
| `int64` | -2^63 to 2^63 | Very large numbers |
| `uint` | 0 to 2^32 | Positive only |

### 2. Floating Point Numbers

```go
var pi float32 = 3.14       // 32-bit precision
var e float64 = 2.71828     // 64-bit (more precise)
```

| Type | Precision | Use Case |
|------|-----------|----------|
| `float32` | ~6 decimals | When memory matters |
| `float64` | ~15 decimals | Default, more accurate |

### 3. Strings

```go
var message string = "Hello"
var emoji string = "🎉"
var quote string = `Line 1
Line 2`  // Raw string (backticks)
```

- Enclosed in double quotes: `"Hello"`
- Support Unicode: `"你好"`, `"🚀"`
- Backticks create raw strings (preserve formatting)

### 4. Booleans

```go
var isTrue bool = true
var isFalse bool = false
```

Used in conditions:

```go
if isTrue {
    // This executes
}
```

## Constants

Values that cannot be changed once declared:

```go
const PI float64 = 3.14159
const AUTHOR = "Go Team"     // Type inferred
const MAX_USERS = 1000

// This would cause an error:
// PI = 3.14 // ERROR: cannot assign to PI
```

## Type Conversion

Converting from one type to another:

```go
var intValue int = 42
var floatValue float64 = float64(intValue)  // int to float
var stringValue string = fmt.Sprintf("%d", intValue) // int to string
```

**Syntax:** `NewType(value)`

## Zero Values

Every data type has a default "zero value" if not initialized:

```go
var defaultInt int           // 0
var defaultString string     // "" (empty)
var defaultBool bool         // false
var defaultFloat float64     // 0.0
```

## Variable Naming Conventions

✅ **Good names:**
```go
var userName string
var userAge int
var isActive bool
var MAX_RETRIES int = 3
```

❌ **Bad names:**
```go
var x string          // Too vague
var data int          // Not descriptive
var TempValue bool    // Avoid mixed case for vars
```

**Rules:**
- Start with letter or underscore
- Can contain letters, digits, underscores
- Case-sensitive: `age` ≠ `Age`
- Use camelCase for variables: `userName`, `userAge`
- Use UPPER_CASE for constants: `MAX_RETRIES`

## Running the Example

```bash
cd 02_variables_datatypes
go run main.go
```

## 🎯 Exercises

### Exercise 1: Create your profile
Declare variables for:
- Your name (string)
- Your age (int)
- Your GPA (float64)
- Are you a student? (bool)

Print them all.

### Exercise 2: Temperature converter
Declare a temperature in Celsius and convert it to Fahrenheit.
Formula: F = (C × 9/5) + 32

### Exercise 3: Circle calculator
Declare PI as a constant and radius as a variable.
Calculate the area: A = πr²

## Common Mistakes

1. **Redeclaring with `:=`**
   ```go
   name := "Alice"
   name := "Bob"  // ERROR: name already declared
   ```
   Use `=` to reassign:
   ```go
   name = "Bob"   // OK
   ```

2. **Using `:=` at package level**
   ```go
   package main
   age := 25  // ERROR: must use var
   ```

3. **Type mismatch**
   ```go
   var age int = "25"  // ERROR: string ≠ int
   ```

## Next Steps

→ [03 - Control Flow](../03_control_flow/)

## Resources

- [Go Types Documentation](https://golang.org/ref/spec#Types)
- [Effective Go - Constants](https://golang.org/doc/effective_go#constants)
