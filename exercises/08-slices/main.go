// Package main — executable program.
package main

import "fmt"

// ======================================================================
// SLICES — Go's most important data structure
// ======================================================================
// A slice is a dynamic, resizable view over an underlying array.
// It has THREE components:
//   - Pointer   → points to the first element of the underlying array
//   - Length    → number of elements currently in the slice
//   - Capacity  → total space in the underlying array from this pointer onward
//
// Think of it like a Python list, but with explicit visibility into its internals.
// Arrays in Go are fixed-size and rarely used directly — slices are the standard.

func main() {
	// ======================================================================
	// PART 1 — Creating slices
	// ======================================================================

	// Slice literal — like an array literal but without a size.
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("fruits:", fruits)        // [apple banana cherry]
	fmt.Println("len:", len(fruits))      // 3 — number of elements
	fmt.Println("cap:", cap(fruits))      // 3 — capacity of underlying array

	fmt.Println()

	// make([]T, length, capacity) — creates a slice with a given length and capacity.
	// All elements are initialised to the zero value of T.
	// Capacity is optional — if omitted, capacity = length.
	nums := make([]int, 3, 5) // length=3, capacity=5
	fmt.Println("nums:", nums)       // [0 0 0]
	fmt.Println("len:", len(nums))   // 3
	fmt.Println("cap:", cap(nums))   // 5

	fmt.Println()

	// ======================================================================
	// PART 2 — append
	// ======================================================================
	// append(slice, elem1, elem2, ...) adds elements to the END of a slice.
	// IMPORTANT: append may return a NEW slice if the underlying array is full.
	// You MUST assign the result back: s = append(s, value)
	// Never write: append(s, value) without capturing the return value — it's lost.

	s := []int{1, 2, 3}
	fmt.Println("before append:", s, "len:", len(s), "cap:", cap(s)) // len:3 cap:3

	// When len == cap, append allocates a new, larger array (usually 2x capacity).
	s = append(s, 4, 5)
	fmt.Println("after append:", s, "len:", len(s), "cap:", cap(s)) // len:5 cap:6 (Go doubled it)

	// Appending one slice to another using the ... spread operator.
	extra := []int{6, 7, 8}
	s = append(s, extra...)  // ... unpacks the slice into individual arguments
	fmt.Println("after append slice:", s) // [1 2 3 4 5 6 7 8]

	fmt.Println()

	// ======================================================================
	// PART 3 — Slicing expressions (sub-slices)
	// ======================================================================
	// s[low:high] creates a new slice header pointing into the SAME array.
	// It includes elements from index low UP TO (not including) high.
	// low defaults to 0, high defaults to len(s).

	letters := []string{"a", "b", "c", "d", "e"}

	fmt.Println("letters[1:3]:", letters[1:3]) // [b c]   — index 1 and 2
	fmt.Println("letters[:2]:", letters[:2])   // [a b]   — from start to index 1
	fmt.Println("letters[3:]:", letters[3:])   // [d e]   — from index 3 to end
	fmt.Println("letters[:]:", letters[:])     // [a b c d e] — full slice (copy of header)

	fmt.Println()

	// ======================================================================
	// PART 4 — Slices share the underlying array (IMPORTANT!)
	// ======================================================================
	// A sub-slice is NOT a copy — it shares memory with the original.
	// Modifying elements through either slice affects BOTH.

	original := []int{10, 20, 30, 40, 50}
	sub := original[1:4] // [20 30 40] — shares original's array

	fmt.Println("original:", original) // [10 20 30 40 50]
	fmt.Println("sub:", sub)           // [20 30 40]

	sub[0] = 999 // modifies the second element of original too!

	fmt.Println("after sub[0] = 999:")
	fmt.Println("  original:", original) // [10 999 30 40 50] — changed!
	fmt.Println("  sub:", sub)           // [999 30 40]

	fmt.Println()

	// ======================================================================
	// PART 5 — copy: making a true independent copy
	// ======================================================================
	// copy(dst, src) copies elements from src into dst.
	// It copies min(len(dst), len(src)) elements.
	// dst and src do NOT share the underlying array after copy.

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src)) // allocate a new slice with the same length
	n := copy(dst, src)
	fmt.Println("copied", n, "elements")
	fmt.Println("src:", src) // [1 2 3 4 5]
	fmt.Println("dst:", dst) // [1 2 3 4 5]

	dst[0] = 999
	fmt.Println("after dst[0] = 999:")
	fmt.Println("  src:", src) // [1 2 3 4 5] — unchanged
	fmt.Println("  dst:", dst) // [999 2 3 4 5] — only dst changed

	fmt.Println()

	// ======================================================================
	// PART 6 — nil slice vs empty slice
	// ======================================================================
	var nilSlice []int          // nil — zero value for a slice
	emptySlice := []int{}       // empty — allocated but has no elements
	madeSlice := make([]int, 0) // also empty

	fmt.Println("nilSlice == nil:", nilSlice == nil)     // true
	fmt.Println("emptySlice == nil:", emptySlice == nil) // false
	fmt.Println("madeSlice == nil:", madeSlice == nil)   // false

	// len() and append() work on nil slices — no panic.
	fmt.Println("len(nilSlice):", len(nilSlice)) // 0
	nilSlice = append(nilSlice, 1, 2)
	fmt.Println("after append:", nilSlice) // [1 2]

	fmt.Println()

	// ======================================================================
	// PART 7 — 2D slices (slice of slices)
	// ======================================================================
	// Go has no built-in matrix type — you build it from slices of slices.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	// Set some positions (like tic-tac-toe).
	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"

	for _, row := range board {
		fmt.Println(row) // [X _ _] then [_ O _] then [_ _ X]
	}

	// ======================================================================
	// EXERCISE FOR YOU
	// 1. Write a function: removeDuplicates(s []int) []int
	//    that returns a new slice with duplicate integers removed.
	//    Hint: use a map[int]bool to track seen values.
	// 2. Write a function: chunk(s []int, size int) [][]int
	//    that splits a slice into sub-slices of the given size.
	//    e.g. chunk([]int{1,2,3,4,5}, 2) → [[1 2] [3 4] [5]]
	// ======================================================================
}

