package helper

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func GetUserInfo(authorizationHeader string) (func() (jwt.MapClaims, error), error) {
	secret := os.Getenv("SECRET")
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("singing method invalid")
		} else if method != jwt.SigningMethodHS256 {

			return nil, errors.New("singing method invalid")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return func() (jwt.MapClaims, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return nil, errors.New(http.StatusText(http.StatusBadRequest))
		}

		return claims, nil
	}, nil

}
