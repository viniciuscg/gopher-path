package main

import (
	"fmt"
	"math/cmplx"
)

// These are global variables, declared at the package level.
// Since they are not initialized, Go assigns them default **zero values**.
// - For booleans: the zero value is 'false'.
var c, python, java bool

// Here, we use a block declaration for global variables with initialization.
var (
	// A boolean explicitly set to false
	ToBe bool = false

	// A uint64 variable set to the maximum possible 64-bit unsigned integer
	// 1 << 64 is 2⁶⁴; subtracting 1 gives the max value: 18446744073709551615
	MaxInt uint64 = 1<<64 - 1

	// A complex number: the square root of (-5 + 12i)
	// cmplx.Sqrt returns a complex128 result
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// Call the 'swap' function, which returns the two strings in reverse.
	// Multiple return values are assigned to 'a' and 'b'.
	a, b := swap("Foo", "Bar")

	// This prints: "Bar Foo"
	fmt.Println(a, b)

	// Call the 'split' function with argument 17.
	// It returns two integers which are printed directly.
	// Output will be: 7 10 (since 17 * 4 / 9 = 7, and 17 - 7 = 10)
	fmt.Println(split(17))

	// Declare a local int variable 'i' without initializing.
	// It gets the default value: 0
	var i int

	// Print the local variable 'i', and the global booleans 'c', 'python', 'java'.
	// Since none of the booleans were assigned values yet, they are all false.
	// Output: 0 false false false
	fmt.Println(i, c, python, java)

	// Declare new **local variables** that shadow the global ones:
	// These are visible only within this scope, and hide the global versions.
	var c, python, java = true, false, "no!"

	// The next line would give an error here because 'j' has not been declared yet.
	// fmt.Println(i, j, c, python, java)

	// Redeclare 'i' and introduce a new variable 'j' in a single statement.
	// This new 'i' shadows the earlier local 'i'.
	var i, j int = 1, 2

	// Declare 'k' using short declaration with inferred type (int).
	k := 3

	// Use short declaration again to redeclare c, python, and java.
	// These now shadow the previous local variables (not the globals anymore).
	// Go allows shadowing in nested scopes, which is useful in some cases but should be done carefully.
	c, python, java := true, false, "no!"

	// Print the most recent versions of all variables in this scope.
	// Output: 1 2 3 true false no!
	fmt.Println(i, j, k, c, python, java)

	// Print types and values using Printf with formatting verbs.
	// %T prints the type of the variable, %v prints the value.

	// Output:
	// Type: bool Value: false
	fmt.Printf("Type: %T Value: %v", ToBe, ToBe)

	// Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v", MaxInt, MaxInt)

	// Type: complex128 Value: (2+3i)
	// This is the square root of (-5 + 12i)
	fmt.Printf("Type: %T Value: %v", z, z)
}

// Function: swap
// Accepts two strings and returns them in reverse order.
// Go supports multiple return values, which is useful in many cases.
func swap(x, y string) (string, string) {
	return y, x
}

// Function: split
// Accepts an integer input and splits it into two values: x and y.
// The return variables are named in the function signature, so we can return them implicitly.
// No need to write: return x, y — just 'return' is enough.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // returns x and y implicitly
}
