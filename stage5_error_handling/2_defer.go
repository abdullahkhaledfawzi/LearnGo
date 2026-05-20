package main

import (
	"fmt"
	"os"
)

/*
===============================================
Stage 5: Error Handling - Part 2
The Defer Statement
===============================================

⚠️ Crucial MCQ Points:
1. `defer` delays the execution of a function until the surrounding function completes (either by returning, or panicking).
2. LIFO Execution: Multiple `defer` statements are executed in Last-In, First-Out order (like a stack).
3. Argument Evaluation: Arguments to a deferred function are evaluated IMMEDIATELY when the defer statement is executed, NOT when the deferred function runs.
4. Panic Immunity: `defer` guarantees execution even if a runtime panic occurs!
5. Named Returns: A deferred function can read and modify named return values before the function actually exits.
*/

func simpleDefer() {
	fmt.Println("======== Simple Defer ========")

	defer fmt.Println("3. Executes LAST (Deferred)")

	fmt.Println("1. Executes FIRST")
	fmt.Println("2. Executes SECOND")
}

/*
===============================================
LIFO (Last In First Out)
===============================================

⚠️ If you have multiple defers, they stack up. The last one declared runs first.
*/

func multipleDeferDemo() {
	fmt.Println("\n======== Multiple Defers (LIFO) ========")

	defer fmt.Println("1. First defer (Runs 3rd)")
	fmt.Println("a")

	defer fmt.Println("2. Second defer (Runs 2nd)")
	fmt.Println("b")

	defer fmt.Println("3. Third defer (Runs 1st)")
	fmt.Println("c")

	fmt.Println("d")
	// Output: a, b, c, d, 3, 2, 1
}

/*
===============================================
Defer Arguments Evaluation
===============================================

⚠️ Highly Tricky MCQ Topic:
Arguments are evaluated immediately.
*/

func deferArgumentsDemo() {
	fmt.Println("\n======== Defer Arguments Evaluation ========")

	x := 10
	
	// `x` is evaluated RIGHT NOW as 10, not later.
	defer fmt.Println("x in defer:", x)

	x = 20
	fmt.Println("x immediately:", x) // 20

	// Output:
	// x immediately: 20
	// x in defer: 10 (Not 20!)
	
	// ⚠️ If you want the deferred function to see the modified value, use a closure!
	y := 10
	defer func() {
		fmt.Println("y in closure defer:", y) // This will capture `y` by reference and print 20!
	}()
	y = 20
}

/*
===============================================
Defer for Resource Cleanup (The Main Use Case)
===============================================

Always use `defer` immediately after successfully opening a resource (file, lock, DB connection).
*/

func fileOperationWithDefer() {
	fmt.Println("\n======== Defer for Cleanup ========")

	// Simulating opening a file
	fmt.Println("1. Opening File")

	defer func() {
		fmt.Println("3. Closing File (Guaranteed via defer)")
	}()

	fmt.Println("2. Processing File Data")
}

/*
===============================================
Database Example
===============================================
*/

type Database struct {
	connected bool
}

func (db *Database) Connect() error {
	db.connected = true
	fmt.Println("Database Connected")
	return nil
}

func (db *Database) Query(sql string) {
	if !db.connected {
		fmt.Println("Error: Not connected")
		return
	}
	fmt.Printf("Executing: %s\n", sql)
}

func (db *Database) Close() error {
	db.connected = false
	fmt.Println("Database Disconnected")
	return nil
}

func databaseExample() {
	fmt.Println("\n======== Database Example ========")

	db := &Database{}

	// Connect
	if err := db.Connect(); err != nil {
		fmt.Println("Connection error:", err)
		return
	}

	// ⚠️ CRITICAL: Defer the close immediately after a successful open!
	defer db.Close()

	// Now you can safely execute queries, and the DB will close even if we return early or panic!
	db.Query("SELECT * FROM users")
}

/*
===============================================
Defer Modifying Named Returns
===============================================

⚠️ MCQ Question: Can a defer statement change the return value of a function?
Yes! BUT ONLY if the function uses Named Return Values.
*/

func namedReturnsWithDefer() (result int, err error) {
	defer func() {
		if err != nil {
			result = -1 // Mutating the return value before it leaves the function!
		}
	}()

	// Simulating an error
	err = fmt.Errorf("something broke")
	result = 100 // We tried to return 100

	return // The defer intercepts this, sees the error, and changes `result` to -1.
}

func namedReturnsDemo() {
	fmt.Println("\n======== Defer with Named Returns ========")

	res, err := namedReturnsWithDefer()
	fmt.Printf("Result: %d, Error: %v\n", res, err)
	// Result: -1 (Modified by defer!)
}

/*
===============================================
Common Mistakes with Defer
===============================================
*/

func deferMistakes() {
	fmt.Println("\n======== Defer Mistakes ========")

	// ❌ Mistake 1: Not checking the error from the `Close` method
	file, err := os.CreateTemp("", "test")
	if err != nil {
		return
	}
	
	// Incorrect (Ignoring errors on close can hide write failures in some systems)
	// defer file.Close() 
	
	// Correct:
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Println("Failed to close file:", closeErr)
		}
	}()

	fmt.Println("Temp file created and deferred close registered.")
}

/*
===============================================
Defer in Loops (A Memory Leak Trap)
===============================================

⚠️ Defer runs at the end of the FUNCTION, not the end of the block/loop!
Using defer inside a large loop will accumulate thousands of deferred functions until the loop finishes, blowing up memory.
*/

func deferInLoopsDemo() {
	fmt.Println("\n======== Defer in Loops ========")

	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d\n", i)
		defer fmt.Printf("Deferred in iteration %d\n", i)
	}
	
	// Output will show that defers wait for the whole function to finish!
	// Iteration 0, Iteration 1, Iteration 2, Deferred 2, Deferred 1, Deferred 0
}
