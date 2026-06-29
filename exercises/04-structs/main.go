// Package main — executable program.
package main

import "fmt"

// ======================================================================
// PERIMETER INTERFACE — a second interface alongside Shape
// ======================================================================
// A type can satisfy multiple interfaces simultaneously with no extra code.
// Any struct that has both Area() and Perimeter() methods satisfies BOTH interfaces.
type Perimeter interface {
	Perimeter() float64
}

// ======================================================================
// STRUCTS — Go's version of classes (but simpler)
// ======================================================================
// Go has NO classes, NO class-based inheritance, NO constructors.
// Instead, you define STRUCTS (plain data containers) and attach METHODS to them.
// This forces composition over inheritance — widely considered a cleaner design.

// Rectangle is a struct with two fields.
// Capital letter = exported (visible outside this package), like "public" in Java.
// Lowercase letter = unexported (private to this package), like "private" in Java.
type Rectangle struct {
	Width  float64 // exported field
	Height float64 // exported field
}

// ======================================================================
// METHODS — functions attached to a type
// ======================================================================
// Syntax: func (receiver TypeName) MethodName(params) returnType { ... }
// The receiver is like "self" in Python or "this" in Java/JS.
// Two kinds of receivers:
//   VALUE receiver  (r Rectangle)  — works on a copy, cannot modify the original
//   POINTER receiver (*Rectangle)  — works on the original, CAN modify it

// Area uses a VALUE receiver because it only reads the struct, never modifies it.
// Go automatically lets you call value-receiver methods on both values AND pointers.
func (r Rectangle) Area() float64 {
	// r here is a COPY of the Rectangle — changing r.Width wouldn't affect the caller.
	return r.Width * r.Height
}

// Perimeter adds a second method — now Rectangle satisfies BOTH Shape AND Perimeter.
// No declaration of "implements" needed for either interface.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// String satisfies the built-in fmt.Stringer interface (from the "fmt" package).
// fmt.Println automatically calls String() if a type has it — no extra code needed.
// This is Go's standard way to give a type a human-readable representation.
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.0f × %.0f)", r.Width, r.Height)
}

// Scale uses a POINTER receiver because it needs to MODIFY the original Rectangle.
// The * before Rectangle means "pointer to Rectangle".
// If you used a value receiver here, Scale() would modify a copy and the original would be unchanged.
func (r *Rectangle) Scale(factor float64) {
	// r is a pointer — r.Width dereferences it automatically (no need for (*r).Width).
	r.Width *= factor // same as r.Width = r.Width * factor
	r.Height *= factor
}

// ======================================================================
// INTERFACES — defining behavior, not data
// ======================================================================
// An interface specifies a SET OF METHODS that a type must have.
// Unlike Java/C#, Go interfaces are IMPLICIT — a type satisfies an interface
// automatically just by having the right methods. You never write "implements Shape".
// This is called "structural typing" or "duck typing with compile-time checking".

// Shape is an interface. Any type with an Area() float64 method satisfies it.
// Rectangle satisfies Shape. Circle (below) satisfies Shape.
// Neither of them declares "implements Shape" — Go figures it out.
type Shape interface {
	Area() float64
}

// Circle is a second struct that also satisfies the Shape interface.
type Circle struct {
	Radius float64
}

// Circle has an Area() method — this makes Circle automatically satisfy Shape.
// Go sees: "Circle has Area() float64" → "Circle implements Shape" — silently, at compile time.
func (c Circle) Area() float64 {
	// math.Pi is more precise, but we use a literal here to keep imports minimal.
	return 3.14159 * c.Radius * c.Radius
}

// ======================================================================
// TRIANGLE — the exercise struct
// ======================================================================
// Triangle also satisfies the Shape interface because it has Area() float64.
// printArea (below) accepts ANY Shape — so it works with Rectangle, Circle, AND Triangle
// without any changes to printArea itself.
type Triangle struct {
	Base   float64
	Height float64
}

