package service

import (
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
