package service

import (
	"context"
	"errors"
	"go-project/internal/orm/models"
	"go-project/internal/orm/repository"
	"log/slog"
)

// UserService 用户服务层
type UserService struct {
	userRepo repository.IUserRepository
}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.GetUserRepository(),
	}
}

// CreateUser 创建用户
func (s *UserService) UserRegister(ctx context.Context, username string, password string) error {
	user := &models.User{Username: username, Password: password}
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		slog.Error("failed to create user", "username", username, "password", password, "err", err)
		return err
	}
	slog.Info("user created successfully", "username", username, "password", password, "id", user.Id)
	return nil
}

func (s *UserService) UserLogin(ctx context.Context, username string, password string) (*models.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		slog.Error("failed to get user by username", "username", username, "err", err)
		return nil, err
	}
	if user.Password != password {
		slog.Error("password is incorrect", "username", username, "password", password)
		return nil, errors.New("password is incorrect")
	}
	return user, nil
}

func (s *UserService) DelUser(ctx context.Context, userId int64) error {
	err := s.userRepo.Delete(ctx, userId)
	if err != nil {
		slog.Error("failed to delete user", "userId", userId, "err", err)
		return err
	}
	return nil
}
