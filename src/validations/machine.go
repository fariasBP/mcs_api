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
	CreateMachineParams struct {
		CompanyId     string `json:"company_id" validate:"required,mongodb"`
		MachineTypeId string `json:"machine_type_id" validate:"required,mongodb"`
		BrandId       string `json:"brand_id" validate:"required,mongodb"`
		Serial        string `json:"serial" validate:"required,min=3,max=50,startsnotwith= ,endsnotwith= "`
		Model         string `json:"model" validate:"required,min=3,max=50,startsnotwith= ,endsnotwith= "`
	}
	GetMachinesParams struct {
		Query string `json:"query" validate:""` // "serial", "model", "company_id", "machine_type_id", "brand_id",  de la maquina
		Limit string `json:"limit" validate:"required,number,gt=0"`
		Page  string `json:"page" validate:"required,number,gt=0"`
	}
	UpdateMachineParams struct {
		MachineId     string `json:"machine_id" validate:"required,mongodb"`
		CompanyId     string `json:"company_id" validate:"required,mongodb"`
		MachineTypeId string `json:"machine_type_id" validate:"required,mongodb"`
		BrandId       string `json:"brand_id" validate:"required,mongodb"`
		Serial        string `json:"serial" validate:"required,min=3,max=50,startsnotwith= ,endsnotwith= "`
		Model         string `json:"model" validate:"required,min=3,max=50,startsnotwith= ,endsnotwith= "`
	}
)

func NewMachineValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &CreateMachineParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &CreateMachineParams{CompanyId: body.CompanyId, MachineTypeId: body.MachineTypeId, BrandId: body.BrandId, Serial: body.Serial, Model: body.Model}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsCompanyById(v.CompanyId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la empresa no existe", ""))
		}
		if !models.ExistsMachineTypeById(v.MachineTypeId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del tipo de maquina no existe", ""))
		}
		if !models.ExistsBrandById(v.BrandId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la marca (fabricante) no existe", ""))
		}
		// verificando que no exista la maquina con el mismo serial
		if models.ExistsMachine(v.Serial) {
			return c.JSON(400, config.SetResError(400, "Error: La maquina serial#"+v.Serial+" ya existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func GetMachinesValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo query
		query := c.QueryParam("query")
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")
		// estableciendo los argumentos de validacion
		v := &GetMachinesParams{Query: query, Limit: limit, Page: page}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// fin del middleware
		return next(c)
	}
}

func UpdateMachineValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &UpdateMachineParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &UpdateMachineParams{MachineId: body.MachineId, CompanyId: body.CompanyId, MachineTypeId: body.MachineTypeId, BrandId: body.BrandId, Serial: body.Serial, Model: body.Model}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsMachineById(v.MachineId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la maquina no existe", ""))
		}
		if !models.ExistsCompanyById(v.CompanyId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la empresa no existe", ""))
		}
		if !models.ExistsMachineTypeById(v.MachineTypeId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del tipo de maquina no existe", ""))
		}
		if !models.ExistsBrandById(v.BrandId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la marca (fabricante) no existe", ""))
		}
		// verificando que no exista otra maquina con el mismo serial
		if models.ExistsOtherMachine(v.MachineId, v.Serial) {
			return c.JSON(400, config.SetResError(400, "Error: La maquina serial#"+v.Serial+" ya existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}
