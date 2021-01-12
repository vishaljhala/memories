package models

import "time"

type User struct {
	ID        uint
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
}
