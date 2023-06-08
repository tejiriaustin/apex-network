package repository

import (
	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
)

type (
	Service struct {
		UserRepo UserRepoInterface
		config   env.Env
		DB       *database.Client
	}
)

type IRepository interface {
}

func NewRepo(env env.Env, db *database.Client) IRepository {
	return &Service{
		UserRepo: NewUserRepo(env),
		config:   env,
		DB:       db,
	}
}
