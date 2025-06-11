package main

import (
	"fmt"
	"yourmodule/mathutils"
)

func main() {
	sum := mathutils.Add(10, 5)
	fmt.Println("Sum:", sum)

	quotient := mathutils.Divide(20, 4)
	fmt.Println("Quotient:", quotient)

	// Uncomment to see panic in action
	// mathutils.Divide(10, 0)
}
