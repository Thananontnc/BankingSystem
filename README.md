# FinPluse Web Application

## Project Overview
FinPluse is a web-based financial application developed as a final project for the Design Thinking Subject. The application provides user authentication functionality and aims to deliver financial services through an intuitive web interface.

## Features
- User Authentication System
  - Login functionality
  - User registration system
  - Secure password handling
- Responsive Web Design
  - Modern UI/UX
  - Mobile-friendly interface
- Database Integration
  - Secure user data storage
  - SQL database backend

## Technical Stack
- **Backend:** Go (Golang)
- **Frontend:** HTML, CSS
- **Database:** SQL
- **Architecture:** MVC Pattern

## Project Structure
```
.
├── main.go              # Application entry point
├── webhandle/          # Web request handlers
├── sqlhandle/          # Database operations
└── static/             # Static files
    ├── css/           # Stylesheet files
    ├── Login.html     # Login page
    ├── Signup.html    # Registration page
    └── dashboard.html # User dashboard
```

## Setup Instructions
1. Clone the repository
2. Ensure Go is installed on your system
3. Set up your SQL database and update connection details
4. Run the following commands:
   ```bash
   go mod tidy
   go run main.go
   ```
5. Access the application at `http://localhost:8080`

## Development Team
This project is developed as part of the Design Thinking Subject coursework.

## Future Enhancements
- Implementation of financial features
- Advanced user dashboard
- Transaction tracking
- Financial analytics

## License
This project is developed for educational purposes.
Presented by 
Thananon Chounudom 
Prathomporn Bunjua
Napatra Hanwari