# Important Notes for the Go MCQ Exam

## Golden Tips for the Exam

### 1️⃣ Edge Cases
- **Zero Values**: `0` for numbers, `""` for strings, `nil` for pointers, maps, slices, interfaces, and channels.
- **Slice vs Array**: Slices are reference descriptors, Arrays are fixed-size value types.
- **Append Overflow**: When a slice exceeds its capacity via `append`, a new underlying array is allocated!
- **Range Copies**: The value returned by a `range` loop is a COPY of the element, not a reference to it.

### 2️⃣ MCQ Hot Spots ⚠️

#### Stage 1: Basics
- The difference between `var` and `:=` (Scope and shadowing).
- Zero values (especially `nil` vs an empty struct/slice).
- Byte (uint8) vs Rune (int32) and how Go strings represent UTF-8 encoded text.
- Type conversions (Go does NOT have implicit type conversions!).

#### Stage 2: Control Flow
- **Fallthrough**: Executes the very next case immediately without checking its condition.
- **Range**: Iterating over a string yields the byte index and the `rune` value, NOT the byte value.
- **Labeled Breaks**: Breaking out of outer loops from within nested loops.

#### Stage 3: Data Structures
- **Slice Capacity**: How and when does it change?
- **nil Slice vs Empty Slice**: Appending to a `nil` slice works perfectly.
- **nil Map**: Writing to a `nil` map causes a PANIC! Reading returns a zero value.
- **Visibility**: Exported (Capitalized) vs Unexported (lowercase) identifiers.

#### Stage 4: Functions & Pointers
- **Value vs Pointer Receiver**: Pointer receivers can mutate the original struct; value receivers mutate a copy.
- **Closures**: Anonymous functions capture outer variables by reference, not by value.
- **Auto-Conversion**: Go automatically takes the address or dereferences pointers when calling methods, provided the receiver is addressable.

#### Stage 5: Error Handling
- **Defer**: Executes in LIFO order (Last In, First Out).
- **Named Returns**: Deferred functions can read and MODIFY named return values before they are actually returned!
- **Panic/Recover**: `recover()` ONLY works when called directly inside a deferred function.

#### Stage 6: Interfaces & Concurrency
- **Implicit Interfaces**: No `implements` keyword. If the methods match, it satisfies the interface.
- **nil Interface Trap**: An interface containing a typed `nil` pointer is NOT a `nil` interface!
- **Deadlock**: An unbuffered channel operation blocks indefinitely if there is no corresponding sender/receiver.
- **Channel Closure**: ONLY the sender should ever close a channel.

#### Stage 7: Standard Library
- **JSON Tags**: `json:"-"` ignores a field, `omitempty` drops zero-values. Unexported fields are ALWAYS ignored.
- **Format Verbs**: `%v` (default), `%T` (Type), `%#v` (Go-syntax representation).

### 3️⃣ Typical MCQ Traps

```go
// Trap 1: Append reallocation
s := make([]int, 3, 5)
s = append(s, 1, 2, 3) // Exceeds capacity!
// Capacity is now 10 (or 12 depending on Go version, but definitely > 5)

// Trap 2: Defer LIFO
defer println("1")
defer println("2")
// Outputs: 2 then 1

// Trap 3: Interface nil check
var ptr *int = nil
var i interface{} = ptr
if i == nil { } // FALSE! (The interface has Type=*int, Value=nil. It is not nil itself.)

// Trap 4: Range on maps
for k, v := range myMap { }
// The iteration order is INTENTIONALLY RANDOMIZED on every run.

// Trap 5: Type assertion failure
value := interface{}("hello")
num, ok := value.(int)
// Does NOT panic because we used the 'ok' idiom. 'num' gets zero value (0), 'ok' is false.

// Trap 6: Writing to nil map
var m map[string]int
m["a"] = 1 // PANIC! You must use make(map[string]int)

// Trap 7: String indexing
str := "Hello"
ch := str[0] // Returns a byte (uint8), not a string!
```

### 4️⃣ Ideal Approach to Answering
1. **Read extremely carefully**: A single missing `*` or `&` changes the entire answer.
2. **Check Scope**: Is the variable shadowed? Is it accessible?
3. **Recall Defaults**: Is a struct field zero-valued? Is a slice `nil`?
4. **Mental Execution**: Walk through the code line by line.
5. **Watch for Panics**: Look for nil map writes, out-of-bounds slice indexing, or closed channel writes.

## Good luck! 🚀
