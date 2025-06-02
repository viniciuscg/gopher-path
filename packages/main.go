package main

import (
	"module/helpers"
)

// main is the program's entry point.
// It calls a helper function to print two lines of text.
func main() {
	helpers.Write(
		"This is the first sentence",
		"This is the second sentence",
	)
}
