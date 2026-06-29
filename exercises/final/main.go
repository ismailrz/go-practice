// Package main — executable program.
// This is the FINAL exercise: a concurrent URL fetcher that ties together
// structs, error handling, goroutines, channels, and the standard library.
package main

import (
	// "bufio"   — buffered I/O: efficiently reads a file line by line
	"bufio"
	// "fmt"     — formatted output (Println, Printf, Fprintf, Errorf)
	"fmt"
	// "net/http" — Go's built-in HTTP client and server (no third-party needed!)
	"net/http"
	// "os"      — operating system interface: files, args, exit
	"os"
	// "sync"    — synchronization: WaitGroup to wait for all goroutines
	"sync"
)

// ======================================================================
// STRUCT — Result holds the outcome of fetching one URL
// ======================================================================
// Instead of returning multiple loose values, we group them in a struct.
// This is a common Go pattern: define a result struct for goroutine output.
type Result struct {
	URL    string // the URL that was fetched
	Status int    // HTTP status code (e.g., 200, 404, 500)
	Err    error  // non-nil if the request failed entirely (network error, etc.)
}

// ======================================================================
// GOROUTINE FUNCTION — fetch one URL and send the Result into a channel
// ======================================================================
// Parameters:
//
//	url     — the URL to fetch
//	wg      — pointer to the WaitGroup so we can call Done() when finished
//	results — a send-only channel (chan<- Result) to return the outcome
//
// "chan<- Result" is a SEND-ONLY channel — this function can only send, not receive.
// This makes the intent clear and prevents accidental reads inside fetch().
func fetch(url string, wg *sync.WaitGroup, results chan<- Result) {
	// defer wg.Done() ensures we always decrement the WaitGroup counter
	// when this goroutine exits — even if it panics or returns early.
	defer wg.Done()

	// http.Get() sends an HTTP GET request.
	// It returns (*http.Response, error).
	// This can fail due to network issues, bad URLs, DNS errors, etc.
	resp, err := http.Get(url)
	if err != nil {
		// Network-level failure (e.g., no internet, DNS failure, refused connection).
		// We send a Result with only the error — Status stays 0 (zero value).
		results <- Result{URL: url, Err: err}
		return // exit the goroutine early
	}

	// defer resp.Body.Close() ensures we always release the connection back to the pool.
	// Forgetting this is a common bug — it leaks connections.
	// This is the classic defer use case: acquire a resource, defer its cleanup.
	defer resp.Body.Close()

	// resp.StatusCode is the HTTP status code: 200 OK, 404 Not Found, 500 Internal Server Error, etc.
	// Send a successful Result back through the channel.
	results <- Result{URL: url, Status: resp.StatusCode}
}

