package router

import (
	"productmicro/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.POST("/api/addCategory", controller.AddCategory)
	router.POST("/api/addProduct", controller.Addproduct)
	router.PUT("/api/UpdateProduct", controller.Updateproduct)

	return router
}
