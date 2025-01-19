package router

import (
	"productmicro/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	//product
	router.POST("/api/addProduct", controller.Addproduct)
	router.PUT("/api/UpdateProduct", controller.Updateproduct)
	router.GET("/api/getall", controller.GetALL)
	router.GET("/api/getwithID", controller.GetProductsWithID)
	router.DELETE("/api/deletedata", controller.Deleterow)

	//category
	router.POST("/api/addCategory", controller.AddCategory)
	router.PUT("/api/UpdateCategory", controller.UpdateCategory)
	router.GET("/api/GetallCategory", controller.GetallCategory)
	router.GET("/api/Categoryget", controller.GetCategoryWithID)

	return router
}
