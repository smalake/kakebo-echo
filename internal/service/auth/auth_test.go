package auth

import (
	"kakebo-echo/internal/repository/auth"
	"kakebo-echo/pkg/errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestServiceLogin(t *testing.T) {
	tests := map[string]struct {
		uid string
		err error
	}{
		"成功：正常に処理できた場合": {
			uid: "testuid",
			err: nil,
		},
		"失敗：エラーが発生した場合": {
			uid: "testuid",
			err: errors.InternalServerError,
		},
		"失敗：ユーザが未登録の場合": {
			uid: "testuid",
			err: errors.ErrUserNotFound,
		},
	}

	// Mock設定
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := auth.NewMockAuthRepository(ctrl)
	mock := New(mockRepo)

	t.Run("テスト", func(t *testing.T) {
		for tn, tt := range tests {
			t.Run(tn, func(t *testing.T) {
				setMockFindUser(mockRepo, tn, tt)
				err := mock.Login(tt.uid)
				assert.Equal(t, err, tt.err)
			})
		}
	})
}

// モックの設定
func setMockFindUser(mockRepo *auth.MockAuthRepository, tn string, tt struct {
	uid string
	err error
}) {
	switch tn {
	case "成功：正常に処理できた場合":
		mockRepo.EXPECT().FindUser(tt.uid).Return(1, tt.err)
	case "失敗：エラーが発生した場合":
		mockRepo.EXPECT().FindUser(tt.uid).Return(-1, tt.err)
	case "失敗：ユーザが未登録の場合":
		mockRepo.EXPECT().FindUser(tt.uid).Return(0, tt.err)
	default:
		mockRepo.EXPECT().FindUser(tt.uid).Return(-1, tt.err)
	}
}

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
