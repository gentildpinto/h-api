package repository

import (
	"github.com/gentildpinto/h-api/internal/domain"
	"gorm.io/gorm"
)

type (
	Orphanages interface {
		All() ([]domain.Orphanage, error)
		FindByID(id string) (domain.Orphanage, error)
		Create(orphanage *domain.Orphanage) error
	}

	Repositories struct {
		Orphanages Orphanages
	}
)

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Orphanages: NewOrphanagesRepo(db),
	}
}