// Area for a triangle = ½ × base × height.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// ======================================================================
// POLYMORPHISM via interfaces
// ======================================================================
// printArea accepts a Shape interface — it doesn't care whether you pass a
// Rectangle, Circle, Triangle, or any future type that has an Area() method.
// This is Go's version of polymorphism — no inheritance needed.
func printArea(s Shape) {
	// s.Area() calls the correct Area() based on the actual type at runtime.
	// This is called dynamic dispatch.
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	// ======================================================================
	// CREATING STRUCTS
	// ======================================================================

	// Struct literal with named fields — the recommended style (order doesn't matter).
	rect := Rectangle{Width: 4, Height: 3}

	// Calling a method — same syntax as any other language.
	fmt.Println("Rectangle area:", rect.Area()) // 12.00

	// Scale is a pointer-receiver method. Go AUTOMATICALLY takes the address of rect
	// when calling a pointer-receiver method on a value. You don't need &rect.Scale(2).
	rect.Scale(2)
	fmt.Println("After Scale(2), area:", rect.Area()) // 48.00

	fmt.Println()

	// ======================================================================
	// fmt.Stringer — automatic String() call
	// ======================================================================
	// fmt.Println checks if a value has a String() method (satisfies fmt.Stringer).
	// If it does, Println calls it automatically — you get a custom representation.
	rect2 := Rectangle{Width: 4, Height: 3}
	fmt.Println(rect2) // prints: Rectangle(4 × 3)  ← calls rect2.String() automatically

	fmt.Println()

	// ======================================================================
	// MULTIPLE INTERFACES — one type, two contracts
	// ======================================================================
	// Rectangle has Area(), Perimeter(), and String() — so it satisfies:
	//   Shape      (requires Area())
	//   Perimeter  (requires Perimeter())
	//   fmt.Stringer (requires String())
	// All at the same time, with zero extra declarations.
	var p Perimeter = Rectangle{Width: 5, Height: 3}
	fmt.Println("Perimeter:", p.Perimeter()) // 16

	fmt.Println()

	// ======================================================================
	// INTERFACE IN ACTION — printArea accepts ANY Shape
	// ======================================================================

	// Passing a Rectangle — printArea doesn't know or care it's a Rectangle.
	printArea(Rectangle{Width: 4, Height: 3}) // Area: 12.00

	// Passing a Circle — same function, different type, correct behavior.
	printArea(Circle{Radius: 5}) // Area: 78.54

	// Passing a Triangle — added AFTER printArea was written, no changes needed.
	printArea(Triangle{Base: 6, Height: 4}) // Area: 12.00

	fmt.Println()

	// ======================================================================
	// SLICE OF INTERFACES — storing mixed types in one collection
	// ======================================================================
	// Because Rectangle, Circle, and Triangle all satisfy Shape,
	// we can put them all in a []Shape slice and loop over them.
	shapes := []Shape{
		Rectangle{Width: 10, Height: 2},
		Circle{Radius: 3},
		Triangle{Base: 8, Height: 5},
	}

	// Range loop: index i and value s for each element in shapes.
	// _ discards the index (like _ in Python). If you need the index, write: for i, s := range shapes
	for _, s := range shapes {
		// s.Area() dispatches to the correct method for each concrete type.
		fmt.Printf("%T → area = %.2f\n", s, s.Area())
		// %T is a format verb that prints the concrete type name (e.g., main.Rectangle)
	}

	// ======================================================================
	// EXERCISE FOR YOU
	// 1. Add a Pentagon struct with a Side float64 field.
	//    Area of a regular pentagon = (Side² × √25+10√5) / 4  ≈ Side² × 1.720
	//    (or just use: 1.720 * side * side as an approximation)
	// 2. Give it an Area() method so it satisfies Shape.
	// 3. Add it to the shapes slice above and run the program again.
	//    printArea and the range loop will work without any other changes.
	// ======================================================================
}

