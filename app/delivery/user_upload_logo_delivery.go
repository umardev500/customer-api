package delivery

import (
	"customer-api/domain"
	"customer-api/helper"
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
)

func (u *appDelivery) UploadLogo(ctx *fiber.Ctx) error {
	authorizationHeader := ctx.Get("Authorization")
	claims, f := helper.GetClaims(ctx, authorizationHeader)
	if claims == nil {
		return f
	}

	file, err := ctx.FormFile("logo")
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, "upload logo", nil)
	}

	rootPath := "./public"
	fileLocation := fmt.Sprintf("/uploads/images/avatars/%s", file.Filename)
	filePath := fmt.Sprintf("%s%s", rootPath, fileLocation)
	out, err := os.Create(filePath)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, "upload logo", nil)
	}

	defer out.Close()

	fileReader, err := file.Open()
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, "upload logo", nil)
	}

	defer fileReader.Close()

	var contents []byte
	var buf = make([]byte, 1024)
	for {
		n, err := fileReader.Read(buf)
		if err != nil && err != io.EOF {
			return helper.HandleResponse(ctx, err, 0, 500, "upload logo", nil)
		}

		if n == 0 {
			break
		}

		contents = append(contents, buf[:n]...)
	}

	_, err = out.Write(contents)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, "upload logo", nil)
	}

	appCtx := ctx.Context()
	userId := claims["user_id"].(string)
	var detailData domain.CustomerDetail = domain.CustomerDetail{Logo: fileLocation}

	resp, err := u.appUsecase.Update(appCtx, userId, detailData)
	if err != nil {
		return helper.HandleResponse(ctx, err, 0, 500, err.Error(), nil)
	}

	return helper.HandleResponse(ctx, err, 200, 0, "Upload logo", resp)
}
