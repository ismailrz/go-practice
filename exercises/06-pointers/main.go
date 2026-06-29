// Package main — executable program.
package main

import "fmt"

// ======================================================================
// POINTERS IN GO
// ======================================================================
// A pointer holds the MEMORY ADDRESS of a value, not the value itself.
// Go is a pass-by-value language — every function argument is a COPY.
// Pointers let you pass a reference so a function can modify the original.
//
// Two operators:
//   & (address-of)   — gives you a pointer to a variable:  p := &x
//   * (dereference)  — reads or writes through a pointer:  *p = 42

// ======================================================================
// WITHOUT POINTERS — function gets a copy, original is unchanged
// ======================================================================

// doubleByValue receives a COPY of n.
// Any changes made here are invisible to the caller.
func doubleByValue(n int) {
	n = n * 2 // modifies the local copy only
	fmt.Println("inside doubleByValue, n =", n) // 20 (local copy)
}

// ======================================================================
// WITH POINTERS — function gets an address, can modify the original
// ======================================================================

// doubleByPointer receives a POINTER to an int (*int).
// Dereferencing with * lets us read and write the original value.
func doubleByPointer(p *int) {
	// *p reads the value at the memory address p points to.
	// *p = ... writes a new value to that address — modifying the original.
	*p = *p * 2
	fmt.Println("inside doubleByPointer, *p =", *p) // 20 (original modified)
}

// ======================================================================
// POINTER TO STRUCT
// ======================================================================
// When you have a pointer to a struct, Go lets you access fields directly
// without writing (*ptr).Field — it auto-dereferences for you.

type Point struct {
	X, Y int
}

// moveRight modifies the Point the pointer points to.
// Without a pointer receiver or pointer parameter, the struct would be copied.
func moveRight(p *Point, dx int) {
	p.X += dx // Go auto-dereferences: same as (*p).X += dx
}

// ======================================================================
// new() — another way to create a pointer
// ======================================================================
// new(T) allocates memory for a value of type T, sets it to its zero value,
// and returns a *T (pointer to T).
// You rarely need new() — most code uses &T{} literals or var declarations.

func main() {
	// ======================================================================
	// DEMO 1 — pass by value (copy)
	// ======================================================================
	x := 10
	fmt.Println("before doubleByValue, x =", x) // 10
	doubleByValue(x)
	fmt.Println("after  doubleByValue, x =", x) // still 10 — copy was modified, not x

	fmt.Println()

	// ======================================================================
	// DEMO 2 — pass by pointer (reference)
	// ======================================================================
	y := 10
	fmt.Println("before doubleByPointer, y =", y) // 10

	// &y gives us the memory address of y (a *int).
	// We pass that address to the function.
	doubleByPointer(&y)

	fmt.Println("after  doubleByPointer, y =", y) // 20 — original was modified

	fmt.Println()

	// ======================================================================
	// DEMO 3 — declaring and using pointers directly
	// ======================================================================
	a := 42
	p := &a // p is of type *int — it holds the address of a

	fmt.Println("a =", a)   // 42
	fmt.Println("p =", p)   // e.g. 0xc0000b4010 — the memory address
	fmt.Println("*p =", *p) // 42 — the value AT that address (dereferencing)

	// Modify a through the pointer.
	*p = 100
	fmt.Println("after *p = 100, a =", a) // 100 — a was changed via the pointer

	fmt.Println()

	// ======================================================================
	// DEMO 4 — nil pointer
	// ======================================================================
	// A pointer's zero value is nil — it points to nothing.
	// Dereferencing a nil pointer causes a panic.
	var ptr *int // ptr is nil
	fmt.Println("ptr is nil:", ptr == nil) // true

	// Uncomment to see the panic:
	// fmt.Println(*ptr)  // panic: runtime error: invalid memory address or nil pointer dereference

	// Always check for nil before dereferencing a pointer you didn't create yourself.
	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("ptr is nil — skipping dereference")
	}

	fmt.Println()

	// ======================================================================
	// DEMO 5 — pointer to struct
	// ======================================================================
	pt := Point{X: 3, Y: 7}
	fmt.Println("before moveRight:", pt) // {3 7}

	moveRight(&pt, 5) // pass address of pt

	fmt.Println("after  moveRight:", pt) // {8 7} — X was modified

	fmt.Println()

	// ======================================================================
	// DEMO 6 — new()
	// ======================================================================
	// new(int) allocates an int, sets it to 0, returns *int.
	counter := new(int)
	fmt.Println("new(int) zero value:", *counter) // 0
	*counter = 5
	fmt.Println("after *counter = 5:", *counter) // 5

	// The common idiom for struct pointers uses &T{} instead of new(T):
	ptNew := &Point{X: 1, Y: 2} // equivalent to new(Point) then setting fields
	fmt.Println("&Point{} pointer:", ptNew)  // &{1 2}
	fmt.Println("dereferenced:", *ptNew)     // {1 2}

	// ======================================================================
	// EXERCISE FOR YOU
	// 1. Write a function: swap(a, b *int) that swaps the values at the
	//    two pointers. Call it from main() and verify both variables changed.
	// 2. Write a function: increment(p *int) that adds 1 to *p.
	//    Call it 3 times on the same variable and print the result (should be 3).
	// ======================================================================
}

