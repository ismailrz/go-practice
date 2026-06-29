// Package main — executable program.
package main

import (
	"fmt"
	"math"
)

// ======================================================================
// TYPE ASSERTIONS & TYPE SWITCHES
// ======================================================================
// When you have a value stored in an interface, the compiler only knows
// the interface type — not the concrete type underneath.
// Type assertions and type switches let you safely recover the concrete type.
//
// This is needed when:
//   - Working with interface{} / any (e.g. from JSON, fmt, or generic functions)
//   - Checking if an interface value also satisfies a second interface
//   - Writing code that behaves differently based on the underlying type

// ======================================================================
// SHAPES — same interfaces as 04-structs, used here for type switch demo
// ======================================================================

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
type Triangle struct{ Base, Height float64 }

func (c Circle) Area() float64    { return math.Pi * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (t Triangle) Area() float64  { return 0.5 * t.Base * t.Height }

// ======================================================================
// A second interface — used to show interface type assertion
// ======================================================================
// Perimeter is NOT implemented by all shapes — only by Rectangle here.
type Perimeter interface {
	Perimeter() float64
}

// Only Rectangle implements Perimeter — Circle and Triangle do not.
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func main() {
	// ======================================================================
	// PART 1 — Type assertion: value.(ConcreteType)
	// ======================================================================
	// Syntax: concrete, ok := interfaceValue.(ConcreteType)
	//   concrete — the value as the concrete type (zero value if failed)
	//   ok       — true if the assertion succeeded, false otherwise
	//
	// ONE-VALUE form: concrete := interfaceValue.(ConcreteType)
	//   Panics if the assertion fails — use only when you are CERTAIN of the type.
	// TWO-VALUE form: concrete, ok := interfaceValue.(ConcreteType)
	//   Never panics — safe for runtime checks.

	var s Shape = Circle{Radius: 5}

	// TWO-VALUE form — safe assertion.
	// This succeeds because s actually holds a Circle.
	if c, ok := s.(Circle); ok {
		fmt.Printf("It's a Circle with radius %.1f and area %.2f\n", c.Radius, c.Area())
	}

	// This fails because s holds a Circle, not a Rectangle.
	if r, ok := s.(Rectangle); ok {
		fmt.Println("It's a Rectangle:", r)
	} else {
		fmt.Println("Not a Rectangle — ok =", ok) // ok = false
	}

	fmt.Println()

	// ONE-VALUE form — use when you're certain (or want a panic on mismatch).
	c := s.(Circle) // panics if s is not a Circle
	fmt.Printf("Asserted Circle radius: %.1f\n", c.Radius)

	// Uncomment to see the panic:
	// _ = s.(Rectangle)  // panic: interface conversion: interface is Circle, not Rectangle

	fmt.Println()

	// ======================================================================
	// PART 2 — Interface assertion: check if a value satisfies a second interface
	// ======================================================================
	// A type assertion can also target an INTERFACE — not just a concrete type.
	// This checks whether the underlying value also implements that second interface.

	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Circle{Radius: 5},
		Triangle{Base: 6, Height: 4},
	}

	fmt.Println("--- checking for Perimeter interface ---")
	for _, sh := range shapes {
		// Does this shape also implement Perimeter?
		if p, ok := sh.(Perimeter); ok {
			// Only Rectangle reaches here — Circle and Triangle don't have Perimeter().
			fmt.Printf("%T has perimeter: %.2f\n", sh, p.Perimeter())
		} else {
			fmt.Printf("%T does not implement Perimeter\n", sh)
		}
	}

	fmt.Println()

	// ======================================================================
	// PART 3 — Type switch: handle multiple concrete types cleanly
	// ======================================================================
	// A type switch is like a regular switch, but each case specifies a TYPE.
	// The variable in each case takes on the concrete type — no assertion needed.
	// Syntax: switch v := interfaceValue.(type) { case T1: ... case T2: ... }
	// The special keyword ".(type)" only works inside a type switch.

	fmt.Println("--- type switch ---")
	for _, sh := range shapes {
		// v takes the concrete type of sh in each case branch.
		switch v := sh.(type) {
		case Circle:
			// Inside this branch, v is a Circle — all Circle fields are accessible.
			fmt.Printf("Circle: radius=%.1f, area=%.2f\n", v.Radius, v.Area())
		case Rectangle:
			// Inside this branch, v is a Rectangle.
			fmt.Printf("Rectangle: %.0fx%.0f, area=%.2f, perimeter=%.2f\n",
				v.Width, v.Height, v.Area(), v.Perimeter())
		case Triangle:
			fmt.Printf("Triangle: base=%.1f height=%.1f, area=%.2f\n",
				v.Base, v.Height, v.Area())
		default:
			// Handles any shape type not listed above.
			fmt.Printf("Unknown shape: %T\n", v)
		}
	}

	fmt.Println()

	// ======================================================================
	// PART 4 — Type switch on any (interface{})
	// ======================================================================
	// The most common use of type switches is with any (interface{}) —
	// a value that could be ANY type. This is common when:
	//   - Parsing JSON (values come back as interface{})
	//   - Writing generic utility functions
	//   - Handling fmt.Println-style variadic arguments

	values := []any{42, "hello", true, 3.14, []int{1, 2, 3}, nil}

	fmt.Println("--- type switch on any ---")
	for _, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("int:    %d\n", val)
		case string:
			fmt.Printf("string: %q\n", val)
		case bool:
			fmt.Printf("bool:   %t\n", val)
		case float64:
			fmt.Printf("float:  %.2f\n", val)
		case []int:
			fmt.Printf("[]int:  %v (len=%d)\n", val, len(val))
		case nil:
			fmt.Println("nil value")
		default:
			// %T prints the concrete type name — useful for debugging unknown types.
			fmt.Printf("unknown type: %T = %v\n", val, val)
		}
	}

	// ======================================================================
	// EXERCISE FOR YOU
	// 1. Add a Square struct (Side float64) with Area() and Perimeter() methods.
	//    Update the type switch in PART 3 to handle Square with its own case.
	// 2. Write a function: describe(v any) string
	//    that returns a human-readable description of any value using a type switch.
	//    Handle: int, float64, string, bool, []int, map[string]int, nil.
	// ======================================================================
}

