package controllers

import (
	"encoding/json"

	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func NewService(c echo.Context) error {
	// obteniendo params
	body := &validations.NewServiceParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// obtener id del usuario
	user := c.Get("id").(string)
	// consultando
	err := models.NewService(body.MachineId, user)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se creo el servicio", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El servicio se ha creado"))
}

func GetServices(c echo.Context) error {
	// obteniendo params
	startedAt := c.QueryParam("started_at")
	endedAt := c.QueryParam("ended_at")
	ascending := c.QueryParam("ascending")
	status := c.QueryParam("status")
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	// convirtiendo params
	startedAtVal, errS := time.Parse(time.RFC3339, startedAt)
	endedAtVal, errE := time.Parse(time.RFC3339, endedAt)
	if errS != nil || errE != nil {
		startedAtVal = time.Time{}
		endedAtVal = time.Time{}
	}
	ascendingVal, err := strconv.ParseBool(ascending)
	if err != nil {
		ascendingVal = true
	}
	statusVal, err := strconv.Atoi(status)
	if err != nil || (statusVal != 0 && statusVal != 1 && statusVal != 2 && statusVal != 3) {
		statusVal = 0
	}
	statusV := models.Status2Service(statusVal)
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	// consultando
	services, count, err := models.GetServices(startedAtVal, endedAtVal, ascendingVal, statusV, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener los servicios", err.Error()))
	}
	// remodelando datos
	servicesRebuild := make([]models.ServiceRebuild, len(services))
	for i, v := range services {
		machine, err := models.GetMachineById(v.MachineId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "No se pudo obtener la maquina del servicio #id:"+v.ID.Hex(), err.Error()))
		}
		company, err := models.GetCompanyById(machine.CompanyId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "No se pudo obtener la empresa de la maquina del servicio #id:"+v.ID.Hex(), err.Error()))
		}
		companyName := company.Name
		brand, err := models.GetBrandById(machine.BrandId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "No se pudo obtener la marca de la maquina del servicio #id:"+v.ID.Hex(), err.Error()))
		}
		brandName := brand.Name
		machineType, err := models.GetMachineTypeById(machine.MachineTypeId)
		if err != nil {
			return c.JSON(500, config.SetResError(500, "No se pudo obtener el tipo de la maquina del servicio #id:"+v.ID.Hex(), err.Error()))
		}
		machineTypeName := machineType.Name

		servicesRebuild[i] = models.ServiceRebuild{
			ID:        v.ID,
			StartedAt: v.StartedAt,
			EndedAt:   v.EndedAt,
			Keepers:   v.Keepers,
			Status:    v.Status,
			Machine: models.MachineRebuild{
				ID:              machine.ID,
				CreatedAt:       machine.CreatedAt,
				UpdatedAt:       machine.UpdatedAt,
				CompanyId:       machine.CompanyId,
				CompanyName:     companyName,
				CompanyManager:  company.Manager,
				CompanyContact:  company.Contact,
				MachineTypeId:   machine.MachineTypeId,
				MachineTypeName: machineTypeName,
				BrandId:         machine.BrandId,
				BrandName:       brandName,
				Serial:          machine.Serial,
				Model:           machine.Model,
			},
		}
	}

	return c.JSON(200, config.SetResJsonCount(200, "Los servicios se han obtenido", count, servicesRebuild))
}

func SleepService(c echo.Context) error {
	// obteniendo params
	body := &validations.SleepServiceParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// actualizamos el servicio
	err := models.SleepService(body.ServiceId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo inactivar el servicio", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El servicio se ha actualizado a inactivo"))
}

func FinishService(c echo.Context) error {
	// obteniendo params
	body := &validations.FinishServiceParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// finalizando el servicio
	err := models.FinishService(body.ServiceId, body.Cancelled)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo finalizar el servicio", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El servicio se ha actualizado a finalizado"))
}

func ProgressService(c echo.Context) error {
	// obteniendo params
	body := &validations.ProgressServiceParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// actualizamos el servicio
	err := models.UpdateProgressService(body.ServiceId, body.Progress)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo actualizar el progreso del servicio", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "Se ha actualizado el progreso del servicio"))
}
