package db

import (
	"fmt"
	"time"

	"github.com/nachoconques0/diagnosis_svc/internal/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New returns a new instance of DB
func New(opts ...Option) (*gorm.DB, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(&options)
	}

	conn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		options.User,
		options.Password,
		options.Host,
		options.Port,
		options.Database,
		options.SSLMode,
	)

	logLevel := logger.Silent
	if options.Debug {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, errors.NewInternalError("error opening connection to DB")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.NewInternalError("error retrieving sql.DB from gorm.DB")
	}

	if options.MaxConnections > 0 {
		sqlDB.SetMaxOpenConns(options.MaxConnections)
		sqlDB.SetMaxIdleConns(options.MaxConnections / 2)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, errors.NewInternalError("error verifying connection to DB")
	}

	return db, nil
}
