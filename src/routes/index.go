package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func IndexRoute(e *echo.Group) {
	// e.GET("/inf", controllers.InfoWeb)
	e.GET("/dataapp", controllers.DataApp)
}
