package test

import (
	"context"
	"go-project/internal/orm/gosql"
	"go-project/internal/orm/repository"
	"testing"

	"github.com/spf13/viper"
)

func setup(t *testing.T) {
	viper.SetConfigFile("../test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitGosql(t *testing.T) {
	setup(t)
	// 手动初始化 gosql 模块
	gosqlModule := gosql.NewModule()
	gosqlModule.Init(viper.GetViper())

	userRepo := repository.GetUserRepository()
	count, err := userRepo.Count(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)

	// 手动关闭 gosql 模块
	gosqlModule.Close()
}
