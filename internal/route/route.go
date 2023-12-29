package route

import (
	"kakebo-echo/internal/appmodels"
	"kakebo-echo/pkg/mysql"
	"log"

	"github.com/labstack/echo/v4"
)

func SetRoute(e *echo.Echo) {

	mc, err := mysql.NewClient()
	if err != nil {
		log.Fatalf("[FATAL]: %+v", err)
	}
	service := appmodels.New(mc)
	e.POST("/login-mail", service.LoginMail)
}
