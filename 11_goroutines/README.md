# Lesson 11: Goroutines

## What are Goroutines?

A **goroutine** is a lightweight thread managed by the Go runtime. It allows concurrent execution of functions.

```go
go functionName()  // Launch goroutine
```

## Creating Goroutines

### Simple Goroutine

```go
func PrintNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}

func main() {
    go PrintNumbers()  // Runs concurrently
    time.Sleep(time.Second)  // Wait for goroutine
}
```

## Waiting for Goroutines

### Using WaitGroup

```go
var wg sync.WaitGroup

func Worker(id int) {
    defer wg.Done()  // Mark as done
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    wg.Add(3)  // Expect 3 goroutines
    
    go Worker(1)
    go Worker(2)
    go Worker(3)
    
    wg.Wait()  // Wait for all
}
```

## Goroutine Scheduling

- Goroutines are multiplexed onto OS threads
- Go scheduler decides which goroutine runs
- Thousands of goroutines can run efficiently

## Common Issues

1. **Race Conditions** - Multiple goroutines accessing same data
2. **Deadlocks** - Goroutines waiting for each other
3. **Resource Leaks** - Goroutines that never finish

## Running the Example

```bash
cd 11_goroutines
go run main.go
```

## Next Steps

→ [12 - Channels](../12_channels/)

## Resources

- [Goroutines Documentation](https://golang.org/doc/effective_go#goroutines)
- [sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup)
