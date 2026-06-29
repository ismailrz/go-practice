// Package main — executable program.
package main

import "fmt"

// ======================================================================
// CONTROL FLOW IN GO
// ======================================================================
// Go has:
//   for   — the ONLY loop keyword (replaces while, do-while, foreach)
//   if / else if / else — standard, but no parentheses around the condition
//   switch — powerful, no automatic fallthrough, works on any type
//   range  — iterates over arrays, slices, maps, strings, channels
//
// Go does NOT have: while, do-while, foreach, ternary operator (?:)

func main() {
	// ======================================================================
	// PART 1 — if / else
	// ======================================================================
	// No parentheses around the condition — that's a compile error in Go.
	// The opening brace MUST be on the same line (no next-line brace).

	score := 75

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B") // this branch runs
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// ======================================================================
	// if with an initialisation statement
	// ======================================================================
	// Go allows a short statement before the condition, separated by ";".
	// The variable declared here (remainder) is scoped to the if/else block.
	// This is extremely common when calling functions that return (value, error).

	if remainder := score % 10; remainder > 5 {
		fmt.Println("Rounds up:", remainder)
	} else {
		fmt.Println("Rounds down:", remainder) // remainder is visible here too
	}
	// fmt.Println(remainder)  // compile error — remainder out of scope here

	fmt.Println()

	// ======================================================================
	// PART 2 — for loop (3 forms)
	// ======================================================================

	// FORM 1: Classic C-style for loop — init; condition; post
	fmt.Println("--- for (classic) ---")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ") // prints: 0 1 2 3 4
	}
	fmt.Println()

	// FORM 2: Condition only — this is Go's "while" loop.
	// Go has no "while" keyword — just omit init and post.
	fmt.Println("--- for (while-style) ---")
	n := 1
	for n < 100 { // same as:  while (n < 100) in other languages
		n *= 2
	}
	fmt.Println("n =", n) // 128

	// FORM 3: Infinite loop — omit everything.
	// Use "break" to exit, "continue" to skip to the next iteration.
	fmt.Println("--- for (infinite with break) ---")
	count := 0
	for {
		count++
		if count == 3 {
			break // exits the loop immediately
		}
	}
	fmt.Println("count =", count) // 3

	// continue — skips the rest of the current iteration
	fmt.Println("--- for with continue (skip evens) ---")
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue // jump to i++ — skip fmt.Println for even numbers
		}
		fmt.Print(i, " ") // prints: 1 3 5
	}
	fmt.Println()

	fmt.Println()

	// ======================================================================
	// PART 3 — range (iterate over collections)
	// ======================================================================
	// "range" works on: slices, arrays, maps, strings, and channels.
	// It returns two values: the index and the value (or key and value for maps).

	// Range over a SLICE — index + value
	fmt.Println("--- range over slice ---")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("  [%d] %s\n", i, fruit)
	}

	// Discard the index with _ if you only need the value.
	fmt.Println("--- range, index discarded ---")
	for _, fruit := range fruits {
		fmt.Println(" ", fruit)
	}

	// Range over a MAP — key + value (order is random every run)
	fmt.Println("--- range over map ---")
	capitals := map[string]string{
		"France":  "Paris",
		"Germany": "Berlin",
		"Japan":   "Tokyo",
	}
	for country, capital := range capitals {
		fmt.Printf("  %s → %s\n", country, capital)
	}

	// Range over a STRING — gives index (byte position) + rune (Unicode character)
	// Runes, not bytes — Go strings are UTF-8; range decodes one rune at a time.
	fmt.Println("--- range over string ---")
	for i, ch := range "Go!" {
		// %c prints the rune as a character, %d as its Unicode code point.
		fmt.Printf("  index %d: %c (U+%04X)\n", i, ch, ch)
	}

	fmt.Println()

	// ======================================================================
	// PART 4 — switch
	// ======================================================================
	// Go's switch does NOT fall through by default (unlike C/Java/JS).
	// Each case is independent — no "break" needed.
	// Use "fallthrough" keyword if you explicitly want the next case to run.

	day := "Monday"

	fmt.Println("--- switch (basic) ---")
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		// Multiple values in one case — comma separated.
		fmt.Println(day, "is a weekday")
	case "Saturday", "Sunday":
		fmt.Println(day, "is the weekend")
	default:
		// default runs if no case matches — like else.
		fmt.Println("Unknown day")
	}

	// ======================================================================
	// switch with no condition — acts like a chain of if/else if
	// ======================================================================
	// Omitting the expression after "switch" makes it equivalent to "switch true".
	// Each case is a boolean expression — the first true case runs.
	fmt.Println("--- switch (no condition) ---")
	temp := 28
	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp < 15:
		fmt.Println("Cold")
	case temp < 25:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot") // 28 >= 25, so this runs
	}

	// ======================================================================
	// switch with an initialisation statement (same as if)
	// ======================================================================
	fmt.Println("--- switch with init ---")
	switch x := 42; {
	case x > 100:
		fmt.Println("big")
	case x > 10:
		fmt.Println("medium") // 42 > 10, this runs
	default:
		fmt.Println("small")
	}

	// ======================================================================
	// EXERCISE FOR YOU
	// 1. Write a for loop that prints the first 10 Fibonacci numbers
	//    (0, 1, 1, 2, 3, 5, 8, 13, 21, 34) using two variables a, b.
	// 2. Use range over a map of student names → scores.
	//    Print only the students who scored above 80.
	// 3. Write a switch that takes a file extension string (".go", ".py", ".js")
	//    and prints the corresponding language name.
	// ======================================================================
}

