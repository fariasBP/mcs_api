package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Permissions(e *echo.Group) {
	router := e.Group("/permissions")
	router.POST("/add", controllers.AddPermission)
}
