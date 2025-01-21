package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Group) {
	router := e.Group("/auth")
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.SignUp)
}
