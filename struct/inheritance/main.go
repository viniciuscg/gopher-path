package main

import "fmt"

// Define a base struct called 'person'.
// This is a simple structure with a single field.
type person struct {
	name string
}

// Define another struct called 'student'.
// Instead of containing a field *named* `person`, it embeds the `person` struct directly.
// This is called **struct embedding** in Go.
//
// Struct embedding allows 'student' to automatically "inherit" the fields and methods of 'person'.
// But this is NOT inheritance like in OOP — it's just a way to promote fields/methods.
type student struct {
	person // Embedded struct — fields of 'person' become part of 'student'
	course string
}

func main() {
	// NOTE: Go is NOT an object-oriented language.
	// It doesn’t support classical inheritance.
	// Instead, Go encourages composition through **struct embedding** like above.

	// Create an instance of 'person' and assign a value to the name field.
	var p person = person{
		name: "Vinicius",
	}

	// Create an instance of 'student', embedding the previously defined person.
	// We assign the entire 'person' struct to the embedded field,
	// and also set a specific course value.
	var s student = student{
		person: p,     // Embed the 'person' instance
		course: "ASD", // Assign course directly
	}

	// Accessing fields of the embedded struct:
	// Because 'person' is embedded, we can access 's.name' directly,
	// as if 'name' were declared on 'student' itself.
	// This is called **field promotion**.
	fmt.Println(s.name) // Output: Vinicius

	// You could also write: fmt.Println(s.person.name)
	// Both are valid; use whichever makes your code clearer.

	// Output the full student struct for reference.
	fmt.Printf("Student Info: %+v", s)
}
