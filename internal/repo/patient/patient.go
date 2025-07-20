package patient

import (
	"context"

	"github.com/google/uuid"
	"github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// NewRepository returns a new diagnosis repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Create inserts a new patient into the database
func (r *Repository) Create(ctx context.Context, p *patient.Entity) (*patient.Entity, error) {
	p.ID = uuid.New()
	if err := r.db.WithContext(ctx).Create(p).Error; err != nil {
		return nil, errors.NewInternalError("could not create patient: " + err.Error())
	}
	return p, nil
}

// Find returns patients, and it can be filtered by patient name
func (r *Repository) Find(
	ctx context.Context,
	filters query.DiagnosisFilters,
	pagination query.Pagination,
) ([]patient.Entity, error) {
	var results []patient.Entity

	tx := r.db
	if filters.PatientName != "" {
		tx = tx.Where("patient.name = ?", filters.PatientName)
	}
	if err := tx.
		Limit(pagination.Limit()).
		Offset(pagination.Offset()).
		Find(&results).Error; err != nil {
		return nil, errors.NewInternalError("could not find patients: " + err.Error())
	}

	return results, nil
}
