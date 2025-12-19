package model

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null;uniqueIndex"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"not null;default:user"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
