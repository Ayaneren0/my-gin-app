package main

import (
	"book-api/handlers"
	"book-api/middleware"
	"book-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.InitDB()
	models.DB.AutoMigrate(&models.Book{})

	r.Use(middleware.Logger())

	api := r.Group("/api")
	{
		books := api.GET("/books")
		{
			books.GET("/", handlers.GetBooks)
			books.GET("/:id", handlers.GetBooksByID)
			books.POST("/", handlers.CreateBooks)
			books.PUT("/:id", handlers.UpdateBooks)
			books.DELETE("/:id", handlers.DeleteBooks)

		}
	}
	r.Run()
}
