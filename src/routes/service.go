package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Service(e *echo.Group) {
	router := e.Group("/service", middlewares.ValidateToken)
	router.POST("/new", controllers.NewService, validations.NewServiceValidate)
	router.GET("/services", controllers.GetServices, validations.GetServicesValidate)
	router.PUT("/sleep", controllers.SleepService, validations.SleepServiceValidate)
	router.PUT("/finish", controllers.FinishService, validations.FinishServiceValidate)
	router.PUT("/progress", controllers.ProgressService, validations.ProgressServiceValidate)
}