func main() {
	// ======================================================================
	// STEP 1 — Read command-line arguments
	// ======================================================================
	// os.Args is a slice of strings: [programName, arg1, arg2, ...]
	// os.Args[0] is always the program binary name.
	// os.Args[1] should be the path to our URLs file.
	if len(os.Args) < 2 {
		// Stderr is better than stdout for error messages — it doesn't mix with program output.
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <urls-file>")
		// os.Exit(1) terminates the program with a non-zero exit code (signals failure to the shell).
		os.Exit(1)
	}

	// ======================================================================
	// STEP 2 — Open and read the URLs file
	// ======================================================================
	// os.Open opens a file for reading. Returns (*os.File, error).
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	// defer file.Close() — even if the rest of main() panics, the file will be closed.
	defer file.Close()

	// bufio.NewScanner wraps the file in a buffered reader.
	// Scanner reads one line at a time — memory-efficient for large files.
	scanner := bufio.NewScanner(file)

	// urls is a slice of strings — initially nil (zero value for slices).
	var urls []string

	// scanner.Scan() advances to the next line and returns true while there are lines.
	// scanner.Text() returns the current line as a string.
	for scanner.Scan() {
		line := scanner.Text()
		// Skip blank lines (empty URLs would cause http.Get to fail anyway).
		if line != "" {
			// append() adds an element to the end of a slice.
			// In Go, slices grow dynamically — similar to Python lists or Java ArrayLists.
			urls = append(urls, line)
		}
	}

	fmt.Printf("Fetching %d URLs concurrently...\n\n", len(urls))

	// ======================================================================
	// STEP 3 — Launch goroutines concurrently
	// ======================================================================

	// BUFFERED channel with capacity = number of URLs.
	// This prevents goroutines from blocking when they try to send results —
	// they can all send immediately without waiting for main to receive.
	// Rule of thumb: buffer capacity = number of senders if you know it in advance.
	results := make(chan Result, len(urls))

	// WaitGroup to wait for ALL goroutines to finish before we read results.
	var wg sync.WaitGroup

	for _, url := range urls {
		// Tell the WaitGroup one more goroutine is starting.
		// Must be called BEFORE "go fetch(...)" — not inside the goroutine.
		wg.Add(1)

		// Launch a goroutine for this URL.
		// All goroutines run concurrently — we don't wait for one to finish before starting the next.
		// For 4 URLs, all 4 HTTP requests happen in parallel!
		go fetch(url, &wg, results)
		// &wg passes a POINTER to the WaitGroup — if we passed wg by value, Done() would
		// modify a copy and the original counter would never reach zero.
	}

	// wg.Wait() blocks here until ALL goroutines have called wg.Done().
	// Only after all goroutines finish do we proceed to read results.
	wg.Wait()

	// close(results) signals that no more values will be sent.
	// This allows the "range results" loop below to terminate automatically.
	// If we didn't close, "for range results" would block forever waiting for more values.
	close(results)

	// ======================================================================
	// STEP 4 — Print all results
	// ======================================================================
	// "for r := range results" receives values from the channel until it's closed and empty.
	for r := range results {
		if r.Err != nil {
			// Network-level error — the request couldn't even be made.
			fmt.Printf("ERROR  %-45s  %v\n", r.URL, r.Err)
		} else {
			// Successful response — print the status code and URL.
			// %d = integer,  %-45s = left-aligned string padded to 45 chars
			fmt.Printf("%d    %-45s\n", r.Status, r.URL)
		}
	}

	// ======================================================================
	// HOW TO RUN THIS:
	//   go run main.go urls.txt
	//
	// EXERCISE FOR YOU:
	// 1. Add a "Duration" field to Result (time.Duration).
	// 2. In fetch(), record the time before and after http.Get using time.Now() and time.Since().
	// 3. Print the duration alongside the status code.
	// 4. At the end, print which URL was the slowest.
	// ======================================================================
}

