package main

import (
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/nachoconques0/diagnosis_svc/internal/app"
	"github.com/nachoconques0/diagnosis_svc/internal/env"
)

func main() {
	maxDBConn, err := strconv.Atoi(env.LoadOrPanic("DB_MAX_CONNECTIONS"))
	if err != nil {
		log.Fatal().Err(err).Msg("could not start application when connecting to DB")
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
		app.WithDBDebug(false),
	}

	err = app.New(options...)
	if err != nil {
		log.Fatal().Err(err).Msg("could not start application when init app")
	}
}
