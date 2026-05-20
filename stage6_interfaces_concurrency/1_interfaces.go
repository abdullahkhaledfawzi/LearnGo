package main

import "fmt"

/*
===============================================
Stage 6: Interfaces and Concurrency - Part 1
Interfaces
===============================================

⚠️ Crucial Concepts (Guaranteed MCQ):
1. An Interface defines a behavior (a set of method signatures). It does NOT dictate how that behavior is implemented or store any state.
2. Implicit Implementation: In Go, you do NOT use the `implements` keyword. If a type has all the methods required by an interface, it automatically implements it! (Duck Typing).
3. The Empty Interface (`interface{}` or `any` in Go 1.18+): It has zero methods. Since every type has at least zero methods, EVERY TYPE satisfies the empty interface.
4. An Interface value under the hood is a tuple of `(Type, Value)`.
5. Interface embedding allows combining smaller interfaces into larger ones (e.g., `io.ReadWriter` combines `io.Reader` and `io.Writer`).

Syntax:
type InterfaceName interface {
    Method1(params) returnType
    Method2(params) returnType
}
*/

// Basic Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implicitly implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implicitly implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func basicInterfaceDemo() {
	fmt.Println("======== Basic Interface ========")

	// Interface variable (Can hold a Rectangle, Circle, or any other Shape)
	var shape Shape

	// Rectangle
	rect := Rectangle{10, 20}
	shape = rect
	fmt.Printf("Rectangle via Shape Interface - Area: %f, Perimeter: %f\n", shape.Area(), shape.Perimeter())

	// Circle
	circle := Circle{5}
	shape = circle
	fmt.Printf("Circle via Shape Interface - Area: %f, Perimeter: %f\n", shape.Area(), shape.Perimeter())

	// Polymorphism: A slice of the Interface type holding different structs
	shapes := []Shape{
		Rectangle{10, 20},
		Circle{5},
		Rectangle{5, 5},
	}

	fmt.Println("\nIterating over all shapes (Polymorphism):")
	for _, s := range shapes {
		fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
	}
}

/*
===============================================
Empty Interface (interface{} or `any`)
===============================================

⚠️ Accepts ANY type. Heavily used in libraries like `fmt.Print` or `encoding/json`.
*/

func emptyInterfaceDemo() {
	fmt.Println("\n======== Empty Interface ========")

	var x interface{} // Can hold anything

	x = 42
	fmt.Printf("x = %v (Dynamic Type: %T)\n", x, x)

	x = "hello"
	fmt.Printf("x = %v (Dynamic Type: %T)\n", x, x)

	x = []int{1, 2, 3}
	fmt.Printf("x = %v (Dynamic Type: %T)\n", x, x)

	x = struct{ Name string }{"Ahmed"}
	fmt.Printf("x = %v (Dynamic Type: %T)\n", x, x)

	printAnything(42)
	printAnything("Go is awesome")
}

func printAnything(x interface{}) {
	fmt.Printf("Received in printAnything: %v (Type: %T)\n", x, x)
}

/*
===============================================
Type Assertions
===============================================

⚠️ How to extract the concrete underlying value from an interface{}.
Syntax: `value, ok := interfaceVar.(TargetType)`
*/

func typeAssertionDemo() {
	fmt.Println("\n======== Type Assertion ========")

	var x interface{} = "hello"

	// 1. Unsafe way (Panics if the type is wrong!)
	// str := x.(string) 
	// fmt.Println(str)

	// 2. Safe way (Comma ok idiom)
	str, ok := x.(string)
	if ok {
		fmt.Printf("Extracted String value: %s\n", str)
	} else {
		fmt.Println("x is not a string")
	}

	// ❌ Panic Example (If uncommented)
	// num := x.(int)  // panic: interface conversion: interface {} is string, not int

	// ✓ Safe Failure Example
	num, ok := x.(int)
	if ok {
		fmt.Printf("Extracted Int value: %d\n", num)
	} else {
		fmt.Println("x is not an int. Type assertion failed safely.")
	}
}

/*
===============================================
Type Switch
===============================================

⚠️ A cleaner way to perform multiple Type Assertions.
*/

func typeSwitchDemo() {
	fmt.Println("\n======== Type Switch ========")

	values := []interface{}{
		42,
		"hello",
		3.14,
		true,
		Rectangle{10, 20},
		nil,
	}

	for _, v := range values {
		switch val := v.(type) { // Special .(type) syntax ONLY works in switch!
		case int:
			fmt.Printf("Type is int, value: %d\n", val)
		case string:
			fmt.Printf("Type is string, value: %s\n", val)
		case float64:
			fmt.Printf("Type is float64, value: %f\n", val)
		case bool:
			fmt.Printf("Type is bool, value: %v\n", val)
		case Shape:
			fmt.Printf("Type satisfies Shape interface, Area: %f\n", val.Area())
		case nil:
			fmt.Println("Type is nil")
		default:
			fmt.Printf("Unknown type for value: %v\n", v)
		}
	}
}

/*
===============================================
Reader and Writer Interfaces (Standard Library)
===============================================

⚠️ The most famous interfaces in Go are from the `io` package.
*/

// Simulated io.Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Simulated io.Writer
type Writer interface {
	Write(p []byte) (n int, err error)
}

// Interface Embedding: Combines Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

func embeddingInterfacesDemo() {
	fmt.Println("\n======== Embedding Interfaces ========")
	fmt.Println("ReadWriter interface successfully embeds Reader and Writer.")
}

/*
===============================================
nil Interface Values (The Trap!)
===============================================

⚠️ Tricky MCQ: A typed nil pointer inside an interface makes the interface NOT nil!
*/

func nilInterfaceDemo() {
	fmt.Println("\n======== nil Interface Trap ========")

	var rect *Rectangle = nil // A typed pointer that is nil
	var shape Shape = rect    // Assigning the typed nil pointer to the interface

	// WHY IS THIS FALSE?
	// Because `shape` now holds (Type=*Rectangle, Value=nil).
	// An interface is ONLY nil if BOTH Type and Value are nil!
	fmt.Printf("Is shape == nil? %v\n", shape == nil) // FALSE!

	// Calling a method on it might panic if the method tries to access struct fields.
	// fmt.Println(shape.Area()) // PANIC: nil pointer dereference
}

/*
===============================================
Interface Satisfaction
===============================================
*/

type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

func (d Dog) Move() string {
	return d.Name + " runs"
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return b.Name + " says: Tweet!"
}

func (b Bird) Move() string {
	return b.Name + " flies"
}

func satisfactionDemo() {
	fmt.Println("\n======== Interface Satisfaction ========")

	animals := []Animal{
		Dog{"Rex"},
		Bird{"Tweety"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
	}
}

/*
===============================================
Pointer Receivers vs Value Receivers in Interfaces
===============================================

⚠️ Critical Rule:
- If a method has a Pointer Receiver (`*T`), ONLY the Pointer Type (`*T`) implements the interface!
- If a method has a Value Receiver (`T`), BOTH the Value Type (`T`) AND the Pointer Type (`*T`) implement the interface.
*/

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

// Pointer Receiver
func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

func pointerReceiverDemo() {
	fmt.Println("\n======== Pointer Receivers in Interfaces ========")

	person := Person{"Ahmed"}

	// ❌ This will not compile!
	// var g Greeter = person  // ERROR: Person does not implement Greeter (Greet method has pointer receiver)

	// ✓ This works perfectly!
	var g Greeter = &person // Passing the pointer
	g.Greet()
}
