package test

import (
	"context"
	"go-project/internal/orm/gorm"
	"go-project/internal/orm/repository"
	"testing"

	"github.com/spf13/viper"
)

func TestWithGorm(t *testing.T) {
	t.Log("  【TestWithGorm】=> Setup: 初始化配置文件...")

	viper.SetConfigFile("../test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}

	gormModule := gorm.NewModule()
	gormModule.Init(viper.GetViper())

	t.Cleanup(func() {
		t.Log("  【TestWithGorm】=> Teardown: 清理测试资源...")
		gormModule.Close()
	})

	t.Log("  【TestWithGorm】=> 开始执行测试逻辑...")

	t.Run("测试gorm用户入库", gormUserRepo)

	t.Log("  【TestWithGorm】=> 测试逻辑执行完毕。")
}

func gormUserRepo(t *testing.T) {
	userRepo := repository.GetUserRepository()
	count, err := userRepo.Count(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}
