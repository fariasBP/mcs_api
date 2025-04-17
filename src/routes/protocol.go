package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Protocol(e *echo.Group) {
	router := e.Group("/protocol", middlewares.ValidateToken)
	router.POST("/new", controllers.NewProtocol, validations.NewProtocolValidate)
	router.GET("/search", controllers.GetProtocols, validations.GetProtocolsValidate)
}
