package main

import "fmt"

func main() {
	// --- BASIC IF STATEMENT ---

	// A simple if condition — checks if number is positive.
	number := 5
	if number > 0 {
		fmt.Println("The number is positive")
	}

	// --- IF-ELSE ---

	// If-else to check even or odd.
	if number%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// --- IF-ELSE IF-ELSE ---

	temperature := 30
	if temperature < 0 {
		fmt.Println("Freezing")
	} else if temperature < 20 {
		fmt.Println("Cold")
	} else if temperature < 30 {
		fmt.Println("Warm")
	} else {
		fmt.Println("Hot")
	}

	// --- IF WITH SHORT STATEMENT ---

	// Go allows you to declare and initialize a variable right inside the `if` statement.
	// The scope of that variable is only inside that if/else block.
	if length := len("golang"); length > 5 {
		fmt.Println("Long word with length:", length)
	} else {
		fmt.Println("Short word")
	}

	// The variable 'length' here is NOT accessible outside the 'if' block
	// fmt.Println(length) // Error: undefined

	// --- IF IS A STATEMENT, NOT AN EXPRESSION ---
	// Obs: omg this is soo cool, go surprise me every time.

	// Unlike languages like JavaScript, Python, or Rust,
	// Go does NOT allow the use of if as an expression (i.e., it doesn't return a value).
	// So this won't work:
	// result := if x > 5 { "big" } else { "small" } // Invalid in Go

	// Instead, you must use full if-else statements.

	// --- EMPTY IF BODY IS INVALID ---

	// You can't have an `if` block without a body — it must do something.
	// if true {
	// } // valid
	// if true // Syntax error, braces required

	// --- NESTED IF STATEMENTS ---

	a := 10
	b := 20
	if a > 5 {
		if b > 15 {
			fmt.Println("Both conditions met")
		}
	}

	age := 25
	name := "Vinicius"

	// --- AND operator (&&): both conditions must be true
	if age > 18 && name == "Vinicius" {
		fmt.Println("Adult named Vinicius")
	}

	// --- OR operator (||): at least one condition must be true
	if age < 18 || name == "Vinicius" {
		fmt.Println("Minor or named Vinicius")
	}

	// --- NOT operator (!): inverts a boolean value
	isStudent := false
	if !isStudent {
		fmt.Println("Not a student")
	}

	// --- OMITTING BRACES IS NOT ALLOWED ---

	// Go forces you to use braces `{}` even for single-line blocks.
	// if a > 5 fmt.Println("too short") // Syntax error

	// This is valid:
	if a > 5 {
		fmt.Println("Braces are required")
	}
}
