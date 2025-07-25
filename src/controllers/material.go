package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func NewMaterial(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.NewMaterialParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando material
	err := models.CreateMaterial(body.ServiceId, body.Name, body.Price, body.Number)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear el material", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El material: '"+body.Name+"' se ha creado"))
}

func UpdateMaterial(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.UpdateMaterialParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// obteniendo material
	material, err := models.GetMaterialById(body.MaterialId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener el material", err.Error()))
	}
	// actualizando material
	material.Name = body.Name
	material.Price = body.Price
	material.Number = body.Number
	// creando material
	err = models.UpdateMaterial(material)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo actualizar el material", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El material se ha actualizado"))
}

func GetMaterials(c echo.Context) error {
	// obteniendo variables de request
	serviceId := c.QueryParam("service_id")
	// obteniendo materiales
	materials, err := models.GetMaterialsByServiceId(serviceId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener los materiales", err.Error()))
	}

	return c.JSON(200, config.SetResJson(200, "Materiales obtenidos", materials))
}
