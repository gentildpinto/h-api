package repository

import (
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func createImages(tx *gorm.DB, ophID uuid.UUID, images []*multipart.FileHeader) (err error) {
	for _, image := range images {
		src, err := image.Open()

		if err != nil {
			return err
		}

		defer src.Close()

		filename := strings.Trim(strings.ToLower(image.Filename), " ")

		filepath := "public/uploads/" + filename
		imageurl := os.Getenv("APP_URL") + "/" + filepath

		dst, err := os.Create(filepath)

		if err != nil {
			return err
		}

		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		if err = tx.Create(&domain.Image{
			OrphanageID: ophID,
			Path:        imageurl,
		}).Error; err != nil {
			return err
		}
	}

	return
}
