package controllers

import (
	"mcs_api/src/config"
	"mcs_api/src/models"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	// obteniendo variables de request
	id := c.Get("id").(string)

	// buscando usuario
	user, err := models.GetUser(id)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo encontrar al usuario", err.Error()))
	}

	return c.JSON(200, config.SetResJson(200, "Token es valido", user))
}
