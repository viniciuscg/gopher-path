package main

import "fmt"

func main() {
	fmt.Println("--- STARTING MAIN FUNCTION ---")

	// --- BASIC CLOSURE ---
	// A function defined inside another function that captures outer variable `x`
	incrementer := makeIncrementer()
	fmt.Println(incrementer()) // Output: 1
	fmt.Println(incrementer()) // Output: 2
	fmt.Println(incrementer()) // Output: 3

	// --- FUNCTION FACTORY WITH PARAMETERS ---
	// A closure that captures its initial parameter
	add5 := makeAdder(5)
	add10 := makeAdder(10)

	fmt.Println("5 + 3 =", add5(3))   // Output: 8
	fmt.Println("10 + 7 =", add10(7)) // Output: 17

	// --- IMMEDIATE INVOCATION ---
	// Define and call a closure right away
	message := func(name string) string {
		return "Hello, " + name
	}("Vinicius")
	fmt.Println(message)

	// --- CLOSURE IN A LOOP (Common Trap) ---
	// Be careful when closures reference loop variables — they can behave unexpectedly.
	funcs := []func(){}
	for i := 0; i < 3; i++ {
		val := i // Capture copy of loop variable
		funcs = append(funcs, func() {
			fmt.Println("Captured i:", val)
		})
	}

	// Now run them — they’ll print 0, 1, 2 correctly (thanks to `val := i`)
	for _, f := range funcs {
		f()
	}

	// --- STATEFUL FUNCTION ---
	// This closure retains state across calls
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("Counter:", counter()) // Output: 1
	fmt.Println("Counter:", counter()) // Output: 2

	// --- ANONYMOUS FUNCTION USED AS ARGUMENT ---
	// Pass a closure directly into a function
	executeWithLogging(func() {
		fmt.Println("Doing some work inside a passed closure.")
	})
}

// --- FUNCTION THAT RETURNS A CLOSURE ---
// This function returns another function that remembers `x`
func makeIncrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// --- FUNCTION FACTORY ---
// Returns a closure that adds a base number to input
func makeAdder(base int) func(int) int {
	return func(n int) int {
		return base + n
	}
}

// --- HIGH-ORDER FUNCTION THAT ACCEPTS A CLOSURE ---
// Takes a closure and runs it, logging around it
func executeWithLogging(task func()) {
	fmt.Println("Starting task...")
	task()
	fmt.Println("Finished task.")
}
