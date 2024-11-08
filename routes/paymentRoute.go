package routes

import (
	"github.com/gesangwidigdo/store-management/controllers"
	"github.com/gin-gonic/gin"
)

func PaymentRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreatePayment)
}
