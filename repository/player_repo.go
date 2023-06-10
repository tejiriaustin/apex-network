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
	CreatePlayer(ctx context.Context, Player models.Player) (*models.Player, error)
	UpdatePlayer(ctx context.Context, playerID string, player models.Player) (*models.Player, error)
	GetPlayerbyID(ctx context.Context, PlayerID string) (*models.Player, error)
}

func NewPlayerRepository(config env.Env, dbClient *database.Client) PlayerRepositoryInterface {
	return &PlayerRepository{db: Repository{
		config: config,
		DB:     dbClient,
	}}
}

func (u *PlayerRepository) CreatePlayer(ctx context.Context, Player models.Player) (*models.Player, error) {
	var newPlayer models.Player
	if err := u.db.WithContext(ctx).
		Table(Player.TableName()).
		Create(&Player).
		Find(&newPlayer, "id = ?", Player.ID.String()).
		Error; err != nil {

		return nil, err
	}
	return &newPlayer, nil
}

func (u *PlayerRepository) UpdatePlayer(ctx context.Context, playerID string, player models.Player) (*models.Player, error) {

	if err := u.db.WithContext(ctx).
		Table(player.TableName()).
		Updates(&player).
		Find(&player, "id = ?", playerID).
		Error; err != nil {

		return nil, err
	}

	return &player, nil
}

func (u *PlayerRepository) GetPlayerbyID(ctx context.Context, playerID string) (*models.Player, error) {
	var player models.Player

	if err := u.db.WithContext(ctx).
		Table(player.TableName()).
		First(&player, "id = ?", playerID).
		Error; err != nil {
		return nil, err
	}

	return &player, nil
}
