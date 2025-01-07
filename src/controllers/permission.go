package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func AddPermission(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.AddPermissionParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// 

	return c.JSON(200, config.SetRes(200, "permiso creado"))
}