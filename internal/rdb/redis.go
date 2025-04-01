package rdb

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var Cli *redis.Client

// ConnectRedis connect to redis, no password set, use default DB 0.
func ConnectRedis(host string, port int) {
	url := host + ":" + strconv.Itoa(port)
	logs.Info("connect to redis: %s", url)
	Cli = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
}

func Set(key, value string) error {
	_, err := Cli.Set(context.Background(), key, value, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) string {
	return Cli.Get(context.Background(), key).Val()
}

func Del(key string) error {
	_, err := Cli.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}
