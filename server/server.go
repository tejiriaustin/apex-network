package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/apex-network/controller"
	"github.com/tejiriaustin/apex-network/repository"
	"github.com/tejiriaustin/apex-network/service"
)

func Start(ctx context.Context,
	service service.IService,
	repo repository.IRepository) {
	router := gin.New()

	controller.Routes(ctx, router, service, repo)

	go func() {
		if err := router.Run(); err != nil {
			log.Fatal("shutting down...")
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := router; err != nil {
		log.Fatal(err)
	}
}
