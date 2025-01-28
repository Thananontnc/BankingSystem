package sqlhandle

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Thananontnc/BankingSystem.git/utils"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect TO Database
func OpenDB() (*sql.DB, error) {
	log.Printf("Attempting to connect to database...")

	dsn := "root:Plai1412@tcp(localhost:3306)/bankingsystem?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	return db, nil
}

// Close Database
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection Closed")
	}
}

// Function to insert a new user (signup)
func RegisterUser(db *sql.DB, email, hashedPassword, username string) error {
	result, err := db.Exec("INSERT INTO users (email, password, full_name) VALUES (?, ?, ?)", email, hashedPassword, username)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error checking rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("No rows were inserted")
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}

// Function to validate user login (check email and password)
func ValidateUser(db *sql.DB, email, password string) bool {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		log.Println("Error fetching user:", err)
		return false
	}

	// Use the secure password comparison function
	return utils.ComparePasswords(storedPassword, password)
}
