package repository

import (
	"context"
	"go-project/internal/orm/models"
)

// UserRepositoryName 用户repository名称
const UserRepositoryName = "user_repository"

// GetUserRepository 获取用户repository（便捷方法）
func GetUserRepository() IUserRepository {
	repo, exists := GetRepository(UserRepositoryName)
	if !exists {
		panic("user repository not registered")
	}
	return repo.(IUserRepository)
}

// IUserRepository 定义了用户数据操作的接口
type IUserRepository interface {
	// Create - 创建一个新用户
	Create(ctx context.Context, user *models.User) error

	// GetByID - 通过 ID 获取用户
	GetByID(ctx context.Context, id int64) (*models.User, error)

	// GetByUsername - 通过 username 获取用户
	GetByUsername(ctx context.Context, username string) (*models.User, error)

	// Count - 统计用户数量
	Count(ctx context.Context) (int64, error)

	// Delete - 删除用户
	Delete(ctx context.Context, id int64) error
}
