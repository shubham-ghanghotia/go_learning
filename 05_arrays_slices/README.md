# Lesson 5: Arrays and Slices

One of Go's most important features! Arrays and slices are collections of elements.

## Arrays vs Slices

| Feature | Array | Slice |
|---------|-------|-------|
| Size | **Fixed** | **Dynamic** |
| Declaration | `[size]type` | `[]type` |
| Memory | Stack | Heap (pointer) |
| Performance | Faster | Flexible |
| Common Use | Known size | Most cases |

## Arrays

### Declaring Arrays

```go
// With specified size
var numbers [3]int
fmt.Println(numbers)  // [0 0 0]

// Initialize with values
var fruits [3]string = [3]string{"apple", "banana", "cherry"}

// Short declaration
colors := [4]string{"red", "green", "blue", "yellow"}

// Let compiler count elements
scores := [...]int{10, 20, 30, 40, 50}
fmt.Println(len(scores))  // 5
```

### Accessing Elements

```go
fruits := [3]string{"apple", "banana", "cherry"}

fmt.Println(fruits[0])              // apple
fmt.Println(fruits[len(fruits)-1])  // cherry

// Modify
fruits[1] = "mango"
fmt.Println(fruits)  // [apple mango cherry]
```

### Iterating Over Arrays

```go
// Traditional loop
for i := 0; i < len(fruits); i++ {
    fmt.Println(fruits[i])
}

// Range loop (preferred)
for index, fruit := range fruits {
    fmt.Printf("Index %d: %s\n", index, fruit)
}

// Ignore index
for _, fruit := range fruits {
    fmt.Println(fruit)
}

// Ignore value (just get index)
for i := range fruits {
    fmt.Println(i)
}
```

## Slices

Slices are flexible, dynamic arrays. They're the most common data structure in Go.

### Declaring Slices

```go
// Empty slice
var emptySlice []int
fmt.Println(len(emptySlice), cap(emptySlice))  // 0 0

// Initialize with values
var numbers []int = []int{1, 2, 3, 4, 5}

// Short declaration
vegetables := []string{"carrot", "broccoli", "spinach"}

// Using make (allocate memory)
mySlice := make([]int, 5)        // Length 5, capacity 5
mySlice2 := make([]int, 3, 10)   // Length 3, capacity 10
```

### Length vs Capacity

- **Length:** How many elements are currently in the slice
- **Capacity:** How many elements can fit before needing to grow

```go
mySlice := make([]int, 3, 10)
fmt.Println(len(mySlice))    // 3
fmt.Println(cap(mySlice))    // 10
```

### Append (Add Elements)

```go
fruits := []string{"apple", "banana"}

// Add one element
fruits = append(fruits, "orange")
fmt.Println(fruits)  // [apple banana orange]

// Add multiple elements
fruits = append(fruits, "grape", "kiwi")
fmt.Println(fruits)  // [apple banana orange grape kiwi]

// Append another slice
morefruits := []string{"mango", "pineapple"}
fruits = append(fruits, morefruits...) // ... unpacks the slice
```

**Important:** You must reassign the result of `append()`!

```go
fruits := []string{"apple"}
append(fruits, "banana")  // This doesn't change fruits!
fruits = append(fruits, "banana")  // Correct way
```

### Slicing (Extract Portion)

Create a new slice from existing array/slice:

```go
numbers := []int{10, 20, 30, 40, 50, 60}

// numbers[start:end] - includes start, excludes end
numbers[1:3]   // [20 30]
numbers[:3]    // [10 20 30]
numbers[3:]    // [40 50 60]
numbers[:]     // [10 20 30 40 50 60]

// Three-index slicing (sets capacity)
numbers[1:3:4]  // Length 2, capacity 3
```

### Copying Slices

```go
original := []int{1, 2, 3}

// ❌ This doesn't copy - just references the same data
copyByReference := original
original[0] = 999
fmt.Println(copyByReference)  // [999 2 3] - affected!

// ✅ This creates a true copy
copyByValue := make([]int, len(original))
copy(copyByValue, original)
original[0] = 999
fmt.Println(copyByValue)  // [1 2 3] - not affected
```

## Slice Functions

### len()
Returns the length:

```go
fruits := []string{"apple", "banana", "orange"}
fmt.Println(len(fruits))  // 3
```

### cap()
Returns the capacity:

```go
mySlice := make([]int, 3, 10)
fmt.Println(cap(mySlice))  // 10
```

### copy()
Copies elements from one slice to another:

```go
src := []int{1, 2, 3}
dst := make([]int, 3)
copy(dst, src)
```

### append()
Adds elements to a slice:

```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4, 5)
```

## Multidimensional Arrays/Slices

### 2D Array

```go
var matrix [2][3]int = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

fmt.Println(matrix[0][1])  // 2
```

### 2D Slice

```go
grid := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

fmt.Println(grid[2][2])  // 9
```

## Common Slice Patterns

### Filter

```go
func filterEven(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

fmt.Println(filterEven([]int{1, 2, 3, 4, 5, 6}))  // [2 4 6]
```

### Map (Transform)

```go
func double(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = num * 2
    }
    return result
}
```

### Sum

```go
func sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

## Running the Example

```bash
cd 05_arrays_slices
go run main.go
```

## 🎯 Exercises

### Exercise 1: Find Maximum
Write a function that finds the maximum value in a slice of integers.

### Exercise 2: Reverse Slice
Write a function that reverses a slice in-place.

### Exercise 3: Remove Duplicates
Write a function that removes duplicate elements from a slice.

### Exercise 4: Merge Slices
Write a function that takes two slices and returns a single merged slice.

## Common Mistakes

1. **Forgetting to reassign append():**
   ```go
   numbers := []int{1, 2, 3}
   append(numbers, 4)  // Wrong! Changes not saved
   numbers = append(numbers, 4)  // Correct
   ```

2. **Array vs Slice confusion:**
   ```go
   var arr [3]int     // Array - fixed size
   var slc []int      // Slice - dynamic
   ```

3. **Slice bounds error:**
   ```go
   s := []int{1, 2, 3}
   fmt.Println(s[5])  // ERROR: index out of range
   ```

4. **Modifying through reference:**
   ```go
   a := []int{1, 2, 3}
   b := a             // b points to same data
   b[0] = 999
   fmt.Println(a[0])  // 999 - a is affected!
   ```

## Next Steps

➜ [06 - Maps](../06_maps/)

## Resources

- [Go Arrays and Slices](https://golang.org/ref/spec#Array_types)
- [Effective Go - Slices](https://golang.org/doc/effective_go#slices)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
