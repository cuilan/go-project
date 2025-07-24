package gorm

import (
	"go-project/internal/orm/repository"

	"gorm.io/gorm"
)

// gormBookRepository 是 IBookRepository 的 GORM 实现
type gormBookRepository struct {
	db *gorm.DB
}

// NewGormBookRepository - 创建一个新的 GORM 用户仓储
func NewGormBookRepository(db *gorm.DB) repository.IBookRepository {
	return &gormBookRepository{db: db}
}
