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

	fmt.Println()

	// ======================================================================
	// DEMO 5 — select: wait on multiple channels at once
	// ======================================================================
	fmt.Println("=== Select demo ===")

	// select is like a switch statement but for channel operations.
	// It waits until ONE of the cases is ready, then executes that case.
	// If multiple cases are ready at the same time, Go picks one at random.
	ch1 := make(chan string, 1) // buffered so sends don't block
	ch2 := make(chan string, 1)

	// Send a value into ch1 from a goroutine after a short delay.
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "result from ch1"
	}()

	// Send a value into ch2 from a goroutine after a longer delay.
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "result from ch2"
	}()

	// select blocks until ONE channel is ready, then handles it.
	// Because ch1 is faster, this will always print ch1's message first.
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			// This case fires when ch1 has a value ready.
			fmt.Println("Received:", msg)
		case msg := <-ch2:
			// This case fires when ch2 has a value ready.
			fmt.Println("Received:", msg)
		}
	}

	// ======================================================================
	// select with a timeout using time.After
	// ======================================================================
	// time.After(d) returns a channel that receives a value after duration d.
	// Pairing it with select gives you a clean timeout pattern.
	slow := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond) // deliberately slow
		slow <- "late value"
	}()

	select {
	case v := <-slow:
		fmt.Println("Got:", v)
	case <-time.After(100 * time.Millisecond):
		// This fires if no value arrives within 100ms.
		fmt.Println("Timed out waiting for slow channel")
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

// ======================================================================
// QUESTIONS — Goroutines & Channels
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  What is a goroutine? How is it different from an OS thread?
//      Why can you run 100,000 goroutines but not 100,000 OS threads?
//
//      A: A goroutine is a lightweight function running concurrently, managed
//         by the Go runtime (not the OS).
//         OS thread: ~1 MB fixed stack, scheduled by the OS kernel — expensive.
//         Goroutine:  ~2 KB initial stack (grows/shrinks as needed), scheduled
//         by Go's own M:N scheduler — many goroutines multiplexed onto few OS threads.
//         100k goroutines ≈ ~200 MB RAM. 100k OS threads ≈ ~100 GB RAM — not feasible.
//
// Q2.  What happens if main() returns while goroutines are still running?
//      How do you prevent this correctly?
//
//      A: When main() returns, the entire program exits immediately — all
//         running goroutines are killed without cleanup.
//         Correct ways to wait:
//           1. sync.WaitGroup — wg.Add/Done/Wait (idiomatic for fan-out work)
//           2. Receive on a channel — block until goroutine sends a signal/result
//         time.Sleep is NOT correct — you're guessing the duration.
//
// Q3.  What is the difference between an unbuffered and a buffered channel?
//      When does each one block on send? When does each block on receive?
//
//      A: Unbuffered (make(chan T)):
//           Send blocks until a receiver is ready.
//           Receive blocks until a sender is ready.
//           Both sides must be ready at the same time — a rendezvous point.
//         Buffered (make(chan T, N)):
//           Send blocks only when the buffer is FULL.
//           Receive blocks only when the buffer is EMPTY.
//           Up to N values can be queued without a receiver being ready.
//
// Q4.  What does closing a channel do?
//      What happens if you try to send to a closed channel?
//      What happens if you receive from a closed channel?
//
//      A: close(ch) signals that no more values will be sent.
//         Sending to a closed channel → panic (runtime error).
//         Receiving from a closed channel → returns immediately with the
//         zero value of the channel's type, and ok=false in the "v, ok := <-ch" form.
//         A "for range ch" loop exits automatically when the channel is closed and drained.
//
// Q5.  What is a "send-only" channel (chan<-) and a "receive-only" channel
//      (<-chan)? Why would you use directional channels in function signatures?
//
//      A: chan<- T  — can only send into this channel (write-only)
//         <-chan T  — can only receive from this channel (read-only)
//         Using directional types in function signatures documents intent clearly
//         and lets the compiler catch mistakes (e.g., accidentally receiving in
//         a producer function). The caller still creates a bidirectional chan T
//         and passes it — Go converts automatically.
//
// Q6.  What is a data race? How do goroutines cause them?
//      Name two ways to avoid data races in Go.
//
//      A: A data race occurs when two goroutines access the same memory location
//         concurrently and at least one of the accesses is a write, with no
//         synchronisation between them. The result is undefined behaviour.
//         Two ways to avoid:
//           1. Channels — pass ownership of data through the channel instead of sharing.
//           2. sync.Mutex — lock before reading/writing shared state, unlock after.
//         Run "go test -race ./..." or "go run -race main.go" to detect races.
//
// Q7.  What is the "closure trap" in goroutine loops, and how do you
//      avoid it? (Hint: relates to variable capture by closures.)
//
//      A: When a goroutine closure captures a loop variable directly, all goroutines
//         may see the same final value of the variable because the variable is shared:
//           for i := 0; i < 3; i++ {
//               go func() { fmt.Println(i) }()  // BUG: all may print 3
//           }
//         Fix: pass the variable as a function argument to create a per-iteration copy:
//           for i := 0; i < 3; i++ {
//               go func(id int) { fmt.Println(id) }(i)  // CORRECT
//           }
//         In Go 1.22+, loop variables are scoped per-iteration by default, fixing this.
//
// PRACTICAL
// ---------
// Q8.  What is the output of this code? Explain why.
//
//      ch := make(chan int)
//      ch <- 1
//      fmt.Println(<-ch)
//
//      A: This DEADLOCKS. "ch <- 1" on an unbuffered channel blocks forever
//         because there is no goroutine ready to receive.
//         The Go runtime detects the deadlock and panics:
//         "fatal error: all goroutines are asleep - deadlock!"
//
// Q9.  Rewrite the above to NOT deadlock, using a goroutine for the send.
//
//      A: ch := make(chan int)
//         go func() { ch <- 1 }()   // send in a goroutine — doesn't block main
//         fmt.Println(<-ch)          // receive in main — prints: 1
//
// Q10. Write a pipeline using two channels:
//      - Generator goroutine: sends integers 1..5 into channel A, then closes A.
//      - Squarer goroutine:   reads from A, squares each, sends into channel B.
//      - main():              reads from B and prints each result.
//      This is the classic Go "pipeline" pattern.
//
//      A: func generate(out chan<- int) {
//             for i := 1; i <= 5; i++ { out <- i }
//             close(out)
//         }
//         func squarer(in <-chan int, out chan<- int) {
//             for v := range in { out <- v * v }
//             close(out)
//         }
//         In main():
//           a := make(chan int)
//           b := make(chan int)
//           go generate(a)
//           go squarer(a, b)
//           for v := range b { fmt.Println(v) }  // 1 4 9 16 25
//
// Q11. Use a select statement to read from two channels (ch1 and ch2)
//      and print which channel delivered a value. What is "select" used for?
//      Hint: select { case v := <-ch1: ... case v := <-ch2: ... }
//
//      A: select lets a goroutine wait on MULTIPLE channel operations simultaneously,
//         proceeding with whichever one is ready first (random if multiple are ready).
//         ch1 := make(chan string, 1)
//         ch2 := make(chan string, 1)
//         ch1 <- "one"
//         select {
//         case v := <-ch1:
//             fmt.Println("from ch1:", v)
//         case v := <-ch2:
//             fmt.Println("from ch2:", v)
//         }
//         Use cases: timeouts (with time.After), cancellation, fan-in from multiple sources.
//
// Q12. What is sync.Mutex and when would you use it INSTEAD of a channel?
//      Write a thread-safe counter using sync.Mutex with Inc() and Value() methods.
//
//      A: sync.Mutex protects shared state with a mutual exclusion lock.
//         Use Mutex when multiple goroutines need to READ and WRITE the same
//         variable and passing ownership via a channel would be awkward.
//         Use channels when goroutines communicate results or pipeline data.
//
//         type Counter struct {
//             mu    sync.Mutex
//             count int
//         }
//         func (c *Counter) Inc() {
//             c.mu.Lock()
//             defer c.mu.Unlock()
//             c.count++
//         }
//         func (c *Counter) Value() int {
//             c.mu.Lock()
//             defer c.mu.Unlock()
//             return c.count
//         }
// ======================================================================
