package services

import (
	"context"
	"imob/internal/entities"
)

type PropertyService interface {
	Create(ctx context.Context, p entities.Property) (entities.Property, error)
	Get(ctx context.Context, id string) (entities.Property, error)
}
