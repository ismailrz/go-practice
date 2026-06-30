package main

import "fmt"

func main() {

	// Slice — dynamic, resizable. Three parts: pointer | length | capacity
	fmt.Println("--- creating slices ---")

	fruits := []string{"apple", "banana", "cherry"} // literal
	fmt.Println("fruits:", fruits, "len:", len(fruits), "cap:", cap(fruits))

	nums := make([]int, 3, 6) // len=3, cap=6, all zeros
	fmt.Println("make:", nums, "len:", len(nums), "cap:", cap(nums))

	var empty []int // nil slice — safe to read and append, not to write index
	fmt.Println("nil:", empty, "== nil:", empty == nil)

	// append — always assign the result back
	// If len == cap, Go allocates a larger array automatically
	fmt.Println("\n--- append ---")

	s := []int{1, 2, 3}
	s = append(s, 4)
	s = append(s, 5, 6)       // multiple at once
	fmt.Println("after append:", s)

	more := []int{7, 8, 9}
	s = append(s, more...)    // spread a slice with ...
	fmt.Println("appended slice:", s)

	// Slicing — s[low:high], shares the same underlying array
	fmt.Println("\n--- slicing ---")

	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println("[1:3]:", letters[1:3]) // [b c]
	fmt.Println("[:2]: ", letters[:2])  // [a b]
	fmt.Println("[3:]: ", letters[3:])  // [d e]

	// Iterating
	fmt.Println("\n--- range ---")

	scores := []int{88, 72, 95, 61, 84}
	total := 0
	for i, score := range scores {
		fmt.Printf("  [%d] %d\n", i, score)
		total += score
	}
	fmt.Printf("  average: %.1f\n", float64(total)/float64(len(scores)))

	// copy — independent copy, no shared memory
	fmt.Println("\n--- copy ---")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)

	dst[0] = 999
	fmt.Println("src:", src) // unchanged
	fmt.Println("dst:", dst) // [999 2 3 4 5]

	// Deleting an element — no built-in, use append
	fmt.Println("\n--- delete element ---")

	d := []string{"a", "b", "c", "d", "e"}
	i := 2                              // delete index 2 ("c")
	d = append(d[:i], d[i+1:]...)
	fmt.Println("after delete:", d)    // [a b d e]

	// EXERCISE
	// 1. Write contains(s []string, target string) bool that returns true
	//    if target exists in s.
	// 2. Write unique(s []int) []int that removes duplicates.
	//    Hint: use a map[int]bool to track seen values.
	// 3. Write rotate(s []int, k int) []int that rotates left by k steps.
	//    rotate([]int{1,2,3,4,5}, 2) → [3 4 5 1 2]
}
