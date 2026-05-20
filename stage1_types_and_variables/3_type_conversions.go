package main

import (
	"fmt"
	"strconv"
)

/*
===============================================
Type Conversions
===============================================

⚠️ Extremely important point for Go exams:
Go DOES NOT have implicit casting (Implicit Conversion).
You MUST explicitly convert types (Explicit Conversion).

Unlike C or Java, where an int can silently become a float.

General syntax: targetType(value)
Example: int(3.14), string(65), int32(100)
*/

func basicTypeConversions() {
	fmt.Println("======== Basic Conversions ========")

	// 1. Conversions between numeric types
	var intVal int = 100
	var int32Val int32 = int32(intVal)
	var uint8Val uint8 = uint8(intVal)

	fmt.Printf("int to int32: %d\n", int32Val)
	fmt.Printf("int to uint8: %d\n", uint8Val)

	// ⚠️ DANGER: Overflow
	// If you try to convert a large value into a smaller type, it silently overflows.
	var largeInt int = 300
	var smallUint uint8 = uint8(largeInt)                        // Takes the last 8 bits only
	fmt.Printf("int(300) to uint8: %d (overflow!)\n", smallUint) // Output: 44 (300 % 256)

	// 2. Convert Float to Int
	var floatVal float64 = 3.99
	var intFromFloat int = int(floatVal)                                  // Truncates decimals, DOES NOT round
	fmt.Printf("float64(3.99) to int: %d (Decimals truncated)\n", intFromFloat) // 3

	// 3. Convert Int to Float
	var intVal2 int = 10
	var floatFromInt float64 = float64(intVal2)
	fmt.Printf("int(10) to float64: %f\n", floatFromInt)

	// ⚠️ MCQ Question: Float precision & accuracy loss
	// Converting a large int64 to float64 might lose precision!
	var largeInt64 int64 = 9223372036854775807
	var floatLarge float64 = float64(largeInt64)
	fmt.Printf("int64 precision loss: %d becomes %f\n", largeInt64, floatLarge)
}

/*
===============================================
String Conversions
===============================================

This is a deep topic with many MCQ questions!

1. string ← []byte (Slice of bytes)
2. string ← []rune (Slice of Runes)
3. Using strconv package for number/string conversions
*/

func stringConversions() {
	fmt.Println("\n======== String Conversions ========")

	// 1. String to []byte
	// This allocates a new byte slice and copies the underlying string data
	str := "Hello"
	bytesSlice := []byte(str)
	fmt.Printf("String to []byte: %v\n", bytesSlice) // [72 101 108 108 111]
	// Each ASCII character is converted to its byte value

	// 2. []byte to String
	// Also allocates a new string and copies data
	byteArr := []byte{72, 101, 108, 108, 111}
	newStr := string(byteArr)
	fmt.Printf("[]byte to String: %s\n", newStr) // Hello

	// 3. String to []rune (Dealing with Unicode)
	arabicStr := "مرحبا"
	runesSlice := []rune(arabicStr)
	fmt.Printf("String to []rune: %v\n", runesSlice)
	fmt.Printf("Number of Runes: %d\n", len(runesSlice)) // 5
	fmt.Printf("Number of Bytes: %d\n", len(arabicStr))  // 10
	// ⚠️ Difference: Non-ASCII characters take 2-4 bytes. Length of string is BYTES.

	// 4. []rune to String
	newRunes := []rune{'W', 'o', 'r', 'l', 'd'}
	newStrFromRunes := string(newRunes)
	fmt.Printf("[]rune to String: %s\n", newStrFromRunes) // World

	// ⚠️ Tricky MCQ Examples
	fmt.Println("\n--- Tricky Examples ---")

	// Example 1: The Wrong Conversion (Integer directly to String)
	wrongConversion := string(65)                   // Converts the number to a Rune/Character!
	fmt.Printf("string(65): %s\n", wrongConversion) // Output: A (ASCII code 65)

	// If you wanted "65" (the string representation):
	correctConversion := "65"
	fmt.Printf("\"65\": %s\n", correctConversion)

	// Example 2: Converting a single Rune
	var singleRune rune = 'A'
	str2 := string(singleRune)
	fmt.Printf("string('A'): %s\n", str2) // A

	// Example 3: Converting a number to text
	number := 123
	// strNumber := string(number) // Allowed, but returns character represented by 123
	// Correct way to get "123":
	strNumber2 := strconv.Itoa(123)
	fmt.Printf("strconv.Itoa(123): %s\n", strNumber2) // "123"
}

