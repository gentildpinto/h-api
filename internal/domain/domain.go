package domain

import (
	"time"

	uuid "github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
