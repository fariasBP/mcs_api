package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Machine(e *echo.Group) {
	router := e.Group("/machine")
	router.POST("/create", controllers.CreateMachine)
	router.GET("/machines-bycompanyandserial", controllers.GetMachinesByCompanyIdAndSerial)
	router.GET("/machines-byserial", controllers.GetMachinesBySerial)
	router.GET("/machines-rebuild-byserial", controllers.GetMachinesRebuildBySerial)
	router.GET("/machines-rebuild-basic-byserial", controllers.GetMachinesRebuildBasicBySerial)
}
