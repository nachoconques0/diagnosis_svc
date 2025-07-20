package service

import (
	"context"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/model"
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
func (s *Service) Create(ctx context.Context, req model.CreatePatientRequest) (*model.PatientResponse, error) {
	p := patient.New(req.Name, req.DNI, req.Email, req.Phone, req.Address)
	if err := p.Valid(); err != nil {
		return nil, err
	}
	res, err := s.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return &model.PatientResponse{
		ID:      res.ID.String(),
		Name:    res.Name,
		DNI:     res.DNI,
		Email:   res.Email,
		Phone:   res.Phone,
		Address: res.Address,
	}, nil
}

// Find patients and filters like name & pagination can be used
func (s *Service) Find(ctx context.Context, filters query.DiagnosisFilters, pagination query.Pagination) ([]model.PatientResponse, error) {
	res, err := s.repo.Find(ctx, filters, pagination)
	if err != nil {
		return []model.PatientResponse{}, err
	}
	var result []model.PatientResponse
	if len(res) > 0 {
		for _, v := range res {
			result = append(result, model.PatientResponse{
				ID:      v.ID.String(),
				Name:    v.Name,
				DNI:     v.DNI,
				Email:   v.Email,
				Phone:   v.Phone,
				Address: v.Address})
		}
	}
	return result, nil
}
