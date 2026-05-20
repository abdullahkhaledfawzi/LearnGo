package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
===============================================
Stage 7: Core Standard Library Packages
Part 3: Time, Math, and Random
===============================================
*/

/*
===============================================
time Package
===============================================

⚠️ Crucial points about Time formatting in Go:
Go does NOT use strftime (like `%Y-%m-%d`).
Instead, Go uses a SPECIFIC REFERENCE DATE to define formats:
Mon Jan 2 15:04:05 MST 2006
(1=Month, 2=Day, 3=Hour 12, 4=Minute, 5=Second, 6=Year)
*/

func timeDemo() {
	fmt.Println("======== time Package ========")

	// 1. Current Time
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)
	fmt.Printf("Date components: %d-%02d-%02d\n", now.Year(), now.Month(), now.Day())

	// 2. Duration (Represents the difference between two instants as nanoseconds)
	fmt.Println("\n--- Duration ---")
	d := 5 * time.Second
	fmt.Printf("Duration: %v\n", d)

	// 3. Time Manipulation (Add / Sub)
	fmt.Println("\n--- Time Math ---")
	future := now.Add(24 * time.Hour) // Adding exactly 24 hours
	fmt.Printf("Tomorrow: %v\n", future)

	past := now.Add(-2 * time.Hour)   // Subtracting 2 hours
	diff := now.Sub(past)             // Returns a Duration
	fmt.Printf("Difference (now - past): %v\n", diff)

	// 4. Formatting (Using the magical 2006 reference date)
	fmt.Println("\n--- Formatting ---")
	// 2006 = Year, 01 = Month, 02 = Day, 15 = Hour (24h), 04 = Min, 05 = Sec
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Printf("Custom Format: %s\n", formatted)

	// Built-in RFC standards
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))

	// 5. Parsing Time from String
	fmt.Println("\n--- Parsing ---")
	str := "2024-12-25"
	parsed, _ := time.Parse("2006-01-02", str)
	fmt.Printf("Parsed Date: %v\n", parsed)

	// 6. Ticker (Fires events repeatedly on a channel)
	fmt.Println("\n--- Ticker ---")
	ticker := time.NewTicker(200 * time.Millisecond)
	count := 0
	for range ticker.C {
		count++
		fmt.Printf("Tick %d\n", count)
		if count >= 3 {
			ticker.Stop() // Prevents memory leaks!
			break
		}
	}
}

/*
===============================================
math Package
===============================================
*/

func mathDemo() {
	fmt.Println("\n======== math Package ========")

	// 1. Constants
	fmt.Printf("Pi: %f\n", math.Pi)

	// 2. Basic Math
	fmt.Println("\n--- Basic Math ---")
	fmt.Printf("Absolute(-5): %f\n", math.Abs(-5))
	fmt.Printf("Square Root(16): %f\n", math.Sqrt(16))
	fmt.Printf("Power(2^3): %f\n", math.Pow(2, 3))
	fmt.Printf("Min(5, 3): %f\n", math.Min(5, 3))
	fmt.Printf("Max(5, 3): %f\n", math.Max(5, 3))

	// 3. Rounding
	fmt.Println("\n--- Rounding ---")
	f := 3.7
	fmt.Printf("Floor(3.7) [Down]: %f\n", math.Floor(f))
	fmt.Printf("Ceil(3.7) [Up]: %f\n", math.Ceil(f))
	fmt.Printf("Round(3.7) [Nearest]: %f\n", math.Round(f))
}

/*
===============================================
math/rand Package
===============================================

⚠️ Note: Starting in Go 1.20, the global random number generator is automatically seeded.
You no longer need to call `rand.Seed(time.Now().UnixNano())`.
*/

func randDemo() {
	fmt.Println("\n======== math/rand Package ========")

	// 1. Random Integer
	fmt.Println("--- Random Numbers ---")
	for i := 0; i < 3; i++ {
		// rand.Intn(N) generates a number from 0 to N-1
		fmt.Printf("Random 1-10: %d\n", rand.Intn(10)+1)
	}

	// 2. Random Float64 (From 0.0 to 1.0)
	fmt.Println("\n--- Random Float ---")
	fmt.Printf("Random Float: %f\n", rand.Float64())

	// 3. Shuffle a slice (Very common in algorithms)
	fmt.Println("\n--- Shuffle ---")
	nums := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i] // Swap syntax
	})
	fmt.Printf("Shuffled Array: %v\n", nums)
}

/*
===============================================
Testing Package Introduction
===============================================
Unit testing is a first-class citizen in Go. 
Tests are written in files ending in `_test.go`.

Example:
func TestMyFunction(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
Execute with: `go test`
*/


