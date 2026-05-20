package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	fmt.Println("Stage 1: Basics and Types")
	fmt.Println("============================================================")

	fmt.Println("\n--- 1. Variables ---")
	variablesDemo()
	typesDemo()

	fmt.Println("\n--- 2. Zero Values ---")
	zeroValuesDemo()
	uninitializedVarsDemo()
	nilDemo()

	fmt.Println("\n--- 3. Type Conversions ---")
	basicTypeConversions()
	stringConversions()
	strconvDemo()
	typeAssertionSimple()

	fmt.Println("\n--- 4. Bytes and Runes ---")
	bytesBasics()
	runesBasics()
	zeroValuesBytes()
	typeAliasesDemo()
	stringIndexing()
}
