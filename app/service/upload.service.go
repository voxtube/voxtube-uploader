package service

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	_ "github.com/joho/godotenv/autoload"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	ctx             context.Context = context.Background()
	cld, err                        = cloudinary.New()
	UploadImageDir  string          = "uploads/images"
	UploadDocsDir   string          = "uploads/docs"
	UploadOthersDir string          = "uploads/others"
)

func init() {
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func FileUploader(file multipart.File, dir string) (*uploader.UploadResult, error) {
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	uploadResult, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{PublicID: id, Folder: dir})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	return uploadResult, nil
}

func DeleteFile(publicID string) (*uploader.DestroyResult, error) {
	return cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
}
