package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"

	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Group) {
	router := e.Group("/auth")
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.SignUp)
	router.GET("/check-token", controllers.CheckToken, middlewares.ValidateToken)
	//router.POST("create-user", controllers.CreateUser)

}
