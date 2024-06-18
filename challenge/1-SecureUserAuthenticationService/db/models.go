package db

import (
	"time"
)

type User struct {
	UserID              int       `gorm:"primaryKey;autoIncrement;column:user_id"`
	FirstName           string    `gorm:"not null;column:first_name"`
	LastName            string    `gorm:"not null;column:last_name"`
	Email               string    `gorm:"not null;column:email;unique"`
	HashPassword        string    `gorm:"not null;column:hash_password"`
	CreatedAt           time.Time `gorm:"not null;column:created_at"`
	PasswordLastChanged time.Time `gorm:"not null;column:password_last_changed"`
}
