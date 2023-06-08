package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
)

type (
	UserRepo struct {
		Config env.Env
	}
)

var _ UserRepoInterface = (*UserRepo)(nil)

func NewUserRepo(config env.Env) *UserRepo {
	return &UserRepo{
		Config: config,
	}
}

type CreateUserInput struct {
	FirstName string
	LastName  string
}

func (u *UserRepo) CreateUser(ctx context.Context,
	user models.User,
) (*models.User, error) {
	return nil, nil
}
