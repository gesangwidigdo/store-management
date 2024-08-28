package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Employee_id      uint      `gorm:"index"`
	Grand_total      float64   `gorm:"not null; type:decimal; default: 0"`
	Transaction_time time.Time `gorm:"not null; type:datetime; default:CURRENT_TIMESTAMP()"`

	Employee Employee `gorm:"foreignKey:Employee_id;references:ID"`
}
