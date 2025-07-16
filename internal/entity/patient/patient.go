package patient

import (
	"regexp"
	"strings"

	"github.com/google/uuid"

	"github.com/nachoconques0/diagnosis_svc/internal/errors"
)

var (
	// ErrMissingName used when name is missing
	ErrMissingName = errors.NewWrongInput("name is required")
	// ErrInvalidName used when name is invalid
	ErrInvalidName = errors.NewWrongInput("name must contain only letters")
	// ErrMissingDNI used when DNI is missing
	ErrMissingDNI = errors.NewWrongInput("DNI is required")
	// ErrInvalidEmail used when email has invalid format
	ErrInvalidEmail = errors.NewWrongInput("invalid email format")
	// ErrInvalidPhone used when phone has invalid format
	ErrInvalidPhone = errors.NewWrongInput("invalid phone format")
	nameRegex       = `^[a-zA-Z]+$`
	emailRegex      = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	phoneRegex      = `^\+?[0-9]{7,15}$`
)

// Entity holds the fields of a patient used for our internal use
type Entity struct {
	ID      uuid.UUID
	Name    string
	DNI     string
	Email   string
	Phone   *string
	Address *string
}

// New returns an instance of a Patient
func New(name, dni, email string, phone, address *string) *Entity {
	return &Entity{
		Name:    strings.TrimSpace(name),
		DNI:     strings.TrimSpace(dni),
		Email:   strings.TrimSpace(email),
		Phone:   phone,
		Address: address,
	}
}

// Valid validates Patient entity fields
func (e *Entity) Valid() error {
	if e.Name == "" {
		return ErrMissingName
	}
	if !regexp.MustCompile(nameRegex).MatchString(e.Name) {
		return ErrInvalidName
	}
	if e.DNI == "" {
		return ErrMissingDNI
	}
	if !regexp.MustCompile(emailRegex).MatchString(e.Email) {
		return ErrInvalidEmail
	}
	if e.Phone != nil {
		if !regexp.MustCompile(phoneRegex).MatchString(*e.Phone) {
			return ErrInvalidPhone
		}
	}
	return nil
}
