package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unownone/osark-server/internal/config"
	"github.com/unownone/osark-server/internal/repository"
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
	sysInfoRepository repository.SysInfoRepository
}

func NewHandler(config *config.Config, db *gorm.DB) Handler {
	sysInfoRepository := repository.NewSysInfoRepository(db, 100)
	return &handler{config: config, db: db, sysInfoRepository: sysInfoRepository}
}
