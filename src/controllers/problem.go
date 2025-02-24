package controllers

import (
	"encoding/json"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func NewProblem(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.NewProblemParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()

	// creando problema
	err := models.NewProblem(body.ServiceId, body.ProtocolId, body.Problem)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se creo el problem", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El problem se ha creado"))
}

func NewSolution(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.NewSolutionParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando problema
	err := models.NewSolution(body.ProblemId, body.Solution)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se creo la solucion", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "La solucion se ha creado"))
}

func UpdateProblem(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.UpdateProblemParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando problema
	err := models.UpdateProblem(body.ProblemId, body.Problem)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo actualizar el problem", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El problem se ha actualizado"))
}

func UpdateSolution(c echo.Context) error {
	// obteniendo variables de request
	body := &validations.UpdateSolutionParams{}
	d := c.Request().Body
	_ = json.NewDecoder(d).Decode(body)
	defer d.Close()
	// creando problema
	err := models.UpdateSolution(body.ProblemId, body.Solution)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo actualizar el solution", err.Error()))
	}

	return c.JSON(200, config.SetRes(200, "El solution se ha actualizado"))
}

func GetProblems(c echo.Context) error {
	// obteniendo variables de request
	serviceId := c.QueryParam("service_id")
	// obteniendo problems
	problems, err := models.GetProblemsByServiceId(serviceId)
	if err != nil {
		return c.JSON(500, config.SetResError(500, "No se pudo obtener los problems", err.Error()))
	}

	return c.JSON(200, config.SetResJson(200, "Problems obtenidos", problems))
}
