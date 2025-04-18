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
	user, err := models.GetUserAndPwd(body.Identifier)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo encontrar al usuario con nick/email: '"+body.Identifier+"'", err.Error()))
	}
	// comparando contraseñas
	if !middlewares.CheckPasswordHash(body.Pwd, user.Pwd) {
		return c.JSON(401, config.SetResError(401, "Error: Contraseña incorrecta", ""))
	}
	// creando token
	token, expiresJWT, err := middlewares.CreateToken(user.ID.Hex(), user.Perm)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error:No se pudo crear el token", err.Error()))
	}

	return c.JSON(200, config.SetResToken(200, "Login exitoso, Token creado", token, expiresJWT))
}
func SignUp(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.SignUpParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()

	// encriptando contrraseña
	hashPwd, err := middlewares.HashPassword(body.Pwd)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear la contraseña", err.Error()))
	}
	// creando usuario
	err = models.CreateUser(body.Nick, body.Name, body.Lname, body.Email, hashPwd, body.Bth, models.Public)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo crear el usuario", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El usuario: '"+body.Nick+"' se ha creado"))
}

func CheckToken(c echo.Context) error {
	// obteniendo variables de request
	id := c.Get("id").(string)

	// buscando usuario
	user, err := models.GetUserById(id)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo encontrar al usuario", err.Error()))
	}

	return c.JSON(200, config.SetResJson(200, "El token es valido", user))
}

func UpgradeToAdmin(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.UpgradeToAdminParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// buscando usuario y verificando permisos
	user, err := models.GetUserById(body.UserId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo encontrar al usuario con id: '"+body.UserId+"'", err.Error()))
	}
	if user.Perm == models.Admin {
		return c.JSON(400, config.SetResError(400, "Error: el usuario: '"+user.Nick+"' ya es 'Admin'", ""))
	}
	// cambiando permisos
	user.Perm = models.Admin
	err = models.UpdateUser(user)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo actualizar el usuario: '"+user.Nick+"'", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El usuario: '"+user.Nick+"' ahora es admin"))
}

func ChangePermission(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.ChangePermissionParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// buscando usuario
	user, err := models.GetUserById(body.UserId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo encontrar al usuario con id: '"+body.UserId+"'", err.Error()))
	}
	// creando texto referencia
	user.Perm = models.Public
	ref := "Public"
	if body.Operator {
		user.Perm = models.Operator
		ref = "Operator"
	}
	// cambiando permisos
	err = models.UpdateUser(user)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "Error: No se pudo actualizar el usuario: '"+user.Nick+"'", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El usuario: '"+user.Nick+"' ahora tiene el permiso: '"+ref+"'"))
}
