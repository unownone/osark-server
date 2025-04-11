package service

import (
	"context"

	"github.com/unownone/osark-server/internal/models"
	"github.com/unownone/osark-server/internal/repository"
)

type Event interface {
	Handle(ctx context.Context, events []*models.LogEvent) error
}

type eventService struct {
	sysRepository repository.SysInfoRepository
	appRepository repository.AppRepository
}

// NewEventService creates a new event service
func NewEventService(sysRepository repository.SysInfoRepository, appRepository repository.AppRepository) Event {
	return &eventService{sysRepository: sysRepository, appRepository: appRepository}
}

// Handle handles an event
func (s *eventService) Handle(ctx context.Context, events []*models.LogEvent) error {
	sysInfo := make([]*models.SystemInfo, 0, len(events))
	appInfo := make([]*models.AppInfo, 0, len(events))
	for _, event := range events {
		sysInfo = append(sysInfo, event.SystemInfo)
		for _, app := range event.AppInfo {
			app.SysInfo = event.SystemInfo
			appInfo = append(appInfo, app)
		}
	}
	if len(sysInfo) > 0 {
		if err := s.sysRepository.Create(ctx, sysInfo...); err != nil {
			return err
		}
	}
	if len(appInfo) > 0 {
		if err := s.appRepository.Create(ctx, appInfo...); err != nil {
			return err
		}
	}
	return nil
}