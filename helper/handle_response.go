package helper

import "github.com/gofiber/fiber/v2"

func HandleResponse(ctx *fiber.Ctx, err error, status int, statusErr int, message string, data interface{}) error {
	errCode := 500
	if statusErr != 0 {
		errCode = statusErr
	}

	if err != nil {
		return ApiResponse(ctx, errCode, err.Error(), nil)
	}
	return ApiResponse(ctx, status, message, data)
}
