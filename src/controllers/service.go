package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"
	"strconv"

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
	machineId := c.QueryParam("machine_id")
	endedAt := c.QueryParam("ended_at")
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
	// verificando que la maquina existe
	if !models.ExistsMachineById(machineId) {
		return c.JSON(400, config.SetResError(400, "La maquina no existe", ""))
	}
	// consultando
	services, count, err := models.GetServices(machineId, endedAt, limitInt, pageInt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener los servicios", err.Error()))
	}

	return c.JSON(200, config.SetResJsonCount(200, "Los servicios se han obtenido", count, services))
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
