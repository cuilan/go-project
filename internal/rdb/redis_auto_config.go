package rdb

import (
	"context"
	"go-project/internal/module"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// RedisConfig 保存 redis 的配置
type RedisConfig struct {
	Network         string        `mapstructure:"network"`
	Addr            string        `mapstructure:"addr"`
	ClientName      string        `mapstructure:"client_name"`
	Password        string        `mapstructure:"password"`
	DB              int           `mapstructure:"db"`
	PoolSize        int           `mapstructure:"pool_size"`
	MinIdleConns    int           `mapstructure:"min_idle_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	PoolTimeout     time.Duration `mapstructure:"pool_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

// redisModule 是 redis 模块的实现
type redisModule struct{}

// NewModule 创建一个新的 redis 模块实例
func NewModule() module.Module {
	return &redisModule{}
}

// Name 返回模块名称，与配置文件中的键 "redis" 对应
func (m *redisModule) Name() string {
	return "redis"
}

// Init 初始化 redis 客户端
func (m *redisModule) Init(v *viper.Viper) error {
	if redisClient != nil {
		slog.Info("redis client already initialized")
		return nil
	}

	var cfg RedisConfig
	// v.Sub(m.Name()) 获取此模块专属的配置树
	if err := v.Sub(m.Name()).Unmarshal(&cfg); err != nil {
		return err
	}
	slog.Debug("redis config loaded", "addr", cfg.Addr, "pool_size", cfg.PoolSize)

	// 初始化 redis 客户端
	redisClient = redis.NewClient(&redis.Options{
		Network:    cfg.Network,
		Addr:       cfg.Addr,
		ClientName: cfg.ClientName,
		Password:   cfg.Password,
		DB:         cfg.DB,
		PoolSize:   cfg.PoolSize,
	})

	// 使用 Ping 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		slog.Error("ping redis failed", "err", err)
		return err
	}

	slog.Info("redis connected successfully")
	return nil
}

// Close 关闭 redis 客户端连接
func (m *redisModule) Close() error {
	if redisClient != nil {
		if err := redisClient.Close(); err != nil {
			slog.Error("failed to close redis client", "err", err)
			return err
		}
		slog.Info("redis client closed successfully")
	}
	return nil
}
