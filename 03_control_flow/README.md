# Lesson 3: Control Flow

Control flow determines which code executes based on conditions and repetition.

## IF Statements

### Basic If

```go
if age >= 18 {
    fmt.Println("You are an adult")
}
```

### If-Else

```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### If-Else If-Else

```go
if temperature > 25 {
    fmt.Println("Hot")
} else if temperature > 15 {
    fmt.Println("Warm")
} else {
    fmt.Println("Cold")
}
```

### If with Variable Initialization

Declare a variable just for the if statement:

```go
if score := 85; score >= 80 {
    fmt.Printf("Great score: %d\n", score)
}
// score is only available inside the if block
```

## For Loops

Go has only ONE loop construct: `for`. It works multiple ways.

### Traditional For Loop

```go
for i := 0; i < 3; i++ {
    fmt.Printf("Iteration %d\n", i)
}
// Output: 0, 1, 2
```

**Syntax:** `for init; condition; increment`

### While-like For Loop

Omit init and increment:

```go
count := 0
for count < 3 {
    fmt.Println(count)
    count++
}
// This acts like a while loop
```

### Infinite Loop

```go
for {
    fmt.Println("Forever")
    break  // Must break or it runs forever!
}
```

### Loop Control: Break and Continue

**Break** - Exit the loop:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Stop here
    }
    fmt.Println(i)  // Prints: 0, 1, 2, 3, 4
}
```

**Continue** - Skip to next iteration:

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue  // Skip printing 2
    }
    fmt.Println(i)  // Prints: 0, 1, 3, 4
}
```

### Range Loop

Iterate over collections:

```go
fruits := []string{"apple", "banana", "cherry"}

for index, fruit := range fruits {
    fmt.Printf("Index %d: %s\n", index, fruit)
}
// Output:
// Index 0: apple
// Index 1: banana
// Index 2: cherry
```

If you only want values (ignore index):

```go
for _, fruit := range fruits {
    fmt.Println(fruit)  // _ ignores the index
}
```

Or just get index:

```go
for i := range fruits {
    fmt.Println(i)  // Just the index
}
```

## Switch Statements

### Basic Switch

```go
day := 3

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")  // This executes
default:
    fmt.Println("Unknown")
}
```

**Note:** No `break` needed! Each case is separate (unlike C/Java).

### Multiple Cases

```go
grade := "B"

switch grade {
case "A":
    fmt.Println("Excellent")
case "B", "C":  // Both B and C execute same code
    fmt.Println("Good")
default:
    fmt.Println("Needs work")
}
```

### Switch with No Expression (Like If-Else)

```go
marks := 75

switch {
case marks >= 80:
    fmt.Println("A")
case marks >= 60:
    fmt.Println("B")  // This executes
case marks >= 40:
    fmt.Println("C")
default:
    fmt.Println("Fail")
}
```

## Logical Operators

### AND (&&)

Both conditions must be true:

```go
if age >= 18 && hasLicense {
    fmt.Println("Can drive")
}
```

### OR (||)

At least one condition must be true:

```go
if isWeekend || isHoliday {
    fmt.Println("No work today")
}
```

### NOT (!)

Reverses the condition:

```go
if !isRaining {
    fmt.Println("Go outside")
}
```

## Comparison Operators

| Operator | Meaning |
|----------|----------|
| `==` | Equal to |
| `!=` | Not equal to |
| `<` | Less than |
| `<=` | Less than or equal |
| `>` | Greater than |
| `>=` | Greater than or equal |

## Running the Example

```bash
cd 03_control_flow
go run main.go
```

## 🎯 Exercises

### Exercise 1: Number Checker
Write a program that:
- Takes a number
- Prints "Positive" if > 0
- Prints "Negative" if < 0
- Prints "Zero" if = 0

### Exercise 2: Loop Countdown
Write a program that counts down from 10 to 1 using a for loop.

### Exercise 3: Multiplication Table
Write a program that prints the 5 times table (5×1 to 5×10).

### Exercise 4: Grade Calculator
Write a program that takes marks (0-100) and prints:
- 90+ = A
- 80-89 = B
- 70-79 = C
- 60-69 = D
- Below 60 = F

## Common Mistakes

1. **Parentheses in conditions** (not needed):
   ```go
   // Wrong:
   if (age > 18) { }
   
   // Right:
   if age > 18 { }
   ```

2. **Semicolons** (Go doesn't use them):
   ```go
   // Wrong:
   for i := 0; i < 5; i++; { }
   
   // Right:
   for i := 0; i < 5; i++ { }
   ```

3. **Single-line if** (not allowed):
   ```go
   // Wrong:
   if age > 18 fmt.Println("Adult")
   
   // Right:
   if age > 18 {
       fmt.Println("Adult")
   }
   ```

## Next Steps

→ [04 - Functions](../04_functions/)

## Resources

- [Go Control Flow](https://golang.org/ref/spec#Statements)
- [Effective Go](https://golang.org/doc/effective_go)
