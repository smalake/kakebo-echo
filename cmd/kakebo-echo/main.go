package main

import (
	// Import Echo v4.

	"kakebo-echo/internal/handler/auth"
	"kakebo-echo/internal/handler/event"
	"kakebo-echo/internal/handler/private"
	"kakebo-echo/internal/handler/setting"
	"kakebo-echo/pkg/database/postgresql"
	mdl "kakebo-echo/pkg/middleware"

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

	authHeader := auth.New(pc)
	e.POST("/login", authHeader.Login)
	e.GET("/login-check", authHeader.LoginCheck)
	e.POST("/register", authHeader.Register)

	api := e.Group("/api/v1")
	// JWT認証
	api.Use(mdl.JwtDecode)

	api.GET("/login-check", authHeader.LoginCheck)
	api.POST("/logout", authHeader.Logout)

	eventHandler := event.New(pc)
	api.POST("/event", eventHandler.Create)
	api.GET("/event", eventHandler.GetAll)
	api.GET("/event/:id", eventHandler.GetOne)
	api.PUT("/event/:id", eventHandler.Update)
	api.DELETE("/event/:id", eventHandler.Delete)

	api.GET("/revision", eventHandler.GetRevision)

	privateHandler := private.New(pc)
	api.POST("/private", privateHandler.Create)
	api.GET("/private", privateHandler.GetAll)
	api.GET("/private/:id", privateHandler.GetOne)
	api.PUT("/private/:id", privateHandler.Update)
	api.DELETE("/private/:id", privateHandler.Delete)

	api.GET("/revision-private", privateHandler.GetRevision)

	settingHandler := setting.New(pc)
	api.GET("/is-admin", settingHandler.AdminCheck)

	// Start an Echo server.
	e.Logger.Fatal(e.Start(":8080"))
}
