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
	ProblemParams struct {
		Problem  string `json:"problem" validate:"required"`
		Solution string `json:"solution" validate:"required"`
	}
	MaterialParams struct {
		Name   string `json:"name" validate:"required"`
		Number int    `json:"number" validate:"required"`
		Price  int    `json:"price" validate:"required"`
	}
	ProtocolParams struct {
		ProtocolId string          `json:"protocol_id" validate:"required"`
		Status     int             `json:"status" validate:"required"`
		Note       string          `json:"note" validate:"required"`
		Problems   []ProblemParams `json:"problems" validate:"required"`
	}

	NewServiceParams struct {
		MachineId string `json:"machine_id" validate:"required,mongodb"`
	}

	AddMaterialToServiceParams struct {
		ServiceId string `json:"service_id" validate:"required,mongodb"`
		Name      string `json:"name" validate:"required,lowercase,min=3,max=30"`
		Number    int    `json:"number" validate:"required,number,gt=0"`
		Price     int    `json:"price" validate:"required,number,gt=0"`
	}
	SleepServiceParams struct {
		ServiceId string `json:"service_id" validate:"required,mongodb"`
	}
	FinishServiceParams struct {
		ServiceId string `json:"service_id" validate:"required,mongodb"`
		Cancelled bool   `json:"cancelled" validate:"boolean"`
	}
	ProgressServiceParams struct {
		ServiceId string `json:"service_id" validate:"required,mongodb"`
		Progress  int    `json:"progress" validate:"required,number,gt=0,lt=4"`
	}
)

func NewServiceValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &NewServiceParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &NewServiceParams{MachineId: body.MachineId}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsMachineById(v.MachineId) {
			return c.JSON(400, config.SetResError(400, "Error: el id de la maquina no existe", ""))
		}
		// verificando que la maquina no tenga un servicio activo
		if models.ExistsServiceActiveFromMachineById(v.MachineId) {
			return c.JSON(400, config.SetResError(400, "Error: la maquina ya tiene un servicio activo", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func AddMaterialToServiceValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &AddMaterialToServiceParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &AddMaterialToServiceParams{
			ServiceId: body.ServiceId,
			Name:      body.Name,
			Number:    body.Number,
			Price:     body.Price,
		}
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
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func SleepServiceValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &SleepServiceParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &SleepServiceParams{
			ServiceId: body.ServiceId,
		}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsServiceById(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del servicio no existe", ""))
		}
		// verificando si el servicio esta activo
		if !models.IsActiveService(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "El servicio ya esta inactivo o finalizado", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func FinishServiceValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &FinishServiceParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &FinishServiceParams{
			ServiceId: body.ServiceId,
			Cancelled: body.Cancelled,
		}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificando que el id exista
		// verificando que el id exista
		if !models.ExistsServiceById(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del servicio no existe", ""))
		}
		// verificando si el servicio esta activo
		if !models.IsActiveService(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: El servicio esta inactivo, tiene que estar activo para poder finalizar", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func ProgressServiceValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &ProgressServiceParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &ProgressServiceParams{
			ServiceId: body.ServiceId,
			Progress:  body.Progress,
		}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// verificando que el id exista
		if !models.ExistsServiceById(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del servicio no existe", ""))
		}
		// verificando si el servicio no ha finalizado
		if models.IsFinishedService(v.ServiceId) {
			return c.JSON(400, config.SetResError(400, "Error: El servicio ya ha finalizado, no puede progresar", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}
