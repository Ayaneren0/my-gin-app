package main

import (
	"bookstore/config"
	"bookstore/internal/handlers"
	"bookstore/internal/middleware"
	"bookstore/internal/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookRepo)
	authHandler := handlers.NewAuthHandler(userRepo)

	// Initialize router
	router := mux.NewRouter()

	// Auth routes
	router.HandleFunc("/api/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/api/login", authHandler.Login).Methods("POST")

	// Book routes (protected by auth middleware)
	router.HandleFunc("/api/books", middleware.AuthMiddleware(bookHandler.CreateBook)).Methods("POST")
	router.HandleFunc("/api/books", middleware.AuthMiddleware(bookHandler.GetBooks)).Methods("GET")
	router.HandleFunc("/api/books/{id}", middleware.AuthMiddleware(bookHandler.GetBook)).Methods("GET")
	router.HandleFunc("/api/books/{id}", middleware.AuthMiddleware(bookHandler.UpdateBook)).Methods("PUT")
	router.HandleFunc("/api/books/{id}", middleware.AuthMiddleware(bookHandler.DeleteBook)).Methods("DELETE")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
