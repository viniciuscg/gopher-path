package main

import "fmt"

func main() {
	// --- DECLARING AND INITIALIZING MAPS ---

	// Declare and initialize a map using map literal syntax.
	// This creates a map with string keys and int values.
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Initial map:", ages)

	// Access values using keys.
	fmt.Println("Alice's age:", ages["Alice"])

	// Add a new key-value pair.
	ages["Charlie"] = 35
	fmt.Println("After adding Charlie:", ages)

	// Update an existing key.
	ages["Bob"] = 26
	fmt.Println("After updating Bob:", ages)

	// Delete a key from the map using 'delete()'
	delete(ages, "Alice")
	fmt.Println("After deleting Alice:", ages)

	// Check if a key exists using the "comma ok" idiom.
	age, ok := ages["Alice"]
	if ok {
		fmt.Println("Alice exists with age:", age)
	} else {
		fmt.Println("Alice not found")
	}

	// --- USING MAKE TO CREATE MAPS ---

	// You can also create an empty map using `make`.
	// Unlike slices, maps created with `make` don't need an initial capacity, but you can provide a hint.
	scores := make(map[string]int) // same as: var scores = map[string]int{}
	scores["Golang"] = 100
	scores["Python"] = 95
	fmt.Println("Scores:", scores)

	// --- MAPS ARE REFERENCE TYPES ---

	// Maps are internally implemented as pointers to runtime hash tables.
	// Assigning a map to a new variable doesn't copy the data — it shares the same reference.
	m1 := map[string]int{"x": 1, "y": 2}
	m2 := m1

	// Modify m2, and m1 also reflects the change.
	m2["x"] = 42
	fmt.Println("m1:", m1) // m1["x"] is now 42
	fmt.Println("m2:", m2)

	// --- MAPS AND ZERO VALUES ---

	// The zero value of a map type is `nil`.
	// Reading from a nil map is safe — it returns the zero value of the value type.
	// But writing to a nil map causes a runtime panic.
	var nilMap map[string]int
	fmt.Println("Nil map read:", nilMap["ghost"]) // Safe; prints 0

	// nilMap["ghost"] = 99 // 🔥 Would panic: assignment to entry in nil map
}
