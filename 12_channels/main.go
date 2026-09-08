package main

import (
	"fmt"
	"time"
)

// Example 1: Simple channel communication
func SendNumbers(ch chan int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Sending: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

// Example 2: Buffered channel
func Producer(ch chan string, id int) {
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("Message %d from Producer %d", i, id)
		fmt.Printf("Producer %d sending: %s\n", id, msg)
		ch <- msg
		time.Sleep(100 * time.Millisecond)
	}
}

// Example 3: Worker with channel
func Worker(id int, jobs chan int, results chan string) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
	}
}

// Example 4: Select with multiple channels
func Task(id int, ch1 chan string, ch2 chan string, result chan string) {
	select {
	case msg := <-ch1:
		result <- fmt.Sprintf("Task %d received from ch1: %s", id, msg)
	case msg := <-ch2:
		result <- fmt.Sprintf("Task %d received from ch2: %s", id, msg)
	case <-time.After(2 * time.Second):
		result <- fmt.Sprintf("Task %d timeout", id)
	}
}

// Example 5: Range over channel
func RangeOverChannel(values []int, ch chan int) {
	for _, v := range values {
		ch <- v
	}
	close(ch)
}

func main() {
	fmt.Println("=== CHANNELS ===\n")

	// Example 1: Unbuffered channel
	fmt.Println("1. Unbuffered Channel:")
	ch := make(chan int)
	go SendNumbers(ch)

	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println()

	// Example 2: Buffered channel
	fmt.Println("2. Buffered Channel:")
	buffered := make(chan string, 2)

	go Producer(buffered, 1)
	go Producer(buffered, 2)

	time.Sleep(2 * time.Second)

	for msg := range buffered {
		fmt.Printf("Received: %s\n", msg)
	}
	fmt.Println()

	// Example 3: Worker pool pattern
	fmt.Println("3. Worker Pool Pattern:")
	jobs := make(chan int, 5)
	results := make(chan string, 5)

	// Start workers
	for w := 1; w <= 2; w++ {
		go Worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
	fmt.Println()

	// Example 4: Select statement
	fmt.Println("4. Select Statement:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	result := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		go Task(i+1, ch1, ch2, result)
	}

	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 2; i++ {
		fmt.Println(<-result)
	}
	fmt.Println()

	// Example 5: Range over channel
	fmt.Println("5. Range Over Channel:")
	numCh := make(chan int)

	go RangeOverChannel([]int{10, 20, 30, 40, 50}, numCh)

	for num := range numCh {
		fmt.Printf("Got: %d\n", num)
	}
	fmt.Println()

	// Example 6: Bidirectional and unidirectional channels
	fmt.Println("6. Channel Directions:")
	bidirCh := make(chan int)

	go func(ch chan int) {
		ch <- 100
	}(bidirCh)

	value := <-bidirCh
	fmt.Printf("Received from bidirectional channel: %d\n", value)
}
