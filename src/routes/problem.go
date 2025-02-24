package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Problem(e *echo.Group) {
	router := e.Group("/problem", middlewares.ValidateToken)
	router.POST("/new", controllers.NewProblem, validations.NewProblemValidate)
	router.PUT("/new-solution", controllers.NewSolution, validations.NewSolutionValidate)
	router.PUT("/update-problem", controllers.UpdateProblem, validations.UpdateProblemValidate)
	router.PUT("/update-solution", controllers.UpdateSolution, validations.UpdateSolutionValidate)
	router.GET("/problems", controllers.GetProblems, validations.GetProblemsValidate)
}
