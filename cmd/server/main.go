package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/nachoconques0/diagnosis_svc/internal/app"
	"github.com/nachoconques0/diagnosis_svc/internal/env"
)

func main() {
	maxDBConn, err := strconv.Atoi(env.LoadOrPanic("DB_MAX_CONNECTIONS"))
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("could not start application: %s", err.Error()))
	}

	options := []app.Option{
		app.WithHTTPPort(env.LoadOrPanic("HTTP_PORT")),
		app.WithDBHost(env.LoadOrPanic("DB_HOST")),
		app.WithDBPort(env.LoadOrPanic("DB_PORT")),
		app.WithDBUser(env.LoadOrPanic("DB_USER")),
		app.WithDBPassword(env.LoadOrPanic("DB_PASSWORD")),
		app.WithDBName(env.LoadOrPanic("DB_NAME")),
		app.WithDBMaxConnections(maxDBConn),
		app.WithSSLMode(env.LoadOrDefault("DB_SSL", "disable")),
		app.WithJWTSecret(env.LoadOrPanic("JWT_SECRET")),
	}

	err = app.New(options...)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("could not start application: %s", err.Error()))
	}
}
