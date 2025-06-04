package main

import "fmt"

// --- BASIC FUNCTION DECLARATION ---

// A simple function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// --- MULTIPLE RETURN VALUES ---

// Function returning two values
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// --- NAMED RETURN VALUES ---

// Return values are named and can be returned implicitly
func rectangleStats(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Implicitly returns named variables
}

// --- VARIADIC FUNCTIONS ---

// Accepts any number of int arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// --- ANONYMOUS FUNCTIONS ---

func demoAnonymous() {
	// Immediately invoked function
	result := func(a, b int) int {
		return a * b
	}(3, 4)
	fmt.Println("Anonymous function result:", result)

	// Function assigned to a variable
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Multiply via variable:", multiply(5, 6))
}

// --- PASSING FUNCTIONS AS PARAMETERS ---

func operate(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// --- RETURNING FUNCTIONS ---

func getMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// --- DEFER WITH FUNCTIONS ---

func withDefer() {
	defer fmt.Println("Deferred: runs last")
	fmt.Println("Normal: runs first")
}

// --- METHODS (FUNCTIONS WITH RECEIVERS) ---

type Circle struct {
	radius float64
}

// Method with value receiver
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Method with pointer receiver
func (c *Circle) scale(factor float64) {
	c.radius *= factor
}

// --- FACTORIAL USING RECURSION ---
func factorial(n int) int {
	if n == 0 {
		return 1 // base case
	}
	return n * factorial(n-1) // recursive case
}

// --- FIBONACCI USING RECURSION ---
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Basic function
	fmt.Println("Add:", add(3, 4))

	// Multiple returns
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	// Named returns
	a, p := rectangleStats(5, 3)
	fmt.Println("Area:", a, "Perimeter:", p)

	// Variadic function
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	// Anonymous functions
	demoAnonymous()

	// Passing functions
	res := operate(10, 5, func(x, y int) int {
		return x - y
	})
	fmt.Println("Operate result:", res)

	// Returning functions
	double := getMultiplier(2)
	fmt.Println("Double 5:", double(5))

	// Defer example
	withDefer()

	// Methods
	c := Circle{radius: 3}
	fmt.Println("Circle area:", c.area())
	c.scale(2)
	fmt.Println("Scaled circle area:", c.area())

	// --- BASIC RECURSION ---

	// Recursion is when a function calls itself.
	// Go fully supports recursion, but make sure to define a clear base case to avoid infinite calls.

	fmt.Println("Factorial of 5:", factorial(5)) // Output: 120
	fmt.Println("Fibonacci of 6:", fibonacci(6)) // Output: 8

	// --- TAIL RECURSION (Go does not optimize it!) ---

	// Go does NOT perform tail call optimization (TCO).
	// So deep recursion can still cause a stack overflow.
}

// NOTES:
// - Parameters are passed by value by default.
// - You can pass pointers to allow modification.
// - You can use `_` to ignore returned values.
// - Functions are first-class citizens: assignable, passable, returnable.
// - Methods can have value or pointer receivers — pointer allows mutation.
// - `defer` delays function execution until the surrounding function returns.
// - Closures (functions capturing outer variables) are fully supported.
// - Variadic functions must have the variadic parameter last.
// - Use `...` to expand a slice into variadic arguments.
