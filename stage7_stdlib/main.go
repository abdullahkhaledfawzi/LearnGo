package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	fmt.Println("Stage 7: Standard Library")
	fmt.Println("============================================================")

	fmt.Println("\n--- 1. Strings & Formatting ---")
	fmtDemo()
	stringsDemo()
	strconvDemo()

	fmt.Println("\n--- 2. OS & JSON ---")
	osDemo()
	ioDemo()
	jsonMarshalDemo()
	jsonUnmarshalDemo()
	jsonSlicesAndMaps()
	jsonGeneric()
	jsonTagsDemo()

	fmt.Println("\n--- 3. Time, Math & Random ---")
	timeDemo()
	mathDemo()
	randDemo()
}
