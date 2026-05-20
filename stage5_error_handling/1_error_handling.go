package main

import (
	"errors"
	"fmt"
)

/*
===============================================
Stage 5: Error Handling - Part 1
The Error Interface
===============================================

⚠️ Critical MCQ Points (100% chance on an exam):

1. `error` is a simple built-in interface.
   type error interface {
       Error() string
   }
2. Any type that implements the `Error() string` method implicitly satisfies the `error` interface!
3. Go does NOT have `try/catch/finally` blocks like Java or Python. Errors are treated as normal values.
4. Standard Pattern: `if err != nil { return err }`
5. The Zero Value of an error is `nil` (meaning no error occurred).
6. ⚠️ Be careful: A typed `nil` pointer assigned to an `error` interface is NOT a `nil` error!
*/

// Example: Custom Error struct
type MyError struct {
	Code    int
	Message string
}

// By implementing this method, MyError automatically becomes an `error`
func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Function returning an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Returning a custom error
func process(value int) (string, error) {
	if value < 0 {
		return "", MyError{
			Code:    400,
			Message: "value must be positive",
		}
	}
	return fmt.Sprintf("processed successfully: %d", value), nil
}

func basicErrorHandling() {
	fmt.Println("======== Error Handling Basics ========")

	// 1. Standard Pattern
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// 2. Error case
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
		fmt.Printf("Returned Value: %.2f (Zero Value)\n", result2)
	}

	// 3. Custom error usage
	res, err3 := process(-5)
	if err3 != nil {
		fmt.Println("Custom error triggered:", err3)
	}

	res2, err4 := process(10)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		fmt.Printf("Result: %s\n", res2)
	}
}

/*
===============================================
Error Wrapping & Unwrapping (Go 1.13+)
===============================================

⚠️ Error wrapping allows you to add context to an error while preserving the original error type!
Syntax: `fmt.Errorf("context: %w", err)`
*/

func innerFunction() error {
	return errors.New("database connection failed")
}

func middleFunction() error {
	err := innerFunction()
	if err != nil {
		return fmt.Errorf("middleFunction failed: %w", err) // %w wraps the error
	}
	return nil
}

func outerFunction() error {
	err := middleFunction()
	if err != nil {
		return fmt.Errorf("outerFunction failed: %w", err)
	}
	return nil
}

func errorWrappingDemo() {
	fmt.Println("\n======== Error Wrapping ========")

	err := outerFunction()
	if err != nil {
		fmt.Println("Wrapped Error Chain:")
		fmt.Println(err) // outerFunction failed: middleFunction failed: database connection failed
	}
	
	// Unwrapping: gets the immediate underlying error
	unwrapped := errors.Unwrap(err)
	fmt.Printf("Unwrapped once: %v\n", unwrapped)
}

/*
===============================================
Type Assertion for Errors vs errors.As
===============================================

Sometimes you need to inspect the *type* of the error to handle it differently.
*/

func handleDifferentErrors(value int) error {
	if value < 0 {
		return MyError{400, "negative value"}
	}
	if value > 100 {
		return errors.New("value too large")
	}
	return nil
}

func typeAssertionForErrors() {
	fmt.Println("\n======== Type Assertion for Errors ========")

	tests := []int{-5, 50, 150}

	for _, val := range tests {
		err := handleDifferentErrors(val)
		if err != nil {
			// Using Type Assertion
			if myErr, ok := err.(MyError); ok {
				fmt.Printf("Custom error for %d: Code=%d, Message=%s\n", val, myErr.Code, myErr.Message)
			} else {
				fmt.Printf("Generic error for %d: %v\n", val, err)
			}
		}
	}
}

/*
===============================================
errors.Is and errors.As (Go 1.13+)
===============================================
These functions gracefully handle wrapped errors, making them vastly superior to `==` or type assertions!
*/

var ErrInvalidInput = errors.New("invalid input")

func validateInput(s string) error {
	if s == "" {
		// We wrap the sentinel error
		return fmt.Errorf("validation failed: %w", ErrInvalidInput)
	}
	return nil
}

func errorComparison() {
	fmt.Println("\n======== errors.Is & errors.As ========")

	err := validateInput("")
	
	// 1. errors.Is: Checks if the specific error exists anywhere in the chain.
	// Never use `err == ErrInvalidInput` if wrapping is involved!
	if errors.Is(err, ErrInvalidInput) {
		fmt.Println("Found ErrInvalidInput inside the wrapped error chain!")
	}

	// 2. errors.As: Extracts a specific error type from the chain.
	var customErr MyError
	wrappedCustom := fmt.Errorf("wrap: %w", MyError{500, "Server Fault"})
	
	if errors.As(wrappedCustom, &customErr) {
		fmt.Printf("Successfully extracted MyError: Code %d\n", customErr.Code)
	}
}

/*
===============================================
Sentinel Errors
===============================================

Pre-defined, globally accessible error variables.
Naming convention: start with `Err` (e.g., `io.EOF`, `os.ErrNotExist`).
*/

var (
	ErrFileNotFound  = errors.New("file not found")
	ErrPermission    = errors.New("permission denied")
	ErrCorruptedFile = errors.New("file corrupted")
)

func readFile(filename string) error {
	if filename == "" {
		return ErrFileNotFound
	}
	if filename == "secret.txt" {
		return ErrPermission
	}
	if filename == "bad.txt" {
		return ErrCorruptedFile
	}
	return nil
}

func sentinelErrorsDemo() {
	fmt.Println("\n======== Sentinel Errors ========")

	files := []string{"", "secret.txt", "good.txt", "bad.txt"}

	for _, f := range files {
		err := readFile(f)
		if err != nil {
			// Because these aren't wrapped, `switch` works perfectly.
			// If they were wrapped, we would use `errors.Is`.
			switch err {
			case ErrFileNotFound:
				fmt.Printf("'%s': File not found\n", f)
			case ErrPermission:
				fmt.Printf("'%s': Permission denied\n", f)
			case ErrCorruptedFile:
				fmt.Printf("'%s': Corrupted file\n", f)
			}
		} else {
			fmt.Printf("'%s': Read successfully\n", f)
		}
	}
}

/*
===============================================
Return Early Pattern (Bouncer Pattern)
===============================================

⚠️ Crucial Go Best Practice:
Instead of nesting successful logic inside `if err == nil`, always handle the error and `return` immediately.
Keep the "happy path" aligned to the left margin.
*/

func multipleOperations(a, b, c int) error {
	// Guard Clause 1
	if a < 0 {
		return fmt.Errorf("a must be positive")
	}

	// Guard Clause 2
	if b == 0 {
		return fmt.Errorf("b cannot be zero")
	}

	// Guard Clause 3
	if c > 100 {
		return fmt.Errorf("c must be less than 100")
	}

	// Happy path (Left aligned)
	fmt.Println("All operations succeeded!")
	return nil
}

func returnEarlyDemo() {
	fmt.Println("\n======== Return Early ========")

	tests := []struct {
		a, b, c int
	}{
		{10, 5, 50},  // ✓
		{-5, 5, 50},  // ✗
		{10, 0, 50},  // ✗
	}

	for _, test := range tests {
		err := multipleOperations(test.a, test.b, test.c)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
