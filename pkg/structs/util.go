package structs

import (
	"github.com/labstack/echo/v4"
)

func ResponseHandler(ctx echo.Context, res HttpResponse) error {
	if res.Error != nil {
		return ctx.JSON(res.Code, res.Error.Error())
	}
	// 成功した場合のレスポンスを設定
	return ctx.JSON(res.Code, res.Data)
}
