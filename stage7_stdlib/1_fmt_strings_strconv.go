package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
===============================================
Stage 7: Core Standard Library Packages
Part 1: fmt, strings, and strconv
===============================================

These three packages are fundamental and used in almost every Go program.
*/

/*
===============================================
fmt Package (Formatting and Printing)
===============================================

⚠️ Crucial Concepts:
- `Print`, `Println` (adds newline), `Printf` (formatted).
- `Sprintf`: Returns the formatted string instead of printing it.
- Format Verbs: `%v` (default format), `%T` (type), `%#v` (Go syntax representation).
*/

func fmtDemo() {
	fmt.Println("======== fmt Package ========")

	// 1. Println (with newline)
	fmt.Println("Hello, World!")

	// 2. Print (no newline)
	fmt.Print("Hello ")
	fmt.Print("World\n")

	// 3. Printf (with format verbs)
	name := "Ahmed"
	age := 25
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	// 4. Sprintf (Returns string)
	formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println("Stored via Sprintf:", formatted)

	// 5. Format Verbs (Must Know for MCQ!)
	fmt.Println("\n--- Format Verbs ---")
	value := 42
	floatVal := 3.14159

	fmt.Printf("%%v (Default): %v\n", value)
	fmt.Printf("%%d (Base 10 / Decimal): %d\n", value)
	fmt.Printf("%%o (Base 8 / Octal): %o\n", value)
	fmt.Printf("%%x (Base 16 / Hexadecimal lowercase): %x\n", value)
	fmt.Printf("%%X (Base 16 / Hexadecimal uppercase): %X\n", value)
	fmt.Printf("%%b (Base 2 / Binary): %b\n", value)

	fmt.Printf("%%f (Float): %f\n", floatVal)
	fmt.Printf("%%.2f (Float precision): %.2f\n", floatVal)
	fmt.Printf("%%e (Scientific notation): %e\n", floatVal)

	fmt.Printf("%%s (String): %s\n", "Hello")
	fmt.Printf("%%q (Quoted String): %q\n", "Hello")

	fmt.Printf("%%T (Type of the variable): %T\n", value)

	// 6. %#v (Go syntax representation - excellent for debugging structs!)
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Ali", 30}
	fmt.Printf("%%#v: %#v\n", p)

	// 7. Errorf (Creates an error with formatting)
	err := fmt.Errorf("HTTP error: %d", 404)
	fmt.Println("Errorf generated:", err)
}

/*
===============================================
strings Package
===============================================

⚠️ Essential for all string manipulations. Note that strings in Go are immutable.
Every function here returns a NEW string or value.
*/

func stringsDemo() {
	fmt.Println("\n======== strings Package ========")

	str := "Hello, World!"

	// 1. Contains
	fmt.Printf("Contains('Hello'): %v\n", strings.Contains(str, "Hello")) // true
	fmt.Printf("Contains('foo'): %v\n", strings.Contains(str, "foo"))     // false

	// 2. Index (Returns -1 if not found)
	fmt.Printf("Index('o'): %d\n", strings.Index(str, "o"))     // 4
	fmt.Printf("Index('xyz'): %d\n", strings.Index(str, "xyz")) // -1

	// 3. Count (Counts non-overlapping instances)
	fmt.Printf("Count('l'): %d\n", strings.Count(str, "l")) // 3

	// 4. Split and Join
	fmt.Println("\n--- Split and Join ---")
	csv := "apple,banana,orange"
	fruits := strings.Split(csv, ",") // Returns []string
	fmt.Printf("Split into slice: %v\n", fruits)

	joined := strings.Join(fruits, " | ")
	fmt.Printf("Joined back together: %s\n", joined)

	// 5. ToUpper and ToLower
	fmt.Println("\n--- Case Conversion ---")
	fmt.Printf("ToUpper: %s\n", strings.ToUpper(str))
	fmt.Printf("ToLower: %s\n", strings.ToLower(str))

	// 6. TrimSpace (Removes leading/trailing whitespace, newlines, tabs)
	trimmed := strings.TrimSpace("  \n hello world \t  ")
	fmt.Printf("TrimSpace: '%s'\n", trimmed)

	// 7. Trim (Removes specific leading/trailing characters)
	fmt.Println("\n--- Trim Functions ---")
	trimmed2 := strings.Trim("***hello***", "*")
	fmt.Printf("Trim: %s\n", trimmed2)

	// 8. TrimPrefix and TrimSuffix
	prefixed := strings.TrimPrefix("prefix_data", "prefix_")
	fmt.Printf("TrimPrefix: %s\n", prefixed)

	suffixed := strings.TrimSuffix("data.txt", ".txt")
	fmt.Printf("TrimSuffix: %s\n", suffixed)

	// 9. Replace (n is the number of replacements. -1 means all)
	replaced := strings.Replace(str, "World", "Go", 1)
	fmt.Printf("Replace: %s\n", replaced)

	// 10. ReplaceAll
	replacedAll := strings.ReplaceAll(str, "l", "L")
	fmt.Printf("ReplaceAll: %s\n", replacedAll)

	// 11. HasPrefix and HasSuffix
	fmt.Println("\n--- Prefix/Suffix ---")
	fmt.Printf("HasPrefix('Hello'): %v\n", strings.HasPrefix(str, "Hello")) // true
	fmt.Printf("HasSuffix('!'): %v\n", strings.HasSuffix(str, "!"))         // true

	// 12. Fields (Splits by any whitespace)
	text := "The   quick\nbrown\tfox"
	fields := strings.Fields(text)
	fmt.Printf("Fields (auto-splits on whitespace): %v\n", fields)
}

/*
===============================================
strconv Package (String Conversions)
===============================================

⚠️ Used to convert Strings to numbers/booleans and vice-versa.
*/

func strconvDemo() {
	fmt.Println("\n======== strconv Package ========")

	// 1. Itoa (Integer to ASCII / String)
	num := 42
	str := strconv.Itoa(num)
	fmt.Printf("Itoa(42): %s (type: %T)\n", str, str)

	// 2. Atoi (ASCII to Integer) - Returns (int, error)
	str2 := "123"
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("Atoi('123'): %d\n", num2)
	}

	// 3. FormatInt (Allows converting int64 to a string in a specific base)
	fmt.Println("\n--- FormatInt (Base Conversions) ---")
	num3 := int64(255)
	fmt.Printf("Binary string: %s\n", strconv.FormatInt(num3, 2))
	fmt.Printf("Octal string: %s\n", strconv.FormatInt(num3, 8))
	fmt.Printf("Hex string: %s\n", strconv.FormatInt(num3, 16))

	// 4. ParseInt (Parse string with specific base)
	str3 := "FF"
	parsed, _ := strconv.ParseInt(str3, 16, 64)
	fmt.Printf("ParseInt('FF', base 16): %d\n", parsed)

	// 5. FormatFloat
	fmt.Println("\n--- FormatFloat ---")
	f := 3.14159
	// 'f' = no exponent, 2 = decimal places, 64 = float64
	fmt.Printf("Format float (2 decimals): %s\n", strconv.FormatFloat(f, 'f', 2, 64))

	// 6. ParseFloat
	str4 := "3.14"
	parsedFloat, _ := strconv.ParseFloat(str4, 64)
	fmt.Printf("ParseFloat('3.14'): %f\n", parsedFloat)

	// 7. FormatBool and ParseBool
	fmt.Println("\n--- Boolean Conversions ---")
	fmt.Printf("FormatBool(true): %s\n", strconv.FormatBool(true))
	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool('true'): %v\n", b)
}

func main() {
	fmtDemo()
	stringsDemo()
	strconvDemo()
}
