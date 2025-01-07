package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Service(e *echo.Group) {
	router := e.Group("/service")
	router.POST("/create", controllers.CreateService)
}
