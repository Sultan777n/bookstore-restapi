package handlers

import (
	"net/http"

	"bookstore/models"
	"github.com/gin-gonic/gin"
)

var (
	Categories    = make(map[int]models.Category)
	CategoriesSeq = 1
)

func GetCategories(c *gin.Context) {
	cats := make([]models.Category, 0, len(Categories))
	for _, c := range Categories {
		cats = append(cats, c)
	}
	c.JSON(http.StatusOK, cats)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = CategoriesSeq
	Categories[CategoriesSeq] = category
	CategoriesSeq++

	c.JSON(http.StatusCreated, category)
}
