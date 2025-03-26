package handlers

import (
	"net/http"

	"bookstore/models"
	"github.com/gin-gonic/gin"
)

var (
	Authors    = make(map[int]models.Author)
	AuthorsSeq = 1
)

func GetAuthors(c *gin.Context) {
	authors := make([]models.Author, 0, len(Authors))
	for _, a := range Authors {
		authors = append(authors, a)
	}
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author.ID = AuthorsSeq
	Authors[AuthorsSeq] = author
	AuthorsSeq++

	c.JSON(http.StatusCreated, author)
}
