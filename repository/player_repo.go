package repository

import (
	"context"
	"fmt"
	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
)

type PlayerRepository struct {
	db Repository
}

type PlayerRepositoryInterface interface {
	CreateUser(ctx context.Context, user models.Player) (*models.Player, error)
	UpdateUser(ctx context.Context, playerID string, player models.Player) (*models.Player, error)
	GetUserbyID(ctx context.Context, userID string) (*models.Player, error)
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
		Table(user.TableName()).
		Create(&user).
		Find(&newUser, "id = ?", user.ID.String()).
		Error; err != nil {

		return nil, err
	}
	return &newUser, nil
}

func (u *PlayerRepository) UpdateUser(ctx context.Context, playerID string, player models.Player) (*models.Player, error) {
	var user models.Player

	if err := u.db.WithContext(ctx).
		Table(user.TableName()).
		Where("id = ?", playerID).
		Updates(&player).
		Find(&user, "id = ?", playerID).
		Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (u *PlayerRepository) GetUserbyID(ctx context.Context, playerID string) (*models.Player, error) {
	var user models.Player

	if err := u.db.WithContext(ctx).
		Table(user.TableName()).
		First(&user, "id = ?", playerID).
		Error; err != nil {
		fmt.Println("asretdyfugi")
		return nil, err
	}

	return &user, nil
}
