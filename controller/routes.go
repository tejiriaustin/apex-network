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
	repo repository.RepositoryInterface,
) {

	controller := NewController()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.FormatResponse(http.StatusOK, "OK", nil))
	})

	gameRouter := r.Group("/game")
	{
		gameRouter.POST("/create-user", controller.CreateUser(sc, repo))
		gameRouter.POST("/fund-wallet", controller.FundWallet(sc, repo))
		gameRouter.GET("/balance", controller.GetWalletBalance(sc, repo))
		gameRouter.POST("/start", controller.StartGameSession(sc, repo))
		gameRouter.DELETE("/end", controller.EndGameSession(sc, repo))
		gameRouter.POST("/roll-dice", controller.RollDice(sc, repo))
	}
}
