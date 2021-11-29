package domain

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	databaseConnection *gorm.DB
	ErrNilDatabase     error = errors.New("database connection is nil")
)

type (
	UUID uuid.UUID

	Base struct {
		ID        UUID      `json:"id" gorm:"primary_key"`
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
		DeletedAt time.Time `json:"-"`
	}
)

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	b.ID = UUID(id)
	return err
}

func (field UUID) String() string {
	return uuid.UUID(field).String()
}

func (field UUID) GormDataType() string {
	return "uuid"
}

func (field UUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(field)
	str := "\"" + s.String() + "\""

	return []byte(str), nil
}

func (field *UUID) UnmarshalJSON(bytes []byte) (err error) {
	s, err := uuid.ParseBytes(bytes)
	*field = UUID(s)

	return
}

func (field *UUID) Scan(value interface{}) (err error) {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*field = UUID(parseByte)

	return
}

func (field UUID) Value() (driver.Value, error) {
	return uuid.UUID(field).MarshalBinary()
}

func ParseUUID(id string) UUID {
	return UUID(uuid.MustParse(id))
}

func Init(dbConn *gorm.DB) error {
	if dbConn == nil {
		return ErrNilDatabase
	}

	databaseConnection = dbConn

	return nil
}
