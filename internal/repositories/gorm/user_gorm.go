package gorm

import (
	"context"
	"imob/internal/entities"
	"imob/internal/logger"

	"gorm.io/gorm"
)

type UserGorm struct {
	db  *gorm.DB
	log logger.Logger
}

func NewUserGorm(db *gorm.DB, log logger.Logger) *UserGorm { return &UserGorm{db: db, log: log} }

func (r *UserGorm) Create(ctx context.Context, u *entities.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *UserGorm) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var u entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
