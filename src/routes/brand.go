package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Brand(e *echo.Group) {
	router := e.Group("/brand", middlewares.ValidateToken)
	router.POST("/new", controllers.NewBrand, validations.NewBrandValidate, middlewares.IsGTEtoAdmin)
	router.GET("/search", controllers.GetBrands, validations.GetBrandsValidate, middlewares.IsGTEtoOperator)
	router.PUT("/update", controllers.UpdateBrand, validations.UpdateBrandValidate, middlewares.IsGTEtoAdmin)
}
