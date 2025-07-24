package gosql

import (
	"go-project/internal/orm/repository"
)

// autowired 自动注入repository到容器
func autowired() {
	repository.RegisterRepository(repository.UserRepositoryName, NewUserRepository())
}
