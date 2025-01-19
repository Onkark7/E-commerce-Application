package controller

import (
	"net/http"
	"productmicro/model"
	"strconv"

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

func UpdateCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindBodyWithJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category Not update"})
		return
	}

	UpdateCategory, err := model.UpdateCategory(category, category.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Category Not Updated", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category Update Successfully !!!",
		"Data":    UpdateCategory,
	})
}

func GetallCategory(c *gin.Context) {

	result, err := model.GetallCategory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Category Not Update", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category List",
		"Data":    result,
	})

}

func GetCategoryWithID(c *gin.Context) {
	Idstr := c.Query("id")

	id, err := strconv.Atoi(Idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
		return
	}

	getID, err := model.GetCategoryWithID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data Not available", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Data Found for ID",
		"Data":    getID,
	})

}
