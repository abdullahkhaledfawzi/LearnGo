package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	fmt.Println("Stage 5: Error Handling, Defer & Panic")
	fmt.Println("============================================================")

	fmt.Println("\n--- 1. Error Handling ---")
	basicErrorHandling()
	errorWrappingDemo()
	typeAssertionForErrors()
	errorComparison()
	sentinelErrorsDemo()
	returnEarlyDemo()

	fmt.Println("\n--- 2. Defer Statement ---")
	simpleDefer()
	multipleDeferDemo()
	deferArgumentsDemo()
	fileOperationWithDefer()
	databaseExample()
	namedReturnsDemo()
	deferMistakes()
	deferInLoopsDemo()

	fmt.Println("\n--- 3. Panic & Recover ---")
	// basicPanic() // Left commented out to avoid crashing the runner
	panicWithRecovery()
	recoveryWithCleanup()
	nestedPanicRecovery()
	whenToUsePanic()
	stackUnwindingDemo()
	panicAssertionDemo()
}
