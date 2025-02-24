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
	NewProblemParams struct {
		ServiceId  string `json:"service_id" validate:"required,mongodb"`
		ProtocolId string `json:"protocol_id" validate:"required,mongodb"`
		Problem    string `json:"problem" validate:"required,startsnotwith= ,endsnotwith= ,min=20"`
	}
	NewSolutionParams struct {
		ProblemId string `json:"problem_id" validate:"required,mongodb"`
		Solution  string `json:"solution" validate:"required,startsnotwith= ,endsnotwith= ,min=20"`
	}
	UpdateProblemParams struct {
		ProblemId string `json:"problem_id" validate:"required,mongodb"`
		Problem   string `json:"problem" validate:"required,startsnotwith= ,endsnotwith= ,min=20"`
	}
	UpdateSolutionParams struct {
		ProblemId string `json:"problem_id" validate:"required,mongodb"`
		Solution  string `json:"solution" validate:"required,startsnotwith= ,endsnotwith= ,min=20"`
	}
	GetProblemsParams struct {
		ServiceId string `json:"service_id" validate:"required,mongodb"`
	}
)

func NewProblemValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &NewProblemParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &NewProblemParams{ServiceId: body.ServiceId, ProtocolId: body.ProtocolId, Problem: body.Problem}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificar id del servicio
		if !models.ExistsServiceById(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del servicio no existe", ""))
		}
		// verificar id del protocolo
		if !models.ExistsProtocolById(v.ProtocolId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del protocolo no existe", ""))
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

func NewSolutionValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &NewSolutionParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &NewSolutionParams{ProblemId: body.ProblemId, Solution: body.Solution}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificar id del problema
		if !models.ExistsProblemById(v.ProblemId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del problema no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))

		return next(c)
	}
}

func UpdateProblemValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &UpdateProblemParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &UpdateProblemParams{ProblemId: body.ProblemId, Problem: body.Problem}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificar id del problema
		if !models.ExistsProblemById(v.ProblemId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del problema no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))

		return next(c)
	}
}

func UpdateSolutionValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &UpdateSolutionParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &UpdateSolutionParams{ProblemId: body.ProblemId, Solution: body.Solution}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificar id del problema
		if !models.ExistsProblemById(v.ProblemId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del problema no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))

		return next(c)
	}
}

func GetProblemsValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo query
		serviceId := c.QueryParam("service_id")
		// estableciendo los argumentos de validacion
		v := &GetProblemsParams{ServiceId: serviceId}
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
