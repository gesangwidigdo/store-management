package routes

import (
	"github.com/gesangwidigdo/store-management/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateProduct)
	r.GET("/", controllers.GetAllProduct)
	r.GET("/:id", controllers.GetProductByID)
	r.PUT("/:id", controllers.UpdateProduct)
	r.DELETE("/:id", controllers.DeleteProduct)
}
