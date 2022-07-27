package controller

import (
	"voxtube-uploader/app/service"

	"github.com/gofiber/fiber/v2"
)

// @Router /accounts/{id} [get] .
func UploadImage(ctx *fiber.Ctx) error {
	// Get image from body
	formFile, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if tempFile, err := formFile.Open(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	} else {
		imageUploader, err := service.FileUploader(tempFile, service.UploadImageDir)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err,
			})
		}
		return ctx.JSON(imageUploader)
	}
}

func UploadVideo(ctx *fiber.Ctx) error {
	// Get video from body
	formFile, err := ctx.FormFile("video")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if tempFile, err := formFile.Open(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	} else {
		videoUploader, err := service.FileUploader(tempFile, service.UploadVideoDir)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err,
			})
		}
		return ctx.JSON(videoUploader)
	}
}

func UploadDocs(ctx *fiber.Ctx) error {
	// Get image from body
	formFile, err := ctx.FormFile("doc")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if tempFile, err := formFile.Open(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	} else {
		docsUploader, err := service.FileUploader(tempFile, service.UploadDocsDir)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err,
			})
		}
		return ctx.JSON(docsUploader)
	}
}

func UploadOthers(ctx *fiber.Ctx) error {
	// Get image from body
	formFile, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if tempFile, err := formFile.Open(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	} else {
		othersUploader, err := service.FileUploader(tempFile, service.UploadOthersDir)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err,
			})
		}
		return ctx.JSON(othersUploader)
	}
}

func Deletefile(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")

	result, err := service.DeleteFile(publicID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(fiber.Map{
		"deleted": true,
		"result":  &result,
	})
}
