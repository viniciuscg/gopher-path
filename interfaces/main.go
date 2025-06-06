package main

import "fmt"

// --- INTERFACE DEFINITION ---
// An interface defines a *behavior*, not a data structure.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// --- STRUCTS THAT IMPLEMENT THE INTERFACE ---
// You don't need to explicitly "implement" in Go.
// If a type satisfies all the methods, it implements the interface automatically.

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// --- COMPOSED INTERFACES ---
// One interface can include others (interface composition)
type DetailedShape interface {
	Shape
	Description() string
}

// --- STRUCT IMPLEMENTING COMPOSED INTERFACE ---
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func (s Square) Description() string {
	return "I am a square."
}

// --- USING INTERFACES ---
func printShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

// --- EMPTY INTERFACE (interface{}) ---
// Accepts *any* type (like `any` in other languages)
func printAnything(v interface{}) {
	fmt.Println("Value:", v)

	// --- TYPE ASSERTION ---
	if str, ok := v.(string); ok {
		fmt.Println("It's a string:", str)
	}

	// --- TYPE SWITCH ---
	switch val := v.(type) {
	case int:
		fmt.Println("It's an int:", val)
	case float64:
		fmt.Println("It's a float:", val)
	case string:
		fmt.Println("It's a string (from switch):", val)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	fmt.Println("=== Interface Basics ===")

	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}
	sq := Square{Side: 2}

	printShapeInfo(c)
	printShapeInfo(r)
	printShapeInfo(sq)

	fmt.Println("Detailed description:", sq.Description())

	// --- Empty interface use cases ---
	printAnything(42)
	printAnything("hello")
	printAnything(3.14)
}
