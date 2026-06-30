package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Boolean
	fmt.Println("--- bool ---")
	var isActive bool = false
	isLoggedIn := true
	fmt.Println(isActive, isLoggedIn)
	fmt.Println("AND:", isActive && isLoggedIn)
	fmt.Println("OR: ", isActive || isLoggedIn)
	fmt.Println("NOT:", !isLoggedIn)

	// Integers
	fmt.Println("\n--- int ---")
	var age int = 25
	var score int64 = 1000000
	count := 10 // type inferred as int

	fmt.Println("age:", age)
	fmt.Println("score:", score)
	fmt.Println("count:", count)

	// Arithmetic
	a, b := 17, 5
	fmt.Println("17 + 5 =", a+b)
	fmt.Println("17 - 5 =", a-b)
	fmt.Println("17 * 5 =", a*b)
	fmt.Println("17 / 5 =", a/b) // integer division — truncates
	fmt.Println("17 % 5 =", a%b) // remainder

	// Float
	fmt.Println("\n--- float64 ---")
	pi := 3.14159
	radius := 5.0
	area := pi * radius * radius
	fmt.Printf("area = %.2f\n", area)

	// String
	fmt.Println("\n--- string ---")
	name := "Go"
	greeting := "Hello, " + name + "!"
	fmt.Println(greeting)
	fmt.Println("length:", len(greeting))

	// Raw string — backticks, no escape sequences
	raw := `Line 1
Line 2
Line 3`
	fmt.Println(raw)

	// byte and rune
	fmt.Println("\n--- byte & rune ---")
	var ch byte = 'A'   // byte = uint8, for ASCII
	var r rune = '😀'   // rune = int32, for Unicode
	fmt.Printf("byte: %c (%d)\n", ch, ch)
	fmt.Printf("rune: %c (%d)\n", r, r)

	// Type conversion — Go requires explicit conversion, no automatic coercion
	fmt.Println("\n--- type conversion ---")
	var x int = 42
	var y float64 = float64(x)  // int → float64
	var z int = int(y)          // float64 → int (truncates)
	fmt.Println("int → float64:", y)
	fmt.Println("float64 → int:", z)

	// int ↔ string needs strconv, NOT a cast
	// string(65) gives "A" (treats 65 as a rune), NOT "65"
	numStr := strconv.Itoa(99)       // int → "99"
	numVal, _ := strconv.Atoi("42") // "42" → 42
	fmt.Println("Itoa(99):", numStr)
	fmt.Println("Atoi(\"42\"):", numVal)

	// Zero values — every type has one, no null/undefined surprises
	fmt.Println("\n--- zero values ---")
	var zeroBool bool
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	fmt.Printf("bool=%v  int=%v  float64=%v  string=%q\n",
		zeroBool, zeroInt, zeroFloat, zeroString)

	// Constants
	fmt.Println("\n--- constants ---")
	const MaxRetries = 3
	const AppName = "GoApp"
	fmt.Println(AppName, "max retries:", MaxRetries)

	// iota — auto-incrementing constant for enumerations
	const (
		Sunday = iota // 0
		Monday        // 1
		Tuesday       // 2
	)
	fmt.Println("Sunday:", Sunday, "Monday:", Monday, "Tuesday:", Tuesday)

	// EXERCISE
	// 1. Declare a variable of each type (bool, int, float64, string) and
	//    print each with its type using fmt.Printf("%T %v\n", v, v)
	// 2. Convert temperature: write celsius to fahrenheit
	//    F = (C * 9 / 5) + 32 — test with 0°C → 32°F, 100°C → 212°F
	// 3. Use strconv.Atoi to parse "123" and add it to an int variable.
}
