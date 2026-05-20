package main

import "fmt"

/*
===============================================
Stage 3: Data Structures
Part 1: Arrays vs Slices
===============================================

⚠️ Crucial Core Differences (99% Guaranteed MCQ Topic):

1. ARRAY:
   - Fixed size (must be defined at compile time or inferred via `...`).
   - VALUE TYPE: Passing an array to a function copies the ENTIRE array (bad for memory if large).
   - Non-dynamic: Cannot grow or shrink.
   - Syntax: `[size]Type` or `[...]Type`
   - Memory: Allocated as a contiguous block on the stack (usually).

2. SLICE:
   - Dynamic size (can grow via `append`).
   - REFERENCE TYPE (Descriptor): It is actually a struct under the hood containing:
     1. Pointer to the underlying Array
     2. Length (int)
     3. Capacity (int)
   - Passing a slice to a function only copies the descriptor (24 bytes on a 64-bit system), NOT the data!
   - Syntax: `[]Type`
   - Memory: Backed by an array, usually allocated on the heap if it escapes or grows dynamically.

3. Every Slice points to an underlying Array. Every Array can be sliced.
*/

func arraysBasics() {
	fmt.Println("======== Arrays ========")

	// 1. Array declaration with fixed size (initialized with zero values)
	var arr [3]int // Array of 3 integers
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Printf("Array: %v\n", arr) // [10 20 30]
	fmt.Printf("Array Length: %d\n", len(arr)) // capacity is also 3, but cap() is rarely used on arrays

	// 2. Declaration with Initialization
	arr2 := [5]string{"Go", "Python", "Java", "Rust", "C++"}
	fmt.Printf("Array with values: %v\n", arr2)

	// 3. Inferring size using (...)
	// The compiler counts the elements for you.
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Array (inferred size): %v, len: %d\n", arr3, len(arr3))

	// ⚠️ Edge Case: Partial Initialization
	arr4 := [5]int{1, 2}                             // Remaining elements get Zero Value (0)
	fmt.Printf("Partial initialization: %v\n", arr4) // [1 2 0 0 0]

	// ⚠️ MCQ Question: Indexed Initialization
	// You can specify elements by index. Unspecified indices get Zero Value.
	arr5 := [5]int{0: 100, 4: 500}
	fmt.Printf("Indexed init: %v\n", arr5) // [100 0 0 0 500]
}

func slicesBasics() {
	fmt.Println("\n======== Slices ========")

	// 1. Creating a Slice from an Array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]                           // Elements from index 1 up to (but not including) 4
	fmt.Printf("Slice from array: %v\n", slice) // [2 3 4]

	// 2. Slice Literal
	// Under the hood, this creates a hidden array and returns a slice pointing to it.
	slice2 := []int{10, 20, 30, 40}
	fmt.Printf("Slice literal: %v\n", slice2)

	// ⚠️ VERY IMPORTANT: Slice Operations (Slicing a Slice)
	fmt.Println("\n--- Slice Operations ---")
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("s[1:3]: %v\n", s[1:3]) // [2 3]
	fmt.Printf("s[:3]: %v\n", s[:3])   // [1 2 3] (From start to index 2)
	fmt.Printf("s[2:]: %v\n", s[2:])   // [3 4 5] (From index 2 to end)
	fmt.Printf("s[:]: %v\n", s[:])     // [1 2 3 4 5] (All elements)

	// ⚠️ DANGER: Negative indices DO NOT work in Go!
	// fmt.Printf("s[-1]: %v\n", s[-1])  // ❌ Compilation Error: invalid slice index

	// 3. The `make` Function
	// make() is used to initialize slices, maps, and channels.
	fmt.Println("\n--- Make Function ---")
	slice3 := make([]int, 3) // length 3, capacity 3
	fmt.Printf("make([]int, 3): %v, len=%d, cap=%d\n", slice3, len(slice3), cap(slice3))

	slice4 := make([]int, 3, 10) // length 3, capacity 10 (Allocates array of size 10)
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n", slice4, len(slice4), cap(slice4))

	// 4. Append
	// `append` adds elements to the end of the slice's length.
	fmt.Println("\n--- Append ---")
	slice5 := []int{1, 2, 3}
	fmt.Printf("Before append: %v, len=%d, cap=%d\n", slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, 4)
	fmt.Printf("After append(4): %v, len=%d, cap=%d\n", slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, 5, 6, 7)
	fmt.Printf("After append(5, 6, 7): %v, len=%d, cap=%d\n", slice5, len(slice5), cap(slice5))

	// ⚠️ Appending another slice using the variadic unpack operator `...`
	slice6 := []int{10, 20}
	slice5 = append(slice5, slice6...) 
	fmt.Printf("After append(slice6...): %v\n", slice5)
}

