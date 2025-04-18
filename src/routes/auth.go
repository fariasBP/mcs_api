package routes

import (
	"mcs_api/src/controllers"
	"mcs_api/src/middlewares"
	"mcs_api/src/validations"

	"github.com/labstack/echo/v4"
)

func Auth(e *echo.Group) {
	router := e.Group("/auth")
	router.POST("/login", controllers.Login, validations.LoginValidate)
	router.POST("/signup", controllers.SignUp, validations.SignUpValidate)
	router.GET("/check-token", controllers.CheckToken, middlewares.ValidateToken)
	router.PUT("/upgrade-to-admin", controllers.UpgradeToAdmin, validations.UpgradeToAdminValidate, middlewares.ValidateToken, middlewares.IsSuper)
	router.PUT("/change-permission", controllers.ChangePermission, validations.ChangePermissionValidate, middlewares.ValidateToken, middlewares.IsGTEtoAdmin)
	//router.POST("create-user", controllers.CreateUser)

}
