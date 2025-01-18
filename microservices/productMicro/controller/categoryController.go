package controller

import (
	"net/http"
	"productmicro/model"

	"github.com/gin-gonic/gin"
)

func AddCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindBodyWithJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	newCategory, err := model.AddCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Category not Added", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category Added Successfully !!!!!",
		"Data":    newCategory,
	})
}
