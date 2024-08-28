package routes

import (
	"github.com/gesangwidigdo/store-management/controllers"
	"github.com/gin-gonic/gin"
)

func ProductTransactionRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateProductTransaction)
	r.GET("/", controllers.GetAllProductTransaction)
	// r.GET("/:id", controllers.GetProductTransactionByID)
	// r.PUT("/:id", controllers.UpdateProductTransaction)
	// r.DELETE("/:id", controllers.DeleteProductTransaction)
}
