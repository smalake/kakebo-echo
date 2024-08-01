package util

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/structs"

	"github.com/labstack/echo/v4"
)

func ResponseHandler(ctx echo.Context, res structs.HttpResponse) error {
	if res.Error != nil {
		return ctx.JSON(res.Code, res.Error.Error())
	}
	// 成功した場合のレスポンスを設定
	return ctx.JSON(res.Code, res.Data)
}

// イベントの内容についてバリデーション
func EventValidation(event model.Event) (bool, string) {
	if event.Amount <= 0 {
		return false, "amount"
	}
	if event.Category < 0 || event.Category > 9 {
		return false, "category"
	}
	return true, ""
}
