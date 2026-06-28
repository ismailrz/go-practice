// Package main — executable program.
package main

// We need two packages:
// "fmt"     — for printing output
// "strconv" — string conversion utilities (strconv.Atoi converts string → int)
import (
	"fmt"
	"strconv"
)

// ======================================================================
// FUNCTIONS IN GO
// ======================================================================
// Syntax:  func functionName(param type) returnType { ... }
// Go supports MULTIPLE return values — this is the idiomatic way to return
// both a result AND an error. There are no exceptions in Go (no try/catch).
// Instead, errors are plain values returned alongside the result.
// The caller is forced to handle (or explicitly ignore) the error.

// divide takes two float64 numbers and returns:
//   - float64: the result (or 0 if there's an error)
//   - error:   nil on success, or a descriptive error value on failure
//
// "error" is a built-in interface in Go — it's not a class or a type you import.
func divide(a, b float64) (float64, error) {
	// Guard clause: check for invalid input first.
	// errors.New() creates a simple error with a message string.
	// We return TWO values: the zero value (0) as the result, and the error.
	if b == 0 {
		// In Python/Java you would: raise ValueError("...")
		// In Go you: return 0, errors.New("...")
		// We use fmt.Errorf here for convenience — it formats a string into an error.
		return 0, fmt.Errorf("cannot divide %.2f by zero", a)
	}

	// Happy path: return the result and nil for the error.
	// "nil" means "no error" — like null/None, but only for pointers, maps, slices, interfaces, etc.
	return a / b, nil
}

// ======================================================================
// DEFER
// ======================================================================
// "defer" schedules a function call to run AFTER the surrounding function returns.
// It's how Go does cleanup (like "finally" in Java/Python), but attached to the
// resource acquisition site rather than a distant try/finally block.
// If multiple defers exist, they run in LIFO order (last in, first out).

func demonstrateDefer() {
	fmt.Println("--- defer demo start ---")

	// This will print LAST, even though it's written first.
	// defer "registers" the call — it doesn't execute it yet.
	defer fmt.Println("deferred: I run after the function returns")

	fmt.Println("step 1")
	fmt.Println("step 2")

	// When demonstrateDefer() is about to return, the deferred call runs.
	fmt.Println("--- defer demo end ---")
}

// ======================================================================
// EXERCISE FUNCTION
// ======================================================================
// parseInt wraps strconv.Atoi to convert a string to an integer.
// strconv.Atoi already returns (int, error), so we just pass it through.
// This pattern — wrapping a stdlib function and returning its error — is
// extremely common in Go codebases.
func parseInt(s string) (int, error) {
	// strconv.Atoi("42")  → 42, nil
	// strconv.Atoi("abc") → 0,  error("strconv.Atoi: parsing \"abc\": invalid syntax")
	return strconv.Atoi(s)
}

func main() {
	// ======================================================================
	// CALLING A FUNCTION THAT RETURNS (value, error)
	// ======================================================================

	// The idiomatic Go pattern:
	//   result, err := someFunction(...)
	//   if err != nil { handle the error }
	//   use result
	//
	// You CANNOT skip this check (well, you can use _ to discard, but you shouldn't).
	// This makes error handling explicit and visible in every function call.

	result, err := divide(10, 2)
	if err != nil {
		// Handle the error — in a real program you might log it, return it, or exit.
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result) // prints: 5
	}

	// Now call divide with b=0 to trigger the error path.
	result, err = divide(7, 0) // note: = not :=  (variables already declared above)
	if err != nil {
		fmt.Println("Error:", err) // prints: cannot divide 7.00 by zero
	} else {
		fmt.Println("7 / 0 =", result)
	}

	fmt.Println() // blank line for readability

	// ======================================================================
	// DEFER IN ACTION
	// ======================================================================
	demonstrateDefer()
	// Output:
	//   --- defer demo start ---
	//   step 1
	//   step 2
	//   --- defer demo end ---
	//   deferred: I run after the function returns

	fmt.Println()

	// ======================================================================
	// EXERCISE — parseInt
	// ======================================================================

	// Test parseInt with a valid number string.
	n, err := parseInt("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed integer:", n) // prints: 42
	}

	// Test parseInt with an invalid string — should return an error.
	n, err = parseInt("abc")
	if err != nil {
		fmt.Println("Error:", err) // prints: strconv.Atoi: parsing "abc": invalid syntax
	} else {
		fmt.Println("Parsed integer:", n)
	}

	// ======================================================================
	// EXERCISE FOR YOU
	// Write a function: safeSqrt(x float64) (float64, error)
	// - If x < 0, return an error: "cannot take square root of negative number"
	// - Otherwise return math.Sqrt(x), nil
	// (You'll need to add "math" to the import block at the top)
	// ======================================================================
}
