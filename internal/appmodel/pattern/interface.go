package pattern

import "github.com/labstack/echo/v4"

type AllPattern struct {
	Id        int    `json:"id"`
	Category  int    `json:"category"`
	StoreName string `json:"storeName"`
}

type Pattern struct {
	Category  int    `json:"category"`
	StoreName string `json:"storeName"`
}

type PatternInterface interface {
	// (GET /pattern)
	Get(ctx echo.Context) error

	// (POST /pattern)
	Register(ctx echo.Context) error

	// (PUT /pattern/${id})
	Update(ctx echo.Context, id int) error

	// (DELETE /pattern/${id})
	Delete(ctx echo.Context, id int) error
}
