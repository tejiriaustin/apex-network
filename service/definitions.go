package service

import (
	"context"
	"github.com/tejiriaustin/apex-network/repository"

	"github.com/tejiriaustin/apex-network/models"
)

type (
	ServiceInterface interface {
		CreatePlayer(ctx context.Context,
			input CreatePlayerInput,
			repo repository.PlayerRepositoryInterface,
		) (*models.Player, error)

		FundWallet(ctx context.Context,
			input FundWalletInput,
			repo repository.PlayerRepositoryInterface) (int, error)

		GetWalletBalance(ctx context.Context,
			input GetWalletBalanceInput,
			repo repository.PlayerRepositoryInterface) (int, error)

		StartGameSession(ctx context.Context,
			input StartGameSessionInput,
			repo repository.PlayerRepositoryInterface,
		) (*models.Player, error)
		EndGameSession(ctx context.Context,
			input EndGameSessionInput,
			repo repository.PlayerRepositoryInterface,
		) error
		RollDice(ctx context.Context,
			input RollDiceInput,
			PlayerRepo repository.PlayerRepositoryInterface,
			walletRepo repository.WalletRepositoryInterface,
		) (*models.Player, int, error)

		GameIsInitialized(ctx context.Context,
			input GameIsInitializedInput,
			PlayerRepo repository.PlayerRepositoryInterface,
		) (*models.Player, error)
		GetWalletTransactions(ctx context.Context,
			input GetTransactionLogsInput,
			walletRepo repository.WalletRepositoryInterface,
		) ([]*models.WalletTransaction, error)
	}
)
