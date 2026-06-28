// Package main — executable program.
package main

// "fmt"  — formatted output
// "sync" — synchronization primitives: WaitGroup, Mutex, etc.
// "time" — time.Sleep, time.Duration, etc.
import (
	"fmt"
	"sync"
	"time"
)

// ======================================================================
// GOROUTINES — lightweight concurrent functions
// ======================================================================
// A goroutine is a function running CONCURRENTLY with other goroutines.
// To start one, just write:  go functionName(args)
// That's it. No threads, no thread pools, no async/await boilerplate.
//
// KEY DIFFERENCE vs OS threads:
//   OS thread  → ~1 MB of stack memory
//   Goroutine  → ~2 KB of stack (grows dynamically)
//   You can easily run 100,000 goroutines on a laptop.
//   The Go runtime schedules them onto OS threads automatically (M:N scheduling).

// worker simulates a task that takes some time.
// id identifies which worker is running (so we can see interleaving in output).
func worker(id int) {
	// This runs concurrently — multiple workers will print at roughly the same time.
	fmt.Printf("Worker %d: starting\n", id)

	// time.Sleep pauses this goroutine (not the whole program or other goroutines).
	// time.Millisecond is a constant — Go's time package uses typed durations.
	// 100 * time.Millisecond = 100ms.  No magic numbers like sleep(0.1).
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Worker %d: done\n", id)
}

// ======================================================================
// CHANNELS — typed pipes between goroutines
// ======================================================================
// A channel lets goroutines COMMUNICATE by passing values.
// Go's mantra: "Do not communicate by sharing memory;
//               share memory by communicating."
// Instead of using a shared variable (and locking it), goroutines
// send values through channels — much safer and easier to reason about.
//
// Create a channel:  ch := make(chan int)
// Send a value:      ch <- 42        (blocks until someone receives)
// Receive a value:   v := <-ch       (blocks until someone sends)

// square computes n*n and sends the result into the channel ch.
// chan int means "a channel that carries int values".
// This function is designed to run as a goroutine.
func square(n int, ch chan int) {
	// The <- operator sends a value into the channel.
	// If no one is receiving yet, this line BLOCKS until a receiver is ready.
	ch <- n * n
}

// ======================================================================
// WAITGROUP — the proper way to wait for goroutines
// ======================================================================
// time.Sleep(1 * time.Second) is a bad way to wait — you're guessing how long
// goroutines take. sync.WaitGroup is the correct tool:
//   wg.Add(1)    — tell the WaitGroup "one more goroutine is starting"
//   wg.Done()    — call when a goroutine finishes (usually via defer)
//   wg.Wait()    — block until all goroutines have called Done()

// fanOut launches n goroutines, each printing their ID, then waits for all to finish.
func fanOut(n int) {
	// WaitGroup tracks how many goroutines are still running.
	// Zero value is fine — no make() needed.
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		// Tell the WaitGroup we're launching one more goroutine.
		// IMPORTANT: call Add BEFORE starting the goroutine, not inside it.
		wg.Add(1)

		// "go func(id int) { ... }(i)" is an immediately-invoked goroutine.
		// We pass "i" as a parameter (named "id") to AVOID the closure trap:
		// if we used "i" directly inside the closure, all goroutines might
		// capture the same final value of i (a classic Go beginner bug).
		go func(id int) {
			// defer wg.Done() runs when this goroutine's function returns.
			// It's the goroutine equivalent of "finally" — guarantees Done() is called.
			defer wg.Done()

			fmt.Printf("Goroutine %d running\n", id)
			time.Sleep(50 * time.Millisecond) // simulate work
		}(i) // <-- we pass i here, so each goroutine gets its own copy
	}

	// Wait blocks here until every goroutine has called wg.Done().
	// Without this, main() might exit before the goroutines finish.
	wg.Wait()
	fmt.Println("All goroutines finished.")
}

func main() {
	// ======================================================================
	// DEMO 1 — Goroutines running concurrently (crude wait with Sleep)
	// ======================================================================
	fmt.Println("=== Goroutine demo ===")

	// "go worker(1)" starts worker in a NEW goroutine — it runs concurrently.
	// main() continues immediately without waiting for worker to finish.
	go worker(1)
	go worker(2)
	go worker(3)

	// Problem: if main() returns, ALL goroutines are killed immediately.
	// Here we sleep as a crude way to give goroutines time to run.
	// In real code, use channels or WaitGroups instead (see below).
	time.Sleep(300 * time.Millisecond)

	fmt.Println()

	// ======================================================================
	// DEMO 2 — Channels: send results back from goroutines
	// ======================================================================
	fmt.Println("=== Channel demo ===")

	// make(chan int) creates an UNBUFFERED channel of integers.
	// Unbuffered means: send blocks until someone receives, and vice versa.
	ch := make(chan int)

	// Start 3 goroutines, each computing the square of 1, 2, and 3.
	// Each goroutine will send its result into ch.
	go square(1, ch) // will send 1
	go square(2, ch) // will send 4
	go square(3, ch) // will send 9

	// Receive 3 values from the channel.
	// Each <-ch blocks until a value is available.
	// We don't know which goroutine finishes first — the order may vary.
	for i := 0; i < 3; i++ {
		result := <-ch // blocks until a goroutine sends a value
		fmt.Println("Received square:", result)
	}

	fmt.Println()

	// ======================================================================
	// DEMO 3 — WaitGroup: the proper way to wait
	// ======================================================================
	fmt.Println("=== WaitGroup demo ===")

	// fanOut launches 5 goroutines and waits for all of them properly.
	fanOut(5)

	fmt.Println()

	// ======================================================================
	// DEMO 4 — Buffered channel
	// ======================================================================
	fmt.Println("=== Buffered channel demo ===")

	// make(chan int, 3) creates a BUFFERED channel with capacity 3.
	// Sends don't block until the buffer is full — useful when you know
	// the number of results in advance (avoids needing a separate goroutine just to drain).
	buffered := make(chan int, 3)

	// These sends don't block because there's room in the buffer.
	buffered <- 10
	buffered <- 20
	buffered <- 30

	// close() signals that no more values will be sent on this channel.
	// Receivers can still drain remaining values after close().
	close(buffered)

	// "range" over a channel receives values until the channel is closed.
	for v := range buffered {
		fmt.Println("From buffered channel:", v)
	}

	// ======================================================================
	// EXERCISE FOR YOU
	// 1. Write a function: sumRange(start, end int, ch chan int)
	//    It should sum all integers from start to end (inclusive) and send the result to ch.
	// 2. In main, split the range 1..100 into two halves:
	//    - Goroutine 1: sum 1..50
	//    - Goroutine 2: sum 51..100
	// 3. Receive both results from the channel and print their sum.
	//    Expected answer: 5050
	// ======================================================================
}
