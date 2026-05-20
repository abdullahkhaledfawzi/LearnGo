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

## How to Use

Run files for each stage:
```bash
go run stage1_types_and_variables/main.go
go run stage2_control_flow/main.go
# And so on...
```

Or run a specific file to focus on a single topic:
```bash
go run stage1_types_and_variables/1_variables.go
```
