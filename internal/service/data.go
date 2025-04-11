package service

import (
	"context"

	"github.com/unownone/osark-server/internal/models"
	"github.com/unownone/osark-server/internal/repository"
)

type Data interface {
	GetAllSystems(ctx context.Context, limit int, offset int) ([]*models.SystemInfo, error)
	GetSystem(ctx context.Context, id string) (*models.SystemInfo, error)
	GetPaginatedAppsForASystem(ctx context.Context, systemID string, limit int, offset int) ([]*models.AppInfo, error)
}

type data struct {
	sysInfoRepository repository.SysInfoRepository
	appRepository     repository.AppRepository
	processRepository repository.ProcessRepository
}

func NewData(sysInfoRepository repository.SysInfoRepository, appRepository repository.AppRepository, processRepository repository.ProcessRepository) Data {
	return &data{sysInfoRepository: sysInfoRepository, appRepository: appRepository, processRepository: processRepository}
}

// GetAllSystems gets all systems
func (d *data) GetAllSystems(ctx context.Context, limit int, offset int) ([]*models.SystemInfo, error) {
	return d.sysInfoRepository.PaginatedGetAllSystems(ctx, limit, offset)
}

// GetSystem gets a system by id
func (d *data) GetSystem(ctx context.Context, id string) (*models.SystemInfo, error) {
	return d.sysInfoRepository.GetSystem(ctx, id)
}

// GetAppsForASystem gets all apps for a system
func (d *data) GetPaginatedAppsForASystem(ctx context.Context, systemID string, limit int, offset int) ([]*models.AppInfo, error) {
	return d.appRepository.GetPaginatedAppsForASystem(ctx, systemID, limit, offset)
}
