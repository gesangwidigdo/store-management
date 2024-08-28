package main

import (
	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	// Custom many2many
	initializers.DB.SetupJoinTable(&models.Transaction{}, "Product", &models.ProductTransaction{})

	initializers.DB.AutoMigrate(
		&models.Employee{},
		&models.Payment{},
		&models.Product{},
		&models.Transaction{},
	)
}
