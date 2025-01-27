package webhandle

import (
	"log"
	"net/http"
)

func SetupRoutes() {
	// Serve static files (CSS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve HTML files for login, signup, and home pages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/Index.html")
	})

	http.HandleFunc("/Login.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/Login.html")
	})

	http.HandleFunc("/Signup.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/Signup.html")
	})

	// Handle form submissions for login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Process login (authentication)
			http.Redirect(w, r, "/Index.html", http.StatusSeeOther)
		} else {
			http.ServeFile(w, r, "static/Login.html")
		}
	})

	// Handle form submissions for registration
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Process signup (create user)
			http.Redirect(w, r, "/dashboard.html", http.StatusSeeOther)
		} else {
			http.ServeFile(w, r, "static/Signup.html")
		}
	})
}
func StartServer() {
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
