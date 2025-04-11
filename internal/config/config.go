package config

import "github.com/caarlos0/env/v11"

type Config struct {
	DBConfig *DBConfig
	PORT     string `env:"PORT"`
}

type DBConfig struct {
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     int    `env:"DB_PORT"`
	DB_USER     string `env:"DB_USER"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_NAME     string `env:"DB_NAME"`
	DB_TYPE     string `env:"DB_TYPE" envDefault:"postgres"`
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

	cfg.DBConfig = &dbcfg
	return &cfg, nil
}

// GetDBConfig returns the DBConfig
func (c *Config) GetDBConfig() *DBConfig {
	return c.DBConfig
}

// GetPort returns the PORT
func (c *Config) GetPort() string {
	return c.PORT
}
