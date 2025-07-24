package test

import (
	"context"
	"go-project/internal/orm/gosql"
	"go-project/internal/service"
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

func TestUserService(t *testing.T) {
	setup(t)
	// 手动初始化 gosql 模块
	gosqlModule := gosql.NewModule()
	gosqlModule.Init(viper.GetViper())

	userService := service.NewUserService()
	ctx := context.Background()

	err := userService.UserRegister(ctx, "zhangyan", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("user register success")
	user, err := userService.UserLogin(ctx, "zhangyan", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("user login success", "user", user.Id)

	err = userService.DelUser(ctx, user.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("user delete success")

	// 手动关闭 gosql 模块
	gosqlModule.Close()
}
