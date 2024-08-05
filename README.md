## 🌐 Socials:
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?logo=linkedin&logoColor=white)]www.linkedin.com/in/ayan-ahmad-ansari) 

# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white) ![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

# Book API

This is a simple RESTful API for managing books, built with Go and the Gin web framework. It demonstrates basic CRUD operations, database, middleware usage, and route grouping.

## Features

- List all books
- Get a specific book by ID
- Create a new book
- Update a book
- Delete a book
- Custom logging middleware
- Grouped routing

## Prerequisites

- Go 1.16 or higher
- Gin web framework

## Installation

1. Clone the repository:
`https://github.com/Ayaneren0/book-api.git`

2. Install dependencies:
go mod tidy

The server will start on `http://localhost:8080`.

## API Endpoints

- `GET /api`: Get all books
- `GET /api/:id`: Get a specific book by ID
- `POST /api`: Create a new book
- `PUT /api`: Update book by give ID 
- `DELETE /api`: Delete book by given ID

## Project Structure
<P>book-api/</br>
├── main.go</br>
├── handlers/</br>
│   └── books.go</br>
├── middleware/</br>
│   └── logger.go</br>
└── models/</br>
└── book.go</br>
</P>

- `main.go`: Entry point of the application, sets up routing and middleware
- `handlers/books.go`: Contains handler functions for book-related operations
- `middleware/logger.go`: Custom logging middleware
- `models/book.go`: Defines the Book struct
- `models/db.go`: Contains database function for store books informations 
  
## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.


