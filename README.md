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

```bash
# Step 1: Clone and enter the project
git clone https://github.com/yourusername/go-task-api.git
cd go-task-api

# Step 2: Download dependencies
go mod tidy

# Step 3: Run the server
go run main.go
Server runs at: http://localhost:8080

🔐 Auth Flow
Register a User
arduino
Copy
Edit
POST /register
{
  "username": "sumit",
  "password": "secret123"
}
Login to Get Token
bash
Copy
Edit
POST /login
{
  "username": "sumit",
  "password": "secret123"
}
Response:

json
Copy
Edit
{ "token": "<your_jwt_token>" }
Add this token as a Bearer token in the Authorization header for all /tasks routes.

📬 Task Routes
Method	Endpoint	Description	Auth Required
POST	/register	Create new user	❌
POST	/login	Login & get token	❌
GET	/tasks	List user’s tasks	✅
POST	/tasks	Create a new task	✅
PUT	/tasks/:id	Update a task	✅
DELETE	/tasks/:id	Delete a task	✅

🧪 Sample Authenticated Request
bash
Copy
Edit
curl -X GET http://localhost:8080/tasks \
  -H "Authorization: Bearer <your_token>"
📄 License
This project is licensed under the MIT License.

🙌 About Me
Built with ❤️ by Sumit Akoliya
🔗 LinkedIn
💻 GitHub

yaml
Copy
Edit

---

✅ Just paste this into your `README.md`. If you'd like a `Dockerfile`, `Makefile`, or Postman collection next, let me know!






You said:
create the entire thing in readme.md format 


ChatGPT said:
Here’s your full README.md — just copy and paste the entire thing into your project’s root as README.md.

markdown
Copy
Edit
# 🧠 Task Manager API – Golang + JWT Auth

A secure, user-specific Task Manager REST API built in **Go**, using **Gin**, **GORM**, **SQLite**, and **JWT Authentication**.  
It demonstrates backend fundamentals like RESTful API design, token-based security, clean project architecture, and user-bound data access — perfect for showcasing in technical interviews or portfolio submissions.

---

## 🚀 Features

- 🔐 **JWT Authentication** – Secure login & registration
- 👤 **User Ownership** – Each user accesses only their tasks
- ✅ **CRUD for Tasks** – Create, read, update, delete your own tasks
- 🧱 **Clean Architecture** – Modular design with separation of concerns
- 💾 **SQLite + GORM** – Lightweight ORM-backed database

---

## 🧰 Tech Stack

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM ORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/index.html)
- [JWT](https://jwt.io/)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

## 📂 Folder Structure

go-task-api/
├── main.go # Entry point
├── models/ # Database models
│ ├── user.go # User model with password methods
│ └── task.go # Task model linked to user
├── controllers/ # Request handling logic
│ ├── auth.go # Register/Login
│ └── task.go # Task CRUD (JWT protected)
├── middleware/ # Route protection
│ └── auth.go # JWT validation
└── utils/ # Token utilities
└── token.go # JWT generation & parsing

yaml
Copy
Edit

---

## ⚙️ Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/yourusername/go-task-api.git
cd go-task-api
2. Install Dependencies
bash
Copy
Edit
go mod tidy
3. Run the Server
bash
Copy
Edit
go run main.go
Server starts at: http://localhost:8080

🔐 Auth Flow
✅ Register a User
http
Copy
Edit
POST /register
Content-Type: application/json

{
  "username": "sumit",
  "password": "secret123"
}
🔑 Login to Get JWT Token
http
Copy
Edit
POST /login
Content-Type: application/json

{
  "username": "sumit",
  "password": "secret123"
}
Response:

json
Copy
Edit
{
  "token": "eyJhbGciOi..."
}
Copy this token and use it in the Authorization header for all task-related endpoints:

makefile
Copy
Edit
Authorization: Bearer <your_token>
📬 Task API Endpoints
Method	Endpoint	Description	Auth Required
POST	/register	Create a new user	❌ No
POST	/login	Login and get token	❌ No
GET	/tasks	List tasks for the user	✅ Yes
POST	/tasks	Create a new task	✅ Yes
PUT	/tasks/:id	Update a task by ID	✅ Yes
DELETE	/tasks/:id	Delete a task by ID	✅ Yes

🧪 Example with curl
bash
Copy
Edit
# Register
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"sumit", "password":"secret123"}'

# Login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"sumit", "password":"secret123"}'

# Authenticated Task Creation
curl -X POST http://localhost:8080/tasks \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Finish Go App", "status":"in-progress"}'
📄 License
This project is licensed under the MIT License.

🙌 Author
Sumit Akoliya
📍 Sheffield, UK
🔗 LinkedIn
💻 GitHub

