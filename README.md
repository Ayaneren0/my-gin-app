## 🌐 Socials:
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?logo=linkedin&logoColor=white)](https://linkedin.com/in/www.linkedin.com/in/ayanahmad15) [![X](https://img.shields.io/badge/X-black.svg?logo=X&logoColor=white)](https://x.com/ayanAhm4d) 

# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
- **Dependencies:**
  - `github.com/dgrijalva/jwt-go` - JWT authentication
  - `github.com/go-sql-driver/mysql` - MySQL driver
  - `github.com/gorilla/mux` - HTTP router
  - `github.com/joho/godotenv` - Environment configuration
  - `golang.org/x/crypto` - Password hashing
# 📊 GitHub Stats:
![](https://github-readme-stats.vercel.app/api?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false)<br/>
![](https://github-readme-streak-stats.herokuapp.com/?user=ayanAhm4d&theme=dark&hide_border=false)<br/>
![](https://github-readme-stats.vercel.app/api/top-langs/?username=ayanAhm4d&theme=dark&hide_border=false&include_all_commits=false&count_private=false&layout=compact)

---
[![](https://visitcount.itsvg.in/api?id=ayanAhm4d&icon=0&color=0)](https://visitcount.itsvg.in)

# Golang Bookstore Management API

A robust and scalable RESTful API built with Go for managing bookstore operations. This project implements a modern, secure, and efficient backend system utilizing MySQL for data persistence and JWT for authentication. The API follows clean architecture principles with the repository pattern, making it maintainable and extensible.

## Features

- **Authentication & Authorization**
  - JWT-based authentication
  - User registration and login
  - Protected routes

- **Book Management**
  - CRUD operations for books
  - Advanced filtering and search
  - Pagination support
  - Sorting capabilities

- **Database**
  - MySQL integration
  - Repository pattern implementation
  - Efficient query handling




## Getting Started

### Prerequisites

- Go 1.21 or higher
- MySQL 5.7 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/ayanAhm4d/Golang-Bookstore-Management-API.git
cd bookstore-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create the database and tables:
```sql
CREATE DATABASE bookstore;
USE bookstore;

CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE books (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

4. Configure environment variables:
Create a `.env` file in the root directory:
```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=bookstore
JWT_SECRET=your_jwt_secret
```

5. Run the application:
```bash
go run cmd/main.go
```

## 📌 API Endpoints

### Authentication

- **Register User**
  ```
  POST /api/register
  ```
  ```json
  {
    "username": "user1",
    "password": "password123",
    "email": "user1@example.com"
  }
  ```

- **Login**
  ```
  POST /api/login
  ```
  ```json
  {
    "username": "user1",
    "password": "password123"
  }
  ```

### Books (Protected Routes)

All book endpoints require JWT token in the Authorization header:
`Authorization: Bearer <your_token>`

- **Create Book**
  ```
  POST /api/books
  ```
  ```json
  {
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "price": 49.99,
    "stock": 100
  }
  ```

- **Get All Books**
  ```
  GET /api/books
  ```
  Query Parameters:
  - `page`: Page number (default: 1)
  - `page_size`: Items per page (default: 10)
  - `search`: Search in title and author
  - `min_price`: Minimum price filter
  - `max_price`: Maximum price filter
  - `author`: Filter by author
  - `sort_by`: Sort field (title, author, price, stock, created_at)
  - `sort_dir`: Sort direction (asc, desc)

- **Get Single Book**
  ```
  GET /api/books/{id}
  ```

- **Update Book**
  ```
  PUT /api/books/{id}
  ```

- **Delete Book**
  ```
  DELETE /api/books/{id}
  ```

## 🔍 Example API Usage

### Get Books with Filtering
```bash
# Get books with pagination
curl "http://localhost:8080/api/books?page=1&page_size=10" \
  -H "Authorization: Bearer <your_token>"

# Search books
curl "http://localhost:8080/api/books?search=golang" \
  -H "Authorization: Bearer <your_token>"

# Filter by price range
curl "http://localhost:8080/api/books?min_price=20&max_price=50" \
  -H "Authorization: Bearer <your_token>"

# Sort books
curl "http://localhost:8080/api/books?sort_by=price&sort_dir=desc" \
  -H "Authorization: Bearer <your_token>"
```

## 📁 Project Structure

```
.
├── cmd
│   └── main.go
├── config
│   └── config.go
├── internal
│   ├── auth
│   │   └── jwt.go
│   ├── handlers
│   │   ├── auth_handler.go
│   │   └── book_handler.go
│   ├── middleware
│   │   └── auth_middleware.go
│   ├── models
│   │   ├── book.go
│   │   ├── pagination.go
│   │   └── user.go
│   └── repository
│       ├── book_repository.go
│       └── user_repository.go
├── go.mod
└── .env
```

## 🔒 Security

- Passwords are hashed using bcrypt
- JWT tokens are used for authentication
- SQL injection prevention through prepared statements
- Input validation and sanitization
- Protected routes using middleware

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request



