package webhandle

import (
	"log"
	"net/http"

	"github.com/Thananontnc/BankingSystem.git/sqlhandle"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func SetupRoutes() {
	// Serve static files (CSS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve HTML files for login, signup, and home pages

	// Handle user registration
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse the form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}

			// Get form values
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Hash the password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Error processing password", http.StatusInternalServerError)
				return
			}

			// Open database connection
			db, err := sqlhandle.OpenDB()
			if err != nil {
				http.Error(w, "Database connection error", http.StatusInternalServerError)
				return
			}
			defer db.Close() // Close the specific connection

			// Register the user
			log.Printf("Attempting to register user with email: %s", email)
			err = sqlhandle.RegisterUser(db, email, string(hashedPassword), username)
			if err != nil {
				log.Printf("Failed to register user: %v", err)
				http.Error(w, "Error registering user", http.StatusInternalServerError)
				return
			}
			log.Printf("Successfully registered user with email: %s", email)

			// Redirect to login page after successful registration
			http.Redirect(w, r, "/Login.html", http.StatusSeeOther)
			return
		} else {
			// Serve the signup page for GET requests
			http.ServeFile(w, r, "static/Signup.html")
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/Index.html")
	})

	http.HandleFunc("/Login.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/Login.html")
	})

	http.HandleFunc("/Signup.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/Signup.html")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Open DB connection
			db, err := sqlhandle.OpenDB()
			if err != nil {
				http.Error(w, "Database connection failed", http.StatusInternalServerError)
				return
			}
			defer db.Close()

			// Validate user login
			if sqlhandle.ValidateUser(db, email, password) {
				// Successful login, redirect to homepage
				http.Redirect(w, r, "/Index.html", http.StatusSeeOther)
			} else {
				// Invalid login credentials
				http.Error(w, "Invalid login", http.StatusUnauthorized)
			}
		} else {
			// Serve the login page for GET requests
			http.ServeFile(w, r, "static/Login.html")
		}
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			email := r.FormValue("email")
			password := r.FormValue("password")
			username := r.FormValue("username")

			// Hash the password (never store plain text passwords)
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Password hashing failed", http.StatusInternalServerError)
				return
			}

			// Open DB connection
			db, err := sqlhandle.OpenDB()
			if err != nil {
				http.Error(w, "Database connection failed", http.StatusInternalServerError)
				return
			}
			defer db.Close()

			// Register new user
			if err := sqlhandle.RegisterUser(db, email, string(hashedPassword), username); err != nil {
				http.Error(w, "Registration failed", http.StatusInternalServerError)
				return
			}

			// Redirect to login page after successful signup
			http.Redirect(w, r, "/Login.html", http.StatusSeeOther)
		} else {
			// Serve the signup page for GET requests
			http.ServeFile(w, r, "static/Signup.html")
		}
	})
}
func StartServer() {
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
