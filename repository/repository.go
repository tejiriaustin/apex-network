package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
)

type (
	Repository struct {
		config env.Env
		DB     *database.Client
	}
)

func NewRepo(env env.Env, db *database.Client) RepositoryInterface {
	return &Repository{
		config: env,
		DB:     db,
	}
}

var _ RepositoryInterface = (*Repository)(nil)

func (u *Repository) CreateUser(ctx context.Context,
	user models.User,
) (*models.User, error) {
	return nil, nil
}
