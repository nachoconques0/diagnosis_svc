package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/user"
	"github.com/nachoconques0/diagnosis_svc/internal/errors"
)

// Repository in charge of managing User repo
type Repository struct {
	db *gorm.DB
}

// NewRepository returns a new diagnosis repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// GetByEmail finds a user by email
func (r *Repository) GetByEmail(ctx context.Context, email string) (*user.Entity, error) {
	var u user.Entity
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFound("user not found")
		}
		return nil, errors.NewInternalError("there was an error obtaining user")
	}
	return &u, nil
}
