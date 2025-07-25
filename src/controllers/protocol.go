package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewProtocol(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.NewProtocolParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando marca
	err := models.CreateProtocol(body.MachineTypeId, body.Acronym, body.Name, body.Description)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear el protocolo", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El protocolo: '"+body.Name+"' se ha creado"))
}

func GetProtocols(c echo.Context) error {
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
	protocols, count, err := models.GetProtocols(query, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo obtener los protocolos", err.Error()))
	}
	// recostruyendo
	protocolsRebuild := make([]models.ProtocolRebuild, len(protocols))
	for i, v := range protocols {
		machineType, err := models.GetMachineTypeById(v.MachineTypeId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "Error: No se pudo obtener el id del tipo de maquina, del protocolo: '"+v.Name+"'", err.Error()))
		}
		machineTypeName := machineType.Name

		protocolsRebuild[i] = models.ProtocolRebuild{
			ID:              v.ID,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
			MachineTypeId:   v.MachineTypeId,
			MachineTypeName: machineTypeName,
			Acronym:         v.Acronym,
			Name:            v.Name,
			Description:     v.Description,
		}
	}

	return c.JSON(200, config.SetResJsonCount(200, "Los protocolos se han obtenido", count, protocolsRebuild))
}

func UpdateProtocol(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.UpdateProtocolParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// obteniendo protocolo
	protocol, err := models.GetProtocolById(body.ProtocolId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo obtener el protocolo: '"+body.Name+"'", err.Error()))
	}
	// actualizando parametros
	protocol.Acronym = body.Acronym
	protocol.Name = body.Name
	protocol.Description = body.Description
	// actualizando protocolo
	err = models.UpdateProtocol(protocol)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo actualizar el protocolo", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El protocolo: '"+body.Name+"' se ha actualizado"))
}
