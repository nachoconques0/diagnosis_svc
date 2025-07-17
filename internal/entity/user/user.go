package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// TableName define user table name
	TableName = "top_doctor.user"
)

// TableName returns table name
func (Entity) TableName() string {
	return TableName
}

type Entity struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Nickname  string
	Password  string
	Email     string `gorm:"not null;unique"`
	CreatedAt time.Time
}
