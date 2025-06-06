package main

import "fmt"

func main() {
	fmt.Println("--- POINTERS IN FUNCTIONS ---")

	// --- MODIFYING VALUES VIA POINTERS ---

	num := 10
	fmt.Println("Before modifyValue:", num)

	modifyValue(&num)                      // Passing the pointer to 'num'
	fmt.Println("After modifyValue:", num) // Output: 20 — modified in-place

	// --- POINTERS TO STRUCTS ---

	user := Person{name: "Alice", age: 30}
	fmt.Println("Before updatePerson:", user)

	updatePerson(&user) // Pass pointer to struct
	fmt.Println("After updatePerson:", user)

	// --- POINTERS AS RETURN VALUES ---

	ptr := createPointer()
	fmt.Println("Pointer returned from function:", *ptr) // Output: 100

	// --- FUNCTION WITH POINTER PARAMETERS RETURNING POINTERS ---

	result := sumPointerValues(&num, ptr)
	fmt.Println("Sum via pointer values:", *result)

	// --- NIL POINTER SAFETY CHECKS ---

	var p *int   // nil
	safePrint(p) // Output: Pointer is nil
	x := 42
	safePrint(&x) // Output: Value: 42
}

// --- FUNCTION THAT MODIFIES A VALUE VIA POINTER ---
func modifyValue(n *int) {
	*n = *n * 2 // Modify value in place
}

// --- STRUCT FOR EXAMPLE ---
type Person struct {
	name string
	age  int
}

// --- FUNCTION THAT MODIFIES STRUCT FIELDS VIA POINTER ---
func updatePerson(p *Person) {
	p.age += 1
	p.name = "Updated " + p.name
}

// --- FUNCTION THAT RETURNS A POINTER TO A LOCAL VARIABLE ---
func createPointer() *int {
	x := 100
	return &x // Escapes to heap — safe to return
}

// --- FUNCTION THAT ACCEPTS AND RETURNS POINTERS ---
func sumPointerValues(a, b *int) *int {
	sum := *a + *b
	return &sum // Escapes — safe
}

// --- NIL SAFETY DEMO ---
func safePrint(p *int) {
	if p == nil {
		fmt.Println("Pointer is nil")
		return
	}
	fmt.Println("Value:", *p)
}
