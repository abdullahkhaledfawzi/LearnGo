package main

import (
	"fmt"
	"sync"
	"time"
)

/*
===============================================
Stage 6: Interfaces and Concurrency - Part 2
Goroutines and Channels
===============================================

⚠️ Crucial Concurrency Concepts:
1. Goroutine: A lightweight "thread" managed by the Go runtime. They cost very little memory (started at 2KB) compared to OS threads (1MB+).
2. The `go` keyword: Starts a new goroutine concurrently.
3. Channel: A typed conduit used to send and receive values between goroutines safely. "Don't communicate by sharing memory; share memory by communicating."
4. Buffered vs Unbuffered: Unbuffered channels block until BOTH sender and receiver are ready (synchronous). Buffered channels block only when full/empty (asynchronous up to capacity).
5. Deadlock: Occurs when all goroutines are asleep waiting for each other (e.g., reading from an empty channel with no writers left).
*/

/*
===============================================
Basic Goroutines
===============================================
*/

func sayHello(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s - Iteration %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func basicGoroutinesDemo() {
	fmt.Println("======== Basic Goroutines ========")

	// Spin up Goroutine 1
	go sayHello("Goroutine 1")

	// Spin up Goroutine 2
	go sayHello("Goroutine 2")

	// The Main goroutine must wait, otherwise it exits immediately and kills all child goroutines!
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Main function finished.")
}

/*
===============================================
Channels (Communication)
===============================================

⚠️ Channels provide built-in synchronization!
*/

func unbufferedChannelDemo() {
	fmt.Println("\n======== Unbuffered Channel ========")

	ch := make(chan int) // Unbuffered (capacity 0)

	go func() {
		fmt.Println("Sender: Preparing to send 42...")
		// This blocks IMMEDIATELY until a receiver is ready to take the value!
		ch <- 42 
		fmt.Println("Sender: 42 sent successfully!")
	}()

	fmt.Println("Main: Doing some work before receiving...")
	time.Sleep(200 * time.Millisecond)
	
	fmt.Println("Main: Waiting to receive...")
	value := <-ch // Receiver pulls the value, unblocking the sender.
	fmt.Printf("Main: Received %d\n", value)
}

func bufferedChannelDemo() {
	fmt.Println("\n======== Buffered Channel ========")

	// Buffered channel with capacity 2
	ch := make(chan int, 2)

	// Sender does NOT block here because the buffer has room.
	ch <- 1
	fmt.Println("Sent 1")

	ch <- 2
	fmt.Println("Sent 2")

	// If we tried to send a 3rd item here without a receiver, it would DEADLOCK!
	// ch <- 3 // ❌ Deadlock! Buffer is full!

	val1 := <-ch
	val2 := <-ch

	fmt.Printf("Received: %d and %d\n", val1, val2)
}

/*
===============================================
Channel Direction
===============================================

⚠️ You can strictly define channel direction in function signatures to enforce read-only or write-only behavior.
- `chan T`: Bidirectional
- `chan<- T`: Send-Only (Write)
- `<-chan T`: Receive-Only (Read)
*/

func send(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i // Write
	}
	close(ch) // ONLY the sender should ever close a channel!
}

func receive(ch <-chan int) {
	// `range` continuously reads from the channel until it is closed.
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}
}

func channelDirectionDemo() {
	fmt.Println("\n======== Channel Direction ========")

	ch := make(chan int)

	go send(ch)
	receive(ch) // Blocks until the sender closes the channel
}

/*
===============================================
Closing Channels
===============================================

⚠️ Critical Rules:
1. ONLY the sender should close the channel, never the receiver.
2. Sending to a closed channel causes a PANIC!
3. Closing an already closed channel causes a PANIC!
4. Reading from a closed channel returns the Zero Value safely immediately.
*/

func closingChannelDemo() {
	fmt.Println("\n======== Closing Channel ========")

	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // Safe close

	// The "Comma ok" idiom to check if channel is closed
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed and drained!")
			break
		}
		fmt.Printf("Received manually: %d\n", val)
	}

	// Range automatically handles the ok check
	fmt.Println("\nUsing range:")
	ch2 := make(chan int, 3)
	ch2 <- 10
	ch2 <- 20
	close(ch2)

	for val := range ch2 {
		fmt.Printf("Received via range: %d\n", val)
	}
}

/*
===============================================
WaitGroup (Synchronization)
===============================================

⚠️ The correct way to wait for multiple goroutines to finish instead of `time.Sleep`.
*/

func waitGroupDemo() {
	fmt.Println("\n======== WaitGroup ========")

	var wg sync.WaitGroup

	tasks := []string{"Task A", "Task B", "Task C"}

	for _, task := range tasks {
		wg.Add(1) // Add 1 to the counter BEFORE starting the goroutine

		go func(name string) {
			defer wg.Done() // Decrement counter when finished

			fmt.Printf("Started %s\n", name)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Finished %s\n", name)
		}(task) // Passing `task` as argument to avoid closure loop trap!
	}

	wg.Wait() // Blocks until the counter reaches exactly 0
	fmt.Println("All tasks completed securely!")
}

/*
===============================================
Select Statement (Multiplexing)
===============================================

⚠️ `select` lets a goroutine wait on MULTIPLE channel operations simultaneously.
It blocks until one of its cases can run, then it executes that case.
If multiple cases are ready, it picks one pseudo-randomly.
*/

func selectDemo() {
	fmt.Println("\n======== Select Multiplexing ========")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from Channel 1 (Fast)"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from Channel 2 (Slow)"
	}()

	// We expect 2 messages, so we loop twice
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

/*
===============================================
Timeout with Select
===============================================
Prevents deadlocks or hanging connections by enforcing a maximum wait time.
*/

func timeoutDemo() {
	fmt.Println("\n======== Timeout ========")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Takes 2 seconds
		ch <- "Data finally arrived!"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second): // Triggers after 1 second
		fmt.Println("Timeout! Data took too long to arrive.")
	}
}

/*
===============================================
Race Conditions and Sync.Mutex
===============================================

⚠️ Race Condition: When two goroutines access the same memory address concurrently, and at least one is writing.
This causes silent, catastrophic data corruption.

Solution: `sync.Mutex` (Mutual Exclusion Lock) or Channels.
*/

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()         // Lock the data
	defer c.mu.Unlock() // Ensure it is unlocked even if a panic occurs
	c.value++           // Safe to modify
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func mutexDemo() {
	fmt.Println("\n======== Mutex (Race Condition Prevention) ========")

	var wg sync.WaitGroup
	counter := &SafeCounter{}

	// 10 goroutines, each incrementing 100 times concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d (Expected exactly 1000 without data races)\n", counter.Value())
}
