package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewMachineType(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.CreateMachineTypeParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando tipo de maquina
	err := models.CreateMachineType(body.Name, body.Description)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear el tipo de maquina", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El Tipo de maquina: '"+body.Name+"' se ha creado"))
}

func GetMachineTypes(c echo.Context) error {
	// obteniendo params
	query := c.QueryParam("query")
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
	machineTypes, count, err := models.GetMachineTypes(query, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener los tipos de maquina", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Se obtuvieron los tipos de maquina", count, machineTypes))
}

func UpdateMachineType(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.UpdateMachineTypeParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// obteniendo tipo de maquina
	machineType, err := models.GetMachineTypeById(body.MachineTypeId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo obtener el tipo de maquina: '"+body.Name+"'", err.Error()))
	}
	// actualizando tipo de maquina
	machineType.Name = body.Name
	machineType.Description = body.Description
	// consultando
	err = models.UpdateMachineType(machineType)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo actualizar el tipo de maquina: '"+body.Name+"'", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El Tipo de maquina: '"+body.Name+"' se ha actualizado"))
}
