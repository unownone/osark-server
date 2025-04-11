package repository

import (
	"context"

	"github.com/unownone/osark-server/internal/models"
	"gorm.io/gorm"
)

type SysInfoRepository interface {
	Create(ctx context.Context, sysInfo ...*models.SystemInfo) error
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
	return r.db.CreateInBatches(sysInfo, r.batchSize).Error
}
