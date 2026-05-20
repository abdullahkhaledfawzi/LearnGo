package main

import "fmt"

/*
===============================================
Stage 3: Data Structures - Part 2
Maps (Dictionaries / Hash Tables)
===============================================

⚠️ Critical Map Concepts:
1. Map is a Reference Type (Like a Slice). Under the hood, it's a pointer to a `hmap` struct.
2. You MUST initialize a map using `make()` or a map literal. A purely declared map is `nil`.
3. Map Keys must be a COMPARABLE type (supports `==` and `!=`).
4. Accessing a non-existent key returns the Zero Value of the value's type. It NEVER panics.
5. Use the "comma ok" idiom (`value, ok := map[key]`) to check if a key actually exists.
6. Map iteration order is INTENTIONALLY RANDOMIZED by the Go runtime to prevent relying on it.
7. Maps are NOT concurrency-safe! (Use `sync.Map` or Mutexes for concurrent access).
*/

func mapsBasics() {
	fmt.Println("======== Maps Basics ========")

	// 1. Declaration (Results in a nil map)
	var emptyMap map[string]int
	fmt.Printf("var emptyMap map[string]int: %v (nil)\n", emptyMap)

	// ⚠️ You must initialize it before writing!
	// emptyMap["apple"] = 1 // ❌ PANIC: assignment to entry in nil map
	
	myMap := make(map[string]int) // Safe to write
	myMap["apple"] = 10
	myMap["banana"] = 20
	fmt.Printf("myMap initialized with make: %v\n", myMap)

	// 2. Map literal
	colors := map[string]string{
		"red":   "Red",
		"green": "Green",
		"blue":  "Blue",
	}
	fmt.Printf("colors literal: %v\n", colors)

	// 3. Accessing Elements
	fmt.Println("\n--- Accessing Elements ---")
	val := myMap["apple"]
	fmt.Printf("myMap[\"apple\"]: %d\n", val) // 10

	// ⚠️ Non-existent key returns Zero Value (Does NOT throw an error like Python's KeyError)
	val2 := myMap["orange"]
	fmt.Printf("myMap[\"orange\"]: %d (Zero Value)\n", val2) // 0

	// 4. Checking if a key exists ("Comma ok" idiom)
	fmt.Println("\n--- Checking Key Existence ---")
	value, ok := myMap["apple"]
	fmt.Printf("myMap[\"apple\"]: value=%d, ok=%v\n", value, ok) // 10, true

	value, ok = myMap["orange"]
	fmt.Printf("myMap[\"orange\"]: value=%d, ok=%v\n", value, ok) // 0, false

	// ⚠️ Using `if` statement for inline checking
	if price, exists := myMap["apple"]; exists {
		fmt.Printf("Apple costs: %d\n", price)
	}
}

func mapOperations() {
	fmt.Println("\n======== Map Operations ========")

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	// 1. Add Element
	fmt.Println("\n--- Add Element ---")
	m["four"] = 4
	fmt.Printf("After adding four: %v\n", m)

	// 2. Update Element
	fmt.Println("\n--- Update Element ---")
	m["one"] = 100
	fmt.Printf("After updating one: %v\n", m)

	// 3. Delete Element
	fmt.Println("\n--- Delete Element ---")
	delete(m, "two")
	fmt.Printf("After deleting two: %v\n", m)

	// ⚠️ Deleting a non-existent key is perfectly SAFE (It does nothing)
	delete(m, "nonexistent")
	fmt.Printf("After deleting nonexistent key: %v\n", m)

	// 4. Length of Map
	fmt.Println("\n--- Length of Map ---")
	fmt.Printf("len(m): %d\n", len(m))
	// Note: cap(m) does NOT exist for maps. You cannot check map capacity.

	// 5. Ranging over a Map
	fmt.Println("\n--- Range over Map ---")
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
	// ⚠️ Remember: Output order is randomized every time!
}

/*
===============================================
Comparable Key Types
===============================================

Keys must support `==` and `!=`.
✓ Allowed: int, float, string, bool, pointers, channels, interfaces, arrays (if elements are comparable), structs (if fields are comparable)
✗ NOT Allowed: Slices, Maps, Functions
*/

