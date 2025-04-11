package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	dbConfig *DBConfig
	port     string `env:"PORT"`
}

type DBConfig struct {
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     int    `env:"DB_PORT"`
	DB_USER     string `env:"DB_USER"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_NAME     string `env:"DB_NAME"`
	DB_TYPE     string `env:"DB_TYPE" envDefault:"postgres"`
}

// GetDSN returns the DSN for the database
func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASSWORD, c.DB_NAME)
}

// NewConfig creates a new Config
func NewConfig() (*Config, error) {
	var dbcfg DBConfig
	if err := env.Parse(&dbcfg); err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	cfg.dbConfig = &dbcfg
	return &cfg, nil
}

// GetDBConfig returns the DBConfig
func (c *Config) GetDBConfig() *DBConfig {
	return c.dbConfig
}

// GetPort returns the PORT
func (c *Config) GetPort() string {
	return c.port
}
