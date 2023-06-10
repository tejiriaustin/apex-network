package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
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
	dieRollCost             = 5
	startGameCost           = 20
)

func NewService(env env.Env) ServiceInterface {
	return &Service{
		Config: env,
	}
}

var _ ServiceInterface = (*Service)(nil)

type (
	CreatePlayerInput struct {
		FirstName string
		LastName  string
	}
	FundWalletInput struct {
		PlayerId string
	}

	GetWalletBalanceInput struct {
		PlayerId string
	}
	StartGameSessionInput struct {
		PlayerId string
	}
	EndGameSessionInput struct {
		PlayerId string
	}
	RollDiceInput struct {
		PlayerId string
	}
	GameIsInitializedInput struct {
		PlayerId string
	}
)

func (u *Service) CreatePlayer(ctx context.Context,
	input CreatePlayerInput,
	repo repository.PlayerRepositoryInterface,
) (*models.Player, error) {

	Player := models.Player{
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		IsPlaying:     false,
		WalletBalance: 0,
	}
	Player.ID = uuid.New()
	Player.CreatedAt = time.Now().UTC()
	Player.FullName = Player.GetFullName()

	return repo.CreatePlayer(ctx, Player)
}

func (u *Service) FundWallet(ctx context.Context,
	input FundWalletInput,
	repo repository.PlayerRepositoryInterface) (int, error) {

	Player, err := repo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return 0, err
	}

	if Player.WalletBalance > 35 {
		return 0, errors.New("player can only fund wallet when balance is less than 35")
	}

	Player.WalletBalance += defaultFundWalletAmount

	Player, err = repo.UpdatePlayer(ctx, Player.ID.String(), *Player)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return Player.WalletBalance, nil
}

func (u *Service) GetWalletBalance(ctx context.Context,
	input GetWalletBalanceInput,
	repo repository.PlayerRepositoryInterface) (int, error) {

	Player, err := repo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return 0, err
	}
	return Player.WalletBalance, nil
}

func (u *Service) StartGameSession(ctx context.Context,
	input StartGameSessionInput,
	repo repository.PlayerRepositoryInterface,
) (*models.Player, error) {
	Player, err := repo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return nil, err
	}

	if Player.WalletBalance < startGameCost {
		return nil, errors.New("insufficient wallet balance")
	}

	if Player.IsPlaying == true {
		return nil, errors.New("can only start a game when no game is in session")
	}

	Player.TargetNumber = genRandomNumber()
	Player.WalletBalance -= startGameCost
	Player.IsPlaying = true

	_, err = repo.UpdatePlayer(ctx, input.PlayerId, *Player)
	if err != nil {
		return nil, err
	}

	return Player, nil
}

func (u *Service) EndGameSession(ctx context.Context,
	input EndGameSessionInput,
	repo repository.PlayerRepositoryInterface,
) error {

	Player, err := repo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return err
	}

	if Player.IsPlaying == false {
		return errors.New("can only end a game if an active game is in session")
	}

	Player.IsPlaying = false
	_, err = repo.UpdatePlayer(ctx, input.PlayerId, *Player)
	if err != nil {
		return err
	}

	return nil
}

func (u *Service) RollDice(ctx context.Context,
	input RollDiceInput,
	PlayerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) (*models.Player, int, error) {

	rolledDie := genRandomNumber()

	Player, err := PlayerRepo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return nil, 0, err
	}

	if Player.IsPlaying == false {
		return nil, 0, errors.New("please start a new session before rolling a dice")
	}

	if Player.HasRolledFirstDie == true {
		// Roll die again but don't get debited
		Player.DiceSum += rolledDie

		if Player.DiceSum == Player.TargetNumber {
			tx := models.WalletTransaction{
				Amount:          10,
				Description:     models.RollCost,
				TransactionType: models.Credit,
			}
			tx.ID = uuid.New()
			_, err = walletRepo.CreateTransaction(ctx, &tx)
			if err != nil {
				return nil, 0, err
			}
		}
		//update hasRolled status to false
		Player.HasRolledFirstDie = false
		_, err = PlayerRepo.UpdatePlayer(ctx, input.PlayerId, *Player)
		if err != nil {
			return nil, 0, err
		}
		return Player, rolledDie, nil
	}

	if Player.WalletBalance < dieRollCost {
		return nil, 0, errors.New("insufficient wallet balance")
	}

	tx := models.WalletTransaction{
		Amount:          dieRollCost,
		Description:     models.RollCost,
		TransactionType: models.Debit,
	}
	tx.ID = uuid.New()

	_, err = walletRepo.CreateTransaction(ctx, &tx)
	if err != nil {
		return nil, 0, err
	}

	Player.DiceSum = rolledDie
	Player.WalletBalance -= dieRollCost
	Player.HasRolledFirstDie = true

	_, err = PlayerRepo.UpdatePlayer(ctx, input.PlayerId, *Player)
	if err != nil {
		return nil, 0, err
	}

	return Player, rolledDie, nil
}

func genRandomNumber() int {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(10) + 2
	return number
}

func (u *Service) GameIsInitialized(ctx context.Context,
	input GameIsInitializedInput,
	PlayerRepo repository.PlayerRepositoryInterface,
) (*models.Player, error) {

	player, err := PlayerRepo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return nil, err
	}
	return player, nil
}
