package service

import (
	"context"
	"errors"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
	"github.com/tejiriaustin/apex-network/repository"
	"math/rand"
	"time"
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
	repo repository.PlayerRepositoryInterface,
) (*models.Player, error) {

	user := models.Player{
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
	repo repository.PlayerRepositoryInterface) (int, error) {
	filter := repository.NewQueryFilter()

	filter.AddFilter("id", input.UserId)

	user, err := repo.GetUser(ctx, filter)
	if err != nil {
		return 0, err
	}
	user.WalletBalance += defaultFundWalletAmount

	user, err = repo.UpdateUser(ctx, filter, user)
	if err != nil {
		return 0, err
	}

	return user.WalletBalance, nil
}

func (u *Service) GetWalletBalance(ctx context.Context,
	input GetWalletBalanceInput,
	repo repository.PlayerRepositoryInterface) (int, error) {

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
	repo repository.PlayerRepositoryInterface,
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

	user.TargetNumber = genRandomNumber()
	user.IsPlaying = true

	_, err = repo.UpdateUser(ctx, filter, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *Service) EndGameSession(ctx context.Context,
	input EndGameSessionInput,
	repo repository.PlayerRepositoryInterface,
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

	user.IsPlaying = false
	_, err = repo.UpdateUser(ctx, filter, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *Service) RollDice(ctx context.Context,
	input RollDiceInput,
	userRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) error {

	filter := repository.NewQueryFilter()

	filter.AddFilter("id", input.UserId)

	user, err := userRepo.GetUser(ctx, filter)
	if err != nil {
		return err
	}

	if user.IsPlaying == false {
		return errors.New("please start a new session before rolling a dice")
	}

	if user.HasRolledFirstDie == true {
		// Roll die again but don't get debited
		user.DiceSum += genRandomNumber()

		if user.DiceSum == user.TargetNumber {
			tx := models.WalletTransaction{
				Amount:          10,
				Description:     models.RollCost,
				TransactionType: models.Credit,
			}
			_, err = walletRepo.CreateTransaction(ctx, &tx)
			if err != nil {
				return err
			}
		}
		//update hasRolled status to false
		user.HasRolledFirstDie = false
		_, err = userRepo.UpdateUser(ctx, filter, user)
		if err != nil {
			return err
		}
		return nil
	}

	tx := models.WalletTransaction{
		Amount:          5,
		Description:     models.RollCost,
		TransactionType: models.Debit,
	}
	tx.Init()

	_, err = walletRepo.CreateTransaction(ctx, &tx)
	if err != nil {
		return err
	}

	user.DiceSum = genRandomNumber()
	_, err = userRepo.UpdateUser(ctx, filter, user)
	if err != nil {
		return err
	}

	return nil
}

func genRandomNumber() int {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(10) + 2
	return number
}
