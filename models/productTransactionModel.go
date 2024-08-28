package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductTransaction struct {
	Transaction_id uint    `gorm:"primaryKey;autoIncrement:false"`
	Product_id     uint    `gorm:"primaryKey;autoIncrement:false"`
	Quantity       int     `gorm:"not null; type:integer"`
	Total          float64 `gorm:"not null; type:decimal; default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
