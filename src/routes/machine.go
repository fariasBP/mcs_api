package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Machine(e *echo.Group) {
	router := e.Group("/machine", middlewares.ValidateToken)
	router.POST("/new", controllers.NewMachine, validations.NewMachineValidate, middlewares.IsGTEtoOperator)
	router.GET("/search", controllers.GetMachines, validations.GetMachinesValidate, middlewares.IsGTEtoOperator)
	router.PUT("/update", controllers.UpdateMachine, validations.UpdateMachineValidate, middlewares.IsGTEtoAdmin)
}
