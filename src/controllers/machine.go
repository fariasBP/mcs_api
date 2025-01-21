package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateMachine(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.CreateMachineParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// verificamos que existe el idOwner (id de la empresa)
	if !models.ExistsCompanyById(body.CompanyId) {
		return c.JSON(400, config.SetResError(400, "el id de la empresa no existe", ""))
	}
	// verificando que exista el tipo de maquina
	if !models.ExistsMachineTypeById(body.MachineTypeId) {
		return c.JSON(400, config.SetResError(400, "el tipo de maquina no existe", ""))
	}
	// creando maquina
	err := models.CreateMachine(body.CompanyId, body.MachineTypeId, body.BrandId, body.Serial, body.Model)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Internal Server Error", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "maquina creada"))
}

func GetMachinesByCompanyId(c echo.Context) error {
	// obteniendo params
	companyId := c.QueryParam("company_id")
	serial := c.QueryParam("serial")
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
	Machines, count, err := models.GetMachinesByCompanyId(companyId, serial, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener las maquinas", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Las maquinas se han obtenido", count, Machines))
}
