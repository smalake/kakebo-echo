package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// テスト用にJWTを発行
func issueToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix()                     // 発行時間
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 有効期限
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}
