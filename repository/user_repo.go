package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/models"
)

type UserRepository struct {
	db Repository
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	UpdateUser(ctx context.Context, filter QueryFilter, column string, value interface{}) (*models.User, error)
	GetUser(ctx context.Context, filter QueryFilter) (*models.User, error)
}

func (u *UserRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	var newUser models.User
	if err := u.db.WithContext(ctx).
		Create(user).
		Find(&newUser, "id = ?", user.ID.String()).
		Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, filter QueryFilter, column string, value interface{}) (*models.User, error) {
	var user models.User

	if err := u.db.WithContext(ctx).
		Where(filter.GetFilter()).
		Update(column, value).
		Find(&user, filter.GetFilter()).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUser(ctx context.Context, filter QueryFilter) (*models.User, error) {
	var user models.User

	if err := u.db.WithContext(ctx).
		Where(filter.GetFilter()).
		Find(&user, filter.GetFilter()).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}