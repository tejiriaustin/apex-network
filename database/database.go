package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/tejiriaustin/apex-network/env"
)

type Client struct {
	*gorm.DB
}

func OpenDatabaseConnection(config env.Env) *Client {

	_ = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		config.GetEnvString(env.DbHost),
		config.GetEnvString(env.DbUsername),
		config.GetEnvString(env.DbPassword),
		config.GetEnvString(env.DbDatabase),
		"12345",
		config.GetEnvString(env.DbTimeZone),
	)

	db, err := gorm.Open(nil, &gorm.Config{})
	if err != nil {
		log.Panicf("faile to connect to database: %v", err)
	}
	return &Client{db}
}
