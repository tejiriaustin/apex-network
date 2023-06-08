package database

import (
	"fmt"
	"github.com/tejiriaustin/apex-network/env"
	"gorm.io/gorm"
	"log"
)

type Client struct {
	*gorm.DB
}

func OpenDatabaseConnection(config env.Env) *Client {

	_ = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v")

	db, err := gorm.Open(nil, &gorm.Config{})
	if err != nil {
		log.Panicf("faile to connect to database: %v", err)
	}
	return &Client{db}
}
