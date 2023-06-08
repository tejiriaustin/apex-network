package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/models"
)

type GameRepository struct {
	db Repository
}

type GameRepositoryInterface interface {
	UpdateGame(ctx context.Context, game models.Game) (*models.Game, error)
}

var _ GameRepositoryInterface = (*GameRepository)(nil)

func (g *GameRepository) UpdateGame(ctx context.Context, game models.Game) (*models.Game, error) {
	if err := g.db.WithContext(ctx).
		Table(game.TableName()).
		Update("", "").
		Error; err != nil {
		return nil, err
	}
	return &game, nil
}
