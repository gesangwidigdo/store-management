package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Employee_id      uint     `gorm:"index"`
	Total_price      float64  `gorm:"not null; type:decimal"`
	Transaction_time string   `gorm:"not null; type:datetime"`

	Employee         Employee `gorm:"foreignKey:Employee_id;references:ID"`
}
