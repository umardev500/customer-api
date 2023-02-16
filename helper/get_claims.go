package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetClaims(ctx *fiber.Ctx, authorizationHeader string) (jwt.MapClaims, error) {
	userData, err := GetUserInfo(authorizationHeader)
	if err != nil {
		return nil, HandleResponse(ctx, err, 0, 0, err.Error(), nil)
	}

	claims, err := userData()
	if err != nil {
		return nil, HandleResponse(ctx, err, 0, 400, err.Error(), nil)
	}

	return claims, nil
}
