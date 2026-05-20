package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	fmt.Println("Stage 2: Control Flow")
	fmt.Println("============================================================")

	fmt.Println("\n--- 1. If/Else and Switch Statements ---")
	basicIfElse()
	switchBasics()
	fallthroughDemo()
	switchWithoutCondition()
	typeSwitchIntro()
	switchOnNil()

	fmt.Println("\n--- 2. Loops and Range ---")
	basicForLoops()
	rangeArraysSlices()
	rangeStrings()
	rangeMaps()
	labeledBreaks()
}
