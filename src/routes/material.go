package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Material(e *echo.Group) {
	router := e.Group("/material", middlewares.ValidateToken)
	router.POST("/create", controllers.CreateMaterial, validations.CreateMaterialValidate)
	router.PUT("/update", controllers.UpdateMaterial, validations.UpdateMaterialValidate)
	router.GET("/materials", controllers.GetMaterials, validations.GetMaterialsValidate)
}