func mapKeyTypes() {
	fmt.Println("\n======== Map Key Types ========")

	// 1. String keys (Most common)
	m1 := map[string]string{"name": "Ahmed"}
	
	// 2. Int keys
	m2 := map[int]string{1: "one"}
	
	// 3. Array keys (Rare but perfectly legal because arrays are comparable)
	m3 := map[[3]int]string{
		{1, 2, 3}: "array key value",
	}
	fmt.Printf("Array keys map: %v\n", m3)

	// ⚠️ Illegal Keys:
	// m4 := map[[]int]string{} // ❌ Error: invalid map key type []int (slices are not comparable)
	// m5 := map[func()]string{} // ❌ Error: invalid map key type func()
	
	_ = m1
	_ = m2
}

/*
===============================================
Maps as Reference Types
===============================================

When you pass a map to a function, you are passing a pointer to the `hmap` header.
Modifications made inside the function WILL affect the original map.
*/

func mapsByValue() {
	fmt.Println("\n======== Maps as Reference Type ========")

	m1 := map[string]int{"a": 1}
	m2 := m1 // m2 now points to the exact same underlying hash table as m1

	m2["b"] = 2
	fmt.Printf("m1: %v\n", m1) // map[a:1 b:2]
	fmt.Printf("m2: %v\n", m2) // map[a:1 b:2]
	// ⚠️ m1 changed because m2 modified the shared memory!

	// Contrast with Arrays (Value Types)
	fmt.Println("\n--- Contrast with Arrays (Value Types) ---")
	arr1 := [2]int{1, 2}
	arr2 := arr1 // Complete copy
	arr2[0] = 99

	fmt.Printf("arr1: %v\n", arr1) // [1 2] (Unchanged)
	fmt.Printf("arr2: %v\n", arr2) // [99 2]
}

/*
===============================================
nil Map vs Empty Map Edge Cases
===============================================
*/

func nilVsEmptyMap() {
	fmt.Println("\n======== nil vs Empty Map ========")

	var nilMap map[string]int
	emptyMap := make(map[string]int)

	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)     // true
	fmt.Printf("emptyMap == nil: %v\n", emptyMap == nil) // false

	fmt.Printf("len(nilMap): %d\n", len(nilMap))     // 0
	fmt.Printf("len(emptyMap): %d\n", len(emptyMap)) // 0

	// ⚠️ DANGEROUS DIFFERENCE: Writing
	// emptyMap["key"] = 1  // ✓ Safe
	// nilMap["key"] = 1    // ❌ PANIC! (Cannot write to a nil map)

	// ✓ Reading from BOTH is completely safe
	val1 := nilMap["key"]   // Returns Zero Value (0)
	val2 := emptyMap["key"] // Returns Zero Value (0)
	fmt.Printf("Reading: nilMap=%d, emptyMap=%d\n", val1, val2)

	// ✓ Iterating over BOTH is completely safe (nil map loop will just bypass)
	for k := range nilMap {
		fmt.Printf("nilMap key: %s\n", k)
	}
}

/*
===============================================
Maps of Structs
===============================================
*/

type Person struct {
	Name string
	Age  int
}

func mapWithStructs() {
	fmt.Println("\n======== Maps with Structs ========")

	people := make(map[string]Person)
	people["ahmed"] = Person{"Ahmed", 25}
	people["fatima"] = Person{Name: "Fatima", Age: 23}

	for name, person := range people {
		fmt.Printf("Key '%s': %s is %d years old\n", name, person.Name, person.Age)
	}
	
	// ⚠️ Advanced Gotcha: You cannot modify a struct field directly inside a map!
	// people["ahmed"].Age = 26 // ❌ Error: cannot assign to struct field people["ahmed"].Age in map
	
	// Correct way to update a struct in a map:
	temp := people["ahmed"]
	temp.Age = 26
	people["ahmed"] = temp // Overwrite the entire struct
	
	// Alternatively, use a map of pointers: `map[string]*Person`
}
