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
	CreateMaterialParams struct {
		ServiceId string `json:"service_id" validate:"required,mongodb"`
		Name      string `json:"name" validate:"required,startsnotwith= ,endsnotwith= ,min=3,max=30"`
		Number    int    `json:"number" validate:"required,number,gt=0"`
		Price     int    `json:"price" validate:"required,number,gte=0"`
	}
	UpdateMaterialParams struct {
		MaterialId string `json:"material_id" validate:"required,mongodb"`
		Name       string `json:"name" validate:"required,startsnotwith= ,endsnotwith= ,min=3,max=30"`
		Number     int    `json:"number" validate:"required,number,gt=0"`
		Price      int    `json:"price" validate:"required,number,gte=0"`
	}
	GetMaterialsParams struct {
		ServiceId string `json:"service_id" validate:"required,mongodb"`
	}
)

func CreateMaterialValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &CreateMaterialParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &CreateMaterialParams{ServiceId: body.ServiceId, Name: body.Name, Number: body.Number, Price: body.Price}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsServiceById(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del servicio no existe", ""))
		}
		// verificando que el servicio este activo
		if !models.IsActiveService(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: el servicio no esta activo", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func UpdateMaterialValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &UpdateMaterialParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &UpdateMaterialParams{MaterialId: body.MaterialId, Name: body.Name, Number: body.Number, Price: body.Price}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsMaterialById(v.MaterialId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del material no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func GetMaterialsValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo query
		serviceId := c.QueryParam("service_id")
		// estableciendo los argumentos de validacion
		v := &GetMaterialsParams{ServiceId: serviceId}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsServiceById(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del servicio no existe", ""))
		}
		// fin del middleware
		return next(c)
	}
}
