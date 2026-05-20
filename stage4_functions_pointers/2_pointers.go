package main

import "fmt"

/*
===============================================
Stage 4: Functions and Pointers - Part 2
Pointers
===============================================

⚠️ Crucial MCQ Points:
1. `&` (Address Operator): Gets the memory address of a variable.
2. `*` (Dereference Operator): Reads or writes the value located at a memory address.
3. Zero Value: The zero value of any pointer is `nil`.
4. Dereferencing `nil`: Causes a severe runtime panic!
5. Value Type vs Reference Type: Pointers are how we simulate pass-by-reference in Go, since everything in Go is technically pass-by-value (a pointer itself is copied, but both copies point to the same memory).

⚠️ Common Mistake: Confusing `*` in type declaration (e.g., `var p *int`) with `*` in usage (e.g., `*p = 5`).
*/

func basicPointers() {
	fmt.Println("======== Basic Pointers ========")

	// 1. Normal Variable
	x := 10
	fmt.Printf("x = %d (Address: %p)\n", x, &x)

	// 2. Getting the Address
	var ptr *int = &x               // `ptr` holds the memory address of `x`
	fmt.Printf("ptr = %p\n", ptr)   // Memory address
	fmt.Printf("*ptr = %d\n", *ptr) // Dereferencing: gets the value of `x`

	// 3. Modifying Value via Pointer
	*ptr = 20
	fmt.Printf("After *ptr = 20: x = %d\n", x) // `x` became 20

	// 4. Pointer to a Pointer
	ptrPtr := &ptr // Pointer pointing to another pointer
	fmt.Printf("ptrPtr = %p\n", ptrPtr)
	fmt.Printf("*ptrPtr = %p\n", *ptrPtr)   // Dereferences to the first pointer (Address of x)
	fmt.Printf("**ptrPtr = %d\n", **ptrPtr) // Double dereference: gets the value of `x`

	// Modifying via pointer to pointer
	**ptrPtr = 30
	fmt.Printf("After **ptrPtr = 30: x = %d\n", x)
}

/*
===============================================
nil Pointers
===============================================
*/

func nilPointersDemo() {
	fmt.Println("\n======== nil Pointers ========")

	var ptr *int
	fmt.Printf("Zero value pointer: %v\n", ptr) // <nil>

	// ⚠️ DANGER: Dereferencing nil
	// fmt.Printf("*ptr: %d\n", *ptr) // ❌ panic: runtime error: invalid memory address or nil pointer dereference

	// Safe checking
	if ptr == nil {
		fmt.Println("Pointer is safely identified as nil.")
	}

	// Assigning a valid address
	x := 42
	ptr = &x
	fmt.Printf("After assignment: *ptr = %d\n", *ptr)
}

/*
===============================================
Why Use Pointers?
===============================================

1. Mutate state: Allow a function to modify a variable from the caller's scope.
2. Performance/Memory: Avoid copying massive structs when passing them to functions.
3. Semantics: A pointer can be `nil`, indicating the "absence" of a value (optional fields in JSON).
*/

// No pointer - receives a COPY of `x`. Caller's `x` is safe.
func incrementNoPointer(x int) {
	x++
}

// Pointer - receives the ADDRESS. Modifies the caller's memory.
func incrementWithPointer(x *int) {
	*x++ // Dereference and increment
}

func functionPointersDemo() {
	fmt.Println("\n======== Function Parameters with Pointers ========")

	x := 10
	
	incrementNoPointer(x)
	fmt.Printf("After incrementNoPointer: x = %d (Unchanged)\n", x)

	incrementWithPointer(&x)
	fmt.Printf("After incrementWithPointer: x = %d (Changed!)\n", x)
}

/*
===============================================
Pointers vs Slices
===============================================

⚠️ Critical Distinction:
Slices are ALREADY reference types (they contain a pointer to the backing array).
You rarely need a pointer to a slice (`*[]int`), unless you need the function to `append` and modify the slice's Length/Capacity for the caller!
*/

func sliceVsPointer() {
	fmt.Println("\n======== Slices vs Pointers ========")

	// Passing Slice directly
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Printf("After modifySlice: slice = %v\n", slice) 
	// Output: [999, 2, 3]. It modified index 0, but the append() was lost!

	// Passing Pointer to Slice
	slice2 := []int{1, 2, 3}
	modifySlicePointer(&slice2)
	fmt.Printf("After modifySlicePointer: slice2 = %v\n", slice2) 
	// Output: [999, 2, 3, 4]. The append() successfully modified the caller's slice!
}

// Modifies underlying data, but cannot change the caller's length/capacity.
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 
	}
	s = append(s, 4) // Reallocates or extends length locally. Caller won't see this!
}

// Full control over the caller's slice descriptor.
func modifySlicePointer(s *[]int) {
	if len(*s) > 0 {
		(*s)[0] = 999
	}
	*s = append(*s, 4) // Modifies the caller's length/capacity directly.
}

/*
===============================================
Escape Analysis (Advanced Memory Topic)
===============================================

In C/C++, returning a pointer to a local variable causes a Segmentation Fault (Dangling Pointer) because the stack frame is destroyed.

In Go, the compiler performs "Escape Analysis". If it detects that a pointer "escapes" the function's scope, it intelligently allocates the variable on the HEAP instead of the stack. It is 100% safe to return pointers to local variables in Go!
*/

func getPointerToNew() *int {
	x := 100 // Escape analysis detects `x` escapes, so `x` is allocated on the Heap.
	return &x // ✓ Perfectly safe in Go.
}

func escapingPointer() {
	fmt.Println("\n======== Escape Analysis ========")
	
	ptr := getPointerToNew()
	fmt.Printf("getPointerToNew: *ptr = %d (Valid memory!)\n", *ptr)
}
