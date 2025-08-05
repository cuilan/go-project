package rdb

import (
	"context"
	"fmt"
	"go-project/internal/rdb"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
)

// TestMain 会在包中所有其他测试运行之前被调用。
func TestMain(m *testing.M) {
	// --- 全局 Setup ---
	// 这部分代码是测试前的准备工作，只会执行一次。
	fmt.Println("【TestMain】=> 全局 Setup: 建立连接，准备测试环境...")

	viper.SetConfigFile("../test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 手动初始化 redis 模块
	redisModule := rdb.NewModule()
	redisModule.Init(viper.GetViper())

	exitCode := m.Run()

	// --- 全局 Teardown ---
	// 这部分代码是测试后的清理工作，只会执行一次。
	fmt.Println("【TestMain】=> 全局 Teardown: 断开连接，清理测试环境...")
	// 手动关闭 redis 模块
	redisModule.Close()

	// 退出测试，并将 m.Run() 的结果作为退出码。
	os.Exit(exitCode)
}

func TestRedis(t *testing.T) {
	redisClient := rdb.GetRedis()
	ctx := context.Background()

	statusCmd := redisClient.Set(ctx, "name", "zhangyan", 0*time.Second)
	t.Log("redis set", statusCmd.Val())

	value := redisClient.Get(ctx, "name")
	t.Log("redis get value:", value.Val())

	intCmd := redisClient.Del(ctx, "name")
	t.Log("redis del", intCmd.Val())
}
