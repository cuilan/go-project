package repository

import (
	"context"
	"go-project/internal/orm/models"
)

// IUserRepository 定义了用户数据操作的接口
type IUserRepository interface {
	// Create - 创建一个新用户
	Create(ctx context.Context, user *models.User) error
	// GetByID - 通过 ID 获取用户
	GetByID(ctx context.Context, id uint) (*models.User, error)
}
