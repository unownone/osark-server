package repository

import (
	"context"

	"github.com/unownone/osark-server/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AppRepository interface {
	Create(ctx context.Context, apps ...*models.AppInfo) error
	CreateOrUpdate(ctx context.Context, apps ...*models.AppInfo) error
	GetPaginatedAppsForASystem(ctx context.Context, systemID string, limit int, offset int) ([]*models.AppInfo, error)
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

func (r *appRepository) CreateOrUpdate(ctx context.Context, apps ...*models.AppInfo) error {
	return r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
		DoNothing: false,
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"created_at": gorm.Expr("app_infos.created_at"),
		}),
	}).CreateInBatches(apps, r.batchSize).Error
}

func (r *appRepository) GetPaginatedAppsForASystem(ctx context.Context, systemID string, limit int, offset int) ([]*models.AppInfo, error) {
	var apps []*models.AppInfo
	if err := r.db.Limit(limit).Offset(offset).Where("sys_info_id = ?", systemID).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}
