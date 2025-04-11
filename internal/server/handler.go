package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unownone/osark-server/internal/config"
)

type Handler interface {
	NewServer() (*fiber.App, error)
	ListenAndServe(*fiber.App) error
}

type handler struct {
	config *config.Config
}

func NewHandler(config *config.Config) Handler {
	return &handler{config: config}
}

