package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProtocol(c echo.Context) error {
	body := &validations.CreateProtocolParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// verificando que no exista la marca
	if models.ExistsProtocol(body.Acronym, body.Name) {
		return c.JSON(400, config.SetResError(400, "El protocolo ya existe", ""))
	}
	// creando marca
	err := models.CreateProtocol(body.Acronym, body.Name, body.Description)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se creo el protocolo", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El protocolo se ha creado"))
}

func GetProtocols(c echo.Context) error {
	// obteniendo params
	name := c.QueryParam("name")
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	// convirtiendo params
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	// consultando
	protocols, count, err := models.GetProtocols(name, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener los protocolos", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Los protocolos se han obtenido", count, protocols))
}
