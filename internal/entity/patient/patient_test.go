package patient_test

import (
	"testing"

	"github.com/nachoconques0/diagnosis_svc/internal/entity/patient"
	"github.com/stretchr/testify/assert"
)

func TestPatient_Valid(t *testing.T) {
	validPhone := "+34673512331"
	invalidPhone := "123"
	address := "carrer de roselloooooo"
	tests := []struct {
		name        string
		patient     *patient.Entity
		expectedErr error
	}{
		{
			name:        "new patient success",
			patient:     patient.New("nachoTopDoctor", "1231231", "nachio@testing.com", &validPhone, &address),
			expectedErr: nil,
		},
		{
			name:        "missing name",
			patient:     patient.New("", "111111", "nachio@testing.com", nil, nil),
			expectedErr: patient.ErrMissingName,
		},
		{
			name:        "invalid name",
			patient:     patient.New("N1231", "X1234567", "jose@example.com", nil, nil),
			expectedErr: patient.ErrInvalidName,
		},
		{
			name:        "missing DNI",
			patient:     patient.New("nachoTopDoctor", "", "nachio@testing.com", nil, nil),
			expectedErr: patient.ErrMissingDNI,
		},
		{
			name:        "email not valid",
			patient:     patient.New("nachoTopDoctor", "1231231", "-", nil, nil),
			expectedErr: patient.ErrInvalidEmail,
		},
		{
			name:        "phone not valid",
			patient:     patient.New("nachoTopDoctor", "1231231", "nachio@testing.com", &invalidPhone, nil),
			expectedErr: patient.ErrInvalidPhone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.patient.Valid()
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}
