package main

import "fmt"

// Pointer receiver — Scale modifies the original Rectangle
type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Pointer parameter — double modifies the caller's variable
func double(n *int) {
	*n = *n * 2
}

func main() {

	// & gives the memory address of a variable
	// * reads or writes the value at that address (dereference)
	fmt.Println("--- basics ---")

	number := 10
	p := &number // p holds the address of number

	fmt.Println("number:", number) // 10
	fmt.Println("p (address):", p) // e.g. 0xc000018090
	fmt.Println("*p (value):", *p) // 10

	*p = 99 // write through the pointer
	fmt.Println("after *p = 99, number:", number) // 99

	// Pass by value vs pass by pointer
	fmt.Println("\n--- pass by value vs pointer ---")

	x := 5
	fmt.Println("before double:", x) // 5
	double(&x)                        // pass address so double can modify x
	fmt.Println("after double:", x)  // 10

	// Without a pointer, the original is unchanged:
	noChange := func(n int) { n = 999 }
	y := 5
	noChange(y)
	fmt.Println("unchanged y:", y) // still 5

	// Nil pointer — zero value of any pointer type
	fmt.Println("\n--- nil pointer ---")

	var ptr *int
	fmt.Println("ptr is nil:", ptr == nil) // true

	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("ptr is nil — skip dereference") // safe
	}
	// *ptr here would panic: invalid memory address or nil pointer dereference

	// Pointer to struct — Go auto-dereferences field access
	fmt.Println("\n--- pointer to struct ---")

	rect := Rectangle{Width: 4, Height: 3}
	fmt.Printf("before: area = %.0f\n", rect.Area()) // 12

	rect.Scale(2) // Go automatically passes &rect for pointer receivers
	fmt.Printf("after Scale(2): area = %.0f\n", rect.Area()) // 48

	// Explicit struct pointer
	rp := &Rectangle{Width: 5, Height: 2}
	rp.Scale(3)
	fmt.Printf("rp after Scale(3): area = %.0f\n", rp.Area()) // 30

	// new() — allocates a zeroed value and returns a pointer
	fmt.Println("\n--- new() ---")

	n := new(int) // allocates int = 0, returns *int
	fmt.Println("new(int):", *n) // 0
	*n = 42
	fmt.Println("after *n = 42:", *n) // 42

	// EXERCISE
	// 1. Write swap(a, b *int) that swaps two values.
	//    x, y := 10, 20 → swap(&x, &y) → x=20, y=10
	// 2. Write increment(p *int) that adds 1 to *p.
	//    Call it 5 times on the same variable. Result should be 5.
	// 3. Add a Reset() method (pointer receiver) to Rectangle that sets
	//    Width and Height to 0. Call it and print the area after.
}
