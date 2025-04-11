package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unownone/osark-server/internal/config"
	"github.com/unownone/osark-server/internal/repository"
	"github.com/unownone/osark-server/internal/service"
	"gorm.io/gorm"
)

type Handler interface {
	NewServer() (*fiber.App, error)
	ListenAndServe(*fiber.App) error
	CaptureEvents(*fiber.Ctx) error
}

type handler struct {
	config            *config.Config
	db                *gorm.DB
	eventService      service.Event
}

// NewHandler creates a new handler
func NewHandler(config *config.Config, db *gorm.DB) Handler {
	sysRepository := repository.NewSysInfoRepository(db, 100)
	appRepository := repository.NewAppRepository(db, 100)
	processRepository := repository.NewProcessRepository(db, 100)
	eventService := service.NewEventService(sysRepository, appRepository, processRepository)
	return &handler{config: config, db: db, eventService: eventService}
}
