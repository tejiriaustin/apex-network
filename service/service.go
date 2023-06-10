package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/tejiriaustin/apex-network/env"
	"github.com/tejiriaustin/apex-network/models"
	"github.com/tejiriaustin/apex-network/repository"
	"github.com/tejiriaustin/apex-network/utils"
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
	winReward               = 10

	False = false
	True  = true
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
	GetTransactionLogsInput struct {
		PlayerId string
	}
)

func (u *Service) CreatePlayer(ctx context.Context,
	input CreatePlayerInput,
	repo repository.PlayerRepositoryInterface,
) (*models.Player, error) {

	Player := models.Player{
		FirstName:         input.FirstName,
		LastName:          input.LastName,
		IsPlaying:         &False,
		WalletBalance:     0,
		HasRolledFirstDie: &False,
	}
	Player.ID = uuid.New()
	Player.CreatedAt = time.Now().UTC()
	Player.FullName = Player.GetFullName()

	return repo.CreatePlayer(ctx, Player)
}

func (u *Service) FundWallet(ctx context.Context,
	input FundWalletInput,
	playerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface) (int, error) {

	player, err := playerRepo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return 0, err
	}

	if player.WalletBalance > 35 {
		return 0, errors.New("player can only fund wallet when balance is less than 35")
	}

	player.WalletBalance += defaultFundWalletAmount

	player, err = playerRepo.UpdatePlayer(ctx, player.ID.String(), *player)
	if err != nil {
		return 0, err
	}
	transaction := models.WalletTransaction{
		PlayerId:        player.ID,
		Amount:          defaultFundWalletAmount,
		Description:     models.FundWallet,
		TransactionType: models.Credit,
	}
	transaction.ID = uuid.New()
	transaction.CreatedAt = time.Now().UTC()
	_, err = walletRepo.CreateTransaction(ctx, &transaction)
	if err != nil {
		return 0, err
	}

	return player.WalletBalance, nil
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
	randFunc utils.RandFunc,
	playerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) (*models.Player, error) {

	Player, err := playerRepo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return nil, err
	}

	if Player.WalletBalance < startGameCost {
		return nil, errors.New("insufficient wallet balance")
	}

	if Player.IsPlaying == &True {
		return nil, errors.New("can only start a game when no game is in session")
	}

	Player.TargetNumber = randFunc()
	Player.WalletBalance -= startGameCost
	Player.IsPlaying = &True

	_, err = playerRepo.UpdatePlayer(ctx, input.PlayerId, *Player)
	if err != nil {
		return nil, err
	}
	transaction := &models.WalletTransaction{
		PlayerId:        Player.ID,
		Amount:          startGameCost,
		Description:     models.StartGame,
		TransactionType: models.Debit,
	}
	transaction.ID = uuid.New()
	transaction.CreatedAt = time.Now().UTC()
	_, err = walletRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return Player, nil
}

func (u *Service) EndGameSession(ctx context.Context,
	input EndGameSessionInput,
	repo repository.PlayerRepositoryInterface,
) error {

	player, err := repo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return err
	}

	if *player.IsPlaying == false {
		return errors.New("can only end a game if an active game is in session")
	}

	player.IsPlaying = &False
	_, err = repo.UpdatePlayer(ctx, input.PlayerId, *player)
	if err != nil {
		return err
	}

	return nil
}

func (u *Service) RollDice(ctx context.Context,
	input RollDiceInput,
	randFunc utils.RandFunc,
	PlayerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) (*models.Player, int, error) {

	rolledDie := randFunc()

	player, err := PlayerRepo.GetPlayerbyID(ctx, input.PlayerId)
	if err != nil {
		return nil, 0, err
	}

	if player.IsPlaying == &False {
		return nil, 0, errors.New("please start a new session before rolling a dice")
	}

	if *player.HasRolledFirstDie == true {
		// Roll die again but don't get debited
		player.DiceSum += rolledDie

		if player.DiceSum == player.TargetNumber {
			err := rewardWin(ctx, *player, PlayerRepo, walletRepo)
			if err != nil {
				return nil, 0, err
			}
			return player, rolledDie, nil
		}

		//update hasRolled status to false
		player.HasRolledFirstDie = &False
		_, err = PlayerRepo.UpdatePlayer(ctx, input.PlayerId, *player)
		if err != nil {
			return nil, 0, err
		}

		return player, rolledDie, nil
	}

	if player.WalletBalance < dieRollCost {
		return nil, 0, errors.New("insufficient wallet balance")
	}

	tx := models.WalletTransaction{
		PlayerId:        player.ID,
		Amount:          dieRollCost,
		Description:     models.RollCost,
		TransactionType: models.Debit,
	}

	tx.ID = uuid.New()
	_, err = walletRepo.CreateTransaction(ctx, &tx)
	if err != nil {
		return nil, 0, err
	}

	player.DiceSum = rolledDie
	player.WalletBalance -= dieRollCost
	player.HasRolledFirstDie = &True

	_, err = PlayerRepo.UpdatePlayer(ctx, input.PlayerId, *player)
	if err != nil {
		return nil, 0, err
	}

	return player, rolledDie, nil
}

func rewardWin(ctx context.Context,
	player models.Player,
	PlayerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface) error {

	tx := models.WalletTransaction{
		PlayerId:        player.ID,
		Amount:          10,
		Description:     models.RollCost,
		TransactionType: models.Credit,
	}
	tx.ID = uuid.New()
	_, err := walletRepo.CreateTransaction(ctx, &tx)
	if err != nil {
		return err
	}

	player.WalletBalance += winReward
	_, err = PlayerRepo.UpdatePlayer(ctx, player.ID.String(), player)
	if err != nil {
		return err
	}
	return nil
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

func (u *Service) GetWalletTransactions(ctx context.Context,
	input GetTransactionLogsInput,
	walletRepo repository.WalletRepositoryInterface,
) ([]*models.WalletTransaction, error) {

	transactions, err := walletRepo.GetWalletTransactions(ctx, input.PlayerId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
