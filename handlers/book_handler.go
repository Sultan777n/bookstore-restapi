package handlers

import (
	"net/http"
	"strconv"

	"bookstore/models"
	"github.com/gin-gonic/gin"
)

var (
	Books    = make(map[int]models.Book)
	BooksSeq = 1
)

type BookResponse struct {
	Data  []models.Book `json:"data"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Total int           `json:"total"`
}

func GetBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	authorID := c.Query("author_id")
	categoryID := c.Query("category_id")

	filteredBooks := make([]models.Book, 0)
	for _, book := range Books {
		if authorID != "" {
			if aid, _ := strconv.Atoi(authorID); book.AuthorID != aid {
				continue
			}
		}
		if categoryID != "" {
			if cid, _ := strconv.Atoi(categoryID); book.CategoryID != cid {
				continue
			}
		}
		filteredBooks = append(filteredBooks, book)
	}

	start := (page - 1) * limit
	if start > len(filteredBooks) {
		start = len(filteredBooks)
	}
	end := start + limit
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	c.JSON(http.StatusOK, BookResponse{
		Data:  filteredBooks[start:end],
		Page:  page,
		Limit: limit,
		Total: len(filteredBooks),
	})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := Authors[book.AuthorID]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author not found"})
		return
	}

	if _, exists := Categories[book.CategoryID]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}

	book.ID = BooksSeq
	Books[BooksSeq] = book
	BooksSeq++

	c.JSON(http.StatusCreated, book)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if book, exists := Books[id]; exists {
		c.JSON(http.StatusOK, book)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
	}
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, exists := Books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.ID = id
	Books[id] = book
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, exists := Books[id]; exists {
		delete(Books, id)
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
	}
}