// ======================================================================
// QUESTIONS — Slices
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  What are the three components of a slice header?
//      What is the difference between length and capacity?
//
//      A: A slice header has:
//           Pointer  — address of the first accessible element in the array
//           Length   — number of elements accessible through this slice (len)
//           Capacity — total elements from the pointer to the end of the array (cap)
//         Length: how many elements you can currently read/write.
//         Capacity: how many you CAN hold before a new array must be allocated.
//         When you append beyond cap, Go allocates a new larger array and copies.
//
// Q2.  Why must you always write "s = append(s, value)" and not just
//      "append(s, value)"?
//
//      A: append() may return a DIFFERENT slice (with a new underlying array)
//         if the original's capacity was exceeded. If you discard the return
//         value, you lose the result entirely — the original s is unchanged.
//         Go will warn you if you call append without using the result,
//         but only as a vet warning, not a compile error.
//
// Q3.  What does the "..." operator do when used with append?
//      e.g.: s = append(s, other...)
//
//      A: "..." unpacks a slice into individual arguments.
//         append(s, other...) is equivalent to appending each element of "other"
//         one by one. Without "...", you'd be appending a []T as a single element
//         to a [][]T (which is a type error for []T).
//
// Q4.  When a sub-slice shares the underlying array with its parent,
//      what unexpected behaviour can occur?
//
//      A: Modifying elements through either slice affects the shared array,
//         so the other slice sees the change. Also, appending to the sub-slice
//         within its capacity overwrites elements in the parent's backing array —
//         a subtle and dangerous bug. Use copy() to get an independent slice
//         when you need to avoid this.
//
// Q5.  What is the difference between a nil slice and an empty slice?
//      Do they behave differently with len(), append(), and range?
//
//      A: nil slice: var s []int       → s == nil is true, len(s) == 0
//         empty slice: s := []int{}   → s == nil is false, len(s) == 0
//         For len(), append(), and range: they behave IDENTICALLY.
//         The only difference is that nil == nil (useful for "was it set?" checks).
//         Prefer returning nil slices from functions (not []T{}) to signal "no data".
//
// PRACTICAL
// ---------
// Q6.  What is the output of this code?
//
//      a := []int{1, 2, 3}
//      b := a
//      b[0] = 99
//      fmt.Println(a[0])
//
//      A: 99
//         Slice assignment (b := a) copies the slice HEADER (pointer, len, cap)
//         — NOT the underlying array. Both a and b point to the same array.
//         Changing b[0] changes the shared array, so a[0] also becomes 99.
//         To get an independent copy: b := make([]int, len(a)); copy(b, a)
//
// Q7.  You have a slice s with len=3, cap=6. How many appends can you do
//      before Go allocates a new underlying array?
//
//      A: 3 more appends (to reach len=6 = cap=6).
//         The 4th append (making len=7 > cap=6) triggers a new allocation.
//         Go typically doubles capacity: new cap would be ~12.
//
// Q8.  Write a one-liner using append to insert value v at index i in slice s.
//
//      A: s = append(s[:i], append([]int{v}, s[i:]...)...)
//         Simpler (and safer — avoids double append aliasing):
//           s = append(s, 0)          // grow by one
//           copy(s[i+1:], s[i:])      // shift right
//           s[i] = v                  // insert
//
// Q9.  What does make([]int, 5, 10) produce? How is it different from
//      make([]int, 5)?
//
//      A: make([]int, 5, 10) → len=5, cap=10. The 5 elements are 0.
//         You can append 5 more times before reallocation.
//         make([]int, 5)     → len=5, cap=5. Equivalent to make([]int, 5, 5).
//         Use make with extra capacity when you know you'll append many times
//         — it avoids repeated reallocations (like ArrayList.ensureCapacity in Java).
// ======================================================================
