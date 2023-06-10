package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/apex-network/repository"
	"github.com/tejiriaustin/apex-network/response"
	"github.com/tejiriaustin/apex-network/service"
)

func Routes(
	ctx context.Context,
	r *gin.Engine,
	sc service.ServiceInterface,
	repo *repository.RepositoryContainer,
) {

	controller := NewController()

	r.GET("/health", func(c *gin.Context) {
		response.FormatResponse(c, http.StatusOK, "OK", nil)
	})

	gameRouter := r.Group("/game")
	{
		gameRouter.POST("/create-player", controller.CreatePlayer(sc, repo.PlayerRepo))
		gameRouter.POST("/fund-wallet/:player_id", controller.FundWallet(sc, repo.PlayerRepo, repo.WalletRpo))
		gameRouter.GET("/balance/:player_id", controller.GetWalletBalance(sc, repo.PlayerRepo))
		gameRouter.POST("/start/:player_id", controller.StartGameSession(sc, repo.PlayerRepo, repo.WalletRpo))
		gameRouter.DELETE("/end/:player_id", controller.EndGameSession(sc, repo.PlayerRepo))
		gameRouter.POST("/roll-dice/:player_id", controller.RollDice(sc, repo.PlayerRepo, repo.WalletRpo))
		gameRouter.GET("/is-playing/:player_id", controller.GameInSession(sc, repo.PlayerRepo))
		gameRouter.GET("/transactions/:player_id", controller.WalletTransactions(sc, repo.WalletRpo))
	}
}
