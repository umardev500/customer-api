package helper

import (
	"customer-api/domain"

	"github.com/gofiber/fiber/v2"
)

func ApiResponse(ctx *fiber.Ctx, code int, message string, data interface{}) error {
	var response domain.APIResponse = domain.APIResponse{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}

	return ctx.JSON(response)
}
