package sqlhandle

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect TO Database
func OpenDB() (*sql.DB, error) {
	// Replace with your DB credentials
	dsn := "username:password@tcp(localhost:3306)/yourdbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
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

func InsertRegister(username, email, password, phone string) error {
	query := "INSERT INTO users (username,email,password,phone) VALUES (?,?,?,?)"
	_, err := DB.Exec(query, username, email, password, phone)
	if err != nil {
		return fmt.Errorf("fail insert register data to users table: %v", err)
	}
	return nil
}

// Function to insert a new user (signup)
func RegisterUser(db *sql.DB, email, hashedPassword string) error {
	_, err := db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, hashedPassword)
	return err
}

// Function to validate user login (check email and password)
func ValidateUser(db *sql.DB, email, password string) bool {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		log.Println("Error fetching user:", err)
		return false
	}
	// Compare the provided password with the stored hashed password
	return storedPassword == password // You should compare hashed passwords in a real scenario
}
