package health

import "github.com/labstack/echo/v4"

type HealthInterface interface {
	// (GET /health-check)
	HealthCheck(ctx echo.Context) error
}
