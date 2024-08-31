package main

import (
	"net/http"

	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"project_name": "store-management",
			"status": "ok",
		})
	})

	routes.Routes(r)

	r.Run()
}
