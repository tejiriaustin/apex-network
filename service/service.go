package service

import (
	"context"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
	"github.com/tejiriaustin/apex-network/repository"
)

type (
	Service struct {
		Config env.Env
	}
)

func NewService(env env.Env) ServiceInterface {
	return &Service{
		Config: env,
	}
}

var _ ServiceInterface = (*Service)(nil)

type (
	CreateUserInput struct {
		FirstName string
		LastName  string
	}
	FundWalletInput struct {
		UserId string
	}

	GetWalletBalanceInput struct {
		UserId string
	}
	StartGameSessionInput struct {
		UserId string
	}
	EndGameSessionInput struct {
		UserId string
	}
	RollDiceInput struct {
		UserId string
	}
)

func (u *Service) CreateUser(ctx context.Context,
	input CreateUserInput,
	repo repository.RepositoryInterface,
) (*models.User, error) {
	return nil, nil
}

func (u *Service) FundWallet(ctx context.Context,
	input FundWalletInput,
	repo repository.RepositoryInterface) (int, error) {
	return 0, nil
}

func (u *Service) GetWalletBalance(ctx context.Context,
	input GetWalletBalanceInput,
	repo repository.RepositoryInterface) (int, error) {
	return 0, nil
}

func (u *Service) StartGameSession(ctx context.Context,
	input StartGameSessionInput,
	repo repository.RepositoryInterface,
) error {
	return nil
}
func (u *Service) EndGameSession(ctx context.Context,
	input EndGameSessionInput,
	repo repository.RepositoryInterface,
) error {
	return nil
}
func (u *Service) RollDice(ctx context.Context,
	input RollDiceInput,
	repo repository.RepositoryInterface,
) {

}
