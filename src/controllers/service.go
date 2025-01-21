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

func CreateService(c echo.Context) error {
	body := &validations.CreateServiceParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// verificando que exista la maquina
	if !models.ExistsMachineById(body.MachineId) {
		return c.JSON(400, config.SetResError(400, "La maquina no existe", ""))
	}
	// creando material
	materials := make([]models.Material, len(body.Materials))
	for i, v := range body.Materials {
		materials[i] = models.Material{
			Name:   v.Name,
			Number: v.Number,
			Price:  v.Price,
		}
	}
	// creando protocolo
	protocols := make([]models.ProtocolData, len(body.Protocols))
	for i, v := range body.Protocols {
		problems := make([]models.Problem, len(v.Problems))
		// verificando que exista el protocolo
		if !models.ExistsProtocolById(v.ProtocolId) {
			return c.JSON(400, config.SetResError(400, "el protocolo no existe", ""))
		}
		for j, p := range v.Problems { // se recorre para problems
			problems[j] = models.Problem{
				Problem:  p.Problem,
				Solution: p.Solution,
			}
			protocols[i] = models.ProtocolData{
				Protocol: v.ProtocolId,
				Status:   config.StatusProtocol(v.Status),
				Note:     v.Note,
				Problems: problems,
			}
		}
	}
	// trabajando en fechas
	started, err := time.Parse(time.DateTime, body.StartedAt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se puede parsear la fecha 'started_at'", err.Error()))
	}
	ended, err := time.Parse(time.DateTime, body.EndedAt)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se puede parsear la fecha 'ended_at'", err.Error()))
	}
	// creando Servicio
	err = models.CreateService(body.MachineId, body.Comments, started, ended, protocols, materials, config.StatusService(body.Status))
	if err != nil {
		return c.JSON(500, config.SetResError(500, "no se creo el servicio", err.Error()))
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
