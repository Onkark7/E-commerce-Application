package controller

import (
	"net/http"
	"productmicro/model"
	mdata "productmicro/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Addproduct(c *gin.Context) {

	var product = mdata.Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	newpro, err := mdata.Addproduct(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to add Product", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Added Successfully !!!!!",
		"Data":    newpro,
	})
}

func Updateproduct(c *gin.Context) {
	var product = mdata.Product{}

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid payload Data"})
		return
	}

	ID := product.Id

	_, err := model.Updateproduct(product, ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update Data", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Updated Successfully !!!",
	})

}

func GetALL(c *gin.Context) {

	produsts, err := model.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to load data", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": produsts,
	})
}

func GetProductsWithID(c *gin.Context) {

	ID := c.Query("id")
	id, err := strconv.Atoi(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not getting"})
		return
	}
	products, err := model.GetProductsWithID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to load Data", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func Deleterow(c *gin.Context) {
	ID := c.Query("id")

	id, err := strconv.Atoi(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not getting"})
	}

	result, err := model.DeleteProductWIthID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in Deleting Data", "message": "err"})
	}

	if result == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID Not Found",
		})
	}

	if result > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data is deleted for following id",
			"id":      result,
		})
	}

}
