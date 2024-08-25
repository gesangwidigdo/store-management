package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Transaction_id uint        `gorm:"not null; type:varchar(255)"`
	Payment_method string      `gorm:"not null; type:varchar(150)"`
	Status         bool        `gorm:"not null; type:boolean"`

	Transaction    Transaction `gorm:"foreignKey:Transaction_id;references:ID"`
	
	Product        []Product   `gorm:"many2many:product_transaction;"`
}
