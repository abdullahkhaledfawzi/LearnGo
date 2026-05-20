package main

import "fmt"

/*
===============================================
Zero Values
===============================================

Every data type in Go has a "Zero Value"
which is automatically assigned when a variable is declared without initialization.

⚠️ This is a very common MCQ topic! Go guarantees memory is initialized.
No garbage values exist in Go.
*/

func zeroValuesDemo() {
	fmt.Println("======== Zero Values ========")

	// 1. Integers -> 0
	var intZero int
	var int32Zero int32
	var uint8Zero uint8

	fmt.Printf("int zero value: %d\n", intZero)
	fmt.Printf("int32 zero value: %d\n", int32Zero)
	fmt.Printf("uint8 zero value: %d\n", uint8Zero)

	// 2. Floats -> 0.0
	var float32Zero float32
	var float64Zero float64

	fmt.Printf("float32 zero value: %f\n", float32Zero)
	fmt.Printf("float64 zero value: %f\n", float64Zero)

	// 3. Strings -> "" (empty string)
	var stringZero string
	fmt.Printf("string zero value: '%s' (empty string)\n", stringZero)
	fmt.Printf("Empty string length: %d\n", len(stringZero))

	// 4. Booleans -> false
	var boolZero bool
	fmt.Printf("bool zero value: %v\n", boolZero)

	// 5. Pointers -> nil
	var pointerZero *int
	fmt.Printf("pointer zero value: %v (nil)\n", pointerZero)

	// 6. Slices -> nil
	var sliceZero []int
	fmt.Printf("slice zero value: %v (nil)\n", sliceZero)
	fmt.Printf("Is slice nil: %v\n", sliceZero == nil)

	// 7. Maps -> nil
	var mapZero map[string]int
	fmt.Printf("map zero value: %v (nil)\n", mapZero)
	fmt.Printf("Is map nil: %v\n", mapZero == nil)
	
	// 8. Channels -> nil
	var chanZero chan int
	fmt.Printf("channel zero value: %v (nil)\n", chanZero)

	// 9. Interfaces -> nil
	var interfaceZero interface{}
	fmt.Printf("interface zero value: %v (nil)\n", interfaceZero)

	// 10. Structs -> nil? NO! Structs get initialized with zero values for all their fields.
	type Person struct {
		Name string
		Age  int
	}
	var structZero Person
	fmt.Printf("struct zero value: %+v (fields are zeroed)\n", structZero)

	// ⚠️ CRUCIAL DIFFERENCE: Empty Slice vs Nil Slice
	// var emptySlice []int          -> nil (No underlying array allocated)
	// emptySlice := []int{}         -> empty slice (Allocated, length 0, capacity 0, points to a zero-byte array)
	// emptySlice := make([]int, 0)  -> empty slice (Allocated, length 0, capacity 0)

	fmt.Println("\n--- CRITICAL: nil vs empty ---")
	var nilSlice []int
	emptySlice := []int{}
	makeSlice := make([]int, 0)

	fmt.Printf("nilSlice == nil: %v (len: %d, cap: %d)\n", nilSlice == nil, len(nilSlice), cap(nilSlice))
	fmt.Printf("emptySlice == nil: %v (len: %d, cap: %d)\n", emptySlice == nil, len(emptySlice), cap(emptySlice))
	fmt.Printf("makeSlice == nil: %v (len: %d, cap: %d)\n", makeSlice == nil, len(makeSlice), cap(makeSlice))

	// ⚠️ Important Edge Case:
	// Appending to a nil slice is perfectly SAFE. `append` will allocate a new underlying array.
	nilSlice = append(nilSlice, 1)
	fmt.Println("After appending to nilSlice:", nilSlice)

	// However, assigning directly to an index of a nil slice will PANIC.
	// nilSlice[0] = 5 // PANIC: index out of range [0] with length 0

	// Also, writing to a nil map will PANIC!
	// mapZero["key"] = 1 // PANIC: assignment to entry in nil map
	// But reading from a nil map is SAFE (returns zero value).
	fmt.Println("Reading from nil map:", mapZero["key"]) // Prints 0
}

/*
===============================================
Uninitialized Values vs Declaration
===============================================

Go does not have "uninitialized" variables with garbage data.
Everything gets a Zero Value upon declaration.
*/

func uninitializedVarsDemo() {
	fmt.Println("\n======== Default Initializations ========")

	// No "random" garbage values in Go
	// Arrays are value types, they are initialized with zero values for their elements
	var numbers [3]int                        // Array of 3 ints
	fmt.Println("Uninitialized array:", numbers) // [0 0 0]

	type Person struct {
		Name string
		Age  int
	}

	var person Person // Struct without explicit initialization
	fmt.Printf("Uninitialized Struct: Name='%s', Age=%d\n", person.Name, person.Age)
	// Output: Name='', Age=0 (Zero values)
}

/*
===============================================
Checking for nil
===============================================

⚠️ MCQ Question: When can a value be nil?
- Pointers
- Slices
- Maps
- Channels
- Interfaces
- Functions

It CANNOT be nil:
- Numbers (int, float), Strings, Arrays, Structs, Booleans
*/

func nilDemo() {
	fmt.Println("\n======== nil vs Zero Value ========")

	// nil or Zero Value?
	var a int
	var p *int
	var s []int
	var m map[string]int
	var i interface{}
	var fn func()

	fmt.Printf("int: %v (is zero value)\n", a)
	fmt.Printf("*int: %v (is nil)\n", p)
	fmt.Printf("[]int: %v (is nil)\n", s)
	fmt.Printf("map: %v (is nil)\n", m)
	fmt.Printf("interface{}: %v (is nil)\n", i)
	fmt.Printf("func: %v (is nil)\n", fn)

	// ⚠️ VERY IMPORTANT: A typed nil is NOT equal to an untyped nil interface
	// Interface under the hood is a tuple (Type, Value).
	// For an interface to be `nil`, BOTH Type and Value must be nil.
	var nilInterface interface{} = nil
	fmt.Printf("nilInterface == nil: %v\n", nilInterface == nil) // true

	var nilPointer *int = nil
	var interfaceWithNilPointer interface{} = nilPointer
	fmt.Printf("interfaceWithNilPointer == nil: %v\n", interfaceWithNilPointer == nil)
	// This will be FALSE! Because interface{} contains (type=*int, value=nil)
	// Since the type is populated, the interface itself is NOT nil!
	// This is a classic Go gotcha and common interview question.
}
