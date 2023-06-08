package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/models"
)

type RepositoryInterface interface {
	CreateUser(ctx context.Context,
		user models.User,
	) (*models.User, error)
}
