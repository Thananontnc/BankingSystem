package webhandle

import (
	"net/http"
)

// Main Page Server
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

// Login Page Server
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/login.html")
}

// Register Page Server
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/Signup.html")
}

// func dashboardHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./static/dashboard.html")
// }
