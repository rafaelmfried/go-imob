package services

import (
	"context"
	"imob/internal/entities"
)

type UserService interface {
	CreateUser(ctx context.Context, u entities.User) (entities.User, error)
	GetUser(ctx context.Context, email string) (entities.User, error)
}
