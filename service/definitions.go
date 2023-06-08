package service

import (
	"context"
	"github.com/tejiriaustin/apex-network/models"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context,
		input CreateUserInput,
	) (*models.User, error)
}
