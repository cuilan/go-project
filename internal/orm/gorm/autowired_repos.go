package gorm

import (
	"go-project/internal/orm/repository"

	"gorm.io/gorm"
)

// autowired 自动注入repository到容器
func autowired(db *gorm.DB) {
	// 注册用户repository到容器
	repository.RegisterRepository(repository.UserRepositoryName, NewGormUserRepository(db))
	// 注册书籍repository到容器
	repository.RegisterRepository(repository.BookRepositoryName, NewGormBookRepository(db))
}
