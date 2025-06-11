package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// --- DATA STRUCTS ---

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Dummy in-memory data store
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// --- MAIN FUNCTION ---

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", usersHandler)         // GET
	http.HandleFunc("/users/create", createHandler) // POST
	http.HandleFunc("/users/query", queryHandler)   // with ?id=...
	http.HandleFunc("/external", externalAPIClient) // client call example

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// --- HANDLER: HOME PAGE ---

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Go HTTP Server!"))
}

// --- HANDLER: GET ALL USERS AS JSON ---

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// --- HANDLER: CREATE USER WITH JSON PAYLOAD ---

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	u.ID = len(users) + 1
	users = append(users, u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// --- HANDLER: QUERY PARAMETER (?id=2) ---

func queryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	for _, u := range users {
		if u.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}

	http.NotFound(w, r)
}

// --- CLIENT: MAKE GET REQUEST TO EXTERNAL API ---

func externalAPIClient(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		http.Error(w, "Failed to call external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
