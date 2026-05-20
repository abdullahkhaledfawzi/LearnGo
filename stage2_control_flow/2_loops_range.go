package main

import "fmt"

/*
===============================================
Stage 2: Control Flow - Part 2
Loops and Range
===============================================

⚠️ Crucial Points:
1. In Go, there is ONLY ONE looping construct: `for`.
2. There are no `while` or `do-while` loops. `for` handles everything.
3. The `range` keyword is used to iterate over data structures.
4. `range` always returns two values: (index, value) or (key, value). You can use `_` to ignore one.
5. MEMORY/POINTER WARNING: The `value` returned by `range` is a COPY. Modifying it does NOT modify the original element.

⚠️ Highly Common MCQ: The difference between `range` behavior and standard loops.
*/

func basicForLoops() {
	fmt.Println("======== Basic For Loops ========")

	// 1. Classic Loop (Standard 3-part loop)
	fmt.Println("\n--- Classic Loop ---")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// Output: 0 1 2 3 4

	// 2. Condition-Only Loop (Acts exactly like a `while` loop)
	fmt.Println("\n--- Condition-Only (While) Loop ---")
	count := 0
	for count < 3 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()
	// Output: 0 1 2

	// 3. Infinite Loop
	fmt.Println("\n--- Infinite Loop ---")
	num := 0
	for { // Same as `while(true)`
		fmt.Printf("%d ", num)
		num++
		if num >= 3 {
			break // Must explicitly break to avoid hanging
		}
	}
	fmt.Println()
	// Output: 0 1 2

	// 4. Continue and Break
	fmt.Println("\n--- continue and break ---")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip this iteration
		}
		if i == 4 {
			break // Exit the loop entirely
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// Output: 0 1 3
}

/*
===============================================
Range (Iterating over Collections)
===============================================

⚠️ Crucial memory and pointer behaviors:
1. `range` works on: array, slice, string, map, channel.
2. Returns (index, copy_of_value).
3. Modifying the copy DOES NOT modify the original collection.
4. With strings, `index` is the BYTE index, and `value` is the RUNE (Unicode character).
*/

func rangeArraysSlices() {
	fmt.Println("\n======== Range on Arrays & Slices ========")

	numbers := []int{10, 20, 30, 40}

	// 1. Full syntax: getting index and value
	fmt.Println("\n--- Full Range Syntax ---")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 2. Ignoring value
	fmt.Println("\n--- Ignoring value ---")
	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
	}

	// 3. Ignoring index (using the blank identifier _)
	fmt.Println("\n--- Ignoring index ---")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// ⚠️ MCQ Question: Modifying `value` inside range (Pass-by-Value behavior)
	fmt.Println("\n--- ⚠️ Modifying value inside range ---")
	slice := []int{1, 2, 3}
	for i, v := range slice {
		v = v * 10 // This ONLY modifies the local copy 'v'
		fmt.Printf("Inside loop: i=%d, v=%d\n", i, v)
	}
	fmt.Printf("Original slice: %v (Unchanged!)\n", slice)
	// Output:
	// Inside loop: i=0, v=10
	// Inside loop: i=1, v=20
	// Inside loop: i=2, v=30
	// Original slice: [1 2 3] (Unchanged!)

	// The CORRECT way to modify elements: use the index
	fmt.Println("\n--- The Correct Way to Modify Elements ---")
	slice2 := []int{1, 2, 3}
	for i := range slice2 {
		slice2[i] = slice2[i] * 10 // Access memory directly
	}
	fmt.Printf("Slice after modification: %v\n", slice2)
	// Output: Slice after modification: [10 20 30]
}

func rangeStrings() {
	fmt.Println("\n======== Range on Strings ========")

	// ⚠️ CRITICAL MCQ TOPIC!
	str := "Hello"

	// `range` on a string returns (byte_index, rune)
	fmt.Println("\n--- range on string (ASCII) ---")
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c (rune value: %d)\n", index, char, char)
	}

	// ⚠️ Edge Case with Multi-byte Characters (e.g., Arabic)
	fmt.Println("\n--- range on string (Arabic) ---")
	arabicStr := "مرحبا"
	for index, char := range arabicStr {
		fmt.Printf("Byte Index: %d, Char: %c\n", index, char)
	}
	// You will notice the index jumps by 2 (e.g., 0, 2, 4, 6, 8) because Arabic characters are 2 bytes each!
	// This proves that the index is a BYTE INDEX, not a character index.

	// Difference between range and slicing to bytes
	fmt.Println("\n--- range vs byte iteration ---")
	str2 := "Go"
	for i, ch := range str2 {
		fmt.Printf("Range: i=%d, ch=%c (Rune)\n", i, ch)
	}
	for i, b := range []byte(str2) {
		fmt.Printf("[]byte Iteration: i=%d, b=%c (Byte)\n", i, b)
	}
}

func rangeMaps() {
	fmt.Println("\n======== Range on Maps ========")

	// 1. Iterating over a map
	fruitsPrice := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 15,
	}

	fmt.Println("\n--- range on map ---")
	for key, value := range fruitsPrice {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 2. Ignoring value
	fmt.Println("\n--- Ignoring value ---")
	for key := range fruitsPrice {
		fmt.Printf("Key: %s\n", key)
	}

	// 3. Ignoring key
	fmt.Println("\n--- Ignoring key ---")
	for _, value := range fruitsPrice {
		fmt.Printf("Value: %d\n", value)
	}

	// ⚠️ CRITICAL: Map Iteration Order is RANDOM
	fmt.Println("\n--- Map Order is Randomized ---")
	// Go intentionally randomizes the starting offset of map iteration.
	// You can NEVER guarantee the order in which map keys will be iterated.
	for key := range fruitsPrice {
		fmt.Printf("Key: %s\n", key)
	}
}

/*
===============================================
Labeled Breaks (Breaking Out of Nested Loops)
===============================================
A powerful feature to escape deep nesting without complex flags.
*/

func labeledBreaks() {
	fmt.Println("\n======== Labeled Breaks ========")

	// 'OuterLoop' is a label pointing to the outer for-loop
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of BOTH loops simultaneously!")
				break OuterLoop // Escapes the entire nested structure
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
