package entity

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
