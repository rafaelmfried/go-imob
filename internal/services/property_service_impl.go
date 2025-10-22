package services

import (
	"context"
	"imob/internal/entities"
	"imob/internal/logger"
	"imob/internal/repositories"
)

type propertyService struct {
	repo repositories.PropertyRepository
	log  logger.Logger
}

func NewPropertyService(r repositories.PropertyRepository, log logger.Logger) PropertyService {
	return &propertyService{repo: r, log: log}
}

func (p *propertyService) Create(ctx context.Context, in entities.Property) (entities.Property, error) {
	if err := p.repo.Create(ctx, &in); err != nil {
		return entities.Property{}, err
	}
	return in, nil
}

func (p *propertyService) Get(ctx context.Context, id string) (entities.Property, error) {
	out, err := p.repo.GetByOwnerID(ctx, id)
	if err != nil {
		return entities.Property{}, err
	}
	if len(out) == 0 {
		return entities.Property{}, nil
	}
	return out[0], nil
}
