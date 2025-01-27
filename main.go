package main

import (
	"log"
	"github.com/Thananontnc/BankingSystem.git/sqlhandle"
	"github.com/Thananontnc/BankingSystem.git/webhandle"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Printf("Initializing application...")
	
	// Test database connection first
	db, err := sqlhandle.OpenDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	log.Printf("Database connection successful")
	
	// Initialize routes
	webhandle.SetupRoutes()
	
	// Start the server
	log.Printf("Starting server on :8080")
	webhandle.StartServer()
}