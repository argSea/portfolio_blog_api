package auth

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// consts
const (
	// permissions
	PERM_USER  = "user"
	PERM_ADMIN = "admin"
)

func HashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if nil != err {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}

	return string(bcryptPassword), nil
}

// check jwt token
func CheckToken(token string) (jwt.MapClaims, bool, error) {
	// parse jwt
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})

	return claims, true, err
}

func AuthorizeRole(token string, roles ...string) (jwt.MapClaims, bool) {
	claims, authorized, err := CheckToken(token)

	if nil != err {
		log.Println("Error checking token: ", err)
		return nil, false
	}

	if !authorized {
		log.Println("Unauthorized")
		return nil, false
	}

	// check if user has required role
	for _, role := range roles {
		if claims["role"] == role {
			return claims, true
		}
	}

	return nil, false
}
