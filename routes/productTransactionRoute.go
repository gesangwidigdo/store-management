package routes

import (
	"github.com/gesangwidigdo/store-management/controllers"
	"github.com/gin-gonic/gin"
)

func ProductTransactionRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateProductTransaction)
	// r.PUT("/:id", controllers.UpdateProductTransaction)
	// r.DELETE("/:id", controllers.DeleteProductTransaction)
}
