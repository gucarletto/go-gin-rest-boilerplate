package models

import "time"

type Transaction struct {
	Id        uint      `json:"id" gorm:"primaryKey;not null"`
	Value     float64   `json:"value" gorm:"not null"`
	User      User      `json:"user" gorm:"embedded" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
