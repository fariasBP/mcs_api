package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"

	"github.com/labstack/echo/v4"
)

func User(e *echo.Group) {
	router := e.Group("/user")
	router.GET("/get", controllers.GetUser, middlewares.ValidateToken)
}
