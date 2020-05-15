package api

import (
	"Online/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", handler.Health)
	admin := e.Group("/admin")
	admin.GET("/plan", handler.Plan)
}
