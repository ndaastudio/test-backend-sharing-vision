package util

import (
	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func ErrorResponse(c echo.Context, code int, message interface{}) error {
	return c.JSON(code, map[string]interface{}{
		"success": false,
		"error":   message,
	})
}
