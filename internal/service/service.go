package service

import (
	"mime/multipart"

	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/gentildpinto/h-api/internal/repository"
)

type (
	Dependencies struct {
		Repos *repository.Repositories
	}

	Services struct {
		Orphanages Orphanages
	}

	Orphanages interface {
		All() ([]domain.Orphanage, error)
		FindByID(id string) (domain.Orphanage, error)
		Create(orphanage *domain.Orphanage, images []*multipart.FileHeader) error
	}
)

func NewServices(deps Dependencies) *Services {
	return &Services{
		Orphanages: NewOrphanagesService(deps.Repos.Orphanages),
	}
}
