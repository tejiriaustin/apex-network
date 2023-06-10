package service

import (
	"context"
	"github.com/tejiriaustin/apex-network/repository"

	"github.com/tejiriaustin/apex-network/models"
)

type (
	ServiceInterface interface {
		CreateUser(ctx context.Context,
			input CreateUserInput,
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
			userRepo repository.PlayerRepositoryInterface,
			walletRepo repository.WalletRepositoryInterface,
		) (*models.Player, int, error)
	}
)
