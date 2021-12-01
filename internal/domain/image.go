package domain

import uuid "github.com/google/uuid"

type Image struct {
	Base
	Path        string    `json:"path"`
	OrphanageID uuid.UUID `json:"orphanage_id"`
}
