# User API

A RESTful User Management API built using Go, Fiber, PostgreSQL, and SQLC.

## Features

* Create User
* Get User By ID
* Update User
* Delete User
* List All Users
* Dynamic Age Calculation from Date of Birth
* Input Validation using Validator
* Structured Logging using Uber Zap
* SQLC Database Layer
* Pagination Support
* Request ID Middleware
* Request Duration Logging
* Docker Support
* Unit Testing

---

## Tech Stack

* Go
* Fiber
* PostgreSQL
* SQLC
* Uber Zap
* Validator
* Docker

---

## Project Structure

```text
user-api/
├── cmd/
│   └── server/
│       └── main.go
├── config/
├── db/
│   ├── migrations/
│   └── sqlc/
├── internal/
│   ├── handler/
│   ├── logger/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   └── service/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE user_api;
```

Create users table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

---

## Environment Variables

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=user_api

SERVER_PORT=8080
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Generate SQLC Files

```bash
sqlc generate
```

---

## Run Application

```bash
go run cmd/server/main.go
```

Server starts on:

```text
http://localhost:8080
```

---

## API Endpoints

### Create User

```http
POST /users
```

Request:

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### Get User By ID

```http
GET /users/1
```

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 36
}
```

---

### Update User

```http
PUT /users/1
```

Request:

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

Response:

```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### Delete User

```http
DELETE /users/1
```

Response:

```http
204 No Content
```

---

### List Users

```http
GET /users
```

Response:

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 36
  }
]
```

---

### Pagination

```http
GET /users?page=1&limit=10
```

---

## Run Tests

```bash
go test ./...
```

---

## Run with Docker

Build and start services:

```bash
docker-compose up --build
```

Stop services:

```bash
docker-compose down
```

---

## Assignment Requirements Completed

* CRUD APIs
* PostgreSQL Integration
* SQLC Usage
* Dynamic Age Calculation
* Validator Integration
* Uber Zap Logging
* Clean HTTP Status Codes
* Pagination
* Unit Test
* Request ID Middleware
* Request Duration Logging
* Docker Support
