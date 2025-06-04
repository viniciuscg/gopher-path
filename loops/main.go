package main

import "fmt"

func main() {
	// --- BASIC FOR LOOP ---

	// A classic for loop with initialization, condition, and post statement

	for i := 0; i < 3; i++ {
		fmt.Println("Basic loop i =", i) // Output: 0 1 2
	}

	// --- WHILE-STYLE LOOP ---

	// Go doesn't have a separate "while" keyword — this is the same as a while loop
	j := 0
	for j < 3 {
		fmt.Println("While-style loop j =", j)
		j++
	}

	// --- INFINITE LOOP (USE WITH BREAK) ---

	// 'for' alone is an infinite loop — use 'break' to exit
	k := 0
	for {
		fmt.Println("Infinite loop k =", k)
		k++
		if k == 2 {
			break // Exit loop when k reaches 2
		}
	}

	// --- USING CONTINUE ---

	// Skip even numbers
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd number:", i) // Output: 1, 3
	}

	// --- LOOPING OVER ARRAYS/SLICES ---

	numbers := []int{10, 20, 30}

	// Use range to loop over slice
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d", index, value)
	}

	// Ignore the index with '_'
	for _, value := range numbers {
		fmt.Println("Value only:", value)
	}

	// Ignore the value with '_'
	for index, _ := range numbers {
		fmt.Println("Index only:", index)
	}

	// --- LOOPING OVER STRINGS (UNICODE SAFE) ---

	// Ranging over strings gives you Unicode-aware runes and their byte positions on the ascii table
	word := "Golang"
	for index, char := range word {
		fmt.Printf("Index: %d, Char: %c", index, char)
	}

	// Note: 'char' is of type rune (alias for int32)

	// --- LOOPING OVER MAPS ---

	// Ranging over a map gives key-value pairs
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d", key, value)
	}

	// Order of iteration over a map is random!

	// --- NESTED LOOPS ---

	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d j=%d", i, j)
		}
	}

	// --- BREAKING MULTIPLE LOOPS WITH LABELS ---

outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // Break out of both loops
			}
			fmt.Println("i:", i, "j:", j)
		}
	}
}
