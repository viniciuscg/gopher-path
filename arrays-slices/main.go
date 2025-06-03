package main

import "fmt"

func main() {
	// --- ARRAYS ---

	// An array in Go has a fixed size and a specific type.
	// Here we declare an array of 5 integers.
	var numbers [5]int

	// Assign values to the array by index.
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Print the entire array.
	fmt.Println("Array:", numbers) // Output: [10 20 30 40 50]

	// You can also declare and initialize an array using a literal:
	grades := [3]float64{7.5, 8.3, 9.0}
	fmt.Println("Grades:", grades)

	// The length of the array is fixed. You can't add or remove elements.
	// Trying to access an out-of-range index will cause a runtime panic.

	// --- SLICES ---

	// A slice is a flexible and more powerful wrapper on top of arrays.
	// It does not have a fixed length — you can append elements to it.
	// Declare a slice using [] without specifying the length.
	names := []string{"Vinicius", "Go", "Language"}

	// Print the slice and its length.
	fmt.Println("Slice:", names)
	fmt.Println("Length:", len(names)) // Output: 3

	// Append a new value to the slice.
	names = append(names, "Gopher")
	fmt.Println("After append:", names) // Output: [Vinicius Go Language Gopher]

	// You can slice a slice — this is like subsetting.
	// Here we take elements from index 1 up to (but not including) index 3.
	subset := names[1:3]
	fmt.Println("Sliced subset:", subset) // Output: [Go Language]

	// --- MAKE FUNCTION ---

	// Create a slice using make. This is useful for dynamic data.
	// make(type, length, capacity)
	numbersSlice := make([]int, 3, 5) // Creates a slice with length 3 and capacity 5
	numbersSlice[0] = 100
	numbersSlice[1] = 200
	numbersSlice[2] = 300
	fmt.Println("Slice from make:", numbersSlice)

	// --- COPY FUNCTION ---

	// Copy one slice into another.
	source := []int{1, 2, 3}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println("Copied slice:", destination)

	// Modifying the source does NOT affect the destination.
	source[0] = 99
	fmt.Println("Source modified:", source)
	fmt.Println("Destination remains unchanged:", destination)
}
