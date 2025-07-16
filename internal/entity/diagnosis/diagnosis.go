package diagnosis

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/nachoconques0/diagnosis_svc/internal/errors"
)

var (
	// ErrMissingDiagnosis used when diagnosis is missing
	ErrMissingDiagnosis = errors.NewWrongInput("diagnosis is required")
	// ErrInvalidPatientID used when patient ID is invalid
	ErrInvalidPatientID = errors.NewWrongInput("invalid patient ID")
	// ErrInvalidDiagnosis used when diagnosis is missing
	ErrInvalidDiagnosis = errors.NewWrongInput("diagnosis must be at least 10 characters")
)

// Entity hold fields of a diagnosis used for our internal use
type Entity struct {
	ID           int
	PatientID    uuid.UUID
	Diagnosis    string
	Prescription *string
	Date         time.Time
}

// New returns an instance of a Patient
func New(patientID string, diagnosis string, prescription *string) (*Entity, error) {
	id, err := uuid.Parse(patientID)
	if err != nil {
		return nil, ErrInvalidPatientID
	}
	return &Entity{
		PatientID:    id,
		Diagnosis:    strings.TrimSpace(diagnosis),
		Prescription: prescription,
		Date:         time.Now(),
	}, nil
}

// Valid validates Diagnosis entity fields
func (e *Entity) Valid() error {
	if e.Diagnosis == "" {
		return ErrMissingDiagnosis
	}
	if len(e.Diagnosis) < 10 {
		return ErrInvalidDiagnosis
	}

	return nil
}
