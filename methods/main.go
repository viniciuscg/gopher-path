package main

import (
	"fmt"
	"strings"
)

// --- STRUCT DEFINITION ---
// Real-world example: A User in an application
type User struct {
	Name  string
	Email string
	Age   int
}

// --- METHOD: VALUE RECEIVER ---
// Value receiver doesn't mutate original
func (u User) Greet() string {
	return fmt.Sprintf("Hello, %s!", u.Name)
}

// --- METHOD: POINTER RECEIVER ---
// Pointer receiver can mutate the struct
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

// --- METHOD: VALIDATION ---
// Useful in production to validate models before DB inserts
func (u User) IsAdult() bool {
	return u.Age >= 18
}

// --- METHOD: FORMATTING DATA ---
// Common for API serialization or printing
func (u User) EmailDomain() string {
	parts := strings.Split(u.Email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return "invalid"
}

// --- EMBEDDED STRUCTS ---
// Like inheritance via composition
type Admin struct {
	User  // Embedded User
	Level int
}

// --- METHOD ON EMBEDDED TYPE ---
func (a Admin) Promote() {
	fmt.Printf("%s has been promoted to admin level %d!", a.Name, a.Level)
}

func main() {
	fmt.Println("=== User Methods Demo ===")

	// Create a user
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   21,
	}

	// Call method with value receiver
	fmt.Println(user.Greet()) // Output: Hello, Alice!

	// Value receiver doesn't change original
	oldEmail := user.Email
	userCopy := user
	userCopy.UpdateEmail("new@change.com")
	fmt.Println("Email should be unchanged:", user.Email == oldEmail) // true

	// Pointer receiver — changes the original
	user.UpdateEmail("alice@newdomain.com")
	fmt.Println("Updated email:", user.Email) // Output: alice@newdomain.com

	// Validation method
	if user.IsAdult() {
		fmt.Println("User is an adult.")
	}

	// Custom formatting
	fmt.Println("Email domain:", user.EmailDomain()) // Output: newdomain.com

	// --- USING EMBEDDED STRUCTS ---
	admin := Admin{
		User:  User{Name: "Bob", Email: "bob@admin.com", Age: 35},
		Level: 5,
	}
	admin.Promote()                            // Uses method on Admin
	fmt.Println("Admin's email:", admin.Email) // Accesses embedded field
}
