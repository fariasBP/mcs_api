package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBrand(c echo.Context) error {
	body := &validations.CreateBrandParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// verificando que no exista la marca
	if models.ExistsBrand(body.Name) {
		return c.JSON(400, config.SetResError(400, "La marca ya existe", ""))
	}
	// creando marca
	err := models.CreateBrand(body.Name)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se creo la marca", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "La marca se ha creado"))
}

func GetBrands(c echo.Context) error {
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
	brands, count, err := models.GetBrands(name, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener las marcas", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Las marcas se han obtenido", count, brands))
}
