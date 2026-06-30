package main

import "fmt"

func main() {

	// Array — fixed size, size is part of the type
	// [3]int and [5]int are two different types
	fmt.Println("--- declaration ---")

	var a [5]int                       // zero value: all elements = 0
	fmt.Println("zero array:", a)

	b := [3]string{"Go", "is", "fun"} // array literal
	fmt.Println("literal:", b)

	c := [...]int{10, 20, 30, 40}     // ... lets compiler count the length
	fmt.Println("auto-length:", c, "len:", len(c))

	// Accessing and modifying elements
	fmt.Println("\n--- access & modify ---")

	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println("original:", nums)

	nums[0] = 100 // index starts at 0
	nums[4] = 999 // last index = len - 1
	fmt.Println("modified:", nums)

	// Iterating with range
	fmt.Println("\n--- range ---")

	grades := [4]int{88, 92, 75, 95}
	sum := 0
	for i, g := range grades {
		fmt.Printf("  student %d: %d\n", i+1, g)
		sum += g
	}
	fmt.Printf("  average: %.1f\n", float64(sum)/float64(len(grades)))

	// Arrays are value types — assignment makes a full copy
	fmt.Println("\n--- arrays are values (copied) ---")

	original := [3]int{1, 2, 3}
	copied := original  // independent copy, not a reference
	copied[0] = 999

	fmt.Println("original:", original) // [1 2 3] — unchanged
	fmt.Println("copied:  ", copied)   // [999 2 3]

	// Array → Slice
	// Slicing an array gives a slice that shares the same memory
	fmt.Println("\n--- array to slice ---")

	arr := [5]int{10, 20, 30, 40, 50}
	sl := arr[1:4] // shares arr's memory
	fmt.Println("arr:", arr)
	fmt.Println("sl: ", sl)

	sl[0] = 999    // modifies arr too!
	fmt.Println("after sl[0]=999 → arr:", arr)

	// EXERCISE
	// 1. Declare [5]float64 of temperatures. Find and print min and max.
	// 2. Declare [5]int, fill it with squares (1,4,9,16,25) using a loop.
	// 3. Copy the array into another and change one value — confirm the
	//    original is unchanged (value semantics).
}
