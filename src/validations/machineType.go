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
	CreateMachineTypeParams struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	GetMachineTypesParams struct {
		Query string `json:"query" validate:""` // "name" (nombre) del tipo de maquina
		Limit string `json:"limit" validate:"required,number,gt=0"`
		Page  string `json:"page" validate:"required,number,gt=0"`
	}
)

func NewMachineTypeValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &CreateMachineTypeParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &CreateMachineTypeParams{Name: body.Name, Description: body.Description}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos.", err.Error()))
		}
		// verificando que no exista el tipo de maquina
		if models.ExistsMachineType(body.Name) {
			return c.JSON(400, config.SetResError(400, "Error: El nombre del tipo de maquina: '"+body.Name+"' ya existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))

		return next(c)
	}
}

func GetMachineTypesValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo query
		query := c.QueryParam("query")
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")
		// estableciendo los argumentos de validacion
		v := &GetMachineTypesParams{Query: query, Limit: limit, Page: page}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// fin del middleware
		return next(c)
	}
}
