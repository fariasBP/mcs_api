package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/middlewares"
	"mcs_api/src/models"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.LoginParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// buscando usuario
	user, err := models.GetUser(body.Identifier)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo encontrar al usuario", err.Error()))
	}
	// comparando contraseñas
	if user.Pwd != body.Pwd {
		return c.JSON(401, config.SetResError(401, "Contraseña incorrecta", ""))
	}
	// creando token
	token, expiresJWT, err := middlewares.CreateToken(user.IdName)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo crear el token", err.Error()))
	}

	return c.JSON(200, config.SetResToken(200, "Token creado", token, expiresJWT))
}
func SignUp(c echo.Context) error {
	// obteniendo variables de request
	body := &models.User{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// verificando que no exista el usuario
	if models.ExistsUser(body.IdName, body.Email) {
		return c.JSON(400, config.SetResError(400, "User already exists", ""))
	}
	// creando usuario
	err := models.CreateUser(body.IdName, body.Name, body.Lname, body.Email, body.Pwd, body.Bth)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Internal Server Error", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "User created"))
}