// ======================================================================
// QUESTIONS — Structs, Methods & Interfaces
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  Go has no classes. What is the Go equivalent, and how do you
//      attach behaviour (methods) to data?
//
//      A: Go uses structs to hold data and methods to define behaviour.
//         A method is just a function with a receiver argument:
//           func (r Rectangle) Area() float64 { ... }
//         There is no constructor either — use a plain function like
//         NewRectangle(w, h float64) Rectangle as a convention.
//
// Q2.  What is the difference between a value receiver and a pointer
//      receiver? Write a rule of thumb for when to use each.
//
//      A: Value receiver  (r Rectangle)  — operates on a COPY. Cannot modify original.
//         Pointer receiver (*Rectangle)  — operates on the ORIGINAL. Can modify it.
//         Rule of thumb:
//           Use pointer receiver if the method modifies the struct.
//           Use pointer receiver if the struct is large (avoids copying).
//           Use value receiver if the method only reads and the struct is small.
//           Be consistent: if any method uses a pointer receiver, use pointer
//           receivers for all methods on that type.
//
// Q3.  Go interfaces are "implicit" — what does that mean?
//      How is this different from Java's "implements" keyword?
//
//      A: In Go, a type satisfies an interface automatically if it has all
//         the required methods — no declaration is needed.
//         In Java you must write: class Circle implements Shape { ... }
//         In Go, if Circle has Area() float64, it satisfies Shape silently.
//         Benefit: you can define an interface for a type you don't own,
//         and third-party types satisfy your interface without modification.
//
// Q4.  What is the empty interface (interface{} or "any") and when
//      would you use it? What is the trade-off?
//
//      A: interface{} (alias: "any" since Go 1.18) is satisfied by every type —
//         it has zero required methods so anything fits.
//         Use it when you genuinely don't know the type at compile time,
//         e.g., fmt.Println(a ...any), JSON unmarshalling into interface{}.
//         Trade-off: you lose type safety — you must use type assertions or
//         type switches to get the value back, which can panic at runtime.
//         Prefer generics (Go 1.18+) or concrete types when possible.
//
// Q5.  Can a struct implement multiple interfaces at the same time?
//      Give an example using the Shape interface and a new
//      Perimeter interface.
//
//      A: Yes — a type can satisfy as many interfaces as it has methods for.
//         type Perimeter interface { Perimeter() float64 }
//         If Rectangle has both Area() and Perimeter() methods, it satisfies
//         both Shape and Perimeter simultaneously with no extra code.
//         This is how Go achieves composable, flexible design.
//
// Q6.  What is "embedding" in Go? How does it differ from inheritance?
//      Example: type ColoredShape struct { Shape; Color string }
//
//      A: Embedding promotes the methods of an inner type to the outer type.
//         type Animal struct { name string }
//         func (a Animal) Speak() string { return a.name }
//         type Dog struct { Animal }       // Dog embeds Animal
//         d := Dog{Animal{"Rex"}}
//         d.Speak()                        // promoted — no redefinition needed
//         Unlike inheritance, there is no is-a relationship. Dog doesn't "extend"
//         Animal — it simply has Animal's methods promoted. You can still override
//         them by defining a method with the same name on Dog.
//
// PRACTICAL
// ---------
// Q7.  What does this print and why?
//
//      type Counter struct{ count int }
//      func (c Counter) Increment() { c.count++ }   // value receiver
//      c := Counter{}
//      c.Increment()
//      c.Increment()
//      fmt.Println(c.count)   // ?
//
//      A: Prints 0.
//         Increment has a VALUE receiver, so it operates on a copy of c.
//         The original c.count is never changed.
//         Fix: change to func (c *Counter) Increment() { c.count++ }
//         Then c.count would be 2.
//
// Q8.  Create a Stringer interface with one method: String() string.
//      Add a String() method to Rectangle so it returns a formatted
//      description like "Rectangle(4x3)".
//      Hint: fmt.Println automatically calls String() if it exists —
//      this is the fmt.Stringer interface from the standard library.
//
//      A: func (r Rectangle) String() string {
//             return fmt.Sprintf("Rectangle(%.0fx%.0f)", r.Width, r.Height)
//         }
//         Now fmt.Println(rect) automatically calls rect.String() and prints:
//         Rectangle(4x3)
//         This works because fmt checks if the value satisfies fmt.Stringer.
//
// Q9.  Write a function totalArea(shapes []Shape) float64 that
//      returns the sum of all shapes' areas. Then call it from main()
//      with a mixed slice of Rectangle, Circle, and Triangle values.
//
//      A: func totalArea(shapes []Shape) float64 {
//             total := 0.0
//             for _, s := range shapes {
//                 total += s.Area()
//             }
//             return total
//         }
//         In main():
//         shapes := []Shape{Rectangle{3,4}, Circle{5}, Triangle{6,4}}
//         fmt.Println(totalArea(shapes))
//
// Q10. What happens if you try to assign a *Rectangle (pointer) to a
//      Shape interface variable, but Rectangle only has value receivers?
//      What if Rectangle has a pointer receiver — does *Rectangle still
//      satisfy Shape? Does Rectangle (non-pointer)?
//
//      A: If Rectangle has only VALUE receivers:
//           Both Rectangle and *Rectangle satisfy Shape.
//           Go automatically dereferences pointers for value-receiver methods.
//         If Rectangle has a POINTER receiver for Area():
//           Only *Rectangle satisfies Shape — NOT Rectangle (value).
//           Because Go cannot take the address of a non-addressable value.
//         Rule: pointer receiver → only pointer satisfies the interface.
//               value receiver  → both value and pointer satisfy the interface.
//
// Q11. Create a Logger interface with Log(message string).
//      Implement it with two structs: ConsoleLogger (prints to stdout)
//      and SilentLogger (does nothing). Write a function
//      process(l Logger) that logs "processing...".
//      Swap between the two implementations in main().
//
//      A: type Logger interface { Log(message string) }
//
//         type ConsoleLogger struct{}
//         func (c ConsoleLogger) Log(msg string) { fmt.Println("[LOG]", msg) }
//
//         type SilentLogger struct{}
//         func (s SilentLogger) Log(msg string) {}   // intentionally empty
//
//         func process(l Logger) { l.Log("processing...") }
//
//         In main():
//           process(ConsoleLogger{})  // prints: [LOG] processing...
//           process(SilentLogger{})   // prints nothing
//         This is the Strategy pattern — swap behaviour without changing process().
// ======================================================================
