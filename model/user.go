package model

import (
	"time"
)

type User struct {
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
