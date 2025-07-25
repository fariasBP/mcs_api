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
	NewBrandParams struct {
		Name string `json:"name" validate:"required,startsnotwith= ,endsnotwith= ,min=3,max=50"` // "name" (nombre) de la marca (fabricante)
	}

	GetBrandsParams struct {
		Query string `json:"query" validate:""` // "name" (nombre) de la marca (fabricante)
		Limit string `json:"limit" validate:"required,number,gt=0"`
		Page  string `json:"page" validate:"required,number,gt=0"`
	}

	UpdateBrandParams struct {
		BrandId string `json:"brand_id" validate:"required,mongodb"`
		Name    string `json:"name" validate:"required,startsnotwith= ,endsnotwith= ,min=3,max=50"`
	}
)

func NewBrandValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &NewBrandParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &NewBrandParams{Name: body.Name}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// verificando que no exista el nombre de la marca
		if models.ExistsBrand(body.Name) {
			return c.JSON(400, config.SetResError(400, "Error: El nombre del marca (fabricante): '"+body.Name+"' ya existe:", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))

		return next(c)
	}
}

func GetBrandsValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo query
		query := c.QueryParam("query")
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")
		// estableciendo los argumentos de validacion
		v := &GetBrandsParams{Query: query, Limit: limit, Page: page}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// fin del middleware
		return next(c)
	}
}

func UpdateBrandValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &UpdateBrandParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &UpdateBrandParams{BrandId: body.BrandId, Name: body.Name}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// verificar id de la marca
		if !models.ExistsBrandById(v.BrandId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la marca (fabricante) no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}
