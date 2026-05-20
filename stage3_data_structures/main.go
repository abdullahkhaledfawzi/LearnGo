package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	fmt.Println("Stage 3: Data Structures")
	fmt.Println("============================================================")

	fmt.Println("\n--- 1. Arrays and Slices ---")
	arraysBasics()
	slicesBasics()
	capacityDemo()
	copyDemo()
	nilVsEmptySlice()

	fmt.Println("\n--- 2. Maps ---")
	mapsBasics()
	mapOperations()
	mapKeyTypes()
	mapsByValue()
	nilVsEmptyMap()
	mapWithStructs()

	fmt.Println("\n--- 3. Structs ---")
	structBasics()
	pointersToStructs()
	exportedVsUnexported()
	embeddedStructs()
	structTags()
	anonymousStructs()
	emptyStruct()
	structMethods()
}
