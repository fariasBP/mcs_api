package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func MachineType(e *echo.Group) {
	router := e.Group("/machinetype", middlewares.ValidateToken)
	router.POST("/new", controllers.NewMachineType, validations.NewMachineTypeValidate)
	router.GET("/search", controllers.GetMachineTypes, validations.GetMachineTypesValidate)
}
