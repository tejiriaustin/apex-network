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
			repo repository.UserRepositoryInterface,
		) (*models.User, error)

		FundWallet(ctx context.Context,
			input FundWalletInput,
			repo repository.UserRepositoryInterface) (int, error)

		GetWalletBalance(ctx context.Context,
			input GetWalletBalanceInput,
			repo repository.UserRepositoryInterface) (int, error)

		StartGameSession(ctx context.Context,
			input StartGameSessionInput,
			repo repository.UserRepositoryInterface,
		) error
		EndGameSession(ctx context.Context,
			input EndGameSessionInput,
			repo repository.UserRepositoryInterface,
		) error
		RollDice(ctx context.Context,
			input RollDiceInput,
			repo repository.RepositoryInterface,
		) error
	}
)
