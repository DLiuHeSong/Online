package handler

import "github.com/labstack/echo/v4"

func Plan(ctx echo.Context) error {
	return ctx.String(200, "我爱你")
}
