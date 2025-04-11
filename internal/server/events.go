package server

import (
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
	sysInfo := new(models.SystemInfo)
	if err := c.BodyParser(sysInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.sysInfoRepository.Create(c.Context(), sysInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(sysInfo)
}
