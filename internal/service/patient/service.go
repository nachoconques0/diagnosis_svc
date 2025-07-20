package service

import (
	"context"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
)

type Service struct {
	repo repository
}

type repository interface {
	Create(ctx context.Context, p *patient.Entity) (*patient.Entity, error)
	Find(
		ctx context.Context,
		filters query.DiagnosisFilters,
		pagination query.Pagination,
	) ([]patient.Entity, error)
}

// New returns a new patient service
func New(repo repository) *Service {
	return &Service{repo: repo}
}

// Create validates and stores a patient
func (s *Service) Create(ctx context.Context, name, email, dni string, phone *string, add *string) (*patient.Entity, error) {
	p := patient.New(name, dni, email, phone, add)
	if err := p.Valid(); err != nil {
		return nil, err
	}
	res, err := s.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Find patients and filters like name & pagination can be used
func (s *Service) Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]patient.Entity, error) {
	res, err := s.repo.Find(ctx, filters, pagination)
	if err != nil {
		return []patient.Entity{}, err
	}
	return res, nil
}
