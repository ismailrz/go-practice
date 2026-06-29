// Every Go file starts with a package declaration.
// "main" is special — it tells Go this file is an executable program, not a library.
// If you're building a reusable package (like a utility), you'd write: package mypackage
package main

// "import" brings in other packages.
// "fmt" is short for "format" — it's part of Go's standard library.
// You don't need to install it; it ships with Go itself.
// Unlike Python's print() or JS's console.log(), Go uses fmt.Println() for output.
import "fmt"

// func main() is the entry point of every Go program.
// Go will call this function automatically when you run the program.
// There is exactly ONE main() per executable — no arguments, no return value.
func main() {
	// fmt.Println() prints text to the terminal followed by a newline.
	// The "fmt." prefix means we're calling the Println function from the fmt package.
	// In Go, exported functions always start with a Capital letter (Println, not println).
	fmt.Println("Hello, Go!")

	// ------------------------------------------------------------------
	// EXERCISE
	// Change the string above to print your own name.
	// Then run it with:   go run main.go
	// And build it with:  go build .      (produces a binary called "01-setup")
	// ------------------------------------------------------------------
}

// ======================================================================
// QUESTIONS — Toolchain & Project Setup
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  What is the difference between "go run main.go" and "go build ."?
//      When would you prefer one over the other?
//
//      A: "go run" compiles and immediately runs — no binary is saved to disk.
//         Use it during development for quick feedback.
//         "go build ." compiles and saves a binary you can distribute and run
//         later without Go installed. Use it for deployment or sharing.
//
// Q2.  What does "package main" mean? What would happen if you renamed it
//      to "package hello" but kept func main() inside?
//
//      A: "package main" is the special package name that tells Go this file
//         is the entry point of an executable program.
//         If you rename it to "package hello", Go treats it as a library package
//         and "go build ." will produce no binary — it won't look for func main().
//
// Q3.  Why do exported identifiers in Go start with a capital letter?
//      Give an example of an exported and an unexported name.
//
//      A: Go uses capitalisation as the visibility rule. Capital = exported (public),
//         lowercase = unexported (package-private). No "public"/"private" keywords needed.
//         Exported:   fmt.Println, http.Get, Rectangle.Width
//         Unexported: myHelper(), internalCounter, rectangle (if lowercase)
//
// Q4.  What is go.mod and why is it needed? What problem did it solve
//      compared to the old $GOPATH approach?
//
//      A: go.mod declares the module path and its dependencies with exact versions.
//         Before go.mod (pre-Go 1.11), all projects had to live inside $GOPATH/src —
//         you couldn't have two projects depend on different versions of the same library.
//         go.mod lets each project live anywhere and manage its own dependency versions.
//
// Q5.  What does "go fmt" do, and why does Go enforce a single formatting
//      style instead of leaving it to the developer?
//
//      A: "go fmt" automatically formats your code to the official Go style
//         (tabs for indentation, specific brace placement, etc.).
//         A single enforced style means no style debates in code reviews,
//         all Go code looks the same regardless of who wrote it, and tools
//         (editors, linters) can rely on a predictable structure.
//
// PRACTICAL
// ---------
// Q6.  Create a second file in this directory called "greet.go"
//      (same package main). Add a function Greet(name string) string
//      that returns "Hello, <name>!". Call it from main().
//      What does this teach you about how Go handles multiple files?
//
//      A: All .go files in the same directory with the same package name are
//         compiled together as one package. You don't need to import them —
//         functions defined in greet.go are directly visible in main.go.
//         The package is the unit of compilation, not the file.
//
// Q7.  Run "go build ." and then "go build -o myapp .".
//      What is the difference in the output binary name?
//
//      A: "go build ." names the binary after the directory (e.g., "01-setup").
//         "go build -o myapp ." lets you specify the output name explicitly.
//         The -o flag is useful in CI/CD scripts where the binary name must be fixed.
//
// Q8.  What command would you run to see all the Go environment variables
//      (like GOPATH, GOROOT, GOOS, GOARCH)?
//      Hint: try "go env"
//
//      A: Run "go env" to see all Go environment variables.
//         Key ones:
//           GOROOT  — where Go is installed
//           GOPATH  — workspace for binaries and module cache
//           GOOS    — target operating system (linux, darwin, windows)
//           GOARCH  — target CPU architecture (amd64, arm64)
//         You can cross-compile for another OS with:
//           GOOS=linux GOARCH=amd64 go build .
// ======================================================================
