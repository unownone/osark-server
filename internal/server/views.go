package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)


func (h *handler) frontend(route *fiber.Router) {
	r := *route
	// System list page
	r.Get("/", h.RenderSystemsPage)
	// System detail page
	r.Get("/systems/:id", h.RenderSystemDetailPage)
	// System apps page
	r.Get("/systems/:id/apps", h.RenderSystemAppsPage)
}

// RenderSystemsPage renders the page showing all systems
func (h *handler) RenderSystemsPage(c *fiber.Ctx) error {
	// Default pagination values
	limit := 20
	offset := 0

	if c.Query("limit") != "" {
		fmt.Sscanf(c.Query("limit"), "%d", &limit)
	}

	if c.Query("offset") != "" {
		fmt.Sscanf(c.Query("offset"), "%d", &offset)
	}

	systems, err := h.data.GetAllSystems(c.Context(), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching systems")
	}

	return c.Render("systems", fiber.Map{
		"Title":   "Systems Overview",
		"Systems": systems,
		"Limit":   limit,
		"Offset":  offset,
	})
}

// RenderSystemDetailPage renders the page showing a single system details
func (h *handler) RenderSystemDetailPage(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("System ID is required")
	}

	system, err := h.data.GetSystem(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching system details")
	}

	return c.Render("system_detail", fiber.Map{
		"Title":  "System Details",
		"System": system,
	})
}

// RenderSystemAppsPage renders the page showing apps for a system
func (h *handler) RenderSystemAppsPage(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("System ID is required")
	}

	// Pagination params (default values)
	limit := 20
	offset := 0

	if c.Query("limit") != "" {
		fmt.Sscanf(c.Query("limit"), "%d", &limit)
	}

	if c.Query("offset") != "" {
		fmt.Sscanf(c.Query("offset"), "%d", &offset)
	}

	apps, err := h.data.GetPaginatedAppsForASystem(c.Context(), id, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching apps")
	}

	system, err := h.data.GetSystem(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching system details")
	}

	return c.Render("system_apps", fiber.Map{
		"Title":  "System Apps",
		"System": system,
		"Apps":   apps,
		"Limit":  limit,
		"Offset": offset,
	})
}
