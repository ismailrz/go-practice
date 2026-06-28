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
