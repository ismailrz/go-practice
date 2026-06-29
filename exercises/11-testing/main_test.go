// The test file is in the SAME package as the code it tests.
// Go convention: name it <file>_test.go — the build tool excludes it from
// normal builds and includes it only when running "go test".
package main

// "testing" is the ONLY import needed for basic tests — no third-party framework.
// "math" is used here for math.Abs in the float comparison helper.
import (
	"math"
	"testing"
)

// ======================================================================
// BASIC TEST FUNCTION
// ======================================================================
// Every test function:
//   1. Must start with "Test" (capital T).
//   2. Must take exactly one argument: t *testing.T
//   3. Lives in a _test.go file.
//
// t.Errorf("message") — marks the test as FAILED and logs a message,
//                        but continues running the rest of the test.
// t.Fatalf("message") — marks as failed AND stops the test immediately.
//                        Use when subsequent code would panic if this check fails.

func TestAdd(t *testing.T) {
	// Basic call: Add(3, 4) should return 7.
	got := Add(3, 4)
	want := 7

	// Convention: name variables "got" (actual result) and "want" (expected result).
	if got != want {
		// t.Errorf formats like fmt.Sprintf. The test is marked failed but continues.
		t.Errorf("Add(3, 4) = %d; want %d", got, want)
	}
}

// ======================================================================
// TABLE-DRIVEN TESTS — Go's idiomatic testing pattern
// ======================================================================
// Instead of writing one TestXxx function per case, define a slice of
// test cases (a "table") and loop over them.
// Benefits:
//   - Adding a new case = adding one line to the table
//   - Failures clearly show which case failed (by name)
//   - All cases run even if one fails (t.Errorf, not t.Fatalf)

func TestAddTable(t *testing.T) {
	// Define a slice of anonymous structs — each struct is one test case.
	tests := []struct {
		name string  // human-readable name shown in test output
		a, b int     // inputs
		want int     // expected output
	}{
		{name: "positive numbers", a: 3, b: 4, want: 7},
		{name: "zero + positive", a: 0, b: 5, want: 5},
		{name: "negative numbers", a: -3, b: -4, want: -7},
		{name: "mixed signs", a: -10, b: 4, want: -6},
		{name: "both zero", a: 0, b: 0, want: 0},
	}

	// Loop over the table.
	for _, tt := range tests {
		// t.Run creates a SUB-TEST for each case.
		// This lets you run a single case with:  go test -run TestAddTable/zero_+_positive
		// (spaces in name are replaced with _ in the test runner)
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// ======================================================================
// TESTING FUNCTIONS THAT RETURN (value, error)
// ======================================================================

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool // true if we expect an error from this call
	}{
		{name: "normal division", a: 10, b: 2, want: 5.0, wantErr: false},
		{name: "divide by zero", a: 5, b: 0, want: 0, wantErr: true},
		{name: "fractional result", a: 1, b: 3, want: 0.3333, wantErr: false},
		{name: "negative dividend", a: -9, b: 3, want: -3.0, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			// Step 1: check whether an error was returned or not.
			if tt.wantErr {
				// We EXPECTED an error — fail if we didn't get one.
				if err == nil {
					t.Errorf("Divide(%v, %v) expected an error but got nil", tt.a, tt.b)
				}
				return // no point checking the value if we expected an error
			}

			// We did NOT expect an error — fail if we got one.
			if err != nil {
				t.Fatalf("Divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
			}

			// Step 2: check the value.
			// Float comparisons need tolerance — 1/3 is never exactly 0.3333.
			// math.Abs(got - want) < epsilon is the standard float equality check.
			const epsilon = 0.0001
			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("Divide(%v, %v) = %.4f; want %.4f (±%.4f)",
					tt.a, tt.b, got, tt.want, epsilon)
			}
		})
	}
}

// ======================================================================
// TESTING BOOLEAN-RETURN FUNCTIONS
// ======================================================================

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"Race Car", true}, // spaces removed, case-insensitive
		{"", true},         // empty string is a palindrome (vacuously)
		{"A", true},        // single character
		{"ab", false},
		{"niveau", false},   // NOT a palindrome (n-i-v-e-a-u reversed = u-a-e-v-i-n)
		{"kayak", true},      // a real palindrome
	}

	for _, tt := range tests {
		// t.Run name is the input string itself — clear in output.
		t.Run(tt.input, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

// ======================================================================
// TESTING MAP RETURN VALUES
// ======================================================================

func TestWordCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "basic sentence",
			input: "go is go",
			want:  map[string]int{"go": 2, "is": 1},
		},
		{
			name:  "mixed case",
			input: "Go GO go",
			want:  map[string]int{"go": 3}, // all normalised to lowercase
		},
		{
			name:  "empty string",
			input: "",
			want:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordCount(tt.input)

			// Compare map length first.
			if len(got) != len(tt.want) {
				t.Errorf("WordCount(%q) returned %d keys; want %d\ngot:  %v\nwant: %v",
					tt.input, len(got), len(tt.want), got, tt.want)
				return
			}

			// Compare each expected key-value pair.
			for word, wantCount := range tt.want {
				if gotCount := got[word]; gotCount != wantCount {
					t.Errorf("WordCount(%q)[%q] = %d; want %d",
						tt.input, word, gotCount, wantCount)
				}
			}
		})
	}
}

