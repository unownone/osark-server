package repository

import (
	"context"

	"github.com/unownone/osark-server/internal/models"
	"gorm.io/gorm"
)

type AppRepository interface {
	Create(ctx context.Context, apps ...*models.AppInfo) error
}

type appRepository struct {
	db        *gorm.DB
	batchSize int
}

func NewAppRepository(db *gorm.DB, batchSize int) AppRepository {
	return &appRepository{db: db, batchSize: batchSize}
}

func (r *appRepository) Create(ctx context.Context, apps ...*models.AppInfo) error {
	return r.db.CreateInBatches(apps, r.batchSize).Error
}
