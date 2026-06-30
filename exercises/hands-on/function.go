package main

import (
	"errors"
	"fmt"
)

// Basic function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Multiple parameters of the same type — write the type once
func add(a, b int) int {
	return a + b
}

// Multiple return values — idiomatic Go for returning a result + error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Variadic — accepts any number of int arguments
// Inside the function, nums is treated as []int
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Function as a parameter — functions are first-class values in Go
func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Closure — a function that captures variables from its surrounding scope
// Each call to makeCounter() returns an independent counter
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// defer — runs when the surrounding function returns, in LIFO order
func demoDefer() {
	fmt.Println("start")
	defer fmt.Println("runs third")
	defer fmt.Println("runs second")
	defer fmt.Println("runs first")
	fmt.Println("end")
}

// Recursion
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	// Basic
	fmt.Println("--- basic ---")
	fmt.Println(greet("Gopher"))
	fmt.Println("3 + 4 =", add(3, 4))

	// Multiple return values + error handling
	fmt.Println("\n--- multiple returns ---")
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Variadic
	fmt.Println("\n--- variadic ---")
	fmt.Println("sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("sum(1..5):", sum(1, 2, 3, 4, 5))

	nums := []int{10, 20, 30}
	fmt.Println("sum(slice...):", sum(nums...)) // spread slice with ...

	// Function as a value
	fmt.Println("\n--- function as value ---")
	multiply := func(a, b int) int { return a * b }
	fmt.Println("apply add:     ", apply(4, 5, add))
	fmt.Println("apply multiply:", apply(4, 5, multiply))

	// Anonymous function called immediately
	fmt.Println("square of 7:", func(x int) int { return x * x }(7))

	// Closure
	fmt.Println("\n--- closure ---")
	counter := makeCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	other := makeCounter() // independent — starts from 0
	fmt.Println(other())   // 1

	// defer
	fmt.Println("\n--- defer ---")
	demoDefer()

	// Recursion
	fmt.Println("\n--- recursion ---")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	// EXERCISE
	// 1. Write power(base, exp int) int using a loop (no math.Pow).
	//    power(2, 10) → 1024
	// 2. Write filter(nums []int, fn func(int) bool) []int that returns
	//    only elements where fn returns true.
	//    Use it to extract even numbers from []int{1,2,3,4,5,6}.
	// 3. Write makeAdder(n int) func(int) int — a closure that adds n.
	//    add5 := makeAdder(5); add5(3) → 8
}
