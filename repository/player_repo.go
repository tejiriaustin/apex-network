package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
)

type PlayerRepository struct {
	db Repository
}

type PlayerRepositoryInterface interface {
	CreateUser(ctx context.Context, user models.Player) (*models.Player, error)
	UpdateUser(ctx context.Context, filter QueryFilter, value interface{}) (*models.Player, error)
	GetUser(ctx context.Context, filter QueryFilter) (*models.Player, error)
}

func NewPlayerRepository(config env.Env, dbClient *database.Client) PlayerRepositoryInterface {
	return &PlayerRepository{db: Repository{
		config: config,
		DB:     dbClient,
	}}
}

func (u *PlayerRepository) CreateUser(ctx context.Context, user models.Player) (*models.Player, error) {
	var newUser models.Player
	if err := u.db.WithContext(ctx).
		Create(user).
		Find(&newUser, "id = ?", user.ID.String()).
		Error; err != nil {

		return nil, err
	}
	return &newUser, nil
}

func (u *PlayerRepository) UpdateUser(ctx context.Context, filter QueryFilter, value interface{}) (*models.Player, error) {
	var user models.Player

	if err := u.db.WithContext(ctx).
		Table(user.TableName()).
		Where(filter.GetFilter()).
		Updates(value).
		Find(&user, filter.GetFilter()).
		Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (u *PlayerRepository) GetUser(ctx context.Context, filter QueryFilter) (*models.Player, error) {
	var user models.Player

	if err := u.db.WithContext(ctx).
		Where(filter.GetFilter()).
		Find(&user, filter.GetFilter()).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}
