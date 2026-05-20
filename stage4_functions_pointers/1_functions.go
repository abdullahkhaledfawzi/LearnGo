package main

import "fmt"

/*
===============================================
Stage 4: Functions and Pointers - Part 1
Functions
===============================================

⚠️ Critical Function Concepts:
1. Multiple Return Values: Functions in Go can return more than one value (extremely common for returning `(result, error)`).
2. Named Return Values: Return variables are predefined and act as local variables.
3. Variadic Functions: Functions that accept a variable number of arguments using `...` (like `fmt.Println`).
4. First-Class Citizens: Functions can be assigned to variables, passed as arguments, and returned from other functions.
5. Anonymous Functions and Closures: Functions without names, and functions that "close over" outer variables.
6. `defer` Statement: Executes a function call *immediately before* the surrounding function returns (LIFO order).
*/

// Basic Function
func add(a, b int) int {
	return a + b
}

// Multiple Return Values
// A hallmark of Go. No need for tuples or wrapper objects.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// Named Return Values (Extremely popular MCQ!)
// The variables `sum` and `product` are automatically initialized with their zero values.
func namedReturns(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // A "naked" return. It automatically returns the current values of `sum` and `product`.
}

/*
⚠️ Naked Returns Warning:
While clean for short functions, they can severely harm readability in long functions. Use with caution.
*/

func basicsDemo() {
	fmt.Println("======== Functions Basics ========")

	// 1. Basic Function
	result := add(5, 3)
	fmt.Printf("add(5, 3): %d\n", result)

	// 2. Multiple Returns (Error Checking Pattern)
	division, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("divide(10, 2): %f\n", division)
	}

	// Example handling an error
	division2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2) // Triggers here
	} else {
		fmt.Printf("divide(10, 0): %f\n", division2)
	}

	// 3. Named Returns
	s, p := namedReturns(4, 5)
	fmt.Printf("namedReturns(4, 5): sum=%d, product=%d\n", s, p)
}

/*
===============================================
Variadic Functions
===============================================

Accept an arbitrary number of arguments.
Inside the function, the variadic parameter becomes a Slice.
*/

// `numbers` is treated as a `[]int` inside the function.
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// The variadic parameter MUST be the LAST parameter in the signature.
func printAll(prefix string, values ...string) {
	fmt.Print(prefix + ": ")
	for _, v := range values {
		fmt.Print(v + " ")
	}
	fmt.Println()
}

func variadicDemo() {
	fmt.Println("\n======== Variadic Functions ========")

	// Passing individual arguments
	fmt.Printf("sum(1, 2, 3): %d\n", sum(1, 2, 3))
	fmt.Printf("sum(1, 2, 3, 4, 5): %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("sum(): %d\n", sum()) // Zero arguments works too!

	// ⚠️ Passing an existing slice to a variadic function using the Unpack Operator `...`
	numbers := []int{10, 20, 30}
	result := sum(numbers...) // The `...` unpacks the slice into individual arguments
	fmt.Printf("sum(numbers...): %d\n", result)
}

/*
===============================================
Anonymous Functions (IIFE)
===============================================
Functions without a name. Often used for short callbacks or goroutines.
*/

func anonymousFunctionsDemo() {
	fmt.Println("\n======== Anonymous Functions ========")

	// 1. Immediately Invoked Function Expression (IIFE)
	result := func(a, b int) int {
		return a * b
	}(4, 5) // Notice the () at the end that invokes it immediately
	fmt.Printf("IIFE (4 * 5): %d\n", result)

	// 2. Assigning function to a variable
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("Assigned multiply(3, 7): %d\n", multiply(3, 7))
}

/*
===============================================
Closures (Highly critical concept)
===============================================

A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
*/

func closuresDemo() {
	fmt.Println("\n======== Closures ========")

	// Outer variable
	counter := 0

	// This anonymous function "closes over" the `counter` variable.
	increment := func() {
		counter++
	}

	decrement := func() {
		counter--
	}

	// Both closures reference the EXACT SAME `counter` memory address!
	fmt.Printf("Initial counter: %d\n", counter)

	increment()
	fmt.Printf("After increment: %d\n", counter) // 1

	increment()
	fmt.Printf("After increment: %d\n", counter) // 2

	decrement()
	fmt.Printf("After decrement: %d\n", counter) // 1

	// ⚠️ DANGEROUS MCQ: Closures inside Loops (Loop Variable Capture)
	fmt.Println("\n--- Dangerous Loop Capture ---")
	var functions []func()

	// Prior to Go 1.22, this would capture the same `i` variable.
	// As of Go 1.22, loop variables are instanced per iteration!
	for i := 0; i < 3; i++ {
		functions = append(functions, func() {
			fmt.Printf("Captured i = %d\n", i)
		})
	}

	for _, f := range functions {
		f() // If Go < 1.22: prints 3, 3, 3. If Go 1.22+: prints 0, 1, 2.
	}
}

/*
===============================================
Higher-Order Functions
===============================================
Functions that take other functions as parameters or return them.
*/

// Function as a Parameter
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Function as a Return Value (Returns a Closure)
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func functionTypesDemo() {
	fmt.Println("\n======== Higher-Order Functions ========")

	add := func(a, b int) int { return a + b }
	fmt.Printf("applyOperation(5, 3, add): %d\n", applyOperation(5, 3, add))

	double := makeMultiplier(2) // `factor` is fixed at 2
	triple := makeMultiplier(3) // `factor` is fixed at 3

	fmt.Printf("double(5): %d\n", double(5)) // 10
	fmt.Printf("triple(5): %d\n", triple(5)) // 15
}

/*
===============================================
Defer Statement (Intro)
===============================================

Defers the execution of a function until the surrounding function returns.
Executed in LIFO (Last-In, First-Out) order. Essential for cleanup (closing files, unlocking mutexes).
*/

func simpleDefer() {
	fmt.Println("\n======== Defer Statement ========")

	defer func() {
		fmt.Println("3. Deferred function 1 executes LAST")
	}()

	fmt.Println("1. Normal execution happens FIRST")

	defer func() {
		fmt.Println("2. Deferred function 2 executes SECOND (LIFO)")
	}()
	
	// Outputs: 1, 2, 3
}