/*
===============================================
strconv Package Conversions
===============================================

strconv provides functions to convert between strings and basic data types.

Important Functions:
- Itoa(i int) -> string (Integer to ASCII)
- Atoi(s string) -> (int, error) (ASCII to Integer)
- FormatInt(i int64, base int) -> string
- ParseInt(s string, base int, bitSize int) -> (int64, error)
- FormatFloat(f float64, fmt byte, prec int, bitSize int) -> string
- ParseFloat(s string, bitSize int) -> (float64, error)
*/

func strconvDemo() {
	fmt.Println("\n======== strconv Package ========")

	// 1. Int to String
	num := 42
	str := strconv.Itoa(num)
	fmt.Printf("Itoa(42): %s (type: %T)\n", str, str)

	// 2. String to Int
	str2 := "123"
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Atoi(\"123\"): %d\n", num2)

	// ⚠️ Error Handling
	invalidStr := "abc"
	num3, err := strconv.Atoi(invalidStr)
	if err != nil {
		fmt.Printf("Atoi(\"abc\"): Error: %v\n", err)
		fmt.Printf("Returned Value on error: %d (Zero value)\n", num3) // 0
	}

	// 3. Int to String with different bases
	fmt.Println("\n--- FormatInt (Different bases) ---")
	num4 := int64(255)
	binary := strconv.FormatInt(num4, 2)   // base 2 (binary)
	octal := strconv.FormatInt(num4, 8)    // base 8 (octal)
	decimal := strconv.FormatInt(num4, 10) // base 10
	hex := strconv.FormatInt(num4, 16)     // base 16 (hexadecimal)

	fmt.Printf("255 in binary: %s\n", binary)   // 11111111
	fmt.Printf("255 in octal: %s\n", octal)     // 377
	fmt.Printf("255 in decimal: %s\n", decimal) // 255
	fmt.Printf("255 in hex: %s\n", hex)         // ff

	// 4. Float to String
	fmt.Println("\n--- Float Conversion ---")
	floatNum := 3.14159
	// 'f' = no exponent, 2 = precision, 64 = float64
	strFloat := strconv.FormatFloat(floatNum, 'f', 2, 64)
	fmt.Printf("FormatFloat(3.14159, 'f', 2, 64): %s\n", strFloat) // 3.14

	// 5. String to Float
	strFloat2 := "3.14"
	floatNum2, err := strconv.ParseFloat(strFloat2, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("ParseFloat(\"3.14\", 64): %f\n", floatNum2)

	// 6. Bool conversion
	fmt.Println("\n--- Bool Conversion ---")
	boolStr := strconv.FormatBool(true)
	fmt.Printf("FormatBool(true): %s\n", boolStr) // "true"

	parsedBool, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool(\"true\"): %v\n", parsedBool) // true
	
	parsedBoolFalse, _ := strconv.ParseBool("0") // "1", "t", "T", "true", "TRUE", "True" -> true
	fmt.Printf("ParseBool(\"0\"): %v\n", parsedBoolFalse) // false
}

/*
===============================================
Type Assertions (Dynamic Type Checking)
===============================================

Checking the dynamic type of an interface{} at runtime.
Will be deeply covered in interfaces, but essential to know.
*/

func typeAssertionSimple() {
	fmt.Println("\n======== Type Assertions Introduction ========")

	// Variable of type interface{} (can hold anything)
	var x interface{} = "hello"

	// Type assertion to extract the underlying string
	// If you omit `ok` and the assertion fails, it will PANIC!
	s, ok := x.(string)
	if ok {
		fmt.Printf("x is a string: %s\n", s)
	}

	// ⚠️ Panic edge case
	// wrongStr := x.(int) // This would PANIC because x is not an int

	// Safe check using switch
	switch v := x.(type) {
	case string:
		fmt.Printf("x is string: %s\n", v)
	case int:
		fmt.Printf("x is int: %d\n", v)
	default:
		fmt.Printf("x is unknown type\n")
	}
}
