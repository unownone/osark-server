package main

import (
	"github.com/unownone/osark-server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initDB initializes the database
func initDB(dbConfig *config.DBConfig) (*gorm.DB, error) {
	postgresConfig := postgres.New(postgres.Config{
		DSN:                  dbConfig.GetDSN(),
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(postgresConfig, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
