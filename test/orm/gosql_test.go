package test

import (
	"context"
	"go-project/internal/orm/gosql"
	"go-project/internal/orm/repository"
	"testing"

	"github.com/spf13/viper"
)

// TestWithGosql 会在包中所有其他测试运行之前被调用。
func TestWithGosql(t *testing.T) {
	// --- 单个测试的 Setup ---
	t.Log("  【TestWithGosql】=> Setup: 测试准备工作...")

	viper.SetConfigFile("../test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
	// 手动初始化 gosql 模块
	gosqlModule := gosql.NewModule()
	gosqlModule.Init(viper.GetViper())

	// --- 使用 t.Cleanup 注册 Teardown 函数 ---
	// t.Cleanup 接收一个函数，这个函数会在当前测试函数执行完毕后被调用。
	// 即使测试失败或 panic，它也保证会被执行。
	// 可以注册多个 Cleanup 函数，它们会以后进先出（LIFO）的顺序执行。
	t.Cleanup(func() {
		t.Log("  【TestWithGosql】=> Teardown: 清理测试资源...")
		gosqlModule.Close()
	})

	t.Log("  【TestWithGosql】=> 开始执行测试逻辑...")

	t.Run("测试gosql用户入库", gosqlUserRepo)

	t.Log("  【TestWithGosql】=> 测试逻辑执行完毕。")
}

func gosqlUserRepo(t *testing.T) {
	userRepo := repository.GetUserRepository()
	count, err := userRepo.Count(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}
