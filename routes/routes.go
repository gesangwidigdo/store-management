package routes

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	// employee route
	employeeRoutes := r.Group("/employee")
	EmployeeRoute(employeeRoutes)

	// product route
	productRoutes := r.Group("product")
	ProductRoute(productRoutes)

	// transaction route
	transactionRoutes := r.Group("/transaction")
	TransactionRoute(transactionRoutes)
}
