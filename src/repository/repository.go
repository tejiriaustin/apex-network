package repository

import (
	"context"
	"github.com/tejiriaustin/apex-network/database"
	"github.com/tejiriaustin/apex-network/env"
	"gorm.io/gorm"
)

type (
	Repository struct {
		config env.Env
		DB     *database.Client
	}
)

func NewRepo(env env.Env, db *database.Client) RepositoryInterface {
	return &Repository{
		config: env,
		DB:     db,
	}
}

var _ RepositoryInterface = (*Repository)(nil)

func (r *Repository) WithContext(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}

func (r *Repository) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.DB.Find(dest, conds)
}

func (r *Repository) Create(value interface{}) *gorm.DB {
	return r.DB.Create(value)
}

func (r *Repository) Delete(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.DB.Delete(dest, conds)
}

func (r *Repository) Update(column string, value interface{}) *gorm.DB {
	return r.DB.Update(column, value)
}
