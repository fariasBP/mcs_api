package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Companies(e *echo.Group) {
	router := e.Group("/company")
	router.POST("/create", controllers.CreateCompany)
	router.GET("/companies", controllers.GetCompanies)
}
