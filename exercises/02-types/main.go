// Package main — this is an executable program.
package main

// We import "fmt" for printing output to the terminal.
import "fmt"

func main() {
	// ======================================================================
	// PART 1 — Variable Declaration
	// ======================================================================

	// := is the SHORT variable declaration operator (only works inside functions).
	// Go INFERS the type from the right-hand side — you don't have to write "int" or "string".
	// This is equivalent to: var x int = 42
	x := 42

	// Go infers this as type string.
	name := "Alice"

	// Printing multiple values: Println accepts any number of arguments.
	fmt.Println("x =", x)
	fmt.Println("name =", name)

	// ======================================================================
	// PART 2 — var keyword (explicit declaration)
	// ======================================================================

	// "var" is used when you want to declare a variable WITHOUT immediately assigning a value,
	// or when you need to be explicit about the type (e.g., at package level, or for clarity).
	var count int     // declares count as int — Go sets it to 0 automatically
	var message string // declares message as string — Go sets it to "" automatically

	// In Python, uninitialized variables cause NameError.
	// In JavaScript, uninitialized variables are "undefined".
	// In Go, EVERY variable always has a well-defined ZERO VALUE from the moment it's declared.
	// This eliminates an entire category of null/undefined bugs.
	fmt.Println("count (zero value) =", count)     // prints: 0
	fmt.Println("message (zero value) =", message) // prints: (empty string)

	// ======================================================================
	// PART 3 — Zero Value table
	// ======================================================================
	// Type         | Zero Value
	// -------------|------------
	// int, float64 | 0
	// string       | ""
	// bool         | false
	// pointer      | nil
	// slice        | nil
	// map          | nil

	var flag bool
	fmt.Println("flag (zero value) =", flag) // prints: false

	// ======================================================================
	// PART 4 — Multiple assignment
	// ======================================================================

	// Go allows assigning multiple variables in one line.
	a, b := 10, 20
	fmt.Println("Before swap: a =", a, ", b =", b)

	// Swap without a temporary variable — Go evaluates the right side fully before assigning.
	// In most languages you'd need: temp := a; a = b; b = temp
	a, b = b, a
	fmt.Println("After swap:  a =", a, ", b =", b)

	// ======================================================================
	// PART 5 — Maps and the "nil map" panic (IMPORTANT!)
	// ======================================================================

	// A map declared with "var" gets the zero value for maps, which is nil.
	// A nil map is READ-safe (reading from it returns zero values without panic),
	// but WRITE-unsafe (writing to a nil map causes a runtime panic).
	var scores map[string]int // scores is nil here

	// Safely reading from a nil map — returns 0 (zero value for int), no panic.
	fmt.Println("Read from nil map:", scores["Alice"]) // prints: 0

	// Uncomment the next line to SEE the panic in action:
	// scores["Alice"] = 10  // panic: assignment to entry in nil map

	// FIX: use make() to allocate memory for the map before writing.
	// make() is a built-in function that initializes maps, slices, and channels.
	scores = make(map[string]int)

	// Now we can safely write to the map.
	scores["Alice"] = 10
	scores["Bob"] = 20
	fmt.Println("scores map:", scores)

	// ======================================================================
	// EXERCISE
	// 1. Declare a map[string]string called "capitals" (cities → country capitals).
	// 2. Add at least 3 entries (e.g., "France" → "Paris").
	// 3. Print the entire map and also look up one specific key.
	// ======================================================================
}
