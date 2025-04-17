package validations

import (
	"mcs_api/src/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	CheckTokenParams struct {
		Token string `json:"token" validate:"required,jwt"`
	}
	LoginParams struct {
		Identifier string `json:"user" validate:"required"`
		Pwd        string `json:"pwd" validate:"required"`
	}
	SignUpParams struct {
		IdName string `json:"id_name" validate:"required,lowercase,min=5,max=30"`
		Name   string `json:"name" validate:"required,lowercase,min=3,max=30"`
		Lname  string `json:"lname" validate:"required,lowercase,min=3,max=30"`
		Email  string `json:"email" validate:"required,email"`
		Pwd    string `json:"pwd" validate:"required"`
		Bth    string `json:"bth" validate:"required,date"`
	}
)

func CheckTokenValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo cabecera token
		tkn := c.Request().Header.Get("Access-Token")
		// estableciendo los argumentos de validacion
		v := &CheckTokenParams{Token: tkn}
		// realizando valdacion
		validate := validator.New()
		if err := validate.Struct(v); err != nil {
			return c.JSON(400, config.SetResError(400, "Error: Valores invalidos.", err.Error()))
		}
		// fin del middleware
		return next(c)
	}
}
