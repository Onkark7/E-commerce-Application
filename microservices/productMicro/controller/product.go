package controller

import (
	"net/http"
	"productmicro/model"
	mdata "productmicro/model"

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
