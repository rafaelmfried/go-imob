package gorm

import (
	"context"
	"imob/internal/entities"
	"imob/internal/logger"

	"gorm.io/gorm"
)

type PropertyGorm struct {
	db  *gorm.DB
	log logger.Logger
}

func NewPropertyGorm(db *gorm.DB, log logger.Logger) *PropertyGorm { return &PropertyGorm{db: db, log: log} }

func (r *PropertyGorm) Create(ctx context.Context, p *entities.Property) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *PropertyGorm) GetByOwnerID(ctx context.Context, ownerID string) ([]entities.Property, error) {
	var properties []entities.Property
	if err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&properties).Error; err != nil {
		return nil, err
	}
	return properties, nil
}