/*
===============================================
⚠️ Popular MCQ: Capacity vs Length and Memory Reallocation
===============================================

- Length (len): Number of elements currently accessible in the slice.
- Capacity (cap): Total number of elements in the underlying array starting from the slice's pointer.

When `append` exceeds the slice's Capacity:
1. Go allocates a NEW underlying array (usually double the old capacity if < 1024 elements).
2. It copies all existing elements to the new array.
3. It appends the new elements.
4. The slice descriptor is updated to point to the new array.
5. ⚠️ DANGER: If other slices were pointing to the OLD array, they will NOT see the new elements!
*/

func capacityDemo() {
	fmt.Println("\n======== Capacity Demo ========")

	// Difference between Array and Slice origins
	arr := [10]int{1, 2, 3, 4, 5}
	slice := arr[:3] // len = 3, cap = 10 (because from index 0 to end of array is 10 spaces)

	fmt.Printf("Array: %v, len=%d\n", arr, len(arr))
	fmt.Printf("Slice: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// Reslicing up to the capacity
	slice = slice[:5]
	fmt.Printf("After reslicing slice[:5]: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
	// Elements 4 and 5 were in the original array

	// ⚠️ HIGHLY TRICKY MCQ: Does Append modify the original array?
	fmt.Println("\n--- Tricky Edge Case: Append modifying Original Array ---")
	original := [5]int{1, 2, 3, 0, 0}
	sliceFromArray := original[:3] // len: 3, cap: 5

	fmt.Printf("Before append: original=%v, slice=%v\n", original, sliceFromArray)

	sliceFromArray = append(sliceFromArray, 99)
	fmt.Printf("After append(99): original=%v, slice=%v\n", original, sliceFromArray)
	// ⚠️ original[3] became 99! Because the slice had enough capacity, it overwrote the original array.

	sliceFromArray = append(sliceFromArray, 100, 101, 102)
	fmt.Printf("After appending more: original=%v, slice=%v\n", original, sliceFromArray)
	// Now, `append` exceeded capacity (5). It created a NEW array. `original` is no longer modified.
}

/*
===============================================
Copy Function & Memory Leaks
===============================================

⚠️ Memory Leak Warning:
If you take a small slice of a HUGE array and keep the slice in memory, the Garbage Collector (GC)
CANNOT free the huge array because the slice holds a pointer to it!

Solution: Use `copy()` to copy the data into a brand new, small slice.
*/

func copyDemo() {
	fmt.Println("\n======== Copy Demo ========")

	// Danger: Modifying slice modifies array
	fmt.Println("\n--- Danger: Shared Memory ---")
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]

	slice[0] = 999
	fmt.Printf("Modified slice, array is now: %v\n", arr)
	// ⚠️ arr[1] became 999!

	// Solution: Use `copy()`
	fmt.Println("\n--- Solution: Using Copy ---")
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3)
	copy(slice2, arr2[1:4]) // copy(dst, src)

	slice2[0] = 999
	fmt.Printf("Modified copied slice, arr2 is still: %v\n", arr2)
	// ✓ arr2 is unchanged because slice2 has its own underlying memory.
}

/*
===============================================
nil Slice vs Empty Slice
===============================================

⚠️ Extremely Common Interview / MCQ Question
*/

func nilVsEmptySlice() {
	fmt.Println("\n======== nil vs Empty Slice ========")

	var nilSlice []int                     // Nil slice (Pointer is nil)
	emptySlice := []int{}                  // Empty slice (Pointer points to a zero-byte memory address)
	makeSlice := make([]int, 0)            // Empty slice

	fmt.Printf("nilSlice == nil: %v\n", nilSlice == nil)     // true
	fmt.Printf("emptySlice == nil: %v\n", emptySlice == nil) // false
	fmt.Printf("makeSlice == nil: %v\n", makeSlice == nil)   // false

	fmt.Printf("len(nilSlice): %d, cap(nilSlice): %d\n", len(nilSlice), cap(nilSlice))         // 0, 0
	fmt.Printf("len(emptySlice): %d, cap(emptySlice): %d\n", len(emptySlice), cap(emptySlice)) // 0, 0
	fmt.Printf("len(makeSlice): %d, cap(makeSlice): %d\n", len(makeSlice), cap(makeSlice))     // 0, 0

	// ⚠️ IMPORTANT:
	// JSON marshaling handles them differently:
	// nilSlice -> outputs `null`
	// emptySlice -> outputs `[]`
	
	// Appending to a nil slice is completely safe!
	nilSlice = append(nilSlice, 1)
	fmt.Println("Appended to nil slice successfully:", nilSlice)
}
