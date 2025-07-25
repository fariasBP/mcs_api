package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Material(e *echo.Group) {
	router := e.Group("/material", middlewares.ValidateToken)
	router.POST("/new", controllers.NewMaterial, validations.CreateMaterialValidate, middlewares.IsGTEtoOperator)
	router.GET("/get", controllers.GetMaterials, validations.GetMaterialsValidate, middlewares.IsGTEtoOperator)
	router.PUT("/update", controllers.UpdateMaterial, validations.UpdateMaterialValidate, middlewares.IsGTEtoOperator)
}
