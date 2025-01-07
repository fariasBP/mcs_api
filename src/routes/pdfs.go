package routes

import (
	"mcs_api/src/controllers"

	"github.com/labstack/echo/v4"
)

func Pdfs(e *echo.Group) {
	router := e.Group("/pdfs")
	router.GET("/service", controllers.GenerateServicePdfs)
}
