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

	// product transaction
	ptRoutes := r.Group("/product_transaction")
	ProductTransactionRoute(ptRoutes)

	// payment route
	paymentRoutes := r.Group("/payment")
	PaymentRoute(paymentRoutes)
}
