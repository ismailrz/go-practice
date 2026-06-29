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
	var count int      // declares count as int — Go sets it to 0 automatically
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

// ======================================================================
// QUESTIONS — Types, Variables & Zero Values
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  What is the difference between ":=" and "="?
//      Can you use ":=" at the package level (outside a function)? Why not?
//
//      A: ":=" is the short variable declaration — it declares AND assigns in one step,
//         with the type inferred from the right-hand side.
//         "=" is plain assignment — the variable must already be declared.
//         ":=" only works inside functions. At package level, Go requires explicit "var"
//         declarations because the compiler processes them in a specific order,
//         and ":=" would be ambiguous.
//
// Q2.  What is a zero value in Go? Why is it safer than null/undefined
//      in languages like JavaScript or Python?
//
//      A: A zero value is the default value every variable gets when declared
//         without an explicit assignment. Go guarantees this — you can never
//         read an uninitialised variable.
//         In JS, accessing an undefined variable can silently propagate NaN or
//         crash at an unexpected point. In Go, you always know exactly what
//         you have: 0, "", false, or nil — no surprises.
//
// Q3.  What is the zero value for each of these types?
//      a) int       b) string      c) bool
//      d) float64   e) []int       f) map[string]int
//
//      A: a) 0        b) ""      c) false
//         d) 0.0      e) nil     f) nil
//
// Q4.  Why does writing to a nil map cause a panic but reading from it
//      does not? What does this tell you about how maps work internally?
//
//      A: A nil map has no underlying hash table allocated.
//         Reading is safe because Go detects nil and returns the zero value
//         for the value type (no memory needed).
//         Writing requires allocating a bucket in the hash table — which
//         can't be done on nil, so Go panics rather than silently corrupt memory.
//         Fix: always initialise maps with make(map[K]V) before writing.
//
// Q5.  What is the difference between a slice and an array in Go?
//      When would you choose one over the other?
//
//      A: Array: fixed size, part of the type — [3]int and [4]int are different types.
//               Copied by value when passed to functions.
//         Slice: dynamic size, a view over an underlying array (pointer + length + capacity).
//               Passed by reference-like header — cheap to pass around.
//         In practice, use slices almost always. Arrays are used when the size is
//         fixed and known at compile time (e.g., [16]byte for a UUID or hash).
//
// Q6.  What does the blank identifier "_" do?
//      Give a situation where you would use it.
//
//      A: "_" discards a value — Go requires every declared variable to be used,
//         so _ lets you intentionally ignore a return value without a compile error.
//         Common uses:
//           result, _ := strconv.Atoi(s)      // ignore the error (risky but sometimes OK)
//           for _, v := range items { ... }   // ignore the index
//
// PRACTICAL
// ---------
// Q7.  What is the output of the following code, and why?
//
//      var s []int
//      fmt.Println(s == nil)   // ?
//      fmt.Println(len(s))     // ?
//      s = append(s, 1, 2, 3)
//      fmt.Println(s)          // ?
//
//      A: true      — a var-declared slice is nil
//         0         — len(nil slice) is 0, not a panic
//         [1 2 3]   — append works on nil slices, allocating a new backing array
//
// Q8.  Declare a map[string][]string that maps a person's name to a list
//      of their hobbies. Add at least 2 people with multiple hobbies each.
//      Loop over the map and print each person and their hobbies.
//
//      A: hobbies := map[string][]string{
//             "Alice": {"reading", "cycling"},
//             "Bob":   {"gaming", "cooking", "hiking"},
//         }
//         for name, list := range hobbies {
//             fmt.Printf("%s: %v\n", name, list)
//         }
//
// Q9.  How do you check if a key exists in a map?
//      Write a short snippet using the "comma ok" idiom.
//      (Hint: val, ok := myMap["key"])
//
//      A: val, ok := scores["Alice"]
//         if ok {
//             fmt.Println("Found:", val)
//         } else {
//             fmt.Println("Key not found")
//         }
//         Without the ok check, a missing key just returns the zero value (0 for int),
//         which is indistinguishable from a key that exists with value 0.
//
// Q10. What is the difference between:
//      a) var s string = "hello"
//      b) s := "hello"
//      c) const s = "hello"
//      In what situations would you use each?
//
//      A: a) Verbose explicit declaration — useful at package level or when
//            you want to be clear about the type.
//         b) Short declaration — idiomatic inside functions when the type is obvious.
//         c) Compile-time constant — the value is baked in at compile time,
//            cannot be changed at runtime, and can be used in array sizes,
//            switch cases, and iota enumerations. Use for magic values like
//            MaxRetries = 3 or Pi = 3.14159.
// ======================================================================
