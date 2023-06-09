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
		gameRouter.POST("/create-user", controller.CreateUser(sc, repo.PlayerRepo))
		gameRouter.POST("/fund-wallet", controller.FundWallet(sc, repo.PlayerRepo))
		gameRouter.GET("/balance", controller.GetWalletBalance(sc, repo.PlayerRepo))
		gameRouter.POST("/start", controller.StartGameSession(sc, repo.PlayerRepo))
		gameRouter.DELETE("/end", controller.EndGameSession(sc, repo.PlayerRepo))
		gameRouter.POST("/roll-dice", controller.RollDice(sc, repo.PlayerRepo, repo.WalletRpo))
	}
}
