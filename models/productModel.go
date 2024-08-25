package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_name string        `gorm:"not null; type:varchar(255)"`
	Price        string        `gorm:"not null; type:decimal"`
	Stock        string        `gorm:"not null; type:integer"`
	
	Transaction  []Transaction `gorm:"many2many:product_transaction;"`
}
