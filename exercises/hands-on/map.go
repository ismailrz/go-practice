package main

import "fmt"

func main() {

	// Map — unordered key-value pairs
	// Zero value is nil — use make() or a literal before writing
	fmt.Println("--- creating maps ---")

	// Literal
	capitals := map[string]string{
		"France":  "Paris",
		"Germany": "Berlin",
		"Japan":   "Tokyo",
	}
	fmt.Println("capitals:", capitals)

	// make() — empty, ready to write
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 82
	fmt.Println("scores:", scores)

	// CRUD
	fmt.Println("\n--- CRUD ---")

	inventory := map[string]int{"apple": 10, "banana": 5}

	inventory["mango"] = 8     // create
	inventory["apple"] = 15   // update
	fmt.Println("after create/update:", inventory)

	fmt.Println("banana:", inventory["banana"]) // read

	delete(inventory, "banana") // delete
	fmt.Println("after delete:", inventory)

	// Comma-ok — distinguish "missing key" from "value is zero"
	// Reading a missing key returns the zero value, not an error
	fmt.Println("\n--- comma-ok ---")

	val, ok := inventory["banana"] // was deleted
	fmt.Printf("banana → val:%d  ok:%v\n", val, ok) // val:0, ok:false

	val, ok = inventory["apple"]
	fmt.Printf("apple  → val:%d  ok:%v\n", val, ok) // val:15, ok:true

	if count, exists := inventory["mango"]; exists {
		fmt.Println("mango count:", count)
	} else {
		fmt.Println("mango not found")
	}

	// Iterating — order is random every run (intentional in Go)
	fmt.Println("\n--- range ---")

	ages := map[string]int{"Alice": 30, "Bob": 25, "Carol": 35}
	for name, age := range ages {
		fmt.Printf("  %s: %d\n", name, age)
	}

	// Keys only
	for name := range ages {
		fmt.Println(" ", name)
	}

	// Map of slices — values can be any type
	fmt.Println("\n--- map of slices ---")

	hobbies := map[string][]string{
		"Alice": {"reading", "cycling"},
		"Bob":   {"gaming", "cooking"},
	}
	hobbies["Alice"] = append(hobbies["Alice"], "chess")

	for person, list := range hobbies {
		fmt.Printf("  %s: %v\n", person, list)
	}

	// Word frequency — common pattern: increment with zero-value default
	fmt.Println("\n--- word count ---")

	words := []string{"go", "is", "fast", "go", "is", "go"}
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++ // 0+1 on first occurrence, no init needed
	}
	fmt.Println("frequency:", freq)

	// EXERCISE
	// 1. Write groupByLength(words []string) map[int][]string that groups
	//    words by their character length.
	//    ["go","is","fun","cool"] → {2:["go","is"], 3:["fun"], 4:["cool"]}
	// 2. Write invertMap(m map[string]int) map[int]string that swaps
	//    keys and values. {"a":1, "b":2} → {1:"a", 2:"b"}
	// 3. Given a sentence, find the most frequently occurring word.
}
