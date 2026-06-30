# Go Language — Practical Session

A hands-on tutorial for developers who already know at least one programming language
(Python, Java, JavaScript, etc.) and want to learn Go efficiently.

Each exercise file is self-contained with:
- **Line-by-line comments** explaining every concept
- **Live runnable demo code**
- **Hands-on exercises** to try during the session
- **Q&A section** with questions and short answers at the bottom

---

## Prerequisites

- Go 1.21 or later — download from [https://go.dev/dl](https://go.dev/dl)
- Verify installation:

```bash
go version
```

- Recommended editor: **VS Code** with the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

---

## Setup

```bash
git clone <this-repo>
cd go-practice
```

No `go get` or dependency installation needed — all exercises use the standard library only.

---

## Running an Exercise

```bash
# From the project root
go run exercises/01-setup/main.go

# Or navigate into a folder
cd exercises/02-types
go run .
```

Running the tests (exercise 11):

```bash
cd exercises/11-testing
go test -v .
go test -cover .
```

---

## Session Plan

Full breakdown with timing, key points, and talking notes: [`PLAN.md`](PLAN.md)

---

## Exercises

| # | Topic | Key Concepts |
|---|-------|-------------|
| [01-setup](exercises/01-setup/main.go) | Toolchain & Project Setup | `go mod`, `go run`, `go build`, `go fmt`, package system |
| [02-types](exercises/02-types/main.go) | Variables & Types | `:=` vs `var`, zero values, maps, `make()`, nil map panic |
| [03-functions](exercises/03-functions/main.go) | Functions & Error Handling | Multiple returns, `(value, error)` pattern, `defer`, no try/catch |
| [04-structs](exercises/04-structs/main.go) | Structs & Interfaces | Value vs pointer receivers, implicit interfaces, polymorphism, `fmt.Stringer` |
| [05-goroutines](exercises/05-goroutines/main.go) | Goroutines & Channels | `go`, unbuffered/buffered channels, `WaitGroup`, `select`, timeouts |
| [06-pointers](exercises/06-pointers/main.go) | Pointers | `&`, `*`, pass-by-value vs pointer, nil pointer, `new()` |
| [07-control-flow](exercises/07-control-flow/main.go) | Control Flow | `for` (3 forms), `range`, `switch`, no `while` keyword |
| [08-slices](exercises/08-slices/main.go) | Slices | `append`, `copy`, length vs capacity, shared backing array, 2D slices |
| [09-strings](exercises/09-strings/main.go) | Strings & Runes | UTF-8, bytes vs runes, `range` on strings, `strings` package, `Builder` |
| [10-type-assertions](exercises/10-type-assertions/main.go) | Type Assertions & Switches | `v.(T)`, `v, ok := i.(T)`, type switch, `any` / `interface{}` |
| [11-testing](exercises/11-testing/) | Unit Testing | `go test`, `*testing.T`, table-driven tests, `t.Run`, subtests, coverage |
| [final](exercises/final/main.go) | Capstone Project | Concurrent URL fetcher — combines structs, errors, goroutines, channels, stdlib |

---

## Hands-on Practice Files

Shorter, single-topic files for quick practice during the session.
Run any file with `go run exercises/hands-on/<file>.go`.

| File | Concepts |
|------|---------|
| [types.go](exercises/hands-on/types.go) | bool, int, float64, string, byte, rune, conversions, iota |
| [loop.go](exercises/hands-on/loop.go) | classic for, while-style, break, continue, range, FizzBuzz |
| [array.go](exercises/hands-on/array.go) | declaration, access, range, value semantics, array→slice |
| [slice.go](exercises/hands-on/slice.go) | make, append, slicing, copy, delete element |
| [map.go](exercises/hands-on/map.go) | literal, make, CRUD, comma-ok, range, word count |
| [function.go](exercises/hands-on/function.go) | multiple returns, variadic, first-class functions, closures, defer, recursion |
| [pointer.go](exercises/hands-on/pointer.go) | `&`, `*`, pass-by-value vs pointer, nil check, `new()` |
| [struct.go](exercises/hands-on/struct.go) | value/pointer receivers, `fmt.Stringer`, embedding, interfaces |
| [error.go](exercises/hands-on/error.go) | `errors.New`, `fmt.Errorf`, sentinel errors, custom error type, `errors.Is/As` |
| [goroutine.go](exercises/hands-on/goroutine.go) | goroutines, unbuffered/buffered channels, WaitGroup, select, Mutex |
| [goroutine-channel.go](exercises/hands-on/goroutine-channel.go) | channel basics — send, receive, buffered channel |
| [goroutine-wait-group.go](exercises/hands-on/goroutine-wait-group.go) | WaitGroup pattern, concurrent tasks |

---

## Coverage vs Go Fundamentals

### Priority 1 — Covered ✅

| Topic | Exercise |
|-------|----------|
| Variables | 02-types, hands-on/types.go |
| Functions | 03-functions, hands-on/function.go |
| Arrays & Slices | 08-slices, hands-on/array.go, hands-on/slice.go |
| Maps | 02-types, hands-on/map.go |
| Structs | 04-structs, hands-on/struct.go |
| Methods | 04-structs, hands-on/struct.go |
| Pointers | 06-pointers, hands-on/pointer.go |
| Interfaces | 04-structs, 10-type-assertions, hands-on/struct.go |
| Error Handling | 03-functions, hands-on/error.go |
| Goroutines | 05-goroutines, hands-on/goroutine.go, hands-on/goroutine-channel.go |
| Channels | 05-goroutines, hands-on/goroutine-channel.go |
| WaitGroup | 05-goroutines, hands-on/goroutine-wait-group.go |
| Mutex | hands-on/goroutine.go |
| Unit Testing | 11-testing |

### Priority 1 — Not Covered

| Topic | Notes |
|-------|-------|
| JSON | `encoding/json`, Marshal/Unmarshal, struct tags |
| Context | Cancellation, deadlines, goroutine lifecycle |

### Priority 2 (Advanced / Follow-up Session)

Generics, Reflection, File Handling, HTTP Server, Dependency Injection, Project Structure

---

## Quick Reference

```bash
go mod init <name>   # create a new module
go run main.go       # compile and run
go build .           # compile to binary
go test ./...        # run all tests
go test -v .         # verbose test output
go test -cover .     # test coverage report
go fmt ./...         # format all code
go vet ./...         # static analysis
go get <pkg>         # add a dependency
go doc fmt.Println   # view documentation for any symbol
```

---

## Go Proverbs (design philosophy in one line each)

- *Don't communicate by sharing memory; share memory by communicating.*
- *Errors are values.*
- *The bigger the interface, the weaker the abstraction.*
- *Clear is better than clever.*
- *A little copying is better than a little dependency.*
