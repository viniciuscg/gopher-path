package main

import (
	"fmt"
	"time"
)

// --- INTRODUCTION TO GOROUTINES ---
//
// A goroutine is a lightweight thread managed by the Go runtime.
// It runs concurrently with other goroutines in the same address space.
//
// Use the keyword `go` before a function call to launch it as a goroutine.
//
// Go manages thousands (even millions) of goroutines using its own scheduler.
//
// GOROUTINES = CONCURRENCY.
// Multiple goroutines may run in parallel depending on GOMAXPROCS setting (runtime.NumCPU()).

func main() {
	// --- BASIC GOROUTINE ---

	// This function will run in the background
	go sayHello()

	// Main goroutine (this function) won't wait for others unless we block/wait
	fmt.Println("Main function continues...")

	// Wait to give the goroutine time to run (not reliable in production)
	time.Sleep(1)

	// --- CONCURRENCY vs PARALLELISM ---

	// Concurrency: Managing multiple tasks at once (interleaved execution)
	// Parallelism: Actually executing multiple tasks at the exact same time (true with multiple CPUs)

	// Example: 2 goroutines running concurrently
	go count("foo")
	go count("bar")

	// Let them run for a while
	time.Sleep(3)

	// --- GOROUTINES IN REAL-WORLD ---
	// Imagine launching web scrapers, downloading files, or handling HTTP requests concurrently.

	// --- USING ANONYMOUS GOROUTINES ---

	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Println("Anonymous goroutine number:", n)
		}(i) // Important: capture value of i
	}

	time.Sleep(1 * time.Second)
}

// --- FUNCTION TO RUN IN A GOROUTINE ---
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

// --- SIMPLE FUNCTION TO SIMULATE WORK ---
func count(label string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(label, i)
		time.Sleep(500)
	}
}
