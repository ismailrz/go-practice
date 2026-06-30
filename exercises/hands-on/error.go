package main

import (
	"errors"
	"fmt"
)

// Basic error — errors.New for a static message
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// fmt.Errorf — formatted error with dynamic values
func getUser(id int) (string, error) {
	users := map[int]string{1: "Alice", 2: "Bob"}
	user, ok := users[id]
	if !ok {
		return "", fmt.Errorf("user %d not found", id)
	}
	return user, nil
}

// Sentinel error — a package-level error value callers can compare against
var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	if id <= 0 {
		// %w wraps ErrNotFound so callers can detect it with errors.Is()
		return fmt.Errorf("findItem(%d): %w", id, ErrNotFound)
	}
	return nil
}

// Custom error type — implement the error interface: Error() string
// Lets callers extract structured data with errors.As()
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

func main() {

	// Basic (value, error) pattern
	fmt.Println("--- basic error ---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err) // cannot divide by zero
	}

	// fmt.Errorf
	fmt.Println("\n--- fmt.Errorf ---")

	name, err := getUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("user:", name) // Alice
	}

	_, err = getUser(99)
	fmt.Println("Error:", err) // user 99 not found

	// Sentinel error + errors.Is
	// errors.Is checks every level of a wrapped error chain
	fmt.Println("\n--- sentinel + errors.Is ---")

	err = findItem(-1)
	fmt.Println("error:", err) // findItem(-1): not found

	if errors.Is(err, ErrNotFound) {
		fmt.Println("root cause: not found") // detected through the wrapper
	}

	fmt.Println("findItem(5):", findItem(5)) // <nil>

	// Custom error type + errors.As
	// errors.As checks the chain and fills the target if a match is found
	fmt.Println("\n--- custom error + errors.As ---")

	err = validateAge(-5)
	fmt.Println("error:", err)

	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("field:", ve.Field)
		fmt.Println("message:", ve.Message)
	}

	fmt.Println("validateAge(25):", validateAge(25)) // <nil>

	// EXERCISE
	// 1. Write parsePositive(s string) (int, error) that:
	//    - Parses s with strconv.Atoi (wrap any parse error with fmt.Errorf + %w)
	//    - Returns error if the number is <= 0
	//    Test with "42", "-5", "abc"
	// 2. Create ErrDivByZero sentinel and update divide() to use it.
	//    Check for it in main with errors.Is.
	// 3. Create InsufficientFundsError{Balance, Amount float64} and use it
	//    in a Withdraw(bal, amt float64) (float64, error) function.
}
