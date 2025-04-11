package service

import (
	"context"
	"fmt"

	"github.com/unownone/osark-server/internal/models"
	"github.com/unownone/osark-server/internal/repository"
)

type Event interface {
	Handle(ctx context.Context, events []*models.LogEvent, identifier string) error
}

type eventService struct {
	sysRepository     repository.SysInfoRepository
	appRepository     repository.AppRepository
	processRepository repository.ProcessRepository
}

// NewEventService creates a new event service
func NewEventService(sysRepository repository.SysInfoRepository, appRepository repository.AppRepository, processRepository repository.ProcessRepository) Event {
	return &eventService{sysRepository: sysRepository, appRepository: appRepository, processRepository: processRepository}
}

// Handle handles an event
func (s *eventService) Handle(ctx context.Context, events []*models.LogEvent, identifier string) error {
	errors := make([]string, 0, len(events))
	sysInfo := make([]*models.SystemInfo, 0, len(events))
	appInfo := make([]*models.AppInfo, 0, len(events))
	processInfo := make([]*models.ProcessInfo, 0, len(events))
	for _, event := range events {
		if event.SystemInfo != nil {
			event.SystemInfo.ID = identifier
			sysInfo = append(sysInfo, event.SystemInfo)
		}
		if event.Error != "" {
			errors = append(errors, event.Error)
		}
		if event.AppInfo != nil {
			for _, app := range event.AppInfo {
				app.SysInfo = &models.SystemInfo{ID: identifier}
				app.ID = &app.BundleID // Use bundle ID as the ID TODO: better ids sometimes
				appInfo = append(appInfo, app)
			}
		}
		if event.Processes != nil {
			for _, process := range event.Processes {
				process.SysInfo = &models.SystemInfo{ID: identifier}
				processInfo = append(processInfo, process)
			}
		}
	}
	fmt.Println("errors", len(errors), "\t\terrors", errors)
	fmt.Println("sysInfo", len(sysInfo))
	fmt.Println("appInfo", len(appInfo))
	fmt.Println("processInfo", len(processInfo))
	if len(sysInfo) > 0 {
		if err := s.sysRepository.Create(ctx, sysInfo...); err != nil {
			return err
		}
	}
	if len(appInfo) > 0 {
		// Deduplicate app entries by ID before saving
		dedupedApps := make(map[string]*models.AppInfo)
		for _, app := range appInfo {
			// Use the ID as key to ensure only one entry per ID
			dedupedApps[*app.ID] = app
		}

		// Convert back to slice
		uniqueApps := make([]*models.AppInfo, 0, len(dedupedApps))
		for _, app := range dedupedApps {
			uniqueApps = append(uniqueApps, app)
		}

		// Now save the deduplicated data
		if err := s.appRepository.CreateOrUpdate(ctx, uniqueApps...); err != nil {
			return err
		}
	}
	if len(processInfo) > 0 {
		if err := s.processRepository.Create(ctx, processInfo...); err != nil {
			return err
		}
	}
	return nil
}
