package service

import (
	"context"

	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
)

type (
	UserService struct {
		Config env.Env
	}
)

var _ UserServiceInterface = (*UserService)(nil)

func NewUserService(config env.Env) *UserService {
	return &UserService{
		Config: config,
	}
}

type CreateUserInput struct {
	FirstName string
	LastName  string
}

func (u UserService) CreateUser(ctx context.Context,
	input CreateUserInput,
) (*models.User, error) {
	return nil, nil
}
