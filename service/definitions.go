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
			repo repository.RepositoryInterface,
		) (*models.User, error)

		FundWallet(ctx context.Context,
			input FundWalletInput,
			repo repository.RepositoryInterface) (int, error)

		GetWalletBalance(ctx context.Context,
			input GetWalletBalanceInput,
			repo repository.RepositoryInterface) (int, error)

		StartGameSession(ctx context.Context,
			input StartGameSessionInput,
			repo repository.RepositoryInterface,
		) error
		EndGameSession(ctx context.Context,
			input EndGameSessionInput,
			repo repository.RepositoryInterface,
		)
		RollDice(ctx context.Context,
			input RollDiceInput,
			repo repository.RepositoryInterface,
		)
	}
)
