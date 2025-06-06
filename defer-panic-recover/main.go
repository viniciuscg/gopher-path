package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	fmt.Println("--- STARTING MAIN FUNCTION ---")

	// --- DEFER TO PROFILE EXECUTION TIME ---
	// This will measure how long the main function takes and print the result at the end.
	// The function returned by `profile("main")` is deferred and runs when `main()` exits.
	defer profile("main")()

	// --- FILE HANDLING ---
	// Open a file and ensure it's closed properly using defer.
	handleFile()

	// --- SAFE CONCURRENT ACCESS WITH MUTEX ---
	// Lock a mutex and ensure it's unlocked even if something fails.
	safeAccess()

	// --- PANIC RECOVERY ---
	// Show how to safely recover from a panic using defer + recover.
	safeFunction()

	// --- MOCK DATABASE CONNECTION CLEANUP ---
	// Simulate a database connection that must be closed with defer.
	queryDatabase()
}

// --- FUNCTION: PROFILE EXECUTION TIME ---
// Returns a function that prints how long the calling function took.
// Usage: defer profile("main")()
func profile(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Execution of %s took %s", name, time.Since(start))
	}
}

// --- FUNCTION: FILE HANDLING EXAMPLE ---
func handleFile() {
	fmt.Println("Opening file...")

	// Try to open a file
	file, err := os.Open("example.txt") // Must exist or return error
	if err != nil {
		log.Println("Could not open file:", err)
		return
	}

	// DEFER ensures the file is closed when this function ends,
	// even if there's a return or an error later in the function.
	defer file.Close()

	fmt.Println("File opened successfully. File will be closed automatically using defer.")
}

// --- FUNCTION: SAFE ACCESS WITH MUTEX ---
func safeAccess() {
	// Lock the mutex
	mu.Lock()

	// DEFER will ensure this is always unlocked, preventing deadlocks.
	defer mu.Unlock()

	fmt.Println("Accessing shared resource safely using mutex and defer.")
}

// --- FUNCTION: PANIC RECOVERY EXAMPLE ---
func safeFunction() {
	// DEFER a function to handle panic using recover
	defer func() {
		if r := recover(); r != nil {
			// recover() catches the panic and allows graceful recovery
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")

	// PANIC is used to abruptly stop the program (usually on unrecoverable errors)
	panic("Something went wrong!")
}

// --- FUNCTION: MOCK DATABASE CONNECTION ---
func queryDatabase() {
	fmt.Println("Connecting to mock database...")

	// Normally you would use:
	// db, err := sql.Open("mysql", "user:pass@/dbname")
	// defer db.Close()

	// Here we just simulate with a print statement
	defer fmt.Println("Closed mock DB connection (simulated)")

	fmt.Println("Running database query (simulated)...")
}
