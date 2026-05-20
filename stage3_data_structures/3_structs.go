package main

import "fmt"

/*
===============================================
Stage 3: Data Structures - Part 3
Structs
===============================================

⚠️ Critical Struct Concepts:
1. Structs are VALUE TYPES. When passed to a function, the ENTIRE struct is copied.
   - Use pointers (`*Struct`) to avoid copying large structs and to allow modification.
2. Exported vs Unexported (Visibility Rules):
   - Fields starting with an UPPERCASE letter are Exported (Public to other packages).
   - Fields starting with a lowercase letter are Unexported (Private to the current package).
3. Go has NO INHERITANCE (no `extends` or `implements` keywords).
   - Instead, Go uses "Composition" (Embedding structs within structs).
4. Struct Tags: Used heavily with libraries (e.g., JSON, ORMs) to define metadata via Reflection.
5. Empty Struct (`struct{}`): Takes exactly 0 bytes of memory. Used for signals or Sets.
*/

// Struct Definition
type Person struct {
	Name string // Exported (Starts with capital letter)
	Age  int    // Exported
	City string // Exported
}

// Mixed Visibility Struct
type Employee struct {
	ID       int     // Exported
	name     string  // Unexported (Only accessible within the `main` package)
	salary   float64 // Unexported
	Position string  // Exported
}

func structBasics() {
	fmt.Println("======== Structs Basics ========")

	// 1. Create instance (Initialized with Zero Values)
	var p1 Person
	fmt.Printf("Zero value struct: %+v\n", p1) // {Name:"" Age:0 City:""}

	// 2. Positional Initialization
	// Must provide ALL fields in the exact order they are defined. Fragile!
	p2 := Person{"Ahmed", 25, "Cairo"}
	fmt.Printf("Positional init: %+v\n", p2)

	// 3. Named Fields Initialization (Highly Recommended)
	// Robust against adding new fields to the struct later.
	p3 := Person{
		Name: "Fatima",
		Age:  23,
		City: "Alexandria", // Trailing comma is required in multi-line formatting!
	}
	fmt.Printf("Named init: %+v\n", p3)

	// ⚠️ Omitted fields get Zero Values
	p4 := Person{
		Name: "Ali",
	}
	fmt.Printf("Partial init: %+v\n", p4) // {Name:Ali Age:0 City:}

	// 4. Accessing and Modifying Fields via dot notation
	fmt.Println("\n--- Access and Modify ---")
	fmt.Printf("Name: %s\n", p2.Name)
	p2.Age = 26
	fmt.Printf("After Modification: %+v\n", p2)
}

/*
===============================================
Pointers to Structs
===============================================

Go allows accessing struct fields through a pointer seamlessly.
You do NOT need to write `(*p).field` like in C/C++.
Go does Implicit Dereferencing: `p.field` works perfectly.
*/

func pointersToStructs() {
	fmt.Println("\n======== Pointers to Structs ========")

	p := Person{"Mohammed", 30, "Riyadh"}

	// 1. Create a pointer
	ptr := &p
	fmt.Printf("Pointer: %p -> Value: %+v\n", ptr, *ptr)

	// 2. Access via pointer (Implicit Dereferencing)
	fmt.Printf("ptr.Name: %s\n", ptr.Name)       // Go automatically does (*ptr).Name
	fmt.Printf("(*ptr).Name: %s\n", (*ptr).Name) // Explicit way (rarely used)

	// 3. Modifying via pointer changes the original struct
	ptr.Age = 31
	fmt.Printf("After pointer modification, original struct: %+v\n", p)

	// ⚠️ Common Idiom: Returning a pointer to a struct literal
	// This allocates the struct and returns its memory address.
	ptr2 := &Person{
		Name: "Sarah",
		Age:  28,
		City: "Jeddah",
	}
	fmt.Printf("Pointer to struct literal: %+v\n", ptr2)
}

/*
===============================================
Exported vs Unexported Fields
===============================================

Rule of thumb: UPPERCASE = Public, lowercase = Private.
This applies to Structs themselves, Functions, Variables, AND Struct Fields!
*/

func exportedVsUnexported() {
	fmt.Println("\n======== Exported vs Unexported ========")

	emp := Employee{
		ID:       101,
		name:     "Ahmed", // Valid here because we are in the SAME package (`main`)
		salary:   5000.0,
		Position: "Engineer",
	}
	fmt.Printf("Employee: %+v\n", emp)

	// If we were in a different package (e.g., package `hr`),
	// trying to access `emp.name` or `emp.salary` would result in a Compilation Error:
	// "emp.name undefined (cannot refer to unexported field or method name)"
}

