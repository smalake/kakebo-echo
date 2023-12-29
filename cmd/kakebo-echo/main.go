package main

import (
	"kakebo-echo/internal/route"

	// Import Echo v4.
	"github.com/labstack/echo/v4"
)

func main() {
	// 初期設定
	// err := env.Load()
	// if err != nil {
	// 	log.Fatalf("[FATAL] env load error: %+v", err)
	// }

	// Create an Echo instance.
	e := echo.New()

	// ルーティング
	route.SetRoute(e)

	// Start an Echo server.
	e.Logger.Fatal(e.Start(":8080"))
}
