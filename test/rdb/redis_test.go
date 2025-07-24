package rdb

import (
	"context"
	"go-project/internal/rdb"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func setup(t *testing.T) {
	viper.SetConfigFile("../test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitRedis(t *testing.T) {
	setup(t)
	// 手动初始化 redis 模块
	redisModule := rdb.NewModule()
	redisModule.Init(viper.GetViper())

	redisClient := rdb.GetRedis()
	ctx := context.Background()

	statusCmd := redisClient.Set(ctx, "name", "zhangyan", 0*time.Second)
	t.Log("redis set", statusCmd.Val())

	value := redisClient.Get(ctx, "name")
	t.Log("redis get value:", value.Val())

	intCmd := redisClient.Del(ctx, "name")
	t.Log("redis del", intCmd.Val())

	// 手动关闭 redis 模块
	redisModule.Close()
}
