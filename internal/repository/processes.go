package repository

import (
	"context"

	"github.com/unownone/osark-server/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProcessRepository interface {
	Create(ctx context.Context, processes ...*models.ProcessInfo) error
}

type processRepository struct {
	db        *gorm.DB
	batchSize int
}

func NewProcessRepository(db *gorm.DB, batchSize int) ProcessRepository {
	return &processRepository{db: db, batchSize: batchSize}
}

func (r *processRepository) Create(ctx context.Context, processes ...*models.ProcessInfo) error {
	return r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
		DoNothing: false,
		Columns:   []clause.Column{},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"created_at": gorm.Expr("processes.created_at"),
		}),
	}).CreateInBatches(processes, r.batchSize).Error
}
