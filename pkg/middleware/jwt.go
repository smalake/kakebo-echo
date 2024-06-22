package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

// JWT認証
func JwtDecode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			log.Printf("invalid or missing JWT")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": errors.New("invalid or missing JWT"),
			})
		}
		tokenString := headerParts[1]

		// Googleの公開鍵を取得
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("Failed to make firebase Newapp: %+v", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to make firebase Newapp",
				"error":   err.Error(),
			})
		}

		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("Failed to Auth: %+v", err)
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Failed to Auth",
				"error":   err.Error(),
			})
		}
		token, err := auth.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			log.Printf("Failed to Verify Token: %+v", err)
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Failed to Verify Token",
				"error":   err.Error(),
			})
		}

		// UIDをContextに設定
		c.Set("uid", token.UID)
		return next(c)

	}
}
