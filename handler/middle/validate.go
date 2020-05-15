package middle

import (
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	// not validate routers
	needSkipperPathKeyword = []string{"/health", "callback", "login"}
	defaultSkipper         = func(c echo.Context) bool {
		for _, word := range needSkipperPathKeyword {
			if strings.Index(c.Request().URL.Path, word) != -1 {
				return true
			}
		}
		return false
	}
)

// ValidateToken token验证
func Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if defaultSkipper(c) {
			return next(c)
		}
		return c.JSON(302, "需要登陆")
	}
}
