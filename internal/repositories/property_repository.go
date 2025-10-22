package repositories

import (
	"context"
	"imob/internal/entities"
)

type PropertyRepository interface {
	Create(ctx context.Context, property *entities.Property) error
	GetByOwnerID(ctx context.Context, ownerID string) ([]entities.Property, error)
}