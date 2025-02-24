package middlewares

import (
	"mcs_api/src/config"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(id string) (string, time.Time, error) {
	// obteniendo variables de entorno
	secretVal, _ := os.LookupEnv("SECRET")
	durationVal, _ := os.LookupEnv("DURATION_JWT")
	// convirtiendo a entero la duracion
	duration, err := strconv.Atoi(durationVal)
	if err != nil {
		duration = 1 // 1 dia(s)
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
		secretVal, _ := os.LookupEnv("SECRET")
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
