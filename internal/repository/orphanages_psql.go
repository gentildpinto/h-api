package repository

import (
	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/gentildpinto/h-api/pkg/logger"
	"gorm.io/gorm"
)

type OrphanagesRepo struct {
	db *gorm.DB
}

func NewOrphanagesRepo(db *gorm.DB) *OrphanagesRepo {
	return &OrphanagesRepo{db: db}
}

func (r *OrphanagesRepo) All() (orphanages []domain.Orphanage, err error) {
	if err = r.db.Find(&orphanages).Error; err != nil {
		logger.Error(err)
	}
	return
}
