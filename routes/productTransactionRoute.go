package routes

import (
	"github.com/gesangwidigdo/store-management/controllers"
	"github.com/gin-gonic/gin"
)

func ProductTransactionRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateProductTransaction)
	r.GET("/:id", controllers.GetProductTransactionByTransactionID)
	// r.PUT("/:id", controllers.UpdateProductTransaction)
	r.DELETE("/:transaction_id/:product_id", controllers.DeleteProductTransaction)
}
