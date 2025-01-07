package controllers

import (
	"mcs_api/src/config"
	"mcs_api/src/middlewares"
	"mcs_api/src/models"

	"github.com/labstack/echo/v4"
)

func GenerateServicePdfs(c echo.Context) error {
	// obteniendo params
	serviceId := c.QueryParam("service_id")
	// verificamos que el id del servicio exista
	if !models.ExistsServiceById(serviceId) {
		return c.JSON(400, config.SetResError(400, "el id del servicio no existe", ""))
	}
	// obtenemos el servicio
	service, err := models.GetServiceById(serviceId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se pudo obtener el servicio", err.Error()))
	}
	// obtenemos la maquina
	machine, err := models.GetMachineById(service.MachineId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se pudo obtener la maquina", err.Error()))
	}
	// obtenemos la compania
	company, err := models.GetCompanyById(machine.CompanyId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se pudo obtener la compania", err.Error()))
	}
	// creamos pdf
	err = middlewares.CreateServicePdf(company, machine, service, "service_"+serviceId+".pdf")
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se pudo crear el pdf", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "pdf creado"))
}
