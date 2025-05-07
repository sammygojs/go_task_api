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
xw
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

## 👀 Code Structure & Flow

This section explains how the project works under the hood and how each file/module fits together so that any developer can understand and extend it easily.

---

### 📋 Core Features Implemented

| Feature                     | Description                                                             |
| --------------------------- | ----------------------------------------------------------------------- |
| 🔐 JWT Authentication       | Users can register and log in. Authenticated users receive a JWT token. |
| 👤 User-specific tasks      | All task operations are tied to the authenticated user.                 |
| 🧱 Modular folder structure | Controllers, middleware, models, and utils are separated by concern.    |
| ✅ Task CRUD                 | Authenticated users can create, view, update, and delete their tasks.   |
| 🔒 Route protection via JWT | Only logged-in users can access `/tasks` routes.                        |

---

### 🗾️ Code Flow Overview

#### 🟩 `main.go` – Entry Point

* Initializes the **SQLite** database.
* Auto-migrates `User` and `Task` models.
* Creates a **Gin** router.
* Defines routes:

  * **Public routes**: `/register`, `/login`
  * **Protected routes**: all `/tasks` routes (require JWT)

```go
r.POST("/register", controllers.Register)
r.POST("/login", controllers.Login)

auth := r.Group("/")
auth.Use(middleware.AuthMiddleware())
{
    auth.GET("/tasks", controllers.GetTasks)
    auth.POST("/tasks", controllers.CreateTask)
    auth.PUT("/tasks/:id", controllers.UpdateTask)
    auth.DELETE("/tasks/:id", controllers.DeleteTask)
}
```

---

#### 📦 `/models`

##### 🔑 `user.go`

Defines the `User` model with:

* Fields: `Username`, `Password`
* Methods:

  * `SetPassword()` – Hashes password using bcrypt
  * `CheckPassword()` – Compares password with stored hash

##### 📄 `task.go`

Defines the `Task` model with:

* Fields: `Title`, `Status`, `UserID` (linked to user)

---

#### ⚙️ `/controllers`

##### 🔐 `auth.go`

* `POST /register` – Creates a new user with a hashed password.
* `POST /login` – Authenticates user and returns a JWT if successful.

##### 📋 `task.go`

All task routes:

* Require a valid JWT
* Read the authenticated `userID` from context
* Only access tasks belonging to that user

---

#### 🛡️ `/middleware/auth.go`

* Validates the `Authorization: Bearer <token>` header.
* Parses the JWT.
* If valid, stores `userID` in `context`.
* If invalid, blocks access with `401 Unauthorized`.

---

#### 🔧 `/utils/token.go`

JWT helper functions:

* `GenerateToken(userID)` – Creates a signed token with expiration.
* `ParseToken(token)` – Validates and decodes token into claims.

Uses `HS256` signing with a static secret key.

---

### 🔀 End-to-End Request Lifecycle

1. **User registers**

   * Calls `POST /register`
   * Password is hashed and saved

2. **User logs in**

   * Calls `POST /login`
   * Returns a JWT

3. **User creates task**

   * Sends `POST /tasks` with token
   * Middleware extracts `userID`
   * Controller creates task with that `userID`

4. **User fetches tasks**

   * `GET /tasks` returns only tasks for the logged-in user

---

### 🔍 Example Token Usage

After logging in:

```http
Authorization: Bearer <your_token_here>
```

Use this header on all `/tasks` routes.

---

## 👀 Code Structure & Flow

This section explains how the project works under the hood and how each file/module fits together so that any developer can understand and extend it easily.

---

### 📋 Core Features Implemented

| Feature                     | Description                                                             |
| --------------------------- | ----------------------------------------------------------------------- |
| 🔐 JWT Authentication       | Users can register and log in. Authenticated users receive a JWT token. |
| 👤 User-specific tasks      | All task operations are tied to the authenticated user.                 |
| 🧱 Modular folder structure | Controllers, middleware, models, and utils are separated by concern.    |
| ✅ Task CRUD                 | Authenticated users can create, view, update, and delete their tasks.   |
| 🔒 Route protection via JWT | Only logged-in users can access `/tasks` routes.                        |

