package health

import "github.com/labstack/echo"

type HealthInterface interface {
	// (GET /health-check)
	HealthCheck(ctx echo.Context) error
}
