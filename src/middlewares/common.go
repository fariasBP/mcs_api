package middlewares

import (
	"fmt"
	"mcs_api/src/config"
	"mcs_api/src/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func Initialization() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	if !models.ExistsSuperUser() {
		fmt.Println("Not exists SuperUser")
		identifier, _ := os.LookupEnv("IDENTIFIER")
		name, _ := os.LookupEnv("NAME")
		lastname, _ := os.LookupEnv("LASTNAME")
		email, _ := os.LookupEnv("EMAIL")
		password, _ := os.LookupEnv("PASSWORD")
		birthday, _ := os.LookupEnv("BIRTHDAY")
		errc := models.CreateSuperUser(identifier, name, lastname, email, password, birthday)
		if errc != nil {
			return errc
		}
		fmt.Println("SuperUser created")
	}
	return nil
}
func CreateToken(id string) (string, time.Time, error) {
	// obteniendo variables de entorno
	secretVal, _ := os.LookupEnv("SECRET_JWT")
	durationVal, _ := os.LookupEnv("DURATION_JWT")
	// convirtiendo a entero la duracion
	duration, err := strconv.Atoi(durationVal)
	if err != nil {
		duration = 30 // 30 dias
	}
	// creando la fecha de expiracion (en dias)
	expiresJWT := time.Now().Add(time.Duration(duration) * 24 * time.Hour)
	// creando claims
	claims := &JwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expiresJWT.Unix(),
		},
	}
	// creando token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenErr := token.SignedString([]byte(secretVal))

	return tokenString, expiresJWT, tokenErr
}

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// obteniendo el header Access-Token
		tkn := c.Request().Header.Get("Access-Token")
		if tkn == "" {
			return c.JSON(401, config.SetRes(401, "No se ha proporcionado el token de acceso"))
		}
		// obteniendo variables de entorno secret
		secretVal, _ := os.LookupEnv("SECRET_JWT")
		// parseando token
		claims := &JwtCustomClaims{}
		token, err := jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretVal), nil
		})
		// validando token
		if err != nil || !token.Valid {
			return c.JSON(401, config.SetRes(401, "Token inválido"))
		}
		// creando variables de sesion
		c.Set("id", claims.Id)
		// retornando token
		return next(c)
	}
}
