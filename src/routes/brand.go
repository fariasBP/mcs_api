package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Brand(e *echo.Group) {
	router := e.Group("/brand")
	router.POST("/create", controllers.CreateBrand)
	router.GET("/brands", controllers.GetBrands)
}
