package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker used in goroutine demos
func printNumbers(label string, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("[%s] %d\n", label, i)
		time.Sleep(40 * time.Millisecond)
	}
}

// Sends n*n into the channel
func square(n int, ch chan int) {
	ch <- n * n
}

// Mutex-protected counter — safe for concurrent access
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	// Goroutine — "go fn()" starts fn concurrently, main continues immediately
	fmt.Println("--- goroutines ---")

	go printNumbers("A", 3)
	go printNumbers("B", 3)
	time.Sleep(300 * time.Millisecond) // crude wait — WaitGroup is better (below)

	// Unbuffered channel — sender blocks until receiver is ready
	fmt.Println("\n--- unbuffered channel ---")

	ch := make(chan int)
	go square(4, ch)
	go square(5, ch)
	go square(6, ch)

	for i := 0; i < 3; i++ {
		fmt.Println("received:", <-ch) // blocks until a goroutine sends
	}

	// Buffered channel — sender only blocks when buffer is full
	fmt.Println("\n--- buffered channel ---")

	buf := make(chan string, 3)
	buf <- "first"
	buf <- "second"
	buf <- "third"
	close(buf) // no more values will be sent

	for msg := range buf { // range exits when channel is closed and empty
		fmt.Println("got:", msg)
	}

	// WaitGroup — correct way to wait for multiple goroutines
	fmt.Println("\n--- WaitGroup ---")

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1) // register one goroutine before starting it
		go func(id int) {
			defer wg.Done() // decrement when done
			fmt.Printf("  task %d done\n", id)
		}(i) // pass i as argument to avoid closure trap
	}

	wg.Wait() // block until all goroutines call Done()
	fmt.Println("all tasks finished")

	// select — wait on multiple channels, handles whichever is ready first
	fmt.Println("\n--- select ---")

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() { time.Sleep(60 * time.Millisecond); ch1 <- "from ch1" }()
	go func() { time.Sleep(30 * time.Millisecond); ch2 <- "from ch2" }()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("received:", msg)
		case msg := <-ch2:
			fmt.Println("received:", msg)
		}
	}

	// select with timeout — time.After returns a channel that fires after d
	slow := make(chan int)
	go func() { time.Sleep(500 * time.Millisecond); slow <- 42 }()

	select {
	case v := <-slow:
		fmt.Println("got:", v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

	// Mutex — protects shared state from concurrent writes
	// Without a lock, 100 goroutines incrementing the same int = data race
	fmt.Println("\n--- Mutex ---")

	counter := &SafeCounter{}
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter.Inc()
		}()
	}

	wg2.Wait()
	fmt.Println("final count:", counter.Value()) // always 100

	// EXERCISE
	// 1. Write concurrentSum(nums []int) int that splits the slice in two,
	//    sums each half in its own goroutine via a channel, returns total.
	//    concurrentSum([]int{1,2,3,4,5,6,7,8,9,10}) → 55
	// 2. Launch 5 goroutines, each sleeping a random short time then sending
	//    its id on a channel. Print the ids in the order they arrive.
	// 3. Remove sync.Mutex from SafeCounter and run with:
	//    go run -race goroutine.go   — observe the data race warning.
}
