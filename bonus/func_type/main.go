package main

import "fmt"

func main() {
	fmt.Println("--- FUNCTION POINTERS IN GO ---")

	// --- FUNCTION AS A VARIABLE ---

	// You can assign a function to a variable
	var op func(int, int) int

	op = add
	fmt.Println("Add:", op(10, 5)) // Output: 15

	op = subtract
	fmt.Println("Subtract:", op(10, 5)) // Output: 5

	// --- FUNCTION TYPE DECLARATION ---

	// You can define your own function type
	type operation func(int, int) int

	var mul operation = multiply
	fmt.Println("Multiply:", mul(4, 3)) // Output: 12

	// --- PASSING FUNCTION AS PARAMETER ---

	result := calculate(8, 2, divide)
	fmt.Println("Divide:", result) // Output: 4

	// --- RETURNING A FUNCTION FROM A FUNCTION ---

	opFunc := chooseOperation("*")
	fmt.Println("Choose operation '*':", opFunc(6, 7)) // Output: 42

	// --- ANONYMOUS FUNCTION AS POINTER VALUE ---

	custom := func(a, b int) int {
		return a*a + b*b
	}
	fmt.Println("Custom logic (a² + b²):", custom(3, 4)) // Output: 25
}

// --- BASIC MATH FUNCTIONS ---
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// --- FUNCTION THAT TAKES A FUNCTION AS ARGUMENT ---
func calculate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// --- FUNCTION THAT RETURNS A FUNCTION ---
func chooseOperation(op string) func(int, int) int {
	switch op {
	case "+":
		return add
	case "-":
		return subtract
	case "*":
		return multiply
	case "/":
		return divide
	default:
		return func(a, b int) int {
			fmt.Println("Unknown op:", op)
			return 0
		}
	}
}
