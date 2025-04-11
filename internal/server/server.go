package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
)

// NewServer creates a new server
func (h *handler) NewServer() (*fiber.App, error) {
	// Initialize template engine
	engine := html.New(h.config.GetViewPath(), ".html")
	
	// Add template functions before loading templates
	engine.AddFunc("add", func(a, b int) int {
		return a + b
	})

	engine.AddFunc("sub", func(a, b int) int {
		if a-b < 0 {
			return 0
		}
		return a - b
	})
	
	engine.AddFunc("currentYear", func() int {
		return time.Now().Year()
	})
	
	// Load templates after adding functions
	if err := engine.Load(); err != nil {
		return nil, err
	}

	// Initialize the application
	config := fiber.Config{
		Views: engine,
	}

	app := fiber.New(config)
	app.Use(requestid.New())
	app.Use(logger.New())
	h.setRoutes(app) // set routes
	return app, nil
}

// ListenAndServe listens and serves the server
func (h *handler) ListenAndServe(server *fiber.App) error {
	serverAddress := ":" + h.config.GetPort()
	return server.Listen(serverAddress)
}
