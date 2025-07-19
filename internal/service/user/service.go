package service

import (
	"context"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/user"
)

type repository interface {
	GetByEmail(ctx context.Context, email string) (*user.Entity, error)
}

type Service struct {
	repo repository
}

func New(repo repository) *Service {
	return &Service{repo: repo}
}

// GetByEmail returns an user by its email
func (s *Service) GetByEmail(ctx context.Context, email string) (*user.Entity, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