// ======================================================================
// QUESTIONS — Bringing It All Together (Final Project)
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  This program uses a buffered channel sized to len(urls).
//      What would happen if we used an unbuffered channel instead,
//      and why? How would you fix the deadlock without using a buffer?
//
//      A: With an unbuffered channel, each fetch() goroutine would block on
//         "results <- Result{...}" until main() receives from the channel.
//         But main() is blocked on wg.Wait() — so no one is receiving.
//         All goroutines and main() are stuck waiting on each other: deadlock.
//         Fix without a buffer: drain the channel in a separate goroutine
//         before calling wg.Wait(), or use wg.Wait() in a goroutine and
//         close results from there.
//
// Q2.  Why do we pass &wg (pointer to WaitGroup) into fetch() instead
//      of passing wg by value? What would go wrong if we passed by value?
//
//      A: sync.WaitGroup contains internal state (a counter). If you pass it
//         by value, fetch() gets its own copy of the WaitGroup — calling
//         Done() on the copy decrements the copy's counter, not the original.
//         The original wg.Wait() in main() would block forever because its
//         counter never reaches zero. Always pass WaitGroup as a pointer.
//
// Q3.  Why is defer resp.Body.Close() important here?
//      What real-world problem does it prevent?
//
//      A: HTTP responses keep a TCP connection open until the body is closed.
//         If you don't close it, the connection is never returned to the pool,
//         eventually exhausting the connection limit and causing new requests
//         to hang or fail. In a loop fetching many URLs (like this program),
//         not closing bodies causes a connection leak that grows until the
//         OS refuses new connections.
//
// Q4.  The fetch() function parameter is "chan<- Result" (send-only).
//      What does this restriction give us? Could you write it as "chan Result"?
//      What is the benefit of being explicit?
//
//      A: Yes, chan Result (bidirectional) also compiles fine here.
//         But chan<- Result documents intent: fetch() only produces results,
//         it never reads from the channel. The compiler enforces this —
//         if someone accidentally adds "<-results" inside fetch(), it won't compile.
//         This prevents bugs and makes the code self-documenting.
//
// Q5.  We call wg.Wait() and then close(results) before the range loop.
//      What would happen if we called close(results) BEFORE wg.Wait()?
//
//      A: The goroutines would still be running and would try to send into
//         a closed channel → panic: "send on closed channel".
//         Always wait for all senders to finish before closing the channel.
//         The pattern is always: wg.Wait() → close(ch) → range ch.
//
// Q6.  What does os.Stderr mean and why do we write error messages there
//      instead of using fmt.Println (which writes to os.Stdout)?
//
//      A: Every process has two standard output streams:
//           stdout (os.Stdout) — normal program output, meant for consumers.
//           stderr (os.Stderr) — diagnostic/error messages, meant for humans.
//         Keeping them separate lets users pipe stdout to a file or another
//         program while still seeing errors in the terminal:
//           go run main.go urls.txt > results.txt   (errors still appear on screen)
//
// PRACTICAL
// ---------
// Q7.  Add a timeout to each HTTP request so that slow URLs don't hang
//      forever. Hint: create a custom http.Client with a Timeout field.
//      http.Client{Timeout: 3 * time.Second}
//      Replace http.Get(url) with client.Get(url).
//
//      A: client := &http.Client{Timeout: 3 * time.Second}
//         resp, err := client.Get(url)
//         After the timeout, http.Get returns an error with "context deadline exceeded".
//         The existing "if err != nil" block handles it automatically — no other changes needed.
//
// Q8.  Add rate limiting: instead of firing all goroutines at once,
//      only allow 2 concurrent requests at a time using a buffered channel
//      as a semaphore. Pattern:
//        sem := make(chan struct{}, 2)
//        sem <- struct{}{}    // acquire
//        defer func() { <-sem }()  // release
//
//      A: sem := make(chan struct{}, 2)  // capacity = max concurrency
//         Inside fetch(), at the top:
//           sem <- struct{}{}            // blocks if 2 goroutines already running
//           defer func() { <-sem }()     // release slot when done
//         Now at most 2 HTTP requests run simultaneously — useful to avoid
//         overwhelming a server or hitting rate limits.
//
// Q9.  Change the program to also write all results to a file called
//      "results.txt" in addition to printing to the terminal.
//      Use os.Create and fmt.Fprintf to write to the file.
//      Use defer to close the file.
//
//      A: outFile, err := os.Create("results.txt")
//         if err != nil { ... }
//         defer outFile.Close()
//         Then in the results loop, write to both:
//           fmt.Printf("%d  %s\n", r.Status, r.URL)
//           fmt.Fprintf(outFile, "%d  %s\n", r.Status, r.URL)
//         Or use io.MultiWriter to write to both with one call:
//           w := io.MultiWriter(os.Stdout, outFile)
//           fmt.Fprintf(w, "%d  %s\n", r.Status, r.URL)
//
// Q10. Add a summary at the end that prints:
//      - Total URLs fetched
//      - How many returned 2xx status codes
//      - How many returned 4xx or 5xx
//      - How many failed with a network error
//
//      A: Collect results into a []Result slice, then:
//         var ok2xx, err4xx5xx, netErr int
//         for _, r := range allResults {
//             switch {
//             case r.Err != nil:                          netErr++
//             case r.Status >= 200 && r.Status < 300:    ok2xx++
//             default:                                    err4xx5xx++
//             }
//         }
//         fmt.Printf("Total: %d  2xx: %d  4xx/5xx: %d  errors: %d\n",
//             len(allResults), ok2xx, err4xx5xx, netErr)
//
// Q11. Right now the order of results is non-deterministic (whichever
//      goroutine finishes first). How would you sort the results by
//      URL alphabetically before printing?
//      Hint: collect results into a []Result slice, then use sort.Slice().
//
//      A: Drain the channel into a slice first:
//         var all []Result
//         for r := range results { all = append(all, r) }
//         Then sort:
//         sort.Slice(all, func(i, j int) bool {
//             return all[i].URL < all[j].URL
//         })
//         Then print the sorted slice.
//         This gives deterministic output regardless of which goroutine finished first.
// ======================================================================
