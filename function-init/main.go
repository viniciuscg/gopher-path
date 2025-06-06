package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	db           *sql.DB // Simulated DB connection
	globalConfig map[string]string
)

// --- INIT FUNCTION ---
// Automatically runs BEFORE main()
// Great for configuration, environment setup, logging, etc.
func init() {
	fmt.Println("=== [init()] STARTUP ===")

	// --- ENVIRONMENT SETUP ---
	globalConfig = map[string]string{
		"ENV":         os.Getenv("APP_ENV"),     // like "development" or "production"
		"DB_CONN_STR": os.Getenv("DB_CONN_STR"), // like "postgres://user:pass@localhost/db"
	}

	if globalConfig["ENV"] == "" {
		globalConfig["ENV"] = "development"
		fmt.Println("Defaulting ENV to:", globalConfig["ENV"])
	}

	// --- SIMULATED DB CONNECTION SETUP ---
	// Normally you'd use: db, err := sql.Open(...)
	mockDBConnect()

	// --- LOGGING SETUP ---
	log.SetPrefix("[APP LOG] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	fmt.Println("Config initialized:", globalConfig)
	fmt.Println("=== [init()] DONE ===")
}

// --- MAIN FUNCTION ---
func main() {
	fmt.Println("=== [main()] RUNNING ===")

	// Use config or db connection already prepared in init()
	fmt.Println("App is running in:", globalConfig["ENV"])
	fmt.Println("Pretending to query the database...")

	if db != nil {
		fmt.Println("Database is ready to use!")
	} else {
		fmt.Println("No database connection available.")
	}

	fmt.Println("=== [main()] DONE ===")
}

// --- MOCK FUNCTION FOR DB CONNECTION ---
// This simulates a DB connection for demonstration purposes.
func mockDBConnect() {
	fmt.Println("Connecting to database (mock)...")

	// Simulating connection success
	db = &sql.DB{} // mock (would be real in actual code)

	// If you wanted to simulate a failure:
	// db = nil
}
