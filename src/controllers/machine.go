package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewMachine(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.CreateMachineParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando maquina
	err := models.CreateMachine(body.CompanyId, body.MachineTypeId, body.BrandId, body.Serial, body.Model)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear la maquina", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "La maquina serial#"+body.Serial+" se ha creada"))
}

func GetMachines(c echo.Context) error {
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
	var machines []models.Machine = nil
	count := int64(0)
	if models.ExistsCompanyById(query) {
		machines, count, err = models.GetMachinesByCompanyOrBrandOrMachineTypeId(query, models.CompanyParam, limitInt, pageInt)
	} else if models.ExistsMachineTypeById(query) {
		machines, count, err = models.GetMachinesByCompanyOrBrandOrMachineTypeId(query, models.MachineTypeParam, limitInt, pageInt)
	} else if models.ExistsBrandById(query) {
		machines, count, err = models.GetMachinesByCompanyOrBrandOrMachineTypeId(query, models.BrandParam, limitInt, pageInt)
	} else {
		machines, count, err = models.GetMachinesByModelAndSerial(query, limitInt, pageInt)
	}
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener las maquinas", err.Error()))
	}
	// recostruyendo
	machinesRebuildBasic := make([]models.MachineRebuild, len(machines))
	for i, v := range machines {
		company, err := models.GetCompanyById(v.CompanyId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "No se pudo obtener la empresa de la maquina serial#"+v.Serial, err.Error()))
		}
		companyName := company.Name
		brand, err := models.GetBrandById(v.BrandId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "No se pudo obtener la marca de la maquina serial#"+v.Serial, err.Error()))
		}
		brandName := brand.Name
		machineType, err := models.GetMachineTypeById(v.MachineTypeId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "No se pudo obtener el tipo de la maquina serial#"+v.Serial, err.Error()))
		}
		machineTypeName := machineType.Name

		machinesRebuildBasic[i] = models.MachineRebuild{
			ID:              v.ID,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
			CompanyId:       v.CompanyId,
			CompanyName:     companyName,
			MachineTypeId:   v.MachineTypeId,
			MachineTypeName: machineTypeName,
			BrandId:         v.BrandId,
			BrandName:       brandName,
			Serial:          v.Serial,
			Model:           v.Model,
		}
	}

	return c.JSON(200, config.SetResJsonCount(200, "Las maquinas se han obtenido", count, machinesRebuildBasic))
}

func UpdateMachine(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.UpdateMachineParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// obteniendo maquina
	machine, err := models.GetMachineById(body.MachineId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo obtener la maquina", err.Error()))
	}
	// actualizando maquina
	machine.CompanyId = body.CompanyId
	machine.MachineTypeId = body.MachineTypeId
	machine.BrandId = body.BrandId
	machine.Serial = body.Serial
	machine.Model = body.Model
	// cnsultando
	err = models.UpdateMachine(machine)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo actualizar la maquina", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "La maquina serial#"+body.Serial+" se ha actualizado"))
}
