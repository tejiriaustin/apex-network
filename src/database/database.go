package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/gorm"

	"github.com/tejiriaustin/apex-network/env"
)

type Client struct {
	*gorm.DB
}

func OpenDatabaseConnection(config env.Env) *Client {

	dsn := fmt.Sprintf(config.GetEnvString(env.DbUrl))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect to database: %v", err)
	}

	//err = db.AutoMigrate(&models.Game{}, &models.WalletTransaction{}, models.Player{})
	//if err != nil {
	//	log.Panicf("failed to run auto migrate: %v", err)
	//}
	return &Client{db}
}
