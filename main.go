package main

import (
	"fmt"

	// "github.com/Thananontnc/BankingSystem.git/sqlhandle"

	"github.com/Thananontnc/BankingSystem.git/webhandle"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database Connection

	webhandle.SetupRoutes()
	fmt.Println("Run at port http://localhost:8080")
	// Start the server
	webhandle.StartServer()

}
