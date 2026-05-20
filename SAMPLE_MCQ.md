# Quick Guide: Sample MCQs and Explanations

## Stage 1: Types & Variables

### Question 1.1
```text
What is the Zero Value of a slice in Go?
A) []
B) nil
C) [0]
D) error
```
**Answer**: B (nil)
**Explanation**: A declared but uninitialized slice is `nil`, not an empty slice `[]`.

---

### Question 1.2
```go
var ptr *int
fmt.Println(ptr)  // What does this print?
// A) 0
// B) <nil>
// C) null
// D) nil
```
**Answer**: B (`<nil>`)
**Explanation**: The `fmt` package prints uninitialized pointers as `<nil>`.

---

### Question 1.3
```go
str := "مرحبا" // Arabic word
fmt.Println(len(str))  // What does this print?
// A) 4
// B) 8
// C) 10
// D) 5
```
**Answer**: C (10)
**Explanation**: The `len()` function returns the number of BYTES, not the number of characters. Arabic characters (UTF-8) take 2 bytes each. 5 chars * 2 bytes = 10 bytes.

---

## Stage 2: Control Flow

### Question 2.1
```go
x := 2
switch x {
case 1:
    fmt.Print("1")
    fallthrough
case 2:
    fmt.Print("2")
case 3:
    fmt.Print("3")
}
// What is the output?
// A) 2
// B) 23
// C) 123
// D) Error
```
**Answer**: A (2)
**Explanation**: `x == 2`, so it executes `case 2`. There is no `fallthrough` inside `case 2`, so it breaks automatically.

---

### Question 2.2
```go
for i := 0; i < 3; i++ {
    defer fmt.Print(i)
}
// What is the output?
// A) 012
// B) 210
// C) 000
// D) Syntax Error
```
**Answer**: B (210)
**Explanation**: `defer` evaluates the argument `i` immediately during loop iteration, but the execution order is LIFO (Last In, First Out).

---

## Stage 3: Data Structures

### Question 3.1
```go
s := make([]int, 3, 5)
s = append(s, 1, 2, 3)
fmt.Printf("len=%d, cap=%d", len(s), cap(s))
// What is the output?
// A) len=6, cap=5
// B) len=6, cap=10
// C) len=5, cap=5
// D) len=6, cap=6
```
**Answer**: B (len=6, cap=10)
**Explanation**: Appending 3 items to a slice with only 2 remaining slots exceeds its capacity (5). Go allocates a new array and doubles the capacity (to 10).

---

### Question 3.2
```go
var m map[string]int
m["a"] = 1  // What happens here?
// A) It assigns 1 to "a" successfully.
// B) It causes a runtime panic.
// C) Compilation error.
// D) It returns -1.
```
**Answer**: B (runtime panic)
**Explanation**: You cannot write to a `nil` map. You must initialize it with `make()`.

---

## Stage 4: Functions & Pointers

### Question 4.1
```go
func modify(x int) {
    x = 10
}
func main() {
    a := 5
    modify(a)
    fmt.Println(a)
}
// What is the output?
// A) 5
// B) 10
// C) error
```
**Answer**: A (5)
**Explanation**: Go passes arguments by value. `modify` receives a copy of `a`, so the original `a` remains unchanged.

---

### Question 4.2
```go
type S struct {
    x int
}
func (s S) modify() {
    s.x = 10
}
func main() {
    obj := S{5}
    obj.modify()
    fmt.Println(obj.x)
}
// What is the output?
// A) 5
// B) 10
// C) Compilation error
```
**Answer**: A (5)
**Explanation**: `modify` uses a Value Receiver (`s S`). It modifies a copy of the struct, not the original `obj`.

---

## Stage 5: Error Handling

### Question 5.1
```go
defer func() {
    fmt.Print("A")
}()
defer func() {
    fmt.Print("B")
}()
defer func() {
    fmt.Print("C")
}()
// What does this print upon function exit?
// A) ABC
// B) CBA
// C) BCA
```
**Answer**: B (CBA)
**Explanation**: Defers execute in Last-In-First-Out (LIFO) order.

---

## Stage 6: Interfaces & Concurrency

### Question 6.1
```go
type Reader interface {
    Read() string
}
type MyType struct {}
func (m MyType) Read() string { return "done" }

func main() {
    var r Reader = MyType{}  // Is this valid?
}
// A) Yes, implicit implementation.
// B) No, requires 'implements' keyword.
```
**Answer**: A (Yes, implicit implementation)
**Explanation**: `MyType` has the required `Read()` method, so it implicitly satisfies the `Reader` interface.

---

### Question 6.2
```go
ch := make(chan int)
ch <- 1  // What happens here?
fmt.Println("done")
// A) Prints "done"
// B) Deadlock panic
// C) Compilation error
```
**Answer**: B (Deadlock panic)
**Explanation**: It's an unbuffered channel. The `ch <- 1` send operation blocks forever because there is no other goroutine ready to receive the value.

---

## Stage 7: Standard Library

### Question 7.1
```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Pass string `json:"-"`
}
// When marshaled to JSON, which fields are included?
// A) All fields
// B) ID and Name only
// C) ID and Pass only
```
**Answer**: B (ID and Name only)
**Explanation**: The `json:"-"` tag explicitly tells the `encoding/json` package to ignore the `Pass` field entirely.
