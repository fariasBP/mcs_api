package controllers

import (
	"encoding/json"
	"strconv"

	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func NewCompany(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.NewCompanyParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()

	// creando compañia
	err := models.CreateCompany(body.Name, body.Manager, body.Latitude, body.Longitude, body.Description, body.Contact)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear la compania", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "La compania: "+body.Name+" se ha creada"))
}

func GetCompanies(c echo.Context) error {
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
	companies, count, err := models.GetCompanies(query, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo obtener las companias", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Las companias se han obtenido", count, companies))
}
