# Lesson 12: Channels

## What are Channels?

A **channel** is a communication mechanism between goroutines. Data is sent and received through channels.

```go
ch := make(chan int)  // Create channel
ch <- 42              // Send
value := <-ch         // Receive
```

## Creating Channels

### Unbuffered Channel
```go
ch := make(chan int)
```
Sender blocks until receiver is ready.

### Buffered Channel
```go
ch := make(chan int, 10)
```
Sender can send up to buffer size without blocking.

## Send and Receive

```go
func Send(ch chan int) {
    ch <- 42  // Send value
}

func Receive(ch chan int) {
    value := <-ch  // Receive value
    fmt.Println(value)
}

func main() {
    ch := make(chan int)
    go Send(ch)
    go Receive(ch)
}
```

## Closing Channels

```go
close(ch)  // Close the channel

// Receiving from closed channel
for value := range ch {
    fmt.Println(value)
}
```

## Select Statement

Choose between multiple channel operations:

```go
select {
case x := <-ch1:
    fmt.Println("Received from ch1:", x)
case y := <-ch2:
    fmt.Println("Received from ch2:", y)
case ch3 <- 42:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No communication ready")
}
```

## Bidirectional Channels

```go
func Bidirectional(ch chan int) {
    ch <- 42     // Send
    value := <-ch // Receive
}
```

## Unidirectional Channels

```go
func SendOnly(ch chan<- int) {
    ch <- 42  // Can only send
}

func ReceiveOnly(ch <-chan int) {
    value := <-ch  // Can only receive
}
```

## Channel Patterns

### Fan-out
```go
func GenerateNumbers(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)
}
```

### Fan-in
```go
func Merge(ch1, ch2 chan int, out chan int) {
    for {
        select {
        case v := <-ch1:
            out <- v
        case v := <-ch2:
            out <- v
        }
    }
}
```

## Running the Example

```bash
cd 12_channels
go run main.go
```

## 🎯 Exercises

### Exercise 1: Producer-Consumer
Create producer and consumer goroutines with channel.

### Exercise 2: Worker Pool
Implement multiple workers reading from a task channel.

### Exercise 3: Timeout
Use select to implement channel receive with timeout.

## Next Steps

→ [13 - File I/O](../13_file_io/)

## Resources

- [Channels Documentation](https://golang.org/doc/effective_go#channels)
- [Select Statement](https://golang.org/ref/spec#Select_statements)
