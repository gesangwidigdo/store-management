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
	initializers.DB.AutoMigrate(
		&models.Employee{},
		&models.Payment{},
		&models.Product{},
		&models.Transaction{},
	)
}
