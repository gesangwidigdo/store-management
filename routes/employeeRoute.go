package routes

import (
	"github.com/gesangwidigdo/store-management/controllers"
	"github.com/gin-gonic/gin"
)

func EmployeeRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateEmployee)
	r.GET("/", controllers.GetAllEmployee)
	r.GET("/:id", controllers.GetEmployeeByID)
	r.PUT("/:id", controllers.UpdateEmployee)
	r.DELETE("/:id", controllers.DeleteEmployee)
}