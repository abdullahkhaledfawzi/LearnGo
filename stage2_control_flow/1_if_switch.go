package main

import "fmt"

/*
===============================================
Stage 2: Control Flow
Part 1: If/Else Statements
===============================================

⚠️ Important Go concepts:
1. No parentheses are needed around the condition (if condition {).
2. The curly braces {} are MANDATORY, even for a single line.
3. You can declare and initialize a variable inside the if statement (scoped only to the if/else blocks).
4. There is no Ternary Operator (x ? y : z) in Go. You must use full if/else.
*/

func basicIfElse() {
	fmt.Println("======== If/Else Basics ========")

	// 1. Simple syntax
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// 2. if/else
	y := 3
	if y > 5 {
		fmt.Println("y is greater than 5")
	} else {
		fmt.Println("y is less than or equal to 5")
	}

	// 3. if/else if/else
	age := 25
	if age < 13 {
		fmt.Println("Child")
	} else if age < 18 {
		fmt.Println("Teenager")
	} else if age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	// ⚠️ MCQ Question: Variable scoped to if statement
	// The variable 'score' is created and evaluated in the same line.
	// It is ONLY available inside the 'if' and 'else' blocks.
	if score := calculateScore(); score > 80 {
		fmt.Printf("Score: %d - Pass!\n", score)
	} else {
		fmt.Printf("Score: %d - Fail\n", score)
	}

	// 'score' is NOT available here (out of scope)
	// fmt.Println(score)  // ❌ Error: undefined: score
}

func calculateScore() int {
	return 85
}

/*
===============================================
Switch Statements
===============================================

⚠️ Extremely Important (Common MCQ Topics):
1. No `break` statement is needed. It breaks automatically at the end of each case!
2. You must use `fallthrough` explicitly to continue to the next case.
3. You can have multiple values in a single case separated by commas.
4. You can write a switch WITHOUT a condition (acts like a long if/else if chain).
5. Switching on `nil` is perfectly legal!
*/

func switchBasics() {
	fmt.Println("\n======== Switch Basics ========")

	// 1. Standard syntax
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// 2. Multiple values in a single case
	char := 'A'
	switch char {
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Printf("%c is a Vowel\n", char)
	case 'B', 'C', 'D':
		fmt.Printf("%c is a Consonant\n", char)
	default:
		fmt.Printf("%c is Unknown\n", char)
	}

	// 3. Default case can be anywhere! (Order doesn't matter for default)
	status := "unknown"
	switch status {
	default:
		fmt.Println("Status Unknown")
	case "active":
		fmt.Println("Active")
	case "inactive":
		fmt.Println("Inactive")
	}
	// Output will be: Status Unknown (even though default is at the top)
}

/*
===============================================
⚠️ Fallthrough (Extremely popular MCQ!)
===============================================

In C/Java, cases fall through by default unless you write `break`.
In Go, it's the EXACT OPPOSITE. They break by default.
You must explicitly write `fallthrough` to jump to the next case.
*/

func fallthroughDemo() {
	fmt.Println("\n======== Fallthrough ========")

	// Example 1: Without fallthrough
	fmt.Println("--- Without fallthrough ---")
	x := 1
	switch x {
	case 1:
		fmt.Println("One")
		// Automatically breaks here
	case 2:
		fmt.Println("Two")
	}
	// Output: One

	// Example 2: With fallthrough
	fmt.Println("\n--- With fallthrough ---")
	y := 1
	switch y {
	case 1:
		fmt.Println("One")
		fallthrough // Executes the body of the NEXT case immediately!
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		// Breaks here
	}
	// Output:
	// One
	// Two
	// Three

	// ⚠️ Tricky Edge Case
	fmt.Println("\n--- Tricky Edge Case ---")
	num := 2
	switch num {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3") // This gets executed because of fallthrough
	default:
		fmt.Println("Default Case")
	}
	// Output:
	// Case 2
	// Case 3
	// (It does NOT print "Default Case" because there's no fallthrough in case 3)
}

/*
===============================================
Switch Without Condition (Acts as if/else chains)
===============================================
*/

func switchWithoutCondition() {
	fmt.Println("\n======== Switch Without Condition ========")

	score := 85
	switch { // No condition here
	case score >= 90:
		fmt.Println("Grade: Excellent")
	case score >= 80:
		fmt.Println("Grade: Very Good")
	case score >= 70:
		fmt.Println("Grade: Good")
	case score >= 60:
		fmt.Println("Grade: Pass")
	default:
		fmt.Println("Grade: Fail")
	}
	// Output: Grade: Very Good
}

/*
===============================================
Type Switch (To be deeply covered later, but good for context)
===============================================
Used to determine the dynamic type of an interface{} variable.
*/

func typeSwitchIntro() {
	fmt.Println("\n======== Type Switch (Intro) ========")

	var value interface{} = "hello"

	switch v := value.(type) { // Special syntax .(type) only works in switch
	case string:
		fmt.Printf("String: %s (Length: %d)\n", v, len(v)) // v is strongly typed as string here
	case int:
		fmt.Printf("Int: %d\n", v) // v is strongly typed as int here
	case bool:
		fmt.Printf("Bool: %v\n", v) // v is strongly typed as bool here
	default:
		fmt.Printf("Unknown type\n")
	}
	// Output: String: hello (Length: 5)
}

/*
===============================================
Switch on nil
===============================================
*/

func switchOnNil() {
	fmt.Println("\n======== Switch on nil ========")

	var ptr *int = nil

	switch ptr {
	case nil:
		fmt.Println("Pointer is nil")
	default:
		fmt.Println("Pointer is NOT nil")
	}

	// Valid with interfaces too
	var i interface{}
	switch i {
	case nil:
		fmt.Println("interface{} is nil")
	}
}
