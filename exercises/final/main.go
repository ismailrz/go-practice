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
//   url     — the URL to fetch
//   wg      — pointer to the WaitGroup so we can call Done() when finished
//   results — a send-only channel (chan<- Result) to return the outcome
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
