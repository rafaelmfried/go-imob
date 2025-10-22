package services

import (
	"context"
	"errors"
	"imob/internal/entities"
	"imob/internal/logger"
	"imob/internal/repositories"
)

type userService struct {
	repo repositories.UserRepository
	log  logger.Logger
}

func NewUserService(r repositories.UserRepository, log logger.Logger) UserService {
	return &userService{repo: r, log: log}
}

func (u *userService) CreateUser(ctx context.Context, in entities.User) (entities.User, error) {
	if e, _ := u.repo.GetByEmail(ctx, in.Email); e != nil {
		return entities.User{}, errors.New("email already exists")
	}
	if err := u.repo.Create(ctx, &in); err != nil {
		return entities.User{}, err
	}
	return in, nil
}

func (u *userService) GetUser(ctx context.Context, email string) (entities.User, error) {
	out, err := u.repo.GetByEmail(ctx, email)
	if err != nil { return entities.User{}, err }
	return *out, nil
}
