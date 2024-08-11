package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	WithContext(ctx context.Context) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Delete(dest interface{}, conds ...interface{}) *gorm.DB
	Update(column string, value interface{}) *gorm.DB
}

type RepositoryContainer struct {
	PlayerRepo PlayerRepositoryInterface
	WalletRpo  WalletRepositoryInterface
}

func NewRepositoryContainer(config env.Env, dbClient *database.Client) *RepositoryContainer {
	return &RepositoryContainer{
		PlayerRepo: NewPlayerRepository(config, dbClient),
		WalletRpo:  NewWalletRepository(config, dbClient),
	}
}
