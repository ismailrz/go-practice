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

	// Uncomment and test once you implement safeSqrt above:
	// root, err := safeSqrt(16)
	// if err != nil {
	//     fmt.Println("Error:", err)
	// } else {
	//     fmt.Println("sqrt(16) =", root) // should print: 4
	// }
	//
	// root, err = safeSqrt(-4)
	// if err != nil {
	//     fmt.Println("Error:", err) // should print: cannot take square root of negative number
	// } else {
	//     fmt.Println("sqrt(-4) =", root)
	// }
}

// ======================================================================
// EXERCISE STUB — implement this function
// ======================================================================
// Step 1: Add "math" to the import block at the top alongside "fmt" and "strconv"
// Step 2: Fill in the body below
//
// func safeSqrt(x float64) (float64, error) {
//     if x < 0 {
//         return 0, fmt.Errorf("???")   // return a descriptive error
//     }
//     return ???, nil                   // hint: math.Sqrt(x)
// }

// ======================================================================
// QUESTIONS — Functions, Error Handling & Defer
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  Go has no try/catch/finally. How does Go handle errors instead?
//      What is the main benefit of this approach over exceptions?
//
//      A: Functions return errors as a plain second return value (value, error).
//         The caller checks "if err != nil" immediately after the call.
//         Benefit: error handling is explicit and visible at every call site —
//         you can never accidentally ignore an error path the way you can
//         forget a try/catch block. It also makes control flow obvious;
//         exceptions can jump unpredictably across the call stack.
//
// Q2.  What is the convention for the last return value when a function
//      can fail? What value indicates "no error"?
//
//      A: By convention the error is always the LAST return value.
//         "nil" means no error (success). A non-nil error means something went wrong.
//         Example: func Open(name string) (*File, error)
//
// Q3.  What does "defer" do? When exactly does a deferred call execute?
//      What happens if a function has multiple defer statements?
//
//      A: defer schedules a function call to run just before the surrounding
//         function returns — regardless of whether it returns normally or panics.
//         Multiple defers execute in LIFO order (last registered, first executed),
//         like a stack. This mirrors resource acquisition order:
//         if you open A then B, you want to close B before A.
//
// Q4.  What is the difference between panic() and returning an error?
//      When is it appropriate to use panic() vs returning an error?
//
//      A: Returning an error is for expected, recoverable failure conditions
//         (file not found, invalid input, network timeout). The caller decides
//         what to do.
//         panic() is for programming errors that should never happen
//         (nil pointer dereference, index out of bounds, violated invariants).
//         Rule of thumb: if the caller can reasonably handle it → return error.
//         If it's a bug that means the program is in an invalid state → panic.
//
// Q5.  What does the blank identifier "_" do when used with multiple
//      return values, e.g.:   result, _ := divide(10, 2)
//      Is this a good practice? When might it be acceptable?
//
//      A: "_" discards the error — the compiler won't complain about an unused variable.
//         It's generally bad practice because it silently ignores failures.
//         Acceptable only when you are certain the error can never occur
//         (e.g., parsing a compile-time constant string: strconv.Atoi("42")),
//         or in quick throwaway scripts and tests.
//
// Q6.  What are "named return values" in Go? What is their purpose?
//      Example: func split(sum int) (x, y int) { ... }
//
//      A: Named returns declare the return variables in the function signature.
//         They are initialised to their zero values at function entry.
//         A bare "return" (with no arguments) returns the current values of those variables.
//         Purpose: useful for short functions where the names clarify meaning,
//         and in deferred functions that modify the return value.
//         Overusing them reduces clarity — avoid in long functions.
//
// PRACTICAL
// ---------
// Q7.  What is the output of this code? Explain why.
//
//      func count() {
//          for i := 0; i < 3; i++ {
//              defer fmt.Println(i)
//          }
//      }
//
//      A: Output:  2
//                  1
//                  0
//         Defers execute in LIFO order. The loop registers three defers:
//         Println(0), Println(1), Println(2). When count() returns, they fire
//         in reverse order: 2, then 1, then 0.
//         Also note: defer captures the VALUE of i at the time of the defer call
//         (not a reference), so each defer captures a different snapshot of i.
//
// Q8.  Write a function: maxOf(a, b int) int
//      Then extend it to: maxOfThree(a, b, c int) int
//      using maxOf inside. This reinforces function composition.
//
//      A: func maxOf(a, b int) int {
//             if a > b { return a }
//             return b
//         }
//         func maxOfThree(a, b, c int) int {
//             return maxOf(a, maxOf(b, c))
//         }
//
// Q9.  Write a function: safeDiv(a, b int) (result int, err error)
//      using NAMED return values. Use a bare "return" at the end.
//      What are the trade-offs of named returns?
//
//      A: func safeDiv(a, b int) (result int, err error) {
//             if b == 0 {
//                 err = fmt.Errorf("division by zero")
//                 return   // bare return: result=0, err=set above
//             }
//             result = a / b
//             return       // bare return: result=set above, err=nil
//         }
//         Trade-off: named returns self-document what is returned, but bare returns
//         in long functions make it hard to see what is actually being returned.
//
// Q10. What is the difference between these two error creations?
//      a) errors.New("something went wrong")
//      b) fmt.Errorf("value %d is invalid", n)
//      When would you use each?
//
//      A: errors.New creates a static error with a fixed message string.
//         fmt.Errorf creates a formatted error message with dynamic values embedded.
//         Use errors.New for sentinel errors (errors you compare against with ==).
//         Use fmt.Errorf when the error message needs runtime values for context.
//         Go 1.13+: fmt.Errorf("...: %w", err) wraps an error so callers can
//         unwrap it with errors.Is() or errors.As().
//
// Q11. Write a function openAndRead(path string) that:
//      - Opens a file with os.Open
//      - Defers file.Close()
//      - Returns the first 100 bytes as a string
//      - Returns an error if anything goes wrong
//      This tests defer + error handling together.
//
//      A: func openAndRead(path string) (string, error) {
//             f, err := os.Open(path)
//             if err != nil {
//                 return "", err
//             }
//             defer f.Close()          // guaranteed to run when function returns
//
//             buf := make([]byte, 100)
//             n, err := f.Read(buf)
//             if err != nil && err != io.EOF {
//                 return "", err
//             }
//             return string(buf[:n]), nil
//         }
// ======================================================================
