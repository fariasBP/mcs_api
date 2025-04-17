package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewBrand(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.NewBrandParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando marca
	err := models.CreateBrand(body.Name)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear la marca (fabricante)", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "La marca (fabricante): "+body.Name+" se ha creado"))
}

func GetBrands(c echo.Context) error {
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
	brands, count, err := models.GetBrands(query, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo obtener las marcas (fabricantes)", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Las marcas (fabricantes) se han obtenido", count, brands))
}
