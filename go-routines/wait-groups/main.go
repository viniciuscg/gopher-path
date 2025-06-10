package main

import (
	"fmt"
	"sync"
	"time"
)

// --- INTRODUCTION TO sync.WaitGroup ---
//
// A WaitGroup waits for a collection of goroutines to finish.
// It's part of the `sync` package and commonly used to:
// Wait for multiple tasks to finish before exiting a function
// Avoid using time.Sleep to "guess" when goroutines are done
// Write safe and reliable concurrent programs

func main() {
	// --- CREATE A WAITGROUP ---

	var wg sync.WaitGroup

	// We're going to launch 3 goroutines.
	wg.Add(3) // Tell the WaitGroup to wait for 3 goroutines

	// --- START GOROUTINES ---
	go doWork("A", 1*time.Second, &wg)
	go doWork("B", 2*time.Second, &wg)
	go doWork("C", 1*time.Second, &wg)

	// --- WAIT FOR ALL GOROUTINES TO FINISH ---
	// This blocks until the counter inside wg reaches 0.
	wg.Wait()

	fmt.Println("All tasks completed ✅")
}

// --- WORKER FUNCTION ---
func doWork(id string, duration time.Duration, wg *sync.WaitGroup) {
	// Always call Done() at the end to decrease the counter
	defer wg.Done()

	fmt.Println("Starting task:", id)
	time.Sleep(duration) // Simulate work
	fmt.Println("Finished task:", id)
}