// ======================================================================
// QUESTIONS — Control Flow
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  Go has no "while" or "do-while" keywords. How do you write each
//      using "for"?
//
//      A: while loop:
//           for condition { ... }           // same as while (condition)
//         do-while equivalent:
//           for {
//               ...
//               if !condition { break }     // check at end of body
//           }
//
// Q2.  What is the difference between "break" and "continue" inside a loop?
//
//      A: break   — exits the loop entirely. Execution continues after the loop.
//         continue — skips the rest of the current iteration and jumps to the
//                    next iteration (the post statement, then condition check).
//
// Q3.  Go's switch does not fall through by default. Why is this considered
//      an improvement over C/Java switch behaviour?
//
//      A: In C/Java, forgetting a "break" causes unintended fallthrough to the
//         next case — a very common source of bugs. Go inverts the default:
//         each case stops automatically. If you genuinely want fallthrough,
//         you must write "fallthrough" explicitly, making intent clear.
//
// Q4.  What does "range" return when iterating over different types?
//
//      A: []T (slice/array) → index int, value T
//         map[K]V           → key K, value V
//         string            → byte-index int, rune (Unicode code point)
//         chan T             → value T (no index)
//         Use _ to discard either return value.
//
// Q5.  What does "switch { case ...: }" (switch with no expression) do?
//      How is it different from "switch true { ... }"?
//
//      A: They are identical. "switch {" is shorthand for "switch true {".
//         Each case evaluates a boolean expression — the first true case runs.
//         This is a clean replacement for long if/else-if chains.
//
// Q6.  What is the scope of a variable declared in an "if" init statement?
//      e.g.: if err := doSomething(); err != nil { ... }
//
//      A: The variable (err) is scoped to the entire if/else block, including
//         all "else if" and "else" branches. It is NOT accessible outside.
//         This is intentional — it keeps error variables local to the check.
//
// PRACTICAL
// ---------
// Q7.  What is the output of this code?
//
//      for i := 0; i < 5; i++ {
//          if i == 3 { continue }
//          if i == 4 { break }
//          fmt.Println(i)
//      }
//
//      A: 0
//         1
//         2
//         (3 is skipped by continue, 4 triggers break so it never prints)
//
// Q8.  Why does iterating over a map with range produce different output
//      order each run? Is this a bug?
//
//      A: No — it's intentional. Go deliberately randomises map iteration order
//         to prevent programs from relying on undefined behaviour (the internal
//         hash layout can change between runs, versions, and architectures).
//         If you need sorted output, collect the keys into a slice and sort it:
//           keys := make([]string, 0, len(m))
//           for k := range m { keys = append(keys, k) }
//           sort.Strings(keys)
//           for _, k := range keys { fmt.Println(k, m[k]) }
//
// Q9.  Write a FizzBuzz using a switch with no condition:
//      For 1..20, print "Fizz" if divisible by 3, "Buzz" if by 5,
//      "FizzBuzz" if both, otherwise the number.
//
//      A: for i := 1; i <= 20; i++ {
//             switch {
//             case i%3 == 0 && i%5 == 0:
//                 fmt.Println("FizzBuzz")
//             case i%3 == 0:
//                 fmt.Println("Fizz")
//             case i%5 == 0:
//                 fmt.Println("Buzz")
//             default:
//                 fmt.Println(i)
//             }
//         }
//
// Q10. What does "fallthrough" do in a switch? When would you use it?
//
//      A: "fallthrough" forces execution to continue into the NEXT case body
//         regardless of whether that case's condition matches.
//         It transfers control unconditionally — the next case's condition
//         is NOT evaluated.
//         Use rarely and with caution — mainly when you need C-style cascading.
//         Example:
//           switch grade {
//           case "A":
//               fmt.Println("Excellent")
//               fallthrough
//           case "B":
//               fmt.Println("Pass")   // runs for both A and B
//           }
// ======================================================================
