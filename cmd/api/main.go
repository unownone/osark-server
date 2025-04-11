package main

import (
	"fmt"
	"os"

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
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("Failed to create config: %w", err)
		os.Exit(1)
	}
	handler := server.NewHandler(cfg)

	if err := runServer(&handler); err != nil {
		fmt.Println("Failed to run server: %w", err)
		os.Exit(1)
	}
}

// runServer runs the server
func runServer(handler *server.Handler) error {
	server, err := (*handler).NewServer()
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	return (*handler).ListenAndServe(server)
}
