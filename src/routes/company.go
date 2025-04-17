package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Companies(e *echo.Group) {
	router := e.Group("/company", middlewares.ValidateToken)
	router.POST("/new", controllers.NewCompany, validations.NewCompanyValidate)
	router.GET("/search", controllers.GetCompanies, validations.GetCompaniesValidate)
}
