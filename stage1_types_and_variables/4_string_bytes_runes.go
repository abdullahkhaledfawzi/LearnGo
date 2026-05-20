package main

import "fmt"

/*
===============================================
Bytes and Runes
===============================================

Highly popular MCQ topic in Go exams!

Core Difference:
- byte: alias for uint8 (represents a value from 0 to 255)
- rune: alias for int32 (represents any Unicode code point)

When to use which:
- byte: Binary data, reading files/networks, ASCII text.
- rune: Text that might contain non-ASCII characters (Arabic, Chinese, Emojis, etc.)
*/

func bytesBasics() {
	fmt.Println("======== Bytes Basics ========")

	// 1. Declaring a single Byte
	var b byte = 65 // 'A' in ASCII
	fmt.Printf("byte value: %d -> character: %c\n", b, b)

	// 2. Byte is just uint8
	var u uint8 = 65
	fmt.Printf("uint8 value: %d -> character: %c\n", u, u)

	// 3. Slice of bytes
	myBytes := []byte{72, 101, 108, 108, 111}
	fmt.Printf("Byte slice: %v\n", myBytes)

	// 4. String to []byte
	str := "Go"
	bytesFromStr := []byte(str)
	fmt.Printf("String '%s' to bytes: %v\n", str, bytesFromStr)
	// G = 71, o = 111

	// 5. []byte to String
	reconstructed := string(myBytes)
	fmt.Printf("Bytes back to string: %s\n", reconstructed)

	// ⚠️ Tricky Examples
	fmt.Println("\n--- Tricky Examples ---")

	// Example 1: Direct number to string (Converts to Rune!)
	wrongWay := string(65)
	fmt.Printf("string(65): %s\n", wrongWay) // A

	// Correct way to build a string from ASCII byte value
	correctWay := string([]byte{65})
	fmt.Printf("string([]byte{65}): %s\n", correctWay) // A

	// Example 2: Byte slices for file IO
	// Standard in Go: when reading from file/network, you deal with []byte
	data := []byte("Hello, World!")
	fmt.Printf("Data from file (simulated): %v\n", data)
	fmt.Printf("As string: %s\n", string(data))
}

/*
===============================================
Runes (Unicode Characters)
===============================================

A rune is an int32, perfectly capable of representing any Unicode character.
*/

func runesBasics() {
	fmt.Println("\n======== Runes Basics ========")

	// 1. Declaring a single Rune
	var r rune = 'A'                                      // Single quotes for rune literals
	fmt.Printf("rune value: %d -> character: %c\n", r, r) // 65 -> A

	// 2. Rune with Unicode characters
	var arabicRune rune = 'أ'
	fmt.Printf("Arabic rune: %d -> character: %c\n", arabicRune, arabicRune)

	var emojiRune rune = '😀'
	fmt.Printf("Emoji rune: %d -> character: %c\n", emojiRune, emojiRune)

	// 3. String to []rune
	arabicStr := "السلام عليكم"
	runes := []rune(arabicStr)
	fmt.Printf("String '%s' to runes: %v\n", arabicStr, runes)
	fmt.Printf("Character count (runes): %d\n", len(runes))
	fmt.Printf("Byte count: %d\n", len(arabicStr))
	// ⚠️ VERY IMPORTANT: len(string) returns the number of BYTES, NOT characters!

	// 4. []rune back to String
	runesArray := []rune{'H', 'e', 'l', 'l', 'o'}
	reconstructed := string(runesArray)
	fmt.Printf("Runes back to string: %s\n", reconstructed)

	// ⚠️ Common MCQ Question: Iterating over a string
	fmt.Println("\n--- Iterate over String ---")
	str := "Hello"

	// Method 1: for ... range (Yields byte index and RUNE!)
	fmt.Println("Using for ... range on string (yields Runes):")
	for i, ch := range str {
		fmt.Printf("Index %d: %c (rune value: %d)\n", i, ch, ch)
	}

	// Method 2: Accessing as []byte
	fmt.Println("\nUsing string as []byte (yields Bytes):")
	bytes := []byte(str)
	for i, b := range bytes {
		fmt.Printf("Index %d: %c (byte value: %d)\n", i, b, b)
	}

	// ⚠️ Critical distinction with non-ASCII text
	fmt.Println("\n--- Tricky Arabic Text Edge Case ---")
	arabicWord := "مرحبا"
	fmt.Printf("Original string: %s\n", arabicWord)
	fmt.Printf("String length (bytes): %d\n", len(arabicWord))
	fmt.Printf("Rune count: %d\n", len([]rune(arabicWord)))

	fmt.Println("Characters when ranging:")
	for i, ch := range arabicWord {
		// Notice the index jumps by 2 for each Arabic character (they are 2 bytes each)
		fmt.Printf("Byte Index %d: %c\n", i, ch)
	}
}

/*
===============================================
Zero Values and Pointers
===============================================
*/

func zeroValuesBytes() {
	fmt.Println("\n======== Zero Values for Bytes & Runes ========")

	var b byte
	var r rune

	fmt.Printf("Zero byte: %d\n", b)
	fmt.Printf("Zero rune: %d\n", r)

	// ⚠️ Difference between byte(0) and nil
	// byte is a numeric type, so it CANNOT be nil.
	// But a pointer to a byte (*byte) CAN be nil.
	var pointerToByte *byte = nil
	fmt.Printf("*byte zero value: %v (nil)\n", pointerToByte)
}

/*
===============================================
Type Aliases Under the Hood
===============================================

byte = uint8
rune = int32
*/

func typeAliasesDemo() {
	fmt.Println("\n======== Type Aliases ========")

	// byte and uint8 are identical
	var b byte = 65
	var u uint8 = 65

	fmt.Printf("byte: %T\n", b)                // uint8
	fmt.Printf("uint8: %T\n", u)               // uint8
	fmt.Printf("Are they equal? %v\n", b == u) // true

	// You can assign them directly without conversion
	b = u
	fmt.Printf("b = u: %d\n", b)

	// rune and int32
	var r rune = 1114111 // Max Unicode value
	var i int32 = 1114111

	fmt.Printf("rune: %T\n", r)                // int32
	fmt.Printf("int32: %T\n", i)               // int32
	fmt.Printf("Are they equal? %v\n", r == i) // true
}

/*
===============================================
String Indexing Warning
===============================================

⚠️ Major MCQ Topic!
When you index a string (e.g. str[0]), you get a BYTE, NOT a rune.
*/

func stringIndexing() {
	fmt.Println("\n======== String Indexing ========")

	str := "Hello"
	// str[0] yields a byte (uint8)
	firstChar := str[0]
	fmt.Printf("str[0] type: %T, value: %d, as char: %c\n", firstChar, firstChar, firstChar)

	// To get the first rune, you must convert to []rune or use strings package
	runes := []rune(str)
	firstCharRune := runes[0]
	fmt.Printf("runes[0] type: %T, value: %d, as char: %c\n", firstCharRune, firstCharRune, firstCharRune)

	// ⚠️ Edge Case with Multi-byte Characters
	fmt.Println("\n--- Multi-byte Example ---")
	arabicStr := "مرحبا"

	// ❌ Wrong Way (Indexing gives you 1 byte of a 2-byte character)
	firstByte := arabicStr[0]
	fmt.Printf("arabicStr[0] (byte): %d (This is not the full character 'م')\n", firstByte)

	// ✓ Correct Way (Convert to runes first)
	arabicRunes := []rune(arabicStr)
	firstRune := arabicRunes[0]
	fmt.Printf("First Arabic letter: %c\n", firstRune)
}
