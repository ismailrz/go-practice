// Package main — executable program.
package main

// "fmt"     — formatted output
// "strings" — string manipulation: Contains, Split, ToUpper, TrimSpace, etc.
// "unicode" — Unicode character classification: IsLetter, IsDigit, etc.
import (
	"fmt"
	"strings"
	"unicode"
)

// ======================================================================
// STRINGS & RUNES IN GO
// ======================================================================
// Go strings are IMMUTABLE sequences of bytes (uint8 values).
// They are always UTF-8 encoded — not arrays of characters.
//
// This matters because:
//   - ASCII characters (a-z, 0-9) are 1 byte each.
//   - Non-ASCII characters (é, 中, 😀) can be 2-4 bytes each.
//   - s[i] gives you a BYTE (uint8), not a character.
//   - range s gives you RUNES (Unicode code points), not bytes.
//
// rune is an alias for int32 — it represents a single Unicode code point.
// Rune literals use single quotes: 'A', '中', '😀'

func main() {
	// ======================================================================
	// PART 1 — String basics
	// ======================================================================

	s := "Hello, Go!"

	// len() on a string returns the number of BYTES, not characters.
	fmt.Println("string:", s)
	fmt.Println("len (bytes):", len(s)) // 10 — all ASCII here, so bytes == chars

	// s[i] gives the byte at index i — type uint8 (same as byte).
	// This is safe for ASCII but gives wrong results for multi-byte characters.
	fmt.Printf("s[0] = %d (%c)\n", s[0], s[0]) // 72 (H) — byte value and its ASCII char

	fmt.Println()

	// ======================================================================
	// PART 2 — Runes and multi-byte characters
	// ======================================================================
	// "café" has 4 visible characters but 5 bytes — "é" is 2 bytes in UTF-8.

	word := "café"
	fmt.Println("word:", word)
	fmt.Println("len (bytes):", len(word))         // 5 — NOT 4!
	fmt.Println("rune count:", len([]rune(word)))   // 4 — correct character count

	fmt.Println()

	// Iterating with a regular index loop gives BYTES — breaks on multi-byte chars.
	fmt.Println("--- byte loop (wrong for non-ASCII) ---")
	for i := 0; i < len(word); i++ {
		fmt.Printf("  byte[%d] = %x\n", i, word[i]) // raw byte values
	}

	fmt.Println()

	// Iterating with range gives RUNES — correct for Unicode.
	// range automatically decodes UTF-8 and gives the rune + byte position.
	fmt.Println("--- range loop (correct for Unicode) ---")
	for byteIndex, r := range word {
		// byteIndex: byte position in the string (not rune index — can skip for multi-byte)
		// r: the full rune (Unicode code point), type rune (int32)
		fmt.Printf("  byte[%d]: %c  (U+%04X)\n", byteIndex, r, r)
	}

	fmt.Println()

	// ======================================================================
	// PART 3 — Converting between string, []byte, and []rune
	// ======================================================================

	original := "Hello"

	// string → []byte: gives you the raw UTF-8 bytes. Use when you need to
	// modify individual bytes (strings are immutable, byte slices are not).
	bytes := []byte(original)
	bytes[0] = 'h'                         // lowercase the first byte
	modified := string(bytes)              // convert back to string
	fmt.Println("modified:", modified)     // hello

	// string → []rune: gives you Unicode code points. Use when you need to
	// index by character position (not byte position).
	runes := []rune("café")
	fmt.Println("runes[3]:", string(runes[3])) // é — correct 4th character

	fmt.Println()

	// ======================================================================
	// PART 4 — String immutability
	// ======================================================================
	// Strings in Go are immutable — you cannot change a character in-place.
	// s[0] = 'h'  // compile error: cannot assign to s[0]
	// To "modify" a string, convert to []byte or []rune, modify, then convert back.

	greeting := "Hello"
	// greeting[0] = 'h'  // compile error
	r := []rune(greeting)
	r[0] = 'h'
	greeting = string(r)
	fmt.Println("greeting:", greeting) // hello

	fmt.Println()

	// ======================================================================
	// PART 5 — strings package (standard library)
	// ======================================================================
	// Go's strings package provides all common string operations.
	// No need for third-party libraries for basic string manipulation.

	text := "  Hello, Go World!  "

	fmt.Println("Contains 'Go':", strings.Contains(text, "Go"))          // true
	fmt.Println("HasPrefix '  Hello':", strings.HasPrefix(text, "  Hello")) // true
	fmt.Println("Count 'o':", strings.Count(text, "o"))                  // 3
	fmt.Println("ToUpper:", strings.ToUpper(text))
	fmt.Println("ToLower:", strings.ToLower(text))
	fmt.Println("TrimSpace:", strings.TrimSpace(text))                   // removes leading/trailing spaces
	fmt.Println("Replace:", strings.Replace(text, "Go", "Gopher", 1))   // replace first occurrence
	fmt.Println("ReplaceAll:", strings.ReplaceAll(text, "o", "0"))

	// Split returns a []string
	csv := "apple,banana,cherry"
	parts := strings.Split(csv, ",")
	fmt.Println("Split:", parts)          // [apple banana cherry]
	fmt.Println("Join:", strings.Join(parts, " | ")) // apple | banana | cherry

	fmt.Println()

	// ======================================================================
	// PART 6 — String building (efficient concatenation)
	// ======================================================================
	// Using + to concatenate in a loop creates a new string each iteration — O(n²).
	// strings.Builder is the efficient way: it writes to a buffer and builds once.

	var builder strings.Builder
	words := []string{"Go", "is", "fast", "and", "simple"}

	for i, w := range words {
		if i > 0 {
			builder.WriteString(" ") // add space between words
		}
		builder.WriteString(w) // append to internal buffer — no allocation
	}

	result := builder.String() // single allocation: convert buffer to string
	fmt.Println("built string:", result) // Go is fast and simple

	fmt.Println()

	// ======================================================================
	// PART 7 — rune utilities (unicode package)
	// ======================================================================
	// The unicode package provides functions to classify runes.

	testChars := []rune{'A', 'z', '5', ' ', '!', 'é'}
	for _, ch := range testChars {
		fmt.Printf("  '%c': letter=%v  digit=%v  space=%v  upper=%v\n",
			ch,
			unicode.IsLetter(ch),
			unicode.IsDigit(ch),
			unicode.IsSpace(ch),
			unicode.IsUpper(ch),
		)
	}

	// ======================================================================
	// EXERCISE FOR YOU
	// 1. Write a function: isPalindrome(s string) bool
	//    that returns true if s reads the same forwards and backwards.
	//    Use []rune so it works correctly with Unicode (e.g. "racecar", "niveau").
	// 2. Write a function: wordCount(s string) map[string]int
	//    that splits a sentence into words and counts how many times each appears.
	//    Use strings.Fields (splits on any whitespace) and a map.
	// ======================================================================
}

