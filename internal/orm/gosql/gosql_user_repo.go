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

// NewUserRepository - 创建一个新的 database/sql 用户仓储
func NewUserRepository() repository.IUserRepository {
	once.Do(func() {
		slog.Info("gosql user repository init")
		userRepo = &gosqlUserRepository{db: GetDB()}
	})
	return userRepo
}

// Create - 创建一个新用户
func (r *gosqlUserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO t_user (username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

// GetByID - 通过 ID 获取用户
func (r *gosqlUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT * FROM t_user WHERE id = ?", id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername - 通过 username 获取用户
func (r *gosqlUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT * FROM t_user WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Count - 统计用户数量
func (r *gosqlUserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM t_user").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Delete - 删除用户
func (r *gosqlUserRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM t_user WHERE id = ?", id)
	return err
}
