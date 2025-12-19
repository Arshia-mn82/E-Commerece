package model

import "time"

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Price       int64  `gorm:"not null"`
	Stock       uint 
	Categories  []Category `gorm:"many2many:product_categories;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
