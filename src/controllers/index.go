package controllers

import (
	"github.com/labstack/echo/v4"
)

func DataApp(c echo.Context) error {
	return c.JSON(200, "Hello, World!")
}
