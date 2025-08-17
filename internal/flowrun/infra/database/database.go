package database

import (
	"context"
	"flow-run/internal/core/domain"
	"flow-run/internal/lib/validator"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	URL             string        `validate:"required,url"`
	MaxOpenConns    int           `validate:"required,min=1,max=1000"`
	MaxIdleConns    int           `validate:"required,min=1,max=100"`
	ConnMaxLifetime time.Duration `validate:"required,min=1m,max=1h"`
}

type Database struct {
	*gorm.DB
}

func NewDatabase(config *DatabaseConfig) (*Database, error) {
	// Validate config before using it
	validatedConfig, err := validator.Struct(config)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(validatedConfig.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(validatedConfig.MaxOpenConns)
	sqlDB.SetMaxIdleConns(validatedConfig.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(validatedConfig.ConnMaxLifetime)

	db.AutoMigrate(&domain.Provider{})

	return &Database{DB: db}, nil
}

func (d *Database) Stop(ctx context.Context) error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
