package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// --- DEFINING STRUCTS FOR JSON ---

type User struct {
	ID       int     `json:"id"`              // Normal field
	Name     string  `json:"name"`            // Renames the field for JSON
	Email    string  `json:"email,omitempty"` // Will be omitted from JSON if empty
	Password string  `json:"-"`               // Will be ignored by JSON
	Active   bool    `json:"active"`
	Balance  float64 `json:"balance,string"` // Marshal as string
	Address  Address `json:"address"`        // Nested struct
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

// --- MAIN FUNCTION ---

func main() {
	// --- CREATING A STRUCT AND MARSHALING TO JSON ---

	user := User{
		ID:       1,
		Name:     "Alice",
		Email:    "",
		Password: "secret123",
		Active:   true,
		Balance:  1050.75,
		Address: Address{
			Street: "123 Go Lane",
			City:   "Gopher City",
		},
	}

	jsonBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling:", err)
	}

	fmt.Println("JSON Output:")
	fmt.Println(string(jsonBytes)) // Pretty-printed JSON

	// --- UNMARSHALING JSON TO STRUCT ---

	jsonStr := `
	{
		"id": 2,
		"name": "Bob",
		"email": "bob@example.com",
		"active": false,
		"balance": "500.25",
		"address": {
			"street": "456 Gopher Rd",
			"city": "GoTown"
		}
	}`

	var user2 User
	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		log.Fatal("Error unmarshaling:", err)
	}

	fmt.Println("Decoded Struct:")
	fmt.Printf("%+v", user2)

	// --- UNMARSHALING TO GENERIC INTERFACE{} ---

	rawJSON := `{"foo": "bar", "number": 123, "active": true}`
	var generic map[string]interface{}

	err = json.Unmarshal([]byte(rawJSON), &generic)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generic Map:")
	for k, v := range generic {
		fmt.Printf("Key: %s, Value: %v, Type: %T", k, v, v)
	}

	// --- WORKING WITH JSON ARRAYS ---

	jsonArr := `[{"id":1,"name":"John"},{"id":2,"name":"Jane"}]`
	var users []User

	err = json.Unmarshal([]byte(jsonArr), &users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Array of Users:")
	for _, u := range users {
		fmt.Printf("User: %+v", u)
	}
}
