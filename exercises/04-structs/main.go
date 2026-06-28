// Package main — executable program.
package main

import "fmt"

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

// Scale uses a POINTER receiver because it needs to MODIFY the original Rectangle.
// The * before Rectangle means "pointer to Rectangle".
// If you used a value receiver here, Scale() would modify a copy and the original would be unchanged.
func (r *Rectangle) Scale(factor float64) {
	// r is a pointer — r.Width dereferences it automatically (no need for (*r).Width).
	r.Width *= factor  // same as r.Width = r.Width * factor
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
	// INTERFACE IN ACTION — printArea accepts ANY Shape
	// ======================================================================

	// Passing a Rectangle — printArea doesn't know or care it's a Rectangle.
	printArea(Rectangle{Width: 4, Height: 3})  // Area: 12.00

	// Passing a Circle — same function, different type, correct behavior.
	printArea(Circle{Radius: 5})               // Area: 78.54

	// Passing a Triangle — added AFTER printArea was written, no changes needed.
	printArea(Triangle{Base: 6, Height: 4})    // Area: 12.00

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
