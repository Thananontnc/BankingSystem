package main

import (
	"fmt"
	"net/http"

	// "github.com/Thananontnc/BankingSystem.git/sqlhandle"

	"github.com/Thananontnc/BankingSystem.git/webhandle"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database Connection

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Server
	http.HandleFunc("/", webhandle.HomeHandler)
	http.HandleFunc("/index.html", webhandle.HomeHandler)
	http.HandleFunc("/login.html", webhandle.LoginHandler)
	http.HandleFunc("/Signup.html", webhandle.RegisterHandler)

	fmt.Println("Run at port http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
