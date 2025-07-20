package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	conn := "postgresql://diagnosis_svc:diagnosis_svc@localhost:5434/diagnosis_svc?sslmode=disable"
	db, err := sql.Open("pgx", conn)
	if err != nil {
		fmt.Printf("Error opening connection to DB: %v\n", err)
		return
	}
	defer db.Close()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Printf("Error starting transaction: %v\n", err)
		return
	}

	if err := insertUser(ctx, tx); err != nil {
		tx.Rollback()
		fmt.Println("error inserting user:", err)
		return
	}

	patientID, err := insertPatient(ctx, tx)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error inserting patient:", err)
		return
	}

	if err := insertDiagnoses(ctx, tx, patientID); err != nil {
		tx.Rollback()
		fmt.Println("Error inserting diagnoses:", err)
		return
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("Error committing transaction: %v\n", err)
		return
	}
}

func insertUser(ctx context.Context, tx *sql.Tx) error {
	nickname := "nacho"
	email := "nacho@gmail.com"
	rawPassword := "testing123123"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	query := `
	INSERT INTO top_doctor.user (id, nickname, password, email, created_at)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err = tx.ExecContext(ctx, query,
		uuid.New(),
		nickname,
		string(passwordHash),
		email,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func insertPatient(ctx context.Context, tx *sql.Tx) (uuid.UUID, error) {
	patientID := uuid.New()

	query := `
	INSERT INTO top_doctor.patient (id, name, dni, email, phone, address, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := tx.ExecContext(ctx, query,
		patientID,
		"nachin patient",
		"12345678",
		"nachinahcin@gmail.com",
		"123456789",
		"barcelona",
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return uuid.Nil, fmt.Errorf("error inserting patient: %w", err)
	}

	return patientID, nil
}

func insertDiagnoses(ctx context.Context, tx *sql.Tx, patientID uuid.UUID) error {
	diagnoses := []struct {
		Diagnosis    string
		Prescription *string
		Date         time.Time
	}{
		{"gripe 1", nil, time.Now()},
		{"gripe 2", ptr("antigripepuntero2"), time.Now().AddDate(0, 0, -1)},
		{"gripe 3", ptr("antigripepuntero3"), time.Now().AddDate(0, 0, -2)},
	}

	query := `
	INSERT INTO top_doctor.diagnosis (id, patient_id, diagnosis, prescription, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	for _, d := range diagnoses {
		_, err := tx.ExecContext(ctx, query,
			uuid.New(),
			patientID,
			d.Diagnosis,
			d.Prescription,
			d.Date,
			time.Now(),
		)
		if err != nil {
			return fmt.Errorf("error inserting diagnosis: %w", err)
		}
	}

	return nil
}

func ptr(s string) *string {
	return &s
}
