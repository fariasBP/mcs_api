package controllers

import (
	"encoding/json"
	"strconv"

	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func CreateCompany(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.CompanyParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// verificando que no exista la empresa
	if models.ExistsCompany(body.Name) {
		return c.JSON(400, config.SetResError(400, "la empresa ya existe", ""))
	}
	// creando empresa
	err := models.CreateCompany(body.Name, body.Manager, body.Location, body.Description, body.Contact)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se creo la empresa", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "empresa creada"))
}

func GetCompanies(c echo.Context) error {
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
	companies, count, err := models.GetCompanies(name, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener las empresas", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Las empresas se han obtenido", count, companies))
}
