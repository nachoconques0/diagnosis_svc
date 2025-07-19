package service

import (
	"context"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
)

type Service struct {
	repo repository
}

type repository interface {
	Create(ctx context.Context, d *diagnosis.Entity) (*diagnosis.Entity, error)
	Find(
		ctx context.Context,
		filters query.DiagnosisFilters,
		pagination query.Pagination,
	) ([]diagnosis.Entity, error)
}

func New(repo repository) *Service {
	return &Service{repo: repo}
}

// Create validates and stores a diagnosis
func (s *Service) Create(ctx context.Context, patientID string, diag string, prescription *string) (*diagnosis.Entity, error) {
	d, err := diagnosis.New(patientID, diag, prescription)
	if err != nil {
		return nil, err
	}

	if err := d.Valid(); err != nil {
		return nil, err
	}
	res, err := s.repo.Create(ctx, d)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Find diagnoses based on filters
func (s *Service) Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]diagnosis.Entity, error) {
	return s.repo.Find(ctx, filters, pagination)
}
