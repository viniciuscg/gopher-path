package main

import "fmt"

func main() {
	// --- ARRAYS ARE FIXED SIZE ---

	// Arrays in Go are value types with fixed size.
	// This creates an array of 5 integers, initialized to 0.
	var arr [5]int
	arr[0] = 42
	fmt.Println("Array:", arr)
	// Output: [42 0 0 0 0]

	// --- SLICES ARE VIEWS ON ARRAYS ---

	// A slice is a dynamic window over an array.
	// It contains a pointer to an array, a length, and a capacity.
	s := []int{1, 2, 3}
	fmt.Println("Slice:", s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))

	// --- USING make() TO CREATE SLICES ---

	// make() allocates an underlying array and returns a slice referencing it.
	// First arg: type, second: length, third (optional): capacity
	s2 := make([]int, 2, 5) // length 2, capacity 5
	fmt.Println("make() slice:", s2)
	fmt.Println("Length:", len(s2))
	fmt.Println("Capacity:", cap(s2))

	// Add elements using append (length increases)
	s2 = append(s2, 10)
	s2 = append(s2, 20)
	s2 = append(s2, 30) // Now length = 5, capacity = 5
	fmt.Println("After appending:", s2)

	// Go will now increase capacity since we exceeded it
	s2 = append(s2, 40) // Triggers reallocation
	fmt.Println("After exceeding capacity:", s2)
	fmt.Println("New Length:", len(s2))
	fmt.Println("New Capacity:", cap(s2)) // Capacity will likely double (e.g., 10)

	// Important: If capacity is exceeded, Go allocates a **new array**, copies the data,
	// and the slice now refers to the new backing array.
	// The old array will remain in memory only if something still references it. Otherwise,
	// Go's GC will eventually reclaim it.

	// --- DEMONSTRATING BACKING ARRAY SHARING ---

	original := make([]int, 3, 3)
	copy(original, []int{1, 2, 3})

	sub := original[:2]
	fmt.Println("Original:", original)
	fmt.Println("Sub-slice:", sub)

	sub[0] = 99
	fmt.Println("Modified sub-slice:", sub)
	fmt.Println("Original affected:", original) // original[0] is now 99 — shared backing array!

	// But after capacity overflow:
	expanded := append(sub, 4, 5)
	expanded[0] = 100
	fmt.Println("Expanded slice:", expanded)
	fmt.Println("Original after append:", original) // Not affected — different backing array

	// --- MEMORY LEAKS: AVOIDING TRAPPED REFERENCES ---

	// Problem: If you take a small slice of a very large array,
	// Go will keep the entire large array in memory **as long as the small slice exists**.

	// Example:
	huge := make([]byte, 1_000_000)
	small := huge[:10] // Only need first 10 bytes
	_ = small          // But huge array stays in memory due to reference

	// Solution: Copy what you need to a new slice:
	safeCopy := make([]byte, len(small))
	copy(safeCopy, small)

	// Now 'huge' can be garbage-collected if not used elsewhere

	// --- BEHIND THE SCENES: HOW ARRAYS & SLICES WORK ---

	// A slice is a lightweight struct:
	// type sliceHeader struct {
	//     Data uintptr  // Pointer to underlying array
	//     Len  int      // Current length
	//     Cap  int      // Capacity (from starting point to end of array)
	// }

	// So appending may allocate new arrays on the heap, depending on growth

	// --- GROWTH STRATEGY OF append() ---

	// Go runtime grows slice capacity intelligently, usually by:
	// - Doubling up to a certain size (e.g., <= 1024)
	// - Slower linear growth after that

	// This makes appending fast — amortized O(1)

}
