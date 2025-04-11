package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// NewServer creates a new server
func (h *handler) NewServer() (*fiber.App, error) {
	app := fiber.New()
	app.Use(requestid.New())
	h.setRoutes(app) // set routes
	return app, nil
}

// ListenAndServe listens and serves the server
func (h *handler) ListenAndServe(server *fiber.App) error {
	serverAddress := ":" + h.config.GetPort()
	return server.Listen(serverAddress)
}