---

### 🗾️ Code Flow Overview

#### 🟩 `main.go` – Entry Point

* Initializes the **SQLite** database.
* Auto-migrates `User` and `Task` models.
* Creates a **Gin** router.
* Defines routes:

  * **Public routes**: `/register`, `/login`
  * **Protected routes**: all `/tasks` routes (require JWT)

```go
r.POST("/register", controllers.Register)
r.POST("/login", controllers.Login)

auth := r.Group("/")
auth.Use(middleware.AuthMiddleware())
{
    auth.GET("/tasks", controllers.GetTasks)
    auth.POST("/tasks", controllers.CreateTask)
    auth.PUT("/tasks/:id", controllers.UpdateTask)
    auth.DELETE("/tasks/:id", controllers.DeleteTask)
}
```

---

#### 📦 `/models`

##### 🔑 `user.go`

Defines the `User` model with:

* Fields: `Username`, `Password`
* Methods:

  * `SetPassword()` – Hashes password using bcrypt
  * `CheckPassword()` – Compares password with stored hash

##### 📄 `task.go`

Defines the `Task` model with:

* Fields: `Title`, `Status`, `UserID` (linked to user)

---

#### ⚙️ `/controllers`

##### 🔐 `auth.go`

* `POST /register` – Creates a new user with a hashed password.
* `POST /login` – Authenticates user and returns a JWT if successful.

##### 📋 `task.go`

All task routes:

* Require a valid JWT
* Read the authenticated `userID` from context
* Only access tasks belonging to that user

---

#### 🛡️ `/middleware/auth.go`

* Validates the `Authorization: Bearer <token>` header.
* Parses the JWT.
* If valid, stores `userID` in `context`.
* If invalid, blocks access with `401 Unauthorized`.

---

#### 🔧 `/utils/token.go`

JWT helper functions:

* `GenerateToken(userID)` – Creates a signed token with expiration.
* `ParseToken(token)` – Validates and decodes token into claims.

Uses `HS256` signing with a static secret key.

---

### 🔀 End-to-End Request Lifecycle

1. **User registers**

   * Calls `POST /register`
   * Password is hashed and saved

2. **User logs in**

   * Calls `POST /login`
   * Returns a JWT

3. **User creates task**

   * Sends `POST /tasks` with token
   * Middleware extracts `userID`
   * Controller creates task with that `userID`

4. **User fetches tasks**

   * `GET /tasks` returns only tasks for the logged-in user

---

### 🔍 Example Token Usage

After logging in:

```http
Authorization: Bearer <your_token_here>
```

Use this header on all `/tasks` routes.

---

### 💪 Testing & Debugging

Use Postman or curl to test the API manually.
Each task is scoped to the logged-in user only.
Unauthenticated access to protected endpoints returns a `401`.

---

### 🔮 What's Next?

You can extend this project with:

* 📁 Projects (1-to-many task grouping)
* 🔖 Tags (many-to-many)
* 📄 Swagger docs (OpenAPI)
* 💪 Docker support
* ✅ Unit and integration tests
* 🌐 Deployment on Railway, Fly.io, or Render


### 💪 Testing & Debugging

Use Postman or curl to test the API manually.
Each task is scoped to the logged-in user only.
Unauthenticated access to protected endpoints returns a `401`.

---

### 🔮 What's Next?

You can extend this project with:

* 📁 Projects (1-to-many task grouping)
* 🔖 Tags (many-to-many)
* 📄 Swagger docs (OpenAPI)
* 💪 Docker support
* ✅ Unit and integration tests
* 🌐 Deployment on Railway, Fly.io, or Render


This project is licensed under the MIT License.

🙌 About Me
Built with ❤️ by Sumit Akoliya
🔗 LinkedIn
💻 GitHub
