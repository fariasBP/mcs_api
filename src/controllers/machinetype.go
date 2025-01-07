package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateMachineType(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.CreateMachineTypeParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// verificando que no exista el tipo de maquina
	if models.ExistsMachineType(body.Name) {
		return c.JSON(400, config.SetResError(400, "el tipo de maquina ya existe", ""))
	}
	// creando tipo de maquina
	err := models.CreateMachineType(body.Name, body.Description)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se creo el tipo de maquina", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "tipo de maquina creado"))
}

func GetMachineTypes(c echo.Context) error {
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
	machineTypes, count, err := models.GetMachineTypes(name, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener los tipos de maquina", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Se obtuvieron los tipos de maquina", count, machineTypes))
}
