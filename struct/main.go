package main

import "fmt"

// Define a custom type 'user' using a struct.
// Structs in Go are like lightweight classes or records.
// This 'user' type has three fields: name, age, and address (which is another struct).
type user struct {
	name    string  `json:"my-name"` // User's name
	age     uint    // User's age (unsigned integer)
	address address // Nested struct for address details
}

// Define a second struct type named 'address'.
// This will be embedded inside the 'user' struct above.
type address struct {
	street string // Street name
	number uint   // House/building number
}

func main() {
	// Declare a variable of type 'user' with its zero values.
	// All fields will be initialized to their zero values:
	// - name: ""
	// - age: 0
	// - address.street: ""
	// - address.number: 0
	var user1 user

	// Assign values to the user1 fields individually.
	user1.name = "Vinicius"
	user1.age = 23

	// Print the struct — Go will automatically format it with field names.
	// Output: {Vinicius 23 { 0}}
	fmt.Println(user1)

	// Declare and initialize a 'user' using a **composite literal**.
	// We omit the address field here; it will remain with zero values.
	user2 := user{
		name: "Golang",
		age:  10,
	}
	// Output: {Golang 10 { 0}}
	fmt.Println(user2)

	// First, create an 'address' instance using a composite literal.
	// This will be assigned to the address field of the next user.
	user3Address := address{
		street: "Some Street",
		number: 123,
	}

	// Now create a full 'user' struct including the nested 'address'.
	user3 := user{
		name:    "Vinicius",
		age:     23,
		address: user3Address,
	}

	// Output: {Vinicius 23 {Some Street 123}}
	fmt.Println(user3)

	// You can also access nested fields directly like this:
	fmt.Println("User3 lives on:", user3.address.street)
}
