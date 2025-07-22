package gosql

import (
	"context"
	"database/sql"
	"go-project/internal/orm/models"
	"go-project/internal/orm/repository"
	"log/slog"
	"sync"
)

var once sync.Once
var userRepo repository.IUserRepository

// gosqlUserRepository 是 IUserRepository 的 database/sql 实现
type gosqlUserRepository struct {
	db *sql.DB
}

// UserRepository - 创建一个新的 database/sql 用户仓储
func UserRepository() repository.IUserRepository {
	once.Do(func() {
		slog.Info("NewGosqlUserRepository")
		userRepo = &gosqlUserRepository{db: GetDB()}
	})
	return userRepo
}

// Create - 创建一个新用户
func (r *gosqlUserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", user.Name)
	return err
}

// GetByID - 通过 ID 获取用户
func (r *gosqlUserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Count - 统计用户数量
func (r *gosqlUserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
