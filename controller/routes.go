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
	service service.IService,
	Repo repository.IRepository,
) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.FormatResponse(http.StatusOK, "OK", nil))
	})
}
