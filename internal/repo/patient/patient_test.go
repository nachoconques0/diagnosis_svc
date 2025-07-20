package patient_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	patientEntity "github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	helpers "github.com/nachoconques0/diagnosis_svc/internal/helpers/db"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/repo/patient"
)

func TestPatientRepository_Create(t *testing.T) {
	db, teardown, err := helpers.NewTestDB()
	if err != nil {
		assert.Nil(t, err)
	}
	defer teardown()
	repo := patient.NewRepository(db)
	ctx := context.Background()

	t.Run("creates a patiente", func(t *testing.T) {
		p := &patientEntity.Entity{
			ID:    uuid.New(),
			Name:  "patientNacho",
			DNI:   "123123",
			Email: "nacho@topdoc.com",
		}
		res, err := repo.Create(ctx, p)
		assert.NoError(t, err)
		assert.Equal(t, p.Email, res.Email)
	})
}

func TestPatientRepository_Find(t *testing.T) {
	db, teardown, err := helpers.NewTestDB()
	if err != nil {
		assert.Nil(t, err)
	}
	defer teardown()
	repo := patient.NewRepository(db)
	ctx := context.Background()

	p := &patientEntity.Entity{
		ID:    uuid.New(),
		Name:  "patientNacho",
		DNI:   "123123",
		Email: "nacho@topdoc.com",
	}
	res, err := repo.Create(ctx, p)
	assert.NoError(t, err)

	t.Run("no filters", func(t *testing.T) {
		results, err := repo.Find(ctx, query.DiagnosisFilters{}, query.Pagination{Page: 1, PageSize: 10})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(results), 1)
	})

	t.Run("by patient name", func(t *testing.T) {
		filters := query.DiagnosisFilters{PatientName: res.Name}
		results, err := repo.Find(ctx, filters, query.Pagination{Page: 1, PageSize: 10})
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, results[0].Name, res.Name)
	})
}
