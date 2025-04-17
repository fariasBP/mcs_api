package validations

import (
	"bytes"
	"encoding/json"
	"io"
	"mcs_api/src/config"
	"mcs_api/src/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	NewProtocolParams struct {
		MachineTypeId string `json:"machine_type_id" validate:"required,mongodb"`
		Acronym       string `json:"acronym" validate:"required"`
		Name          string `json:"name" validate:"required,startsnotwith= ,endsnotwith= ,min=3,max=30"`
		Description   string `json:"description" validate:"required, startsnotwith= ,endsnotwith= "`
	}
	GetProtocolsParams struct {
		Query string `json:"query" validate:""` // "name" (nombre) del protocolo
		Limit string `json:"limit" validate:"required,number,gt=0"`
		Page  string `json:"page" validate:"required,number,gt=0"`
	}
)

func NewProtocolValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &NewProtocolParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &NewProtocolParams{MachineTypeId: body.MachineTypeId, Acronym: body.Acronym, Name: body.Name, Description: body.Description}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsMachineTypeById(v.MachineTypeId) {
			return c.JSON(400, config.SetResError(400, "Error: El id del tipo de maquina no existe", ""))
		}
		// verificando que no exista la marca
		if models.ExistsProtocol(body.MachineTypeId, body.Name, body.Acronym) {
			return c.JSON(400, config.SetResError(400, "Error: El Acronimo: '"+body.Acronym+"' o Nombre: '"+body.Name+"' del protocolo ya existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func GetProtocolsValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo query
		query := c.QueryParam("query")
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")
		// estableciendo los argumentos de validacion
		v := &GetProtocolsParams{Query: query, Limit: limit, Page: page}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// fin del middleware
		return next(c)
	}
}
