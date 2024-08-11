package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
)

type WalletRepository struct {
	db Repository
}

type WalletRepositoryInterface interface {
	GetWalletTransactions(ctx context.Context, playerID string) ([]*models.WalletTransaction, error)
	CreateTransaction(ctx context.Context, tx *models.WalletTransaction) (*models.WalletTransaction, error)
}

func NewWalletRepository(config env.Env, dbClient *database.Client) WalletRepositoryInterface {
	return &WalletRepository{db: Repository{
		config: config,
		DB:     dbClient,
	}}
}

var _ WalletRepositoryInterface = (*WalletRepository)(nil)

func (w *WalletRepository) GetWalletTransactions(ctx context.Context, playerID string) ([]*models.WalletTransaction, error) {
	transaction := make([]*models.WalletTransaction, 0)

	if err := w.db.WithContext(ctx).
		Table("wallet_transaction").
		Find(&transaction, "player_id = ?", playerID).
		Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (w *WalletRepository) CreateTransaction(ctx context.Context, tx *models.WalletTransaction) (*models.WalletTransaction, error) {
	if err := w.db.WithContext(ctx).Create(&tx).Error; err != nil {
		return nil, err
	}
	return tx, nil
}
