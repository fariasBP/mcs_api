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
	LoginParams struct {
		Identifier string `json:"user" validate:"required"` // "nick" o "email" del usuario
		Pwd        string `json:"pwd" validate:"required"`
	}
	SignUpParams struct {
		Nick  string `json:"nick" validate:"required,lowercase,min=3,max=30"`
		Name  string `json:"name" validate:"required,lowercase,min=3,max=30"`
		Lname string `json:"lname" validate:"required,lowercase,min=3,max=30"`
		Email string `json:"email" validate:"required,email"`
		Pwd   string `json:"pwd" validate:"required"`
		Bth   string `json:"bth" validate:"required,datetime=2006-01-02"`
	}
	CheckTokenParams struct {
		Token string `json:"token" validate:"required,jwt"`
	}
	UpgradeToAdminParams struct {
		UserId string `json:"user_id" validate:"required,mongodb"`
	}
	ChangePermissionParams struct {
		UserId   string `json:"user_id" validate:"required,mongodb"`
		Operator bool   `json:"operator" validate:"boolean"` // true: operador, false: public
	}
)

func LoginValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &LoginParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &LoginParams{Identifier: body.Identifier, Pwd: body.Pwd}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos.", err.Error()))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func SignUpValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &SignUpParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &SignUpParams{Nick: body.Nick, Name: body.Name, Lname: body.Lname, Email: body.Email, Pwd: body.Pwd, Bth: body.Bth}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos.", err.Error()))
		}
		// verificando que no exista el usuario
		if models.ExistsUser(body.Nick, body.Email) {
			return c.JSON(400, config.SetResError(400, "Error: El nick o email del usuario: '"+body.Nick+"'/'"+body.Email+"' ya existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func CheckTokenValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo cabecera token
		tkn := c.Request().Header.Get("Access-Token")
		// estableciendo los argumentos de validacion
		v := &CheckTokenParams{Token: tkn}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos.", err.Error()))
		}
		// fin del middleware
		return next(c)
	}
}

func UpgradeToAdminValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &UpgradeToAdminParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &UpgradeToAdminParams{UserId: body.UserId}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos.", err.Error()))
		}
		// verificar id del usuario
		if !models.ExistsUserById(v.UserId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del usuario no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}

func ChangePermissionValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo body json
		body := &ChangePermissionParams{}
		data, _ := io.ReadAll(c.Request().Body)
		reader := bytes.NewReader(data)
		_ = json.NewDecoder(reader).Decode(body)
		// estableciendo los argumentos de validacion
		v := &ChangePermissionParams{UserId: body.UserId, Operator: body.Operator}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Parametros invalidos.", err.Error()))
		}
		// verificar id del usuario
		if !models.ExistsUserById(v.UserId) {
			return c.JSON(400, config.SetResError(400, "Error: el id del usuario no existe", ""))
		}
		// fin del middleware
		c.Request().Body = io.NopCloser(bytes.NewReader([]byte(data)))
		return next(c)
	}
}
