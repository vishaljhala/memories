package models

import "time"

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Password  string
	Email     string
}
