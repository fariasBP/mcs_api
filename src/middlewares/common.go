package middlewares

import (
	"fmt"
	"mcs_api/src/models"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type JwtCustomClaims struct {
	Id   string            `json:"id"`
	Perm models.Permission `json:"perm"`
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
		hashPwd, err := HashPassword(password)
		if err != nil {
			return err
		}
		err = models.CreateSuperUser(identifier, name, lastname, email, hashPwd, birthday)
		if err != nil {
			return err
		}
		fmt.Println("SuperUser created")
	}
	return nil
}

// hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
