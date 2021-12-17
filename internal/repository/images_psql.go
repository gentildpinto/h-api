package repository

import (
	"mime/multipart"

	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/gentildpinto/h-api/pkg/storage"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func createImages(tx *gorm.DB, ophID uuid.UUID, images []*multipart.FileHeader) (err error) {
	for _, image := range images {

		filename, err := storage.UploadImage(image)
		if err != nil {
			return err
		}

		filepath := "https://firebasestorage.googleapis.com/v0/b/happyngrb.appspot.com/o/images%2F" + filename + "?alt=media"

		if err = tx.Create(&domain.Image{
			OrphanageID: ophID,
			Path:        filepath,
		}).Error; err != nil {
			return err
		}
	}

	return
}
