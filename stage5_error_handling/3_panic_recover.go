package main

import "fmt"

/*
===============================================
Stage 5: Error Handling - Part 3
Panic and Recover
===============================================

⚠️ Crucial Points:
1. `panic`: Immediately halts the normal execution of the current function. All deferred functions are run, and the function returns to its caller. This bubbles up the stack until the program crashes and prints a stack trace.
2. `recover`: A built-in function that regains control of a panicking goroutine.
3. ⚠️ `recover` is ONLY useful when called inside a `defer` function! Calling it anywhere else does absolutely nothing.
4. If `recover` catches a panic, it returns the value passed to the `panic` call. The program resumes normal execution AFTER the function that panicked.
5. Panic is NOT like Exceptions (try/catch)! Use it ONLY for truly unrecoverable programming errors, NOT for expected errors (like bad user input).
*/

func basicPanic() {
	fmt.Println("======== Basic Panic ========")

	fmt.Println("Before panic")
	// panic("Critical failure!") // This would crash the entire application
	fmt.Println("After panic (Will never execute if panicked)")
}

/*
===============================================
Recovery with Defer
===============================================

⚠️ `recover` stops the panic from crashing the program.
*/

func panicWithRecovery() {
	fmt.Println("\n======== Panic with Recovery ========")

	defer func() {
		// Calling recover() inside defer
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic! The panic message was: %v\n", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Test error!") 
	fmt.Println("After panic (Will never execute, but the program survives!)")
}

/*
===============================================
Recovery with Cleanup
===============================================

Real-world use case: Ensuring a web server doesn't crash entirely if a single request handler panics.
*/

type Server struct {
	Running bool
}

func (s *Server) Start() {
	s.Running = true
	fmt.Println("Server started")
}

func (s *Server) Stop() {
	s.Running = false
	fmt.Println("Server stopped")
}

func (s *Server) HandleRequest() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Request handler panicked: %v\n", r)
			s.Stop() // Cleanup state
		}
	}()

	fmt.Println("Processing request...")
	panic("Null pointer exception during request processing!")
	fmt.Println("After panic (Ignored)") // Unreachable
}

func recoveryWithCleanup() {
	fmt.Println("\n======== Recovery with Cleanup ========")

	server := &Server{}
	server.Start()

	server.HandleRequest()

	fmt.Printf("Server state: Running=%v\n", server.Running)
}

/*
===============================================
Nested Panic Recovery and Re-panicking
===============================================
Sometimes you catch a panic, realize you can't handle it, and throw it again.
*/

func nestedPanicRecovery() {
	fmt.Println("\n======== Nested Panic Recovery ========")

	// Outer Defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Outer recover caught: %v\n", r)
		}
	}()

	// Inner Anonymous Function
	func() {
		// Inner Defer
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Inner recover caught: %v\n", r)
				fmt.Println("Inner recover re-panicking...")
				panic(r) // Re-throwing the panic!
			}
		}()

		panic("Error inside inner function!")
	}()

	fmt.Println("This line will NEVER print because the inner function re-panicked.")
}

/*
===============================================
When to use Panic?
===============================================

✓ DO use Panic for:
- Assertion failures (e.g., regex compilation on startup)
- Impossible situations (index out of bounds)
- Irrecoverable programming errors (nil pointer dereference)

✗ DO NOT use Panic for:
- Expected errors (Network timeouts, bad user input, missing files)
- Control flow

Always use the `error` interface for expected failures!
*/

func whenToUsePanic() {
	fmt.Println("\n======== When to Use Panic ========")

	// Example of a valid panic scenario:
	data := make([]int, 5)
	idx := 10
	
	// Simulating an impossible index
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Caught out of bounds panic:", r)
		}
	}()

	if idx < 0 || idx >= len(data) {
		panic(fmt.Sprintf("Index %d out of bounds for slice of length %d", idx, len(data)))
	}
}

/*
===============================================
Stack Unwinding
===============================================

When a panic occurs, Go traverses UP the call stack:
1. It executes all `defer` statements in the current function.
2. It returns to the caller and executes its `defer` statements.
3. This repeats until it hits `main()` (and crashes) OR hits a `recover()`.
*/

func level3() {
	defer fmt.Println("Defer in level3")
	fmt.Println("Inside level3")
	panic("Catastrophic failure in level3!")
}

func level2() {
	defer fmt.Println("Defer in level2")
	fmt.Println("Inside level2")
	level3()
}

func level1() {
	defer func() {
		fmt.Println("Defer in level1 (Checking for recover)")
		if r := recover(); r != nil {
			fmt.Printf("Caught panic in level1: %v\n", r)
		}
	}()

	fmt.Println("Inside level1")
	level2()
}

func stackUnwindingDemo() {
	fmt.Println("\n======== Stack Unwinding ========")

	level1()

	fmt.Println("Program continues normally after stack unwinding and recovery!")
}

/*
===============================================
Panic Type Assertions
===============================================

`panic` can take ANY type (it accepts `interface{}`).
Therefore, `recover()` returns `interface{}`. You can use Type Assertions to find out what was panicked.
*/

func panicAssertionDemo() {
	fmt.Println("\n======== Panic Assertion ========")

	func() {
		defer func() {
			if r := recover(); r != nil {
				// Type Assertion Switch
				switch v := r.(type) {
				case string:
					fmt.Printf("String panic: %s\n", v)
				case error:
					fmt.Printf("Error panic: %v\n", v)
				case int:
					fmt.Printf("Integer panic: %d\n", v)
				default:
					fmt.Printf("Unknown panic type: %T\n", v)
				}
			}
		}()

		panic(404) // Panicking with an integer!
	}()
}
