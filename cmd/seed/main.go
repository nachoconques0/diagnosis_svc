package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	conn := "postgres://postgres:postgres@localhost:5432/diagnosisdb?sslmode=disable"
	db, err := sql.Open("pgx", conn)
	if err != nil {
		fmt.Errorf("error opening connection to DB")
	}
	defer db.Close()

	nickname := "nacho"
	email := "nacho@gmail.com"
	rawPassword := "testing123123"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Errorf("error hashing password: %v", err)
	}

	query := `
	INSERT INTO top_doctor.user (id, nickname, password, email, created_at)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err = db.ExecContext(context.Background(), query,
		uuid.New(),
		nickname,
		string(passwordHash),
		email,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		fmt.Errorf("error inserting user: %v", err)
	}
	fmt.Println("user inserted successfully")
}