// ======================================================================
// QUESTIONS — Pointers
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  Go is "pass by value". What does this mean?
//      When does pass-by-value cause problems, and how do pointers solve it?
//
//      A: Every function argument in Go is a COPY of the original value.
//         Changes inside the function do not affect the caller's variable.
//         Problem: if you pass a large struct to 10 functions, you copy it 10 times
//         (wasted memory), and you can't modify the original.
//         Pointers solve both: pass &myStruct (just an 8-byte address),
//         and the function can modify the original through *p.
//
// Q2.  What is the zero value of a pointer? What happens if you dereference it?
//
//      A: The zero value of any pointer type is nil — it points to no address.
//         Dereferencing nil (*p) causes a runtime panic:
//         "invalid memory address or nil pointer dereference"
//         Always guard with: if p != nil { ... } before dereferencing
//         a pointer you didn't create yourself.
//
// Q3.  What is the difference between *T and &T in Go?
//
//      A: *T is a TYPE — "pointer to T". e.g., var p *int declares p as a pointer to int.
//         & is an OPERATOR — "address of". e.g., p := &x gives p the address of x.
//         * is also an OPERATOR when used on a pointer value — "dereference".
//         e.g., *p reads or writes the value at the address p holds.
//
// Q4.  When you access a field on a pointer to a struct (p.Field), does Go
//      require you to write (*p).Field? Why or why not?
//
//      A: No. Go auto-dereferences pointer-to-struct for field access.
//         p.Field is syntactic sugar for (*p).Field.
//         Same applies for calling methods: p.Method() works even if Method
//         is defined on the value type — Go takes the address automatically.
//
// Q5.  What is the difference between new(T) and &T{}?
//
//      A: Both allocate a zeroed T on the heap and return a *T.
//         new(T) gives you a pointer to a zeroed T — you must set fields separately.
//         &T{field: val} allocates and initialises fields in one step (more common).
//         In practice, &T{} is preferred for structs; new() is occasionally used
//         for primitives like new(int) or new(sync.Mutex).
//
// PRACTICAL
// ---------
// Q6.  What is the output of this code?
//
//      func add(a, b int) int { return a + b }
//      x, y := 3, 4
//      result := add(x, y)
//      fmt.Println(x, y, result)
//
//      A: 3 4 7
//         x and y are unchanged — add() received copies, not pointers.
//
// Q7.  Fix this function so it actually zeroes the caller's variable:
//
//      func zero(x int) { x = 0 }
//
//      A: func zero(x *int) { *x = 0 }
//         Call with: zero(&myVar)
//
// Q8.  Why do methods like (r *Rectangle) Scale() use pointer receivers
//      while (r Rectangle) Area() uses a value receiver?
//      Connect this to what you learned about pointers.
//
//      A: Scale() needs to modify the Rectangle's fields (Width, Height).
//         With a value receiver, Scale() gets a COPY — changes are lost.
//         With a pointer receiver (*Rectangle), Scale() gets the address of the
//         original struct and can modify it through *r.
//         Area() only reads — no modification needed — so a value receiver
//         (cheap copy) is fine and signals "this method is read-only".
//
// Q9.  What does this print and why?
//
//      a := 1
//      b := &a
//      c := &a
//      *b = 10
//      fmt.Println(*c)
//
//      A: 10
//         Both b and c point to the SAME variable a.
//         Modifying *b modifies a, so *c (which also reads a) sees the change.
//         This is pointer aliasing — two pointers referring to the same memory.
// ======================================================================
