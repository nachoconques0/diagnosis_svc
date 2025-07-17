package diagnosis_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	diagnosisEntity "github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	helpers "github.com/nachoconques0/diagnosis_svc/internal/helpers/db"
	"github.com/nachoconques0/diagnosis_svc/internal/helpers/query"
	"github.com/nachoconques0/diagnosis_svc/internal/repo/diagnosis"
)

func TestDiagnosisRepository_Create(t *testing.T) {
	db, teardown, err := helpers.NewTestDB()
	if err != nil {
		assert.Nil(t, err)
	}
	defer teardown()
	repo := diagnosis.NewRepository(db)
	ctx := context.Background()

	t.Run("creates a diagnosis", func(t *testing.T) {
		patientID := insertTestPatient(t, db)
		d := &diagnosisEntity.Entity{
			PatientID: patientID,
			Diagnosis: "un poco de fiebre lol",
			CreatedAt: time.Now(),
		}
		res, err := repo.Create(ctx, d)
		assert.NoError(t, err)
		assert.NotNil(t, res.ID)
		assert.Equal(t, d.Diagnosis, res.Diagnosis)
	})

	t.Run("invalid patient id", func(t *testing.T) {
		patientID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
		d := &diagnosisEntity.Entity{
			PatientID: patientID,
			Diagnosis: "",
			CreatedAt: time.Now(),
		}
		_, err := repo.Create(ctx, d)
		assert.Error(t, err)
	})
}

func TestDiagnosisRepository_Find(t *testing.T) {
	db, teardown, err := helpers.NewTestDB()
	if err != nil {
		assert.Nil(t, err)
	}
	defer teardown()
	repo := diagnosis.NewRepository(db)
	ctx := context.Background()

	patientID := insertTestPatient(t, db)
	insertMultipleDiagnoses(t, db, patientID, 5)

	t.Run("no filters", func(t *testing.T) {
		results, err := repo.Find(ctx, query.DiagnosisFilters{}, query.Pagination{Page: 1, PageSize: 10})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(results), 1)
	})

	t.Run("by patient name", func(t *testing.T) {
		filters := query.DiagnosisFilters{PatientName: "patientNacho"}
		results, err := repo.Find(ctx, filters, query.Pagination{Page: 1, PageSize: 10})
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
	})

	t.Run("by date", func(t *testing.T) {
		date := time.Now()
		filters := query.DiagnosisFilters{Date: &date}
		results, err := repo.Find(ctx, filters, query.Pagination{Page: 1, PageSize: 10})
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
	})

	t.Run("by patient name and date", func(t *testing.T) {
		date := time.Now()
		filters := query.DiagnosisFilters{PatientName: "patientNacho", Date: &date}
		results, err := repo.Find(ctx, filters, query.Pagination{Page: 1, PageSize: 10})
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
	})
}

func insertTestPatient(t *testing.T, db *gorm.DB) uuid.UUID {
	patient := &patient.Entity{
		ID:    uuid.New(),
		Name:  "patientNacho",
		DNI:   "123123",
		Email: "nacho@topdoc.com",
	}
	err := db.Create(patient).Error
	assert.NoError(t, err)
	return patient.ID
}

func insertMultipleDiagnoses(t *testing.T, db *gorm.DB, patientID uuid.UUID, count int) {
	for i := 0; i < count; i++ {
		d := &diagnosisEntity.Entity{
			ID:        uuid.New(),
			PatientID: patientID,
			Diagnosis: "fiebre numero? " + strconv.Itoa(i),
			CreatedAt: time.Now().Add(time.Duration(-i) * time.Hour),
		}
		err := db.Create(d).Error
		assert.NoError(t, err)
	}
}
