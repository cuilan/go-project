package rdb

import (
	"context"
	"go-project/internal/module"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func setup(t *testing.T) {
	viper.SetConfigFile("../../configs/examples/redis.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitRedis(t *testing.T) {
	setup(t)
	module.InitModules(viper.GetViper())

	redisClient := GetRedis()
	redisClient.Set(context.Background(), "name", "zhangyan", 0*time.Second)

	value := redisClient.Get(context.Background(), "name")
	t.Log(value)

	redisClient.Del(context.Background(), "name")

	module.CloseModules()
}
