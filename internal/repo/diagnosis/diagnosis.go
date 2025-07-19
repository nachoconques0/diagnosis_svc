package diagnosis

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
)

type Repository struct {
	db *gorm.DB
}

// NewRepository returns a new diagnosis repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Create inserts a new diagnosis into the database
func (r *Repository) Create(ctx context.Context, d *diagnosis.Entity) (*diagnosis.Entity, error) {
	d.ID = uuid.New()
	if err := r.db.WithContext(ctx).Create(d).Error; err != nil {
		return nil, errors.NewInternalError("could not create diagnosis: " + err.Error())
	}
	return d, nil
}

// Find returns patient diagnosis, and it can be filtered by patient name or diagnosis created_at
func (r *Repository) Find(
	ctx context.Context,
	filters query.DiagnosisFilters,
	pagination query.Pagination,
) ([]diagnosis.Entity, error) {
	var results []diagnosis.Entity

	tx := r.db.WithContext(ctx).
		Table("top_doctor.diagnose AS d").
		Joins("JOIN top_doctor.patient p ON p.id = d.patient_id")

	if filters.PatientName != "" {
		tx = tx.Where("p.name = ?", filters.PatientName)
	}

	if filters.Date != nil {
		tx = tx.Where("DATE(d.created_at) = ?", filters.Date.Format("2006-01-02"))
	}

	if err := tx.
		Order("d.created_at DESC").
		Limit(pagination.Limit()).
		Offset(pagination.Offset()).
		Find(&results).Error; err != nil {
		return nil, errors.NewInternalError("could not find diagnoses: " + err.Error())
	}

	return results, nil
}
