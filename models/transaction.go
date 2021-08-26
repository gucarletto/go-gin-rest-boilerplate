package models

type Transaction struct {
	Value     float64 `json:"value"`
	User      User    `json:"user"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
