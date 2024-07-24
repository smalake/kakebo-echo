package model

import "time"

type Event struct {
	Amount    int       `json:"amount" db:"amount"`
	Category  int       `json:"category" db:"category"`
	Memo      string    `json:"memo" db:"memo"`
	StoreName string    `json:"store_name" db:"store_name"`
	Date      time.Time `json:"date" db:"date"`
}

type EventCreate struct {
	Amount1   int    `json:"amount1"`
	Amount2   int    `json:"amount2"`
	Category1 int    `json:"category1"`
	Category2 int    `json:"category2"`
	Memo1     string `json:"memo1"`
	Memo2     string `json:"memo2"`
	StoreName string `json:"store_name"`
	Date      string `json:"date"`
}

// TODO: Event構造体を使った形にするかも？
type EventGet struct {
	ID        int       `json:"id" db:"id"`
	Amount    int       `json:"amount" db:"amount"`
	Category  int       `json:"category" db:"category"`
	StoreName string    `json:"store_name" db:"store_name"`
	Date      time.Time `json:"date" db:"date"`
}

type EventOne struct {
	ID            int       `json:"id" db:"id"`
	Amount        int       `json:"amount" db:"amount"`
	Category      int       `json:"category" db:"category"`
	Memo          string    `json:"memo" db:"memo"`
	StoreName     string    `json:"store_name" db:"store_name"`
	Date          time.Time `json:"date" db:"date"`
	CreateUser    string    `json:"create_user" db:"create_user"`
	UpdateUser    string    `json:"update_user" db:"update_user"`
	CreatedAtDate time.Time `db:"created_at"`
	UpdatedAtDate time.Time ` db:"updated_at"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

type GetIDs struct {
	ID      int `json:"id" db:"id"`
	GroupID int `json:"group_id" db:"group_id"`
}

type EventResponse struct {
	ID        int    `json:"id"`
	Amount    int    `json:"amount"`
	Category  int    `json:"category"`
	StoreName string `json:"store_name"`
	Date      string `json:"date"`
}
