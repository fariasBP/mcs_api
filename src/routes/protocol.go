package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Protocol(e *echo.Group) {
	router := e.Group("/protocol")
	router.POST("/create", controllers.CreateProtocol)
	router.GET("/protocols", controllers.GetProtocols)
}
