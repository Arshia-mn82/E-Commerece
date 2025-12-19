package model

import "time"

type Order struct {
	ID        uint         `gorm:"primaryKey"`
	UserID    uint         `gorm:"not null;index"`
	Status    string       `gorm:"not null;default:pending"`
	Price     int64        `gorm:"not null"`
	Address   string       `gorm:"not null"`
	Items     []OrderItem `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
}

type OrderItem struct {
	ID        uint  `gorm:"primaryKey"`
	OrderID   uint  `gorm:"not null;index"`
	ProductID uint  `gorm:"not null"`
	Price     int64 `gorm:"not null"`
	Quantity  uint  `gorm:"not null"`
}
