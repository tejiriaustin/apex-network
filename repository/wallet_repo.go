package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/models"
)

type WalletRepository struct {
	db Repository
}

type WalletRepositoryInterface interface {
	GetWalletTransactions(ctx context.Context, filter QueryFilter) ([]*models.WalletTransaction, error)
}

var _ WalletRepositoryInterface = (*WalletRepository)(nil)

func (w *WalletRepository) GetWalletTransactions(ctx context.Context, filter QueryFilter) ([]*models.WalletTransaction, error) {
	transaction := make([]*models.WalletTransaction, 0)
	if err := w.db.WithContext(ctx).
		Find(transaction, filter.GetFilter()).
		Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
