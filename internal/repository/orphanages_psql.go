package repository

import (
	"mime/multipart"

	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/gentildpinto/h-api/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrphanagesRepo struct {
	db *gorm.DB
}

func NewOrphanagesRepo(db *gorm.DB) *OrphanagesRepo {
	return &OrphanagesRepo{db: db}
}

func (r *OrphanagesRepo) All() (orphanages []domain.Orphanage, err error) {
	if err = r.db.Preload("Images").Find(&orphanages).Error; err != nil {
		logger.Error(err)
	}
	return
}

func (r *OrphanagesRepo) FindByID(id string) (orphanage domain.Orphanage, err error) {
	uid, _ := uuid.Parse(id)

	if err = r.db.Preload("Images").First(&orphanage, uid).Error; err != nil {
		logger.Error(err)
		return domain.Orphanage{}, err
	}
	return
}

func (r *OrphanagesRepo) Create(orphanage *domain.Orphanage, images []*multipart.FileHeader) (err error) {
	tx := r.db.Begin()

	if err = tx.Create(&orphanage).Error; err != nil {
		logger.Error(err)
		tx.Rollback()
		return
	}

	if err = createImages(tx, orphanage.ID, images); err != nil {
		logger.Error(err)
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}
