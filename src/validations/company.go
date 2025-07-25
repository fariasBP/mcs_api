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
	NewCompanyParams struct {
		Name        string `json:"name" validate:"required,min=2,max=50,startsnotwith= ,endsnotwith= "`
		Manager     string `json:"manager" validate:"required,min=3,max=50,startsnotwith= ,endsnotwith= "`
		Latitude    string `json:"latitude" validate:"required,latitude"`
		Longitude   string `json:"longitude" validate:"required,longitude"`
		Description string `json:"description" validate:"required,startsnotwith= ,endsnotwith= ,min=4"`
		Contact     string `json:"contact" validate:"required,min=5"`
	}
	GetCompaniesParams struct {
		Query string `json:"query" validate:""` // "name" (nombre) de la compania
		Limit string `json:"limit" validate:"required,number,gt=0"`
		Page  string `json:"page" validate:"required,number,gt=0"`
	}
	UpdateCompanyParams struct {
		CompanyId   string `json:"company_id" validate:"required,mongodb"`
		Name        string `json:"name" validate:"required,min=2,max=50,startsnotwith= ,endsnotwith= "`
		Manager     string `json:"manager" validate:"required,min=3,max=50,startsnotwith= ,endsnotwith= "`
		Latitude    string `json:"latitude" validate:"required,latitude"`
		Longitude   string `json:"longitude" validate:"required,longitude"`
		Description string `json:"description" validate:"required,startsnotwith= ,endsnotwith= ,min=4"`
		Contact     string `json:"contact" validate:"required,min=5"`
	}
)

func NewCompanyValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &NewCompanyParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &NewCompanyParams{Name: body.Name, Manager: body.Manager, Latitude: body.Latitude, Longitude: body.Longitude, Description: body.Description, Contact: body.Contact}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// verificando que no exista el nombre de la compania
		if models.ExistsCompany(body.Name) {
			return c.JSON(400, config.SetResError(400, "Error: El nombre: '"+body.Name+"' de la compania ya existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func GetCompaniesValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo query
		query := c.QueryParam("query")
		limit := c.QueryParam("limit")
		page := c.QueryParam("page")
		// estableciendo los argumentos de validacion
		v := &GetCompaniesParams{Query: query, Limit: limit, Page: page}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// fin del middleware
		return next(c)
	}
}

func UpdateCompanyValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &UpdateCompanyParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &UpdateCompanyParams{CompanyId: body.CompanyId, Name: body.Name, Manager: body.Manager, Latitude: body.Latitude, Longitude: body.Longitude, Description: body.Description, Contact: body.Contact}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos", err.Error()))
		}
		// verificar id de la compania
		if !models.ExistsCompanyById(v.CompanyId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la compania no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}
