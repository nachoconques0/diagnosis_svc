package diagnosis_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nachoconques0/diagnosis_svc/internal/entity/diagnosis"
	"github.com/stretchr/testify/assert"
)

func TestNewDiagnosis(t *testing.T) {
	validUUID := uuid.New().String()
	invalidUUID := "invalidID"
	prescription := "paracetamollol"
	crazyDiagnosis := "el loco este tiene dolor de cabeza de tanto buscar trabajo XD"

	tests := []struct {
		name         string
		patientID    string
		diagnosis    string
		prescription *string
		expectedErr  error
	}{
		{
			name:         "success diagnosis",
			patientID:    validUUID,
			diagnosis:    crazyDiagnosis,
			prescription: &prescription,
			expectedErr:  nil,
		},
		{
			name:         "invalid patient id",
			patientID:    invalidUUID,
			diagnosis:    crazyDiagnosis,
			prescription: &prescription,
			expectedErr:  diagnosis.ErrInvalidPatientID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := diagnosis.New(tt.patientID, tt.diagnosis, tt.prescription)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)

			}
		})
	}
}

func TestDiagnosis_Valid(t *testing.T) {
	validUUID := uuid.New().String()
	crazyDiagnosis := "el loco este tiene dolor de cabeza de tanto buscar trabajo XD"

	tests := []struct {
		name        string
		diagnosis   string
		expectedErr error
	}{
		{
			name:        "success diagnosis valid",
			diagnosis:   crazyDiagnosis,
			expectedErr: nil,
		},
		{
			name:        "missing diagnosis",
			diagnosis:   "",
			expectedErr: diagnosis.ErrMissingDiagnosis,
		},
		{
			name:        "invalid diagnosis",
			diagnosis:   "undolor",
			expectedErr: diagnosis.ErrInvalidDiagnosis,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := diagnosis.New(validUUID, tt.diagnosis, nil)
			if err != nil {
				t.Fatalf("unexpected constructor error: %v", err)
			}

			err = d.Valid()
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)

			}
		})
	}
}
