package main

import (
	"bookstore/handlers"
	"bookstore/models"
	"github.com/gin-gonic/gin"
)

func initData() {
	handlers.Authors[1] = models.Author{ID: 1, Name: "Абай Құнанбайұлы"}
	handlers.Authors[2] = models.Author{ID: 2, Name: "Ахмет Байтұрсынұлы"}

	handlers.Categories[1] = models.Category{ID: 1, Name: "Қара сөз"}
	handlers.Categories[2] = models.Category{ID: 2, Name: "Қазақ әдебиеті"}

	handlers.Books[1] = models.Book{
		ID:         1,
		Title:      "7 қара сөз",
		AuthorID:   1,
		CategoryID: 1,
		Price:      1200.50,
	}
	handlers.Books[2] = models.Book{
		ID:         2,
		Title:      "Маса",
		AuthorID:   2,
		CategoryID: 1,
		Price:      950.00,
	}
}

func main() {
	r := gin.Default()
	initData()

	// Книги
	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	// Авторы
	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)

	// Категории
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	r.Run(":8080")
}
