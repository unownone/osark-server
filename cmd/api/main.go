package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/unownone/osark-server/internal/config"
	"github.com/unownone/osark-server/internal/server"
)

// @title OSARK API
// @version 1.0
// @description OSARK API to write and read data to the database
// @contact.name API Support
// @contact.email imon@ikr.one
// @host localhost:3000
// @BasePath /
func main() {
	// Setup structured logging
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	if err := godotenv.Load("../../.env"); err != nil {
		slog.Info("No .env file found, using system env")
	}
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("Failed to create config", "error", err)
		os.Exit(1)
	}
	db, err := initDB(cfg.GetDBConfig())
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	handler := server.NewHandler(cfg, db)

	if err := runServer(&handler); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}

// runServer runs the server
func runServer(handler *server.Handler) error {
	server, err := (*handler).NewServer()
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	
	// Log server start
	slog.Info("Starting server", "port", (*handler).GetConfig().GetPort())
	
	return (*handler).ListenAndServe(server)
}
