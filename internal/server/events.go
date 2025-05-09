package server

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/unownone/osark-server/internal/models"
)

// @Summary Captures events and saves them into the database
// @Description Captures events and saves them into the database
// @Accept json
// @Produce json
// @Param
// @Param events body []models.Event true "Events"
// @Success 201 {object} models.Response
// @Router /api/events [post]
func (h *handler) CaptureEvents(c *fiber.Ctx) error {
	var events []*models.LogEvent
	if err := c.BodyParser(&events); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	identifier := c.Get("X-Identifier")
	if identifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "X-Identifier header is required"})
	}

	go func() {
		if err := h.eventService.Handle(c.Context(), events, identifier); err != nil {
			slog.Error("Error capturing events", "error", err)
		}
	}()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Events captured successfully"})
}
