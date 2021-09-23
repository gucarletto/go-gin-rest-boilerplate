package models

import "time"

type Transaction struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Value     float64   `json:"value"`
	User      User      `json:"user" gorm:"embedded"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
