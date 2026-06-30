package main

import "fmt"

func main() {

	// Classic for loop: init; condition; post
	fmt.Println("--- classic for ---")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While-style: condition only (Go has no "while" keyword)
	fmt.Println("\n--- while-style ---")
	n := 1
	for n <= 32 {
		fmt.Println(n)
		n *= 2
	}

	// Infinite loop with break
	fmt.Println("\n--- infinite + break ---")
	count := 0
	for {
		count++
		if count == 4 {
			break
		}
		fmt.Println("count:", count)
	}

	// continue — skip to next iteration
	fmt.Println("\n--- continue (skip even) ---")
	for i := 1; i <= 8; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// range over slice
	fmt.Println("\n--- range slice ---")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("[%d] %s\n", i, fruit)
	}

	// range over map (order is random each run)
	fmt.Println("\n--- range map ---")
	scores := map[string]int{"Alice": 90, "Bob": 85, "Carol": 92}
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// range over string — gives runes (Unicode code points), not bytes
	fmt.Println("\n--- range string ---")
	for i, ch := range "Go!" {
		fmt.Printf("index %d: %c\n", i, ch)
	}

	// Nested loops
	fmt.Println("\n--- nested loops (multiplication table) ---")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%-3d", i, j, i*j)
		}
		fmt.Println()
	}

	// FizzBuzz
	fmt.Println("\n--- FizzBuzz 1-15 ---")
	for i := 1; i <= 15; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// EXERCISE
	// 1. Print the sum of all numbers 1..100 using a for loop.
	// 2. Use range to find the largest number in []int{3,7,1,9,4,6}.
	// 3. Print a 5-row triangle of stars using nested loops:
	//    *
	//    **
	//    ***
	//    ****
	//    *****
}
