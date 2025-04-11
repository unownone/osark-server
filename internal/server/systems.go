package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unownone/osark-server/internal/models"
)

// @Summary Get all systems
// @Description Get all systems
// @Accept json
// @Produce json
// @Param limit query int false "limit of systems to return" default(10)
// @Param offset query int false "offset of systems to return" default(0)
// @Success 200 {object} models.GetPaginatedResponse[*models.SystemInfo]
func (h *handler) GetSystems(c *fiber.Ctx) error {
	queryParams := models.QueryParams{}
	if err := c.QueryParser(&queryParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := queryParams.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	systems, err := h.data.GetAllSystems(c.Context(), queryParams.Limit, queryParams.GetOffset())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(models.GetPaginatedResponse[*models.SystemInfo]{
		Total:  len(systems),
		Limit:  queryParams.Limit,
		Offset: queryParams.GetOffset(),
		Data:   systems,
	})
}

// @Summary Get a system by id
// @Description Get a system by id
// @Accept json
// @Produce json
// @Param id path string true "id of the system to get"
// @Success 200 {object} models.SystemInfo
func (h *handler) GetSystem(c *fiber.Ctx) error {
	id := c.Params("id")
	system, err := h.data.GetSystem(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(system)
}

// @Summary Get apps for a system
// @Description Get apps for a system
// @Accept json
// @Produce json
// @Param id path string true "id of the system to get apps for"
// @Param limit query int false "limit of apps to return" default(10)
// @Param offset query int false "offset of apps to return" default(0)
// @Success 200 {object} models.GetPaginatedResponse[*models.AppInfo]
func (h *handler) GetAppsForASystem(c *fiber.Ctx) error {
	id := c.Params("id")
	queryParams := models.QueryParams{}
	if err := c.QueryParser(&queryParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := queryParams.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	apps, err := h.data.GetPaginatedAppsForASystem(c.Context(), id, queryParams.Limit, queryParams.GetOffset())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(models.GetPaginatedResponse[*models.AppInfo]{
		Total:  len(apps),
		Limit:  queryParams.Limit,
		Offset: queryParams.GetOffset(),
		Data:   apps,
	})
}
