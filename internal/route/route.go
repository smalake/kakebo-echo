package route

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"

	"kakebo-echo/internal/appmodels"
	"kakebo-echo/internal/service"
	"kakebo-echo/pkg/postgresql"
)

func SetRoute(e *echo.Echo) {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339_nano}] method=${method}, uri=${uri}, status=${status}\n",
	}))
	// CORS
	e.Use(newCors())

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
	api.Use(jwtDecode)
	api.GET("/login-check", service.LoginCheckHandler)
	api.POST("/logout", service.LogoutHandler)

	// イベント
	// api.GET("/event", sevice.GetAllEvent)
	// api.PUT("/event/:id", sevice.UpdateEvent)
	// api.GET("/event/:id", sevice.GetOneEvent)
	// api.POST("/event", sevice.CreateEvent)
	// api.DELETE("/event", sevice.DeleteEvent)
}

// CORSの設定
func newCors() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			c.Response().Writer.Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
			c.Response().Header().Set("Access-Control-Max-Age", "12h0m0s")
			c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Authorization")
			c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Length")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	}
}

type CustomClaims struct {
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Iss      string `json:"iss"`
	Aud      string `json:"aud"`
	AuthTime int64  `json:"auth_time"`
	UserId   string `json:"user_id"`
	Sub      string `json:"sub"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// JWT認証
func jwtDecode(next echo.HandlerFunc) echo.HandlerFunc {
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
