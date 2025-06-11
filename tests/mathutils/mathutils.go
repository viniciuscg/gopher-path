package mathutils

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Divide returns the result of integer division. Panics if dividing by zero.
func Divide(a, b int) int {
	if b == 0 {
		panic("cannot divide by zero")
	}
	return a / b
}
