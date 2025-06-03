package main

import "fmt"

func main() {
	// You can't add 'n1' and 'n2' directly in Go like this:
	// var n1 int8 = 1
	// var n2 int16 = 10
	// sum := n1 + n2

	// This won't compile, because Go is very strict about **type safety**.
	// Even though both are integers, they are different types (`int8` and `int16`),
	// and Go does NOT automatically convert one to match the other.
	// This is intentional to avoid silent bugs caused by unintended conversions.
	// You need to convert one to the other explicitly.

	var n1 int8 = 1
	var n2 int16 = 10

	// Example: convert n1 to int16 so both types match
	// But this maybe is not a good practice
	sum := int16(n1) + n2
	fmt.Println(sum) // Output: 11

	// ----------------------------------------------

	// Variable declaration with full syntax
	var text string = "text"

	// Short declaration with inferred type (also string)
	text2 := "text"

	// Both are valid and result in the same value and type.
	// Use `:=` when you're inside a function and want concise syntax.

	// ----------------------------------------------

	// In Go, the ternary conditional operator (e.g., 5 > 3 ? "Yes" : "No") does NOT exist.
	// This is by design — Go favors **simplicity and clarity** over compactness.
	// The equivalent logic must be written using a regular `if` statement:

	if 5 > 3 {
		fmt.Println("Sim") // "Yes"
	} else {
		fmt.Println("Não") // "No"
	}

	// This encourages consistent and readable code, especially in larger codebases.
}
