package main

import (
	"github.com/unownone/osark-server/internal/config"
	"github.com/unownone/osark-server/internal/models"
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
	// TODO: improve this , automigrate is not a good idea in production
	if err := migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

// migrate migrates the database
func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.AppInfo{},
		&models.ProcessInfo{},
		&models.SystemInfo{},
	)
}
