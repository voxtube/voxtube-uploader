package route

import (
	"voxtube-uploader/app/controller"

	"github.com/gofiber/fiber/v2"
)

func UploadRoute(app fiber.Router) {
	uplaod := app.Group("/upload")

	uplaod.Post("/image", controller.UploadImage)
	uplaod.Post("/doc", controller.UploadDocs)
	uplaod.Post("/other", controller.UploadOthers)
	uplaod.Delete("/file/:id", controller.Deletefile)
}