// ======================================================================
// QUESTIONS — Strings & Runes
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  What is the difference between a byte and a rune in Go?
//
//      A: byte   = uint8 — a single raw byte (0-255). Go strings are []byte underneath.
//         rune   = int32 — a Unicode code point (can represent any character globally).
//         For ASCII text, they're equivalent (one character = one byte).
//         For Unicode text (é, 中, 😀), one rune can be 2-4 bytes.
//         Use rune when you care about characters; byte when you care about raw data.
//
// Q2.  Why does len("café") return 5 instead of 4?
//
//      A: len() counts BYTES, not characters. "café" is encoded in UTF-8:
//           c=1 byte, a=1 byte, f=1 byte, é=2 bytes → total 5 bytes.
//         To get the character count, convert to []rune: len([]rune("café")) == 4.
//
// Q3.  Why should you use "for i, r := range s" instead of "for i := 0; i < len(s); i++"
//      when iterating over a string?
//
//      A: s[i] gives you the byte at position i — for multi-byte Unicode characters,
//         you get a partial byte, not the intended character.
//         range decodes the UTF-8 automatically and gives you the full rune at each
//         CHARACTER boundary. The byte index still advances by the byte width of each rune.
//
// Q4.  Why are Go strings immutable? How do you "modify" a string?
//
//      A: Immutability allows strings to be safely shared between goroutines and used
//         as map keys without copying. The compiler can also optimise string handling.
//         To modify: convert to []byte (for byte manipulation) or []rune
//         (for character manipulation), make changes, then convert back with string().
//
// Q5.  What is strings.Builder and why is it better than "+" for
//      concatenating many strings in a loop?
//
//      A: Each "a + b" allocates a new string (Go strings are immutable — you can't
//         append in-place). In a loop of N concatenations, that's O(N²) allocations.
//         strings.Builder maintains a []byte buffer internally. WriteString() appends
//         to the buffer cheaply. String() does ONE final allocation.
//         Result: O(N) total allocations instead of O(N²).
//
// PRACTICAL
// ---------
// Q6.  What is the output of this code?
//
//      s := "Go"
//      fmt.Println(len(s))
//      fmt.Println(s[0])
//      fmt.Printf("%c\n", s[0])
//
//      A: 2       — len counts bytes, "Go" has 2 ASCII bytes
//         71      — s[0] is a byte (uint8), 'G' has ASCII value 71
//         G       — %c formats the byte as its character representation
//
// Q7.  How do you check if a string contains only ASCII characters?
//
//      A: func isASCII(s string) bool {
//             for i := 0; i < len(s); i++ {
//                 if s[i] > 127 { return false }
//             }
//             return true
//         }
//         Or: len(s) == len([]rune(s))
//         (true when every character is 1 byte, i.e., all ASCII)
//
// Q8.  Write a function that reverses a string correctly for Unicode input.
//      (Reversing bytes gives wrong results for multi-byte characters.)
//
//      A: func reverse(s string) string {
//             runes := []rune(s)
//             for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//                 runes[i], runes[j] = runes[j], runes[i]
//             }
//             return string(runes)
//         }
//         reverse("café") → "éfac"  (correct Unicode reversal)
//
// Q9.  What is the difference between strings.Split(s, "") and []rune(s)?
//
//      A: strings.Split(s, "") splits by UTF-8 characters → []string of 1-char strings.
//         []rune(s) converts to a slice of int32 Unicode code points.
//         For most purposes, []rune is more useful (you can index, modify, reverse).
//         strings.Split is useful when you need []string for further string processing.
// ======================================================================
