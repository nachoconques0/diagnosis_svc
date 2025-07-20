package model

// CreatePatientRequest type that holds fields when creating a patient
type CreatePatientRequest struct {
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required,email"`
	DNI     string  `json:"dni" binding:"required"`
	Phone   *string `json:"phone,omitempty"`
	Address *string `json:"address,omitempty"`
}

// PatientResponse type that holds fields when creating a response for a patient
type PatientResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	DNI     string  `json:"dni"`
	Phone   *string `json:"phone,omitempty"`
	Address *string `json:"address,omitempty"`
}
