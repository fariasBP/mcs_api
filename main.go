package main

import (
	"os"

	"mcs_api/src/middlewares"
	"mcs_api/src/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// midlewares
	e.Use(middleware.CORS())
	err := middlewares.Initialization()
	if err != nil {
		e.Logger.Fatal(err)
	} else {
		// estableciendo rutas
		router := e.Group("/api")
		routes.Auth(router)
		// routes.Permissions(router)
		routes.MachineType(router)
		routes.Brand(router)
		routes.Companies(router)
		routes.Machine(router)
		routes.Protocol(router)
		routes.Service(router)
		routes.Pdfs(router)
		// iniciando server
		portApi, _ := os.LookupEnv("PORT")
		e.Logger.Fatal(e.Start(":" + portApi))
	}
}
