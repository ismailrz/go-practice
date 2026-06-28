# Go Language — Practical Session Plan

**Audience:** Developers who know at least one other language (Python/Java/JS/etc.)  
**Duration:** 1–2 hours  
**Format:** Live coding + short exercises. Explain a concept, then immediately run it.

---

## Session Philosophy

Don't teach Go *from scratch* — teach Go *as different*.  
Experienced developers already understand loops, functions, and variables.  
Focus on **what Go does that surprises you**, and **why it was designed that way**.

---

## Topics & Time Allocation

| # | Topic | Time | Key Point |
|---|-------|------|-----------|
| 1 | Toolchain & Project Setup | 10 min | go mod, go run, go build |
| 2 | Types, Variables & Zero Values | 10 min | := vs var, no null surprises |
| 3 | Functions & Multiple Returns | 10 min | errors as values, not exceptions |
| 4 | Structs & Interfaces | 20 min | no inheritance, implicit interfaces |
| 5 | Goroutines & Channels | 25 min | concurrency as a first-class citizen |
| 6 | Wrap-up: Standard Library Tour | 5 min | net/http in ~10 lines |

---

## Topic 1 — Toolchain & Project Setup (10 min)

### What to show
```bash
mkdir hello && cd hello
go mod init hello
```

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

```bash
go run main.go     # compile + run
go build .         # produce binary
./hello            # run binary (no runtime needed)
```

### Key talking points
- **No `npm install` or `pip install` step.** Single binary output.
- `go fmt` is built-in. Formatting is not a debate in Go.
- `go mod` replaced `$GOPATH` — projects live anywhere now.

### Hands-on (2 min)
> "Initialize a module called `workshop`, print your name."

---

## Topic 2 — Types, Variables & Zero Values (10 min)

### What to show
```go
// Short declaration (inside functions)
x := 42
name := "Alice"

// Explicit (package level or when type matters)
var count int
var message string

// Zero values — no null/undefined surprises
fmt.Println(count)    // 0
fmt.Println(message)  // ""

// Multiple assignment
a, b := 1, 2
a, b = b, a  // swap — no temp variable
```

### Key talking point
In Python/JS, uninitialized variables are `None`/`undefined` — runtime surprises.  
In Go, every type has a **zero value**. You always know what you have.

| Type | Zero Value |
|------|-----------|
| int, float | 0 |
| string | "" |
| bool | false |
| pointer/slice/map | nil |

### Hands-on (2 min)
> "Declare a map of string→int without initializing it. Try to write to it. See what happens. Then fix it with `make`."

---

## Topic 3 — Functions & Multiple Returns (10 min)

### What to show

**Multiple return values (the Go way to handle errors):**
```go
import (
    "errors"
    "fmt"
    "strconv"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
}
```

**defer — runs when the function exits:**
```go
func readFile() {
    f, _ := os.Open("file.txt")
    defer f.Close()  // guaranteed cleanup, even on early return
    // ... work with f
}
```

### Key talking points
- Go has **no exceptions** (no try/catch). Errors are return values.
- This forces you to *decide* what to do with an error at every call site.
- `defer` is how you do `finally` — but cleaner.

### Hands-on (3 min)
> "Write a function `parseInt(s string) (int, error)` that wraps `strconv.Atoi`. Call it with `"42"` and `"abc"` and handle both cases."

---

## Topic 4 — Structs & Interfaces (20 min)

### What to show

**Structs with methods:**
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver (can mutate)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println(rect.Area())  // 50
    rect.Scale(2)
    fmt.Println(rect.Area())  // 200
}
```

**Interfaces — implicit, not explicit:**
```go
// No "implements" keyword needed
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Rectangle and Circle BOTH satisfy Shape — automatically
func printArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
    printArea(Rectangle{Width: 4, Height: 3})
    printArea(Circle{Radius: 5})
}
```

### Key talking points
- **No inheritance.** Composition over inheritance is enforced by the language.
- Interfaces are satisfied *implicitly* — a type doesn't need to declare which interface it implements.
- This means you can define an interface for code you didn't write (e.g., stdlib types).

### Hands-on (5 min)
> "Add a `Triangle` struct with a `Base` and `Height`. Give it an `Area()` method. Pass it to `printArea` — it works without changing `printArea`."

---

## Topic 5 — Goroutines & Channels (25 min)

This is Go's superpower. Take your time here.

### Part A — Goroutines (10 min)

```go
import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Println("Hello,", name)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go sayHello("Alice")   // runs concurrently
    go sayHello("Bob")     // runs concurrently
    time.Sleep(1 * time.Second)  // wait (crude — channels fix this)
}
```

**Key talking point:** A goroutine costs ~2KB of stack (vs ~1MB for an OS thread).  
You can run **100,000 goroutines** on a laptop.

### Part B — Channels (15 min)

```go
import "fmt"

func sum(nums []int, ch chan int) {
    total := 0
    for _, n := range nums {
        total += n
    }
    ch <- total  // send result into channel
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

    ch := make(chan int)

    go sum(nums[:4], ch)  // goroutine 1: sum first half
    go sum(nums[4:], ch)  // goroutine 2: sum second half

    a, b := <-ch, <-ch    // receive both results (blocks until ready)
    fmt.Println(a + b)    // 36
}
```

**WaitGroup — proper way to wait for goroutines:**
```go
import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Println("Worker", id, "done")
        }(i)
    }

    wg.Wait()
    fmt.Println("All workers done")
}
```

### Key talking points
- Channels are *typed pipes* between goroutines.
- **"Do not communicate by sharing memory; share memory by communicating."** — Go proverb.
- `sync.WaitGroup` is the idiomatic way to wait (not `time.Sleep`).

### Hands-on (5 min)
> "Launch 3 goroutines. Each should compute the square of a number (1, 2, 3) and send it back on a shared channel. Collect all 3 results in main."

---

## Topic 6 — Wrap-up: Standard Library Tour (5 min)

Show how far the standard library gets you — no third-party dependencies needed for common tasks.

```go
// A working HTTP server in 10 lines
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Go! Path: %s", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

```bash
go run main.go
curl http://localhost:8080/hello
```

**Highlight other stdlib packages:**
- `encoding/json` — JSON marshal/unmarshal
- `os`, `io` — file and I/O operations
- `testing` — built-in test framework (`go test ./...`)
- `context` — cancellation and deadlines across goroutines

---

## Final Exercise — Bring It All Together (if time allows)

> Build a small CLI tool that:
> 1. Reads a list of URLs from a text file
> 2. Fetches each URL **concurrently** using goroutines
> 3. Prints the status code and URL for each response

This uses: structs, error handling, goroutines, channels, and the standard library.

Starter skeleton is in `exercises/final/main.go`.

---

## Running Order Checklist

- [ ] Go installed: `go version` should print 1.21+
- [ ] Editor ready (VS Code + Go extension recommended)
- [ ] `go mod init workshop` done in a fresh directory
- [ ] Each topic: explain → live code → run → hands-on

---

## What to Skip (given 1-2 hours)

- Generics (Go 1.18+) — mention exists, skip details
- `reflect` package — advanced, rarely needed
- `cgo` — C interop, out of scope
- Detailed memory model / `sync/atomic` — too deep for intro

---

## Quick Reference Card

```
go mod init <name>     create a new module
go run main.go         compile and run
go build .             compile to binary
go test ./...          run all tests
go fmt ./...           format all code
go vet ./...           static analysis
go get <pkg>           add a dependency
```
