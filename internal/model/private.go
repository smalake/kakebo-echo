package model

import "time"

type PrivateOne struct {
	ID            int       `json:"id" db:"id"`
	Amount        int       `json:"amount" db:"amount"`
	Category      int       `json:"category" db:"category"`
	Memo          string    `json:"memo" db:"memo"`
	StoreName     string    `json:"store_name" db:"store_name"`
	Date          time.Time `json:"date" db:"date"`
	CreatedAtDate time.Time `db:"created_at"`
	UpdatedAtDate time.Time ` db:"updated_at"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}
