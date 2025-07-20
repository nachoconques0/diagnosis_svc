package service

import (
	"context"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
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

// New returns a new diagnosis service
func New(repo repository) *Service {
	return &Service{repo: repo}
}

// Create validates and stores a diagnosis
func (s *Service) Create(ctx context.Context, req model.CreateDiagnosisRequest) (*model.DiagnosisResponse, error) {
	d, err := diagnosis.New(req.PatientID, req.Diagnosis, req.Prescription)
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
	return &model.DiagnosisResponse{
		ID:           res.ID.String(),
		PatientID:    res.PatientID.String(),
		Diagnosis:    res.Diagnosis,
		Prescription: res.Prescription,
		CreatedAt:    res.CreatedAt,
	}, nil
}

// Find diagnoses based on filters
func (s *Service) Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]model.DiagnosisResponse, error) {
	res, err := s.repo.Find(ctx, filters, pagination)
	if err != nil {
		return []model.DiagnosisResponse{}, err
	}
	var result []model.DiagnosisResponse
	if len(res) > 0 {
		for _, v := range res {
			result = append(result, model.DiagnosisResponse{
				ID:           v.ID.String(),
				PatientID:    v.PatientID.String(),
				Diagnosis:    v.Diagnosis,
				Prescription: v.Prescription,
				CreatedAt:    v.CreatedAt,
			})
		}
	}
	return result, nil
}
