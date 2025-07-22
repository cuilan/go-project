package gorm

import (
	"context"
	"go-project/internal/orm/models"
	"go-project/internal/orm/repository"

	"gorm.io/gorm"
)

// gormUserRepository 是 IUserRepository 的 GORM 实现
type gormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository - 创建一个新的 GORM 用户仓储
func NewGormUserRepository(db *gorm.DB) repository.IUserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *gormUserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
