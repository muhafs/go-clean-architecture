package entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