// ======================================================================
// QUESTIONS — Type Assertions & Type Switches
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  What is a type assertion? What are the two forms and when do you
//      use each?
//
//      A: A type assertion extracts the concrete value from an interface.
//         ONE-VALUE:  v := i.(T)        — panics if i does not hold a T.
//                                         Use when you are certain of the type.
//         TWO-VALUE:  v, ok := i.(T)   — ok=false if mismatch, never panics.
//                                         Use for safe runtime checks.
//         The two-value form is almost always preferred — panics are hard to debug.
//
// Q2.  What is a type switch? How is it different from a regular switch?
//
//      A: A regular switch compares VALUES (switch x { case 1: case 2: }).
//         A type switch compares TYPES (switch v := x.(type) { case int: case string: }).
//         In each case branch of a type switch, the variable v automatically has
//         the concrete type of that case — no additional assertion needed.
//         The keyword ".(type)" is special syntax only valid in a type switch.
//
// Q3.  What is "any" in Go? Is it the same as interface{}?
//
//      A: Yes — "any" is an alias for interface{} introduced in Go 1.18.
//         interface{} is satisfied by EVERY type (it has zero method requirements).
//         "any" is just cleaner to write. Both work identically.
//         Use sparingly — it loses compile-time type safety. When you use any,
//         you must do type assertions or type switches at runtime.
//
// Q4.  Can you do a type assertion to an interface type (not a struct)?
//      What does that check?
//
//      A: Yes. i.(SomeInterface) checks whether the concrete value stored in i
//         also satisfies SomeInterface.
//         Example: if p, ok := shape.(Perimeter); ok { ... }
//         This is how you check for "optional" behaviour — a type that may or may
//         not implement an additional interface beyond the one you have in hand.
//
// Q5.  When would you use a type switch over multiple if/ok assertions?
//
//      A: When you need to handle 3+ types differently, a type switch is cleaner:
//         - Single expression tested once (not repeated i.(T) calls)
//         - Each branch gets the correctly-typed variable automatically
//         - "default" handles unexpected types without a catch-all if
//         Use if/ok assertions when you only need to check 1-2 specific types.
//
// PRACTICAL
// ---------
// Q6.  What is the output of this code?
//
//      var i interface{} = "hello"
//      s, ok := i.(string)
//      n, ok2 := i.(int)
//      fmt.Println(s, ok)
//      fmt.Println(n, ok2)
//
//      A: hello true
//         0 false
//         First assertion succeeds (i holds "hello", a string) → s="hello", ok=true
//         Second assertion fails (i is not an int) → n=0 (zero value), ok2=false
//
// Q7.  What happens here?
//
//      var s Shape = Circle{Radius: 3}
//      r := s.(Rectangle)
//      fmt.Println(r)
//
//      A: PANIC at runtime: "interface conversion: interface is main.Circle, not main.Rectangle"
//         The one-value form panics when the concrete type doesn't match.
//         Fix: use r, ok := s.(Rectangle) and check ok before using r.
//
// Q8.  Write a function: sumInts(values []any) int
//      that sums only the int values from a mixed []any slice,
//      ignoring all other types.
//
//      A: func sumInts(values []any) int {
//             total := 0
//             for _, v := range values {
//                 if n, ok := v.(int); ok {
//                     total += n
//                 }
//             }
//             return total
//         }
//         sumInts([]any{1, "x", 2, true, 3}) → 6
// ======================================================================
