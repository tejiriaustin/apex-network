package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tejiriaustin/apex-network/repository"
	"github.com/tejiriaustin/apex-network/requests"
	"github.com/tejiriaustin/apex-network/response"
	"github.com/tejiriaustin/apex-network/service"
	"github.com/tejiriaustin/apex-network/utils"
	"net/http"
	"strconv"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) FundWallet(sc service.ServiceInterface,
	playerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.FundWalletInput{
			PlayerId: ctx.Param("player_id"),
		}

		walletBalance, err := sc.FundWallet(ctx, input, playerRepo, walletRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		resp := struct {
			Balance int    `json:"balance"`
			Asset   string `json:"asset"`
		}{
			Balance: walletBalance,
			Asset:   "sats",
		}

		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) CreatePlayer(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreatePlayerRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		input := service.CreatePlayerInput{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		}

		Player, err := sc.CreatePlayer(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "OK", Player)
	}
}

func (c *Controller) GetWalletBalance(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.GetWalletBalanceInput{
			PlayerId: ctx.Param("player_id"),
		}

		balance, err := sc.GetWalletBalance(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		resp := struct {
			Balance string `json:"balance"`
			Asset   string `json:"asset"`
		}{
			Balance: strconv.Itoa(balance),
			Asset:   "sats",
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) StartGameSession(sc service.ServiceInterface,
	playerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.StartGameSessionInput{
			PlayerId: ctx.Param("player_id"),
		}

		randFunc := utils.GetRandFunc()

		player, err := sc.StartGameSession(ctx, input, randFunc, playerRepo, walletRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		resp := struct {
			Target string `json:"target"`
			Asset  string `json:"asset"`
		}{
			Target: strconv.Itoa(player.TargetNumber),
			Asset:  "sats",
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) RollDice(sc service.ServiceInterface,
	PlayerRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.RollDiceInput{
			PlayerId: ctx.Param("player_id"),
		}

		randFunc := utils.GetRandFunc()
		player, rolledDie, err := sc.RollDice(ctx, input, randFunc, PlayerRepo, walletRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		resp := struct {
			TargetNumber string `json:"target_number"`
			Draw         string `json:"draw"`
			Asset        string `json:"asset"`
			Status       string `json:"status"`
		}{
			TargetNumber: strconv.Itoa(player.TargetNumber),
			Draw:         strconv.Itoa(rolledDie),
			Asset:        "sats",
			Status:       status(player.TargetNumber, player.DiceSum),
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func status(target, cast int) string {
	if cast == target {
		return "WON"
	}
	return "LOST"
}

func (c *Controller) EndGameSession(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.EndGameSessionInput{
			PlayerId: ctx.Param("player_id"),
		}

		err := sc.EndGameSession(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", nil)

	}
}

func (c *Controller) GameInSession(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := service.GameIsInitializedInput{
			PlayerId: ctx.Param("player_id"),
		}

		player, err := sc.GameIsInitialized(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		resp := struct {
			GameInSession bool `json:"game_in_session"`
		}{
			GameInSession: *player.IsPlaying,
		}

		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) WalletTransactions(sc service.ServiceInterface,
	walletRepo repository.WalletRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := service.GetTransactionLogsInput{
			PlayerId: ctx.Param("player_id"),
		}

		transactions, err := sc.GetWalletTransactions(ctx, input, walletRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", transactions)
	}
}
