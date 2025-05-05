# 🧠 Task Manager API – Golang + JWT Auth

A secure, user-specific Task Manager REST API built in **Go**, using **Gin**, **GORM**, **SQLite**, and **JWT Authentication**. This project demonstrates modern backend development practices including clean code architecture, token-based security, and user-bound data access.

---

## 🚀 Features

- 🔐 **JWT Authentication** – Secure login & registration
- 👤 **User Ownership** – Each user accesses only their tasks
- ✅ **CRUD for Tasks** – Create, read, update, delete tasks
- 🧱 **Clean Architecture** – Modular with models, controllers, middleware
- 💾 **SQLite + GORM** – Lightweight ORM-powered storage

---

## 🧰 Tech Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/index.html)
- [JWT](https://jwt.io/)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

## 📂 Folder Structure

go-task-api/
├── main.go # App entry point
├── models/ # User and Task models
│ ├── user.go
│ └── task.go
├── controllers/ # Business logic
│ ├── auth.go
│ └── task.go
├── middleware/ # JWT middleware
│ └── auth.go
└── utils/ # Token utilities
└── token.go

yaml
Copy
Edit

---

## ⚙️ How to Run

# Step 1: Clone and enter the project
git clone https://github.com/yourusername/go-task-api.git
cd go-task-api

# Step 2: Download dependencies
go mod tidy

# Step 3: Run the server
go run main.go
Server runs at: http://localhost:8080

## 🔐 Auth Flow
Register a User

POST /register
{
  "username": "sumit",
  "password": "secret123"
}

Login to Get Token

POST /login
{
  "username": "sumit",
  "password": "secret123"
}

Response:

{ "token": "<your_jwt_token>" }

# Add this token as a Bearer token in the Authorization header for all /tasks routes.

## 📬 Task Routes
Method	Endpoint	Description	Auth Required
POST	/register	Create new user	❌
POST	/login	Login & get token	❌
GET	/tasks	List user’s tasks	✅
POST	/tasks	Create a new task	✅
PUT	/tasks/:id	Update a task	✅
DELETE	/tasks/:id	Delete a task	✅

# 🧪 Sample Authenticated Request

curl -X GET http://localhost:8080/tasks \
  -H "Authorization: Bearer <your_token>"

This project is licensed under the MIT License.

🙌 About Me
Built with ❤️ by Sumit Akoliya
🔗 LinkedIn
💻 GitHub
