package service

import (
	"mime/multipart"

	"github.com/gentildpinto/h-api/internal/domain"
	"github.com/gentildpinto/h-api/internal/repository"
)

type (
	OrphanagesService struct {
		repo repository.Orphanages
	}
)

func NewOrphanagesService(repo repository.Orphanages) *OrphanagesService {
	return &OrphanagesService{
		repo: repo,
	}
}

func (s *OrphanagesService) All() ([]domain.Orphanage, error) {
	return s.repo.All()
}

func (s *OrphanagesService) FindByID(id string) (domain.Orphanage, error) {
	return s.repo.FindByID(id)
}

func (s *OrphanagesService) Create(orphanage *domain.Orphanage, images []*multipart.FileHeader) error {
	return s.repo.Create(orphanage, images)
}
