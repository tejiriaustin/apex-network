package repository

import (
	"context"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	WithContext(ctx context.Context) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Delete(dest interface{}, conds ...interface{}) *gorm.DB
	Update(column string, value interface{}) *gorm.DB
}
