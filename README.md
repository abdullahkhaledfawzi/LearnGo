# Go Exam Preparation - Comprehensive Study Guide

A comprehensive project for Go MCQ exam preparation focusing on:
- Standard Library ONLY (No third-party dependencies)
- Edge Cases and Traps
- Zero Values and Initialization
- Memory Behavior and Pointers

## Project Structure

### Stage 1: Basics and Types (`stage1_types_and_variables/`)
- Variable Declarations and Scope
- Zero Values
- Type Conversions
- Strings, Bytes, and Runes

### Stage 2: Control Flow (`stage2_control_flow/`)
- If/Else Statements
- Switch Statements (with `fallthrough`)
- Loops (`for` and `range`)

### Stage 3: Data Structures (`stage3_data_structures/`)
- Arrays vs. Slices
- Maps
- Structs and Composition

### Stage 4: Functions and Pointers (`stage4_functions_pointers/`)
- Functions and Multiple Returns
- Pointers and Memory Addresses
- Methods (Value vs. Pointer Receivers)

### Stage 5: Error Handling (`stage5_error_handling/`)
- The `error` Interface
- The `defer` Statement
- `panic` and `recover`

### Stage 6: Interfaces and Concurrency (`stage6_interfaces_concurrency/`)
- Interfaces and Implicit Implementation
- Type Assertions
- Goroutines and Channels

### Stage 7: Standard Library (`stage7_stdlib/`)
- `fmt`, `strings`, `strconv`
- `io`, `os`, `encoding/json`
- `time`, `math`, `math/rand`

## How to Use This Repository

To run the examples in any stage, you **must navigate into that stage's directory** and run the entire package together using `go run .`.

Do NOT run individual files (like `go run 1_variables.go`) because the functions are shared across the package.

**Example Usage:**
```bash
# 1. Open your terminal and navigate to the project directory (LearnGo)

# 2. Enter a specific stage
cd stage1_types_and_variables

# 3. Run all files in that stage together
go run .
```

After running the code, open the files in your editor and read the detailed English comments to understand the "Why" behind the code.

## Essential Files
- [**EXAM_NOTES.md**](./EXAM_NOTES.md) - Golden tips, edge cases, and traps. Review this right before the exam.
- [**SAMPLE_MCQ.md**](./SAMPLE_MCQ.md) - Practice questions with explanations. Test your knowledge.
