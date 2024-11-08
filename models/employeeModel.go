package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Employee_name    string   `gorm:"not null; type:varchar(150)"`
	Gender           string   `gorm:"not null; type:varchar(10)"`
	Telephone_number string   `gorm:"not null; type:varchar(15)"`
	Username         string   `gorm:"not null; type:varchar(100)"`
	Password         string   `gorm:"not null; type: varchar(100)"`
}
