package event

import "github.com/labstack/echo/v4"

type CreateEvent struct {
	Amount1   int    `json:"amount1"`
	Amount2   int    `json:"amount2"`
	Category1 int    `json:"category1"`
	Category2 int    `json:"category2"`
	Memo1     string `json:"memo1"`
	Memo2     string `json:"memo2"`
	StoreName int    `json:"storeName"`
	Date      string `json:"date"`
}

type EventId struct {
	EventId int `json:"eventId"`
}

type AllEvent struct {
	Id        int    `json:"id"`
	Amount    int    `json:"amount"`
	Category  int    `json:"category"`
	StoreName string `json:"storeName"`
	Date      string `json:"date"`
}

type OneEvent struct {
	Amount     int    `json:"amount"`
	Category   int    `json:"category"`
	Memo       string `json:"memo"`
	StoreName  string `json:"storeName"`
	Date       string `json:"date"`
	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type UpdateEvent struct {
	Amount    int    `json:"amount"`
	Category  int    `json:"category"`
	Memo      string `json:"memo"`
	StoreName string `json:"storeName"`
	Date      string `json:"date"`
}

type Revision struct {
	Revision int `json:"revision"`
}

type EventInterface interface {
	// (GET /event)
	GetAll(ctx echo.Context) error

	// (POST /event)
	Create(ctx echo.Context) error

	// (PUT /event/${id})
	Update(ctx echo.Context, id int) error

	// (GET /event/${id})
	GetOne(ctx echo.Context, id int) error

	// (DELETE /event/${id})
	Delete(ctx echo.Context, id int) error

	// (GET /revision)
	Revision(ctx echo.Context) error
}
