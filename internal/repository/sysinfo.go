package repository

import (
	"context"

	"github.com/unownone/osark-server/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SysInfoRepository interface {
	Create(ctx context.Context, sysInfo ...*models.SystemInfo) error
	PaginatedGetAllSystems(ctx context.Context, limit int, offset int) ([]*models.SystemInfo, error)
	GetSystem(ctx context.Context, id string) (*models.SystemInfo, error)
}

type sysInfoRepository struct {
	db        *gorm.DB
	batchSize int
}

// NewSysInfoRepository creates a new system info repository
func NewSysInfoRepository(db *gorm.DB, batchSize int) SysInfoRepository {
	return &sysInfoRepository{db: db, batchSize: batchSize}
}

// Create creates a new system info
func (r *sysInfoRepository) Create(ctx context.Context, sysInfo ...*models.SystemInfo) error {
	return r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
		DoNothing: false,
		Columns:   []clause.Column{},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"created_at": gorm.Expr("system_infos.created_at"),
		}),
	}).CreateInBatches(sysInfo, r.batchSize).Error
}

// PaginatedGetAllSystems gets all systems
func (r *sysInfoRepository) PaginatedGetAllSystems(ctx context.Context, limit int, offset int) ([]*models.SystemInfo, error) {
	var sysInfos []*models.SystemInfo
	if err := r.db.Limit(limit).Offset(offset).Find(&sysInfos).Error; err != nil {
		return nil, err
	}
	return sysInfos, nil
}

// GetSystem gets a system by id
func (r *sysInfoRepository) GetSystem(ctx context.Context, id string) (*models.SystemInfo, error) {
	var sysInfo models.SystemInfo
	if err := r.db.Where("id = ?", id).First(&sysInfo).Error; err != nil {
		return nil, err
	}
	return &sysInfo, nil
}
