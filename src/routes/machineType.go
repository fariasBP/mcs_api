package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func MachineType(e *echo.Group) {
	router := e.Group("/machinetype")
	router.POST("/create", controllers.CreateMachineType)
	router.GET("/machinetypes", controllers.GetMachineTypes)
}