// ======================================================================
// QUESTIONS — Testing
// ======================================================================
//
// CONCEPTUAL
// ----------
// Q1.  What is the naming convention for test files and test functions in Go?
//
//      A: Test files must end in "_test.go" — the Go build tool excludes them
//         from normal builds and includes them only for "go test".
//         Test function names must start with "Test" followed by a capital letter.
//         e.g.: func TestAdd(t *testing.T) {}
//         Benchmark functions start with "Benchmark", example functions with "Example".
//
// Q2.  What is the difference between t.Errorf and t.Fatalf?
//
//      A: t.Errorf("msg") — marks the test as failed and logs the message,
//                           but the test function CONTINUES running.
//                           Use for independent checks within a test.
//         t.Fatalf("msg") — marks as failed, logs, and STOPS the test immediately.
//                           Use when continuing would panic or give meaningless results
//                           (e.g., when err != nil and the next line uses the result).
//
// Q3.  What is a table-driven test? Why is it the idiomatic Go testing pattern?
//
//      A: A table-driven test defines test cases as a slice of structs (the "table"),
//         then loops over them calling t.Run for each one.
//         Idiomatic because:
//           - Adding cases is a single line in the table (not a new function)
//           - t.Run names each sub-test → failures clearly show which case broke
//           - All cases run independently (t.Errorf, not panic)
//           - Input, expected output, and name are co-located — easy to review
//
// Q4.  How do you run only one specific test or sub-test from the command line?
//
//      A: go test -run TestName .           — runs tests matching "TestName"
//         go test -run TestName/sub_name .  — runs a specific sub-test (spaces → _)
//         go test -v .                       — verbose: shows all test names and PASS/FAIL
//         go test -cover .                   — shows statement coverage percentage
//
// Q5.  Why can't you use == to compare floats in tests? What do you do instead?
//
//      A: Floating-point arithmetic is imprecise — 1.0/3.0 is not exactly 0.3333...
//         Two calculations that should give the same result may differ by tiny amounts.
//         Instead, check that the absolute difference is within a small tolerance (epsilon):
//           if math.Abs(got - want) > 0.0001 { t.Errorf(...) }
//         Some codebases use a helper like almostEqual(a, b, epsilon float64) bool.
//
// PRACTICAL
// ---------
// Q6.  Write a test for a Multiply(a, b int) int function using the
//      table-driven pattern. Include at least 5 cases.
//
//      A: func TestMultiply(t *testing.T) {
//             tests := []struct{ name string; a, b, want int }{
//                 {"positive", 3, 4, 12},
//                 {"by zero", 5, 0, 0},
//                 {"negatives", -3, -4, 12},
//                 {"mixed signs", -3, 4, -12},
//                 {"by one", 7, 1, 7},
//             }
//             for _, tt := range tests {
//                 t.Run(tt.name, func(t *testing.T) {
//                     if got := Multiply(tt.a, tt.b); got != tt.want {
//                         t.Errorf("Multiply(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
//                     }
//                 })
//             }
//         }
//
// Q7.  What does "go test -cover ." tell you? What does 100% coverage mean?
//
//      A: -cover reports the percentage of source lines executed by at least one test.
//         100% coverage means every line was executed — but it does NOT mean the
//         code is correct or all edge cases are tested. A line can be executed
//         with the wrong input and still "pass". Coverage is a floor, not a ceiling.
//         Aim for high coverage on business logic; don't obsess over 100%.
//
// Q8.  What is a benchmark function? How do you write and run one?
//
//      A: Benchmark functions start with "Benchmark" and take *testing.B.
//         The framework calls the function body b.N times and measures elapsed time.
//         func BenchmarkAdd(b *testing.B) {
//             for i := 0; i < b.N; i++ {
//                 Add(3, 4)
//             }
//         }
//         Run with: go test -bench=. .
//         Output:   BenchmarkAdd-8    1000000000    0.31 ns/op
// ======================================================================
