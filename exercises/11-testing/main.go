// Package main — executable program.
// This file contains the functions we want to test.
// The tests themselves live in main_test.go (same package, same directory).
//
// Go's testing framework is BUILT INTO the language — no third-party library needed.
// Run all tests in this directory with:   go test .
// Run with verbose output:               go test -v .
// Run a specific test:                   go test -run TestAdd .
// Run with coverage report:              go test -cover .
package main

import (
	"fmt"
	"strings"
)

// ======================================================================
// FUNCTIONS UNDER TEST
// ======================================================================
// These are the functions we will write tests for in main_test.go.
// They are deliberately simple so the focus stays on TESTING patterns.

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Divide returns a/b and an error if b is zero.
// Returns (0, error) on division by zero — the (value, error) pattern from 03-functions.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// IsPalindrome reports whether s reads the same forwards and backwards.
// Uses runes so it handles Unicode correctly.
func IsPalindrome(s string) bool {
	// Normalise: lowercase and remove spaces so "Race Car" == "racecar".
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// WordCount returns a map of word → frequency for the words in s.
func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	// strings.Fields splits on any whitespace and handles multiple spaces.
	for _, word := range strings.Fields(s) {
		counts[strings.ToLower(word)]++
	}
	return counts
}

// ======================================================================
// main — demo of the functions (tests are in main_test.go)
// ======================================================================
func main() {
	fmt.Println("Add(3, 4) =", Add(3, 4))

	result, err := Divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Divide(10, 3) = %.4f\n", result)
	}

	_, err = Divide(5, 0)
	fmt.Println("Divide(5, 0) error:", err)

	fmt.Println("IsPalindrome(racecar):", IsPalindrome("racecar"))
	fmt.Println("IsPalindrome(hello):", IsPalindrome("hello"))
	fmt.Println("IsPalindrome(Race Car):", IsPalindrome("Race Car"))

	fmt.Println("WordCount:", WordCount("go is go and go is great"))

	// ======================================================================
	// Now open main_test.go to see how we test all of the above.
	// Run the tests with:  go test -v .
	// ======================================================================
}
