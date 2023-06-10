package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tejiriaustin/apex-network/models"
	"github.com/tejiriaustin/apex-network/repository"
	"github.com/tejiriaustin/apex-network/requests"
	"github.com/tejiriaustin/apex-network/response"
	"github.com/tejiriaustin/apex-network/service"
	"net/http"
	"strconv"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) FundWallet(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.FundWalletInput{
			UserId: ctx.Param("user_id"),
		}
		walletBalance, err := sc.FundWallet(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		resp := struct {
			Balance int `json:"balance"`
		}{
			Balance: walletBalance,
		}

		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) CreateUser(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateUserRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		input := service.CreateUserInput{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		}

		user, err := sc.CreateUser(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "OK", user)
	}
}

func (c *Controller) GetWalletBalance(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.GetWalletBalanceInput{
			UserId: ctx.Param("user_id"),
		}

		balance, err := sc.GetWalletBalance(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		resp := struct {
			Balance string `json:"balance"`
			Asset   string
		}{
			Balance: strconv.Itoa(balance),
			Asset:   "sats",
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) StartGameSession(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.StartGameSessionInput{
			UserId: ctx.Param("user_id"),
		}

		player, err := sc.StartGameSession(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		resp := struct {
			Player *models.Player `json:"user"`
			Target string         `json:"target"`
			Asset  string         `json:"asset"`
		}{
			Player: player,
			Target: strconv.Itoa(player.TargetNumber),
			Asset:  "sats",
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) RollDice(sc service.ServiceInterface,
	userRepo repository.PlayerRepositoryInterface,
	walletRepo repository.WalletRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.RollDiceInput{
			UserId: ctx.Param("user_id"),
		}

		player, rolledDie, err := sc.RollDice(ctx, input, userRepo, walletRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		resp := struct {
			Player *models.Player `json:"user"`
			Draw   string         `json:"target"`
			Asset  string         `json:"asset"`
		}{
			Player: player,
			Draw:   strconv.Itoa(rolledDie),
			Asset:  "sats",
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", resp)
	}
}

func (c *Controller) EndGameSession(sc service.ServiceInterface,
	repo repository.PlayerRepositoryInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := service.EndGameSessionInput{
			UserId: ctx.Param("user_id"),
		}

		err := sc.EndGameSession(ctx, input, repo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		response.FormatResponse(ctx, http.StatusOK, "OK", nil)

	}
}
