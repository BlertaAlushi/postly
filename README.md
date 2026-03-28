# 🚀 Postly API

Postly is a RESTful backend API built with **Go (Gin)** and **PostgreSQL**, providing core features of a social media platform such as authentication, posts, likes, comments, follow system, and personalized feeds.

---

## 📌 Features

* JWT Authentication (access + refresh tokens)
* User registration & login
* Create, update, delete posts
* Like and comment on posts
* Follow / unfollow users
* Feed, explore, and user posts
* Search users

---

## 🛠 Tech Stack

* Language: Go
* Framework: Gin
* Database: PostgreSQL
* Auth: JWT (github.com/golang-jwt/jwt/v5)
* Validation: go-playground/validator
* Migrations: golang-migrate

---

## 📁 Project Structure

```
postly/
│
├── configs/        # DB & app configuration
├── controllers/    # HTTP handlers
├── interfaces/     # Interfaces (e.g. Normalizer)
├── middlewares/    # JWT auth middleware
├── migrations/     # Database migrations
├── models/         # Request & response models
├── repositories/   # Database queries
├── routes/         # Route definitions
├── services/       # Business logic
│
├── .env
├── go.mod
└── main.go
```
---

## ⚙️ Setup & Installation

### 1. Clone project

```
https://github.com/BlertaAlushi/postly.git
cd postly
```

### 2. Create `.env`

```
PORT=8080

DB_CONNECTION=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=yourpassword
DB_DATABASE=postly

JWT_SECRET=your_secret
```

### 3. Run migrations

```
migrate -path migrations -database "postgres://user:password@localhost:5432/postly?sslmode=disable" up
```

### 4. Run the server

```
go run main.go
```

---

## 🔐 Authentication

* Uses JWT Access Token + Refresh Token
* Access token required for protected routes
* Refresh token used to generate new access tokens
* Logout invalidates refresh token

---

## API Routes

### Public

* `POST /api/register`
* `POST /api/login`
* `POST /api/token/refresh`

---

### Protected (require JWT)

#### User

* `POST /api/logout`
* `GET /api/users`

#### Posts

* `POST /api/posts`
* `GET /api/posts/:id`
* `PUT /api/posts/:id`
* `DELETE /api/posts/:id`

#### Feed

* `GET /api/feed`
* `GET /api/explore`
* `GET /api/users/:id/posts`

#### Likes

* `GET /api/posts/:id/likes`
* `POST /api/posts/:id/like`
* `DELETE /api/posts/:id/like`

#### Comments

* `GET /api/posts/:id/comments`
* `POST /api/posts/:id/comments`
* `GET /api/posts/:id/comments/:comment_id`
* `PUT /api/posts/:id/comments/:comment_id`
* `DELETE /api/posts/:id/comments/:comment_id`

#### Follow

* `POST /api/follow/:follow_id`
* `DELETE /api/follow/:follow_id`
* `GET /api/users/:id/followers`
* `GET /api/users/:id/following`

---

## 👩‍💻 Author

Blerta Alushi
