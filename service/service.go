package service

import (
	"context"
	"errors"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
	"github.com/tejiriaustin/apex-network/repository"
)

type (
	Service struct {
		Config env.Env
	}
)

var (
	defaultFundWalletAmount = 155
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
	repo repository.UserRepositoryInterface,
) (*models.User, error) {

	user := models.User{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		IsPlaying:     false,
		WalletBalance: 0,
	}
	user.Init()
	user.FullName = user.GetFullName()

	return repo.CreateUser(ctx, user)
}

func (u *Service) FundWallet(ctx context.Context,
	input FundWalletInput,
	repo repository.UserRepositoryInterface) (int, error) {
	filter := repository.NewQueryFilter()

	filter.AddFilter("id", input.UserId)

	user, err := repo.UpdateUser(ctx, filter, models.FieldUserBalance, defaultFundWalletAmount)
	if err != nil {
		return 0, err
	}

	return user.WalletBalance, nil
}

func (u *Service) GetWalletBalance(ctx context.Context,
	input GetWalletBalanceInput,
	repo repository.UserRepositoryInterface) (int, error) {

	filter := repository.NewQueryFilter()

	filter.AddFilter("id", input.UserId)

	user, err := repo.GetUser(ctx, filter)
	if err != nil {
		return 0, err
	}
	return user.WalletBalance, nil
}

func (u *Service) StartGameSession(ctx context.Context,
	input StartGameSessionInput,
	repo repository.UserRepositoryInterface,
) error {

	filter := repository.NewQueryFilter()

	filter.AddFilter("id", input.UserId)

	user, err := repo.GetUser(ctx, filter)
	if err != nil {
		return err
	}

	if user.IsPlaying == true {
		return errors.New("can only start a game when no game is in session")
	}

	_, err = repo.UpdateUser(ctx, filter, models.FieldUserIsPlaying, true)
	if err != nil {
		return err
	}

	return nil
}

func (u *Service) EndGameSession(ctx context.Context,
	input EndGameSessionInput,
	repo repository.UserRepositoryInterface,
) error {

	filter := repository.NewQueryFilter()

	filter.AddFilter("id", input.UserId)

	user, err := repo.GetUser(ctx, filter)
	if err != nil {
		return err
	}

	if user.IsPlaying == false {
		return errors.New("can only end a game if an active game is in session")
	}

	_, err = repo.UpdateUser(ctx, filter, models.FieldUserIsPlaying, false)
	if err != nil {
		return err
	}

	return nil
}

func (u *Service) RollDice(ctx context.Context,
	input RollDiceInput,
	repo repository.RepositoryInterface,
) error {
	return nil
}
