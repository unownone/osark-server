package server

import "github.com/gofiber/fiber/v2"

func (h *handler) setRoutes(route *fiber.App) {
	route.Get("/ping", ping)

	api := route.Group("/api")
	api.Get("/ping", ping) // health check/ ping route
	api.Post("/events", h.CaptureEvents)

	api.Get("/systems", h.GetSystems)
	api.Get("/systems/:id", h.GetSystem)

	api.Get("systems/:id/apps", h.GetAppsForASystem)

	frontendRouter := route.Group("/")
	h.frontend(&frontendRouter)
}

// @Summary Ping the server
// @Description Ping the server to check if it is running
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func (h *handler) frontend(route *fiber.Router) {

}
