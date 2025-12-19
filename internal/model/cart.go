package model

import "time"

type Cart struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"not null;uniqueIndex"`
	Items     []CartItem `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey"`
	CartID    uint    `gorm:"not null;index;uniqueIndex:uniq_cart_product"`
	ProductID uint    `gorm:"not null;uniqueIndex:uniq_cart_product"`
	Quantity  uint    `gorm:"not null"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
