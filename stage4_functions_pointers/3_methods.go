package main

import "fmt"

/*
===============================================
Stage 4: Functions and Pointers - Part 3
Methods
===============================================

⚠️ Crucial Concepts (Guaranteed MCQ):
1. A Method is simply a function with a special "Receiver" argument.
2. The Receiver appears between the `func` keyword and the method name.
3. You can define methods on ANY user-defined type (Structs, Aliased basic types).
4. Value Receiver: Operates on a COPY of the data. Modifying it does NOT affect the original.
5. Pointer Receiver: Operates on a POINTER to the data. Modifying it affects the original.

Syntax:
func (receiver Type) methodName(params) returnType
*/

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Value Receiver (Safe, Read-Only typically)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Value Receiver - Modification fails to affect original
func (r Rectangle) Scale(factor float64) {
	r.Width *= factor // Only modifies the local copy `r`
	r.Height *= factor
}

// Pointer Receiver - Successfully modifies original
func (r *Rectangle) ScalePointer(factor float64) {
	r.Width *= factor // Modifies the original memory
	r.Height *= factor
}

// Circle methods
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) Grow(amount float64) {
	c.Radius += amount
}

func valueVsPointerReceiver() {
	fmt.Println("======== Value vs Pointer Receiver ========")

	rect := Rectangle{10, 20}
	fmt.Printf("Original rect: %+v\n", rect)
	fmt.Printf("Area: %f\n", rect.Area())

	// Value Receiver
	fmt.Println("\n--- Value Receiver ---")
	rect.Scale(2)
	fmt.Printf("After rect.Scale(2): %+v (Unchanged!)\n", rect)

	// Pointer Receiver
	fmt.Println("\n--- Pointer Receiver ---")
	rect.ScalePointer(2)
	fmt.Printf("After rect.ScalePointer(2): %+v (Changed!)\n", rect)
}

/*
===============================================
Methods on Custom Basic Types
===============================================

You can attach methods to `int`, `float64`, `string`, etc., by aliasing them.
Must be defined in the SAME package as the type definition.
*/

type Distance float64
type Speed float64

// Stringer interface implementation for custom formatting
func (d Distance) String() string {
	return fmt.Sprintf("%.2f km", d)
}

func (s Speed) String() string {
	return fmt.Sprintf("%.2f km/h", s)
}

// Method on Distance
func (d Distance) TimeToTravel(s Speed) float64 {
	return float64(d) / float64(s)
}

func customTypesMethods() {
	fmt.Println("\n======== Methods on Custom Types ========")

	distance := Distance(100)
	speed := Speed(50)

	fmt.Printf("Distance: %s\n", distance)
	fmt.Printf("Speed: %s\n", speed)
	fmt.Printf("Time to travel: %.2f hours\n", distance.TimeToTravel(speed))
}

/*
===============================================
Method Chaining (Fluent Interface)
===============================================

If a method returns its receiver (usually a pointer), you can chain method calls.
Highly common in Builders, ORMs (like GORM), and configuration patterns.
*/

type Builder struct {
	value string
}

func (b *Builder) Add(s string) *Builder {
	b.value += s
	return b // Return the pointer to allow chaining
}

func (b *Builder) AddSpace() *Builder {
	b.value += " "
	return b
}

func (b *Builder) String() string {
	return b.value
}

func methodChainingDemo() {
	fmt.Println("\n======== Method Chaining ========")

	builder := &Builder{}
	result := builder.Add("Hello").AddSpace().Add("World").AddSpace().Add("!")
	fmt.Printf("Chained result: %s\n", result)
}

/*
===============================================
⚠️ MCQ Trap: Auto-conversion of Receivers
===============================================

Go is exceptionally forgiving with method calls to improve developer experience.

1. Value Receiver Method:
   - Call on Value: Valid
   - Call on Pointer: Valid! Go automatically dereferences it `(*ptr).Method()`

2. Pointer Receiver Method:
   - Call on Pointer: Valid
   - Call on Value: Valid! Go automatically takes the address `(&value).Method()`
     - ⚠️ EXCEPTION: Go can only do this if the value is "addressable". You cannot call a pointer method on a literal or unaddressable value.
*/

func autoConversionDemo() {
	fmt.Println("\n======== Auto-conversion by Go ========")

	rect := Rectangle{10, 20} // Value
	rectPtr := &rect          // Pointer

	// Area() has a Value Receiver
	fmt.Printf("rect.Area(): %f\n", rect.Area())       // ✓ Standard
	fmt.Printf("rectPtr.Area(): %f\n", rectPtr.Area()) // ✓ Go Auto-Dereferences

	// ScalePointer() has a Pointer Receiver
	rect.ScalePointer(2)    // ✓ Go Auto-References (&rect)
	rectPtr.ScalePointer(2) // ✓ Standard
	
	// ⚠️ Unaddressable Exception:
	// Rectangle{10, 20}.ScalePointer(2) // ❌ ERROR: Cannot call pointer method on unaddressable value
}

/*
===============================================
Methods vs Functions
===============================================

Under the hood, a method is just a function where the receiver is passed as the first argument.
*/

// Function
func RectangleArea(r Rectangle) float64 {
	return r.Width * r.Height
}

// Method
func (r Rectangle) AreaMethod() float64 {
	return r.Width * r.Height
}

func methodsVsFunctions() {
	fmt.Println("\n======== Methods vs Functions ========")

	rect := Rectangle{10, 20}
	fmt.Printf("Function: %f\n", RectangleArea(rect))
	fmt.Printf("Method: %f\n", rect.AreaMethod())
}