/*
===============================================
Embedded Structs (Composition over Inheritance)
===============================================

Go does not have classes or inheritance. Instead, it uses Embedding.
*/

type Address struct {
	Street string
	City   string
	ZIP    string
}

type Employee2 struct {
	ID      int
	Name    string
	Address Address // Named Embedding (Composition)
}

type Employee3 struct {
	ID      int
	Name    string
	Address // Anonymous Embedding (Promotes fields)
}

func embeddedStructs() {
	fmt.Println("\n======== Embedded Structs ========")

	// 1. Named Embedding
	emp1 := Employee2{
		ID:   101,
		Name: "Ahmed",
		Address: Address{
			Street: "Tahrir St",
			City:   "Cairo",
			ZIP:    "11111",
		},
	}
	fmt.Printf("Named Embedded City: %s\n", emp1.Address.City)

	// 2. Anonymous Embedding (Field Promotion)
	emp2 := Employee3{
		ID:   102,
		Name: "Fatima",
		Address: Address{
			Street: "Nile St",
			City:   "Alexandria",
			ZIP:    "21111",
		},
	}
	
	// ⚠️ Field Promotion Feature:
	// Because `Address` was embedded anonymously, its fields are "promoted" to the parent struct!
	fmt.Printf("Promoted City access: %s\n", emp2.City)      // Direct access!
	fmt.Printf("Explicit City access: %s\n", emp2.Address.City) // Still works
}

/*
===============================================
Struct Tags (Metadata)
===============================================

Used extensively for JSON, XML, Database mapping (GORM), and Validation.
Tags are read at runtime using the `reflect` package.
*/

type Product struct {
	ID    int     `json:"id" db:"product_id"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price,omitempty"` // omitempty: omit from JSON if zero value
	
	// ⚠️ Unexported fields CANNOT be marshaled to JSON, even with tags!
	internalCode string `json:"internal_code"` // This will be IGNORED by encoding/json
}

func structTags() {
	fmt.Println("\n======== Struct Tags ========")

	prod := Product{
		ID:    1,
		Name:  "Laptop",
		Price: 999.99,
		internalCode: "XYZ123",
	}
	fmt.Printf("Product Struct: %+v\n", prod)
	// Note: Tags don't affect regular Go code behavior, they are read by external libraries.
}

/*
===============================================
Anonymous Structs
===============================================

Useful for one-off data structures, especially in unit tests (Table-Driven Tests)
or parsing a specific JSON payload without defining a global struct type.
*/

func anonymousStructs() {
	fmt.Println("\n======== Anonymous Structs ========")

	// Defining and initializing instantly
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Anonymous struct: %+v\n", config)
}

/*
===============================================
Empty Struct struct{} (Advanced Memory Topic)
===============================================

An empty struct uses exactly 0 bytes of memory.
Used for:
1. Channels acting strictly as signals (no data transmission needed).
2. Implementing "Set" data structures using maps (`map[string]struct{}`).
*/

func emptyStruct() {
	fmt.Println("\n======== Empty Struct ========")

	// Uses 0 bytes
	var empty struct{}
	fmt.Printf("Empty struct value: %v\n", empty)

	// Building a Set (unique values only, no memory wasted on map values)
	setOfNumbers := make(map[int]struct{})
	
	// The `struct{}{}` is the instantiation of the empty struct
	setOfNumbers[1] = struct{}{} 
	setOfNumbers[2] = struct{}{}

	if _, exists := setOfNumbers[1]; exists {
		fmt.Println("1 exists in the Set!")
	}
}

/*
===============================================
Struct Methods (Introduction)
===============================================
Methods are functions with a special "receiver" argument.
*/

// Value Receiver (Operates on a COPY of the struct)
// Safe, but modifications won't affect the original struct.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s", p.Name)
}

// Pointer Receiver (Operates on the ORIGINAL struct memory)
// Use this if you need to modify the struct or if the struct is very large.
func (p *Person) HaveBirthday() {
	p.Age++
}

func structMethods() {
	fmt.Println("\n======== Struct Methods (Intro) ========")

	p := Person{"Mohammed", 25, "Dammam"}

	fmt.Println(p.Greet())

	p.HaveBirthday() // Modifies the actual struct
	fmt.Printf("After Birthday: %+v\n", p)
}
