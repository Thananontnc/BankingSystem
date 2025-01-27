package sqlhandle

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect TO Database
func ConnectToDB(username, password, hostname, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbname)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Verify the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully!")
	return nil
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
