package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	fmt.Println("Stage 4: Functions, Pointers & Methods")
	fmt.Println("============================================================")

	fmt.Println("\n--- 1. Functions ---")
	basicsDemo()
	variadicDemo()
	anonymousFunctionsDemo()
	closuresDemo()
	functionTypesDemo()
	simpleDefer()

	fmt.Println("\n--- 2. Pointers ---")
	basicPointers()
	nilPointersDemo()
	functionPointersDemo()
	sliceVsPointer()
	escapingPointer()

	fmt.Println("\n--- 3. Methods ---")
	valueVsPointerReceiver()
	customTypesMethods()
	methodChainingDemo()
	autoConversionDemo()
	methodsVsFunctions()
}
