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

type EventGet struct {
	ID int `json:"id" db:"id"`
	Event
}
