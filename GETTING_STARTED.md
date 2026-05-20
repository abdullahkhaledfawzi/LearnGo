# 🚀 Getting Started - Go Exam Prep

## Welcome! 👋

This project contains a **comprehensive Go MCQ exam preparation program** featuring heavily detailed practical examples for every core concept.

---

## 📁 Project Structure

```text
go-exam-prep/
├── stage1_types_and_variables/     # Stage 1: Types & Variables
│   ├── 1_variables.go              # Declarations and scope
│   ├── 2_zero_values.go            # Zero values
│   ├── 3_type_conversions.go       # Type conversions
│   └── 4_string_bytes_runes.go     # Bytes and Runes
│
├── stage2_control_flow/            # Stage 2: Control Flow
│   ├── 1_if_switch.go              # If/Else and Switch (Fallthrough)
│   └── 2_loops_range.go            # For and Range loops
│
├── stage3_data_structures/         # Stage 3: Data Structures
│   ├── 1_arrays_slices.go          # Arrays vs Slices
│   ├── 2_maps.go                   # Maps
│   └── 3_structs.go                # Structs and Composition
│
├── stage4_functions_pointers/      # Stage 4: Functions & Pointers
│   ├── 1_functions.go              # Functions and closures
│   ├── 2_pointers.go               # Pointers
│   └── 3_methods.go                # Methods and receivers
│
├── stage5_error_handling/          # Stage 5: Error Handling
│   ├── 1_error_handling.go         # Error interface
│   ├── 2_defer.go                  # Defer statement
│   └── 3_panic_recover.go          # Panic and Recover
│
├── stage6_interfaces_concurrency/  # Stage 6: Interfaces & Concurrency
│   ├── 1_interfaces.go             # Implicit Interfaces
│   └── 2_goroutines_channels.go    # Goroutines and Channels
│
├── stage7_stdlib/                  # Stage 7: Standard Library
│   ├── 1_fmt_strings_strconv.go    # fmt, strings, strconv
│   ├── 2_io_os_json.go             # io, os, encoding/json
│   └── 3_misc_stdlib.go            # time, math, rand
│
├── README.md                       # Project overview
├── START_HERE.md                   # Quick start guide
├── SUMMARY.md                      # Comprehensive summary
├── EXAM_NOTES.md                   # Golden exam tips
├── SAMPLE_MCQ.md                   # Sample MCQ questions
└── go.mod                          # Go module file
```

---

## 🎯 How to Study Effectively

### ✅ Recommended Study Steps:

#### 1️⃣ Stage 1 (Basics) - 1 Day
```bash
cd stage1_types_and_variables
go run 1_variables.go
go run 2_zero_values.go
go run 3_type_conversions.go
go run 4_string_bytes_runes.go
```
**Focus On**:
- `var` vs `:=` declarations.
- Zero values for every type.
- Differences between Bytes (uint8) and Runes (int32).

---

#### 2️⃣ Stage 2 (Control Flow) - 4-6 Hours
```bash
cd stage2_control_flow
go run 1_if_switch.go
go run 2_loops_range.go
```
**Focus On**:
- ⚠️ `fallthrough` in `switch` blocks (Guaranteed Question).
- Iterating over strings using `range`.

---

#### 3️⃣ Stage 3 (Data Structures) - 1 Day
```bash
cd stage3_data_structures
go run 1_arrays_slices.go
go run 2_maps.go
go run 3_structs.go
```
**Focus On**:
- ⚠️ Slice capacity growth and append overflows.
- Writing to a `nil` map (Panics!).
- Exported (Capital) vs Unexported (lowercase) struct fields.

---

#### 4️⃣ Stage 4 (Functions & Pointers) - 1 Day
```bash
cd stage4_functions_pointers
go run 1_functions.go
go run 2_pointers.go
go run 3_methods.go
```
**Focus On**:
- ⚠️ Value vs Pointer Receivers on structs.
- Closures modifying outer variables.
- Auto-conversion of pointers when calling methods.

---

#### 5️⃣ Stage 5 (Error Handling) - 4-6 Hours
```bash
cd stage5_error_handling
go run 1_error_handling.go
go run 2_defer.go
go run 3_panic_recover.go
```
**Focus On**:
- ⚠️ `defer` LIFO order execution.
- Mutating named return values via `defer`.
- `panic` mechanics and `recover()` limitations.

---

#### 6️⃣ Stage 6 (Interfaces & Concurrency) - 2 Days
```bash
cd stage6_interfaces_concurrency
go run 1_interfaces.go
go run 2_goroutines_channels.go
```
**Focus On**:
- ⚠️ Implicit Interface implementation (Duck Typing).
- The "Nil Interface" trap.
- Deadlocks in unbuffered channels.

---

#### 7️⃣ Stage 7 (Standard Library) - 1 Day
```bash
cd stage7_stdlib
go run 1_fmt_strings_strconv.go
go run 2_io_os_json.go
go run 3_misc_stdlib.go
```
**Focus On**:
- ⚠️ JSON Tags (`json:"-"` and `omitempty`).
- String formatting verbs (`%v`, `%T`).

---

## 🔍 How to Retain Knowledge

### 🎓 The Best Learning Loop:

1. **Read the Code and Comments**:
   Read the highly detailed English comments inside the `.go` files carefully.
2. **Run the Code**:
   Execute it to verify the output matches your expectations.
3. **Break the Code**:
   Change values, create panics, or remove pointers to see how the compiler reacts.
4. **Test Yourself**:
   Answer the questions in `SAMPLE_MCQ.md`.

Good luck!
