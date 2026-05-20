package main

import (
	"fmt"
	"unsafe"
)

/*
===============================================
Stage 1: Basics and Types
Part 1: Variables
===============================================

In Go, there are 3 main ways to declare a variable:
1. var x int           - Long form (Declaration without initialization, assigns zero value)
2. var x = 5           - Long form with type inference
3. x := 5              - Short assignment (Declaration and initialization) - ONLY inside functions!

⚠️ Important MCQ Notes & Edge Cases:
- `:=` only works inside functions. Using it at the package level will cause a compile-time error.
- `var` works everywhere (global/package level and local/function level).
- Variables declared but not used will cause a compile-time error (except for package-level variables).
- Grouping variables in a single `var` block is idiomatic for package-level variables.
- Memory: Package-level variables are allocated on the data segment. Local variables are allocated on the stack, but if they escape the function (e.g., returning a pointer to a local variable), Go's escape analysis will allocate them on the heap.
*/

func variablesDemo() {
	// Method 1: Long declaration
	var a int
	a = 10
	fmt.Println("Method 1 - var a int:", a)

	// Method 2: Declaration with initialization
	var b int = 20
	fmt.Println("Method 2 - var b int = 20:", b)

	// Method 3: Type Inference
	var c = 30 // Go infers that c is of type int
	fmt.Println("Method 3 - var c = 30:", c)

	// Method 4: Short assignment ⚠️ ONLY inside functions
	d := 40
	fmt.Println("Method 4 - d := 40:", d)

	// Method 5: Grouping variables
	var (
		x int    = 50
		y string = "Hello"
		z bool   = true
	)
	fmt.Printf("Method 5 - Grouping: x=%d, y=%s, z=%v\n", x, y, z)

	// Method 6: Multiple declarations on a single line
	var p, q, r = 100, 200, 300
	fmt.Printf("Method 6 - Multiple: p=%d, q=%d, r=%d\n", p, q, r)

	// ⚠️ Common MCQ Question: Difference between = and :=
	// = is used for reassignment on an existing variable.
	// := is used for declaration and initialization.

	// This will cause a compilation error:
	// d := 40  // Error: no new variables on left side of :=

	// But this is valid (reassignment):
	d = 50
	fmt.Println("Reassignment: d =", d)

	// ⚠️ Edge Case: Shadowing
	// You can shadow a variable in an inner scope using :=
	{
		d := 60 // This creates a NEW 'd' in this inner scope, shadowing the outer 'd'
		fmt.Println("Inner scope d (shadowed):", d)
	}
	fmt.Println("Outer scope d remains unchanged:", d)

	// ⚠️ MCQ Question: Blank Identifier (_)
	// Used to ignore unwanted values, bypassing the "declared and not used" compile error.
	fmt.Println("\n--- Blank Identifier ---")
	x1, _ := 100, 200 // Ignoring the second value
	fmt.Println("x1 =", x1, "(second value ignored)")
}

/*
===============================================
Basic Types
===============================================

1. Integers:
   - int, int8, int16, int32, int64
   - uint, uint8, uint16, uint32, uint64
   - uintptr: an unsigned integer large enough to store the uninterpreted bits of a pointer value.

2. Floats:
   - float32, float64

3. Complex Numbers:
   - complex64 (two float32s), complex128 (two float64s)

4. Strings:
   - string (immutable slice of bytes, read-only). Memory-wise, it's a struct with a pointer to the backing array and a length.

5. Characters:
   - byte (alias for uint8)
   - rune (alias for int32, represents a Unicode code point)

6. Booleans:
   - bool (true or false only)
*/

func typesDemo() {
	fmt.Println("\n======== Data Types ========")

	// Integers
	var intVal int = 42
	var int32Val int32 = 100
	var uint8Val uint8 = 255 // Max value for uint8

	fmt.Printf("int: %d, int32: %d, uint8: %d\n", intVal, int32Val, uint8Val)

	// Floats
	var float32Val float32 = 3.14
	var float64Val float64 = 2.71828

	fmt.Printf("float32: %f, float64: %f\n", float32Val, float64Val)

	// Complex Numbers
	var comp complex64 = 1 + 2i
	fmt.Printf("complex64: %v, Real: %f, Imaginary: %f\n", comp, real(comp), imag(comp))

	// Strings
	var str string = "Welcome to Go"
	fmt.Println("String:", str)
	fmt.Println("String length (bytes, NOT characters):", len(str))

	// Booleans
	var isTrue bool = true
	var isFalse bool = false

	fmt.Printf("bool: %v and %v\n", isTrue, isFalse)

	// ⚠️ MCQ Question: Size of different types
	// Using unsafe.Sizeof to check the type size (depends on architecture for int/uint/uintptr)
	fmt.Println("\n--- Memory Size (unsafe.Sizeof) ---")
	fmt.Printf("Size of int: %d bytes (depends on 32/64-bit OS)\n", unsafe.Sizeof(intVal))
	fmt.Printf("Size of int32: %d bytes\n", unsafe.Sizeof(int32Val))
	fmt.Printf("Size of bool: %d bytes\n", unsafe.Sizeof(isTrue))
	fmt.Printf("Size of string header: %d bytes\n", unsafe.Sizeof(str)) // 16 bytes on 64-bit (pointer + length)
}

// Removing main func from here as it will clash with other files in the same directory (stage1_types_and_variables)
// We will have a separate main.go file to run everything.
