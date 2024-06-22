package main

import (
	// Import Echo v4.

	"kakebo-echo/internal/appmodels"
	"kakebo-echo/internal/service"
	mdl "kakebo-echo/pkg/middleware"
	"kakebo-echo/pkg/postgresql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 初期設定
	// err := env.Load()
	// if err != nil {
	// 	log.Fatalf("[FATAL] env load error: %+v", err)
	// }

	// Create an Echo instance.
	e := echo.New()

	// ログ設定
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339_nano}] method=${method}, uri=${uri}, status=${status}\n",
	}))
	// CORS
	e.Use(mdl.NewCors())

	pc, err := postgresql.NewClient()
	if err != nil {
		e.Logger.Fatalf("[FATAL]: %+v", err)
	}
	// defer mc.Close()

	appModel := appmodels.New(pc)
	service := service.New(appModel)
	e.POST("/login", service.LoginHandler)
	e.POST("/register", service.RegisterUserHandler)

	api := e.Group("/api/v1")

	// JWT認証
	api.Use(mdl.JwtDecode)
	api.GET("/login-check", service.LoginCheckHandler)
	api.POST("/logout", service.LogoutHandler)

	// Start an Echo server.
	e.Logger.Fatal(e.Start(":8080"))
}
