package main

import "fmt"

func main() {
	fmt.Println("============================================================")
	fmt.Println("Stage 6: Interfaces & Concurrency")
	fmt.Println("============================================================")

	fmt.Println("\n--- 1. Interfaces ---")
	basicInterfaceDemo()
	emptyInterfaceDemo()
	typeAssertionDemo()
	typeSwitchDemo()
	embeddingInterfacesDemo()
	nilInterfaceDemo()
	satisfactionDemo()
	pointerReceiverDemo()

	fmt.Println("\n--- 2. Goroutines and Channels ---")
	basicGoroutinesDemo()
	unbufferedChannelDemo()
	bufferedChannelDemo()
	channelDirectionDemo()
	closingChannelDemo()
	waitGroupDemo()
	selectDemo()
	timeoutDemo()
	mutexDemo()
}
