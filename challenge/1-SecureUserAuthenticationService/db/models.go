package db

import (
	"time"
)

type User struct {
	UserID              int       `gorm:"primaryKey;column:user_id"`
	FirstName           string    `gorm:"first_name"`
	LastName            string    `gorm:"last_name"`
	Email               string    `gorm:"email"`
	HashPassword        string    `gorm:"hash_password"`
	CreatedAt           time.Time `gorm:"created_at"`
	PasswordLastChanged time.Time `gorm:"password_last_changed"`
}
