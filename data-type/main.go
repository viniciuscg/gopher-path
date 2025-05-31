package main

import (
	"errors"
	"fmt"
)

func main() {
	// In Go, we have several integer types:
	// `int` uses the default size (32 or 64 bits depending on the architecture).
	// `int8`, `int16`, `int32`, `int64` specify exact sizes.
	var a int = 42
	var b int8 = 120
	var c int64 = 9000000000000000000

	// `rune` is an alias for int32, commonly used to represent Unicode characters.
	var r rune = 'あ'

	// `byte` is an alias for uint8.
	var by byte = 255

	// Unsigned integers (do not allow negative values)
	var u uint = 1000
	var u8 uint8 = 255
	var u64 uint64 = 18446744073709551615

	// Implicit typing with `:=`
	x := 1000000         // inferred as int
	// I Dont no if thism make sense but you can:
	ux := uint(1234567)  // need explicit cast for unsigned

	// Floating-point numbers
	var f32 float32 = 123.456
	var f64 float64 = 123456.789012
	floatInferred := 3.14 // inferred as float64

	// Strings
	var s string = "Hello, Go!"
	stringInferred := "Olá, Gopher!"

	// Characters (runes) use single quotes and represent Unicode points.
	char := 'B' // This is rune, underlying type is int32. 'B' == 66

	// Variables declared without initialization get zero values:
	var emptyString string
	var zeroInt int
	var zeroBool bool
	var zeroError error

	fmt.Println("Empty string:", emptyString)
	fmt.Println("Zero int:", zeroInt)
	fmt.Println("Zero bool:", zeroBool)
	fmt.Println("Zero error:", zeroError) // <nill>

	// Booleans
	var isTrue bool = true
	var isFalse bool = false
	inferredBool := true

	// Errors in Go are values of type `error`
	var err error = errors.New("Something went wrong")
	fmt.Println("Error:", err)

	// You can declare a nil error (default zero value of error)
	var nilError error
	fmt.Println("Nil error:", nilError)

	// A function is also a type. Example of assigning a function to a variable:
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println("Function result:", add(3, 4))
}
