package main

import "fmt"

func main() {
	// --- BASIC POINTER USAGE ---

	// Declare a variable 'a' of type int and assign it the value 10.
	a := 10

	// Create a pointer 'ptr' that holds the memory address of 'a'.
	// '&a' means "address of a"
	var ptr *int = &a

	// Print the value of 'a', the address stored in 'ptr', and the value at that address.
	fmt.Println("Value of a:", a)       // Output: 10
	fmt.Println("Address of a:", ptr)   // Example: 0xc000014088
	fmt.Println("Value via ptr:", *ptr) // Output: 10 — *ptr dereferences the pointer

	// Modify the value at the memory location pointed to by 'ptr'.
	*ptr = 20
	fmt.Println("New value of a:", a) // Output: 20 — original 'a' changed via pointer

	// --- POINTERS WITH FUNCTIONS ---

	value := 50
	fmt.Println("\nBefore increment:", value)

	// Call a function passing a pointer to 'value'
	increment(&value)
	fmt.Println("After increment:", value) // Output: 51

	// --- POINTER TYPES ---

	// Go is statically typed, so *int, *string, *float64 are all distinct pointer types.
	var pi *int     // pointer to int
	var ps *string  // pointer to string
	var pf *float64 // pointer to float64

	// All of them have a default zero value of 'nil'
	fmt.Println("\nUninitialized pointers:", pi, ps, pf)

	// --- MEMORY ALLOCATION WITH `new` ---

	// The `new` keyword allocates zeroed storage for a type and returns a pointer to it.
	numPtr := new(int)                               // *int, initially 0
	fmt.Println("\nValue via new pointer:", *numPtr) // Output: 0
	*numPtr = 99
	fmt.Println("Updated value via pointer:", *numPtr) // Output: 99

	// Note: `new(T)` is similar to `&T{}` for structs — allocates memory and returns a pointer.

	// --- NIL CHECKS ---

	var p *int
	fmt.Println("\nNil pointer:", p) // Output: <nil>

	if p == nil {
		fmt.Println("Safe to skip dereferencing a nil pointer")
	}

	// --- POINTER TO STRUCT ---

	// Define a simple struct
	type user struct {
		name string
	}

	u := user{name: "Vinicius"}

	// Get a pointer to the struct
	up := &u

	// Go allows direct field access via struct pointer
	fmt.Println("\nUser name via pointer:", up.name)

	// Update the field via the pointer
	up.name = "Gopher"
	fmt.Println("Updated name:", u.name)

	// --- STACK vs HEAP ALLOCATION (High Level) ---

	// Go uses a garbage-collected runtime. It will automatically determine whether a value
	// should live on the stack (fast and temporary) or the heap (longer-lived).
	//
	// Example:
	x := createPointer()
	fmt.Println("\nPointer from function:", *x) // Output: 42
}

// --- FUNCTION THAT MODIFIES A VALUE VIA POINTER ---
func increment(num *int) {
	*num += 1
}

// --- FUNCTION THAT RETURNS A POINTER TO A LOCAL VARIABLE ---
// Go handles this safely — escapes to the heap if needed.
func createPointer() *int {
	val := 42
	return &val // safe, Go ensures this is allocated on heap if necessary
}
