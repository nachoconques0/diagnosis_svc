package model

import (
	"time"
)

// CreatePatientRequest type that holds fields when creating a diagnosis
type CreateDiagnosisRequest struct {
	PatientID    string  `json:"patient_id" binding:"required"`
	Diagnosis    string  `json:"diagnosis" binding:"required"`
	Prescription *string `json:"prescription,omitempty"`
}

// PatientResponse type that holds fields when creating a response for diagnosis
type DiagnosisResponse struct {
	ID           string    `json:"id,omitempty"`
	PatientID    string    `json:"patient_id,omitempty"`
	Diagnosis    string    `json:"diagnosis,omitempty"`
	Prescription *string   `json:"prescription,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
