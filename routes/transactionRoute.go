package routes

import (
	"github.com/gesangwidigdo/store-management/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateTransaction)
	r.GET("/", controllers.GetAllTransaction)
	r.GET("/:id", controllers.GetTransactionByID)
	r.DELETE("/:id", controllers.DeleteTransaction)
}
