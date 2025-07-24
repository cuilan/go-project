package gorm

import (
	"fmt"
	"go-project/internal/module"
	"go-project/internal/orm/models"
	"go-project/internal/utils"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// GormConfig - GORM 配置结构体
type GormConfig struct {
	Driver string `mapstructure:"driver"`
	Dsn    string `mapstructure:"dsn"`
	// 连接池配置
	MaxOpenConns    int `mapstructure:"max_open_conns"`
	MaxIdleConns    int `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime int `mapstructure:"conn_max_idle_time"`
	// GORM 特定配置
	Config struct {
		AutoMigrate                              bool   `mapstructure:"auto_migrate"`
		TablePrefix                              string `mapstructure:"table_prefix"`
		DisableForeignKeyConstraintWhenMigrating bool   `mapstructure:"disable_foreign_key_constraint_when_migrating"`
		EnableLogger                             bool   `mapstructure:"enable_logger"`
		EnableSlog                               bool   `mapstructure:"enable_slog"`
		LogLevel                                 string `mapstructure:"log_level"`
		SlowThresholdMs                          int    `mapstructure:"slow_threshold_ms"`
	} `mapstructure:"config"`
}

// gormModule 是 gorm 模块的实现
type gormModule struct{}

// NewModule 创建一个新的 gorm 模块实例
func NewModule() module.Module {
	return &gormModule{}
}

// Name 返回模块名称，与配置文件中的键 "gorm" 对应
func (m *gormModule) Name() string {
	return "gorm"
}

// Init 初始化 GORM 数据库连接并注册仓储服务
func (m *gormModule) Init(v *viper.Viper) error {
	var cfg GormConfig
	// v.Sub(m.Name()) 获取此模块专属的配置树
	if err := v.Sub(m.Name()).Unmarshal(&cfg); err != nil {
		return err
	}

	// 设置默认值
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 20
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 10
	}
	if cfg.ConnMaxLifetime == 0 {
		cfg.ConnMaxLifetime = 3600
	}
	if cfg.ConnMaxIdleTime == 0 {
		cfg.ConnMaxIdleTime = 3600
	}
	if cfg.Config.SlowThresholdMs == 0 {
		cfg.Config.SlowThresholdMs = 200
	}
	if cfg.Config.LogLevel == "" {
		cfg.Config.LogLevel = "info"
	}

	// 初始化数据库连接
	db, err := NewGormDB(&cfg)
	if err != nil {
		return err
	}

	// 自动注入repository到容器
	autowired(db)

	return nil
}

// Close 关闭数据库连接
func (m *gormModule) Close() error {
	// GORM 会自动管理连接池，这里可以添加清理逻辑
	return nil
}

// NewGormDB - 根据配置初始化 GORM
func NewGormDB(c *GormConfig) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch c.Driver {
	case "mysql":
		dialector = mysql.Open(c.Dsn)
	case "postgres":
		dialector = postgres.Open(c.Dsn)
	case "sqlite":
		dialector = sqlite.Open(c.Dsn)
	default:
		slog.Error("unsupported database driver", "driver", c.Driver)
		return nil, fmt.Errorf("unsupported database driver: %s", c.Driver)
	}
	maskedDsn, maskErr := utils.MaskDsn(c.Driver, c.Dsn)
	if maskErr != nil {
		slog.Error("failed to mask dsn", "err", maskErr)
		return nil, maskErr
	}
	slog.Debug("gorm config loaded", "driver", c.Driver, "dsn", maskedDsn)

	// 配置日志
	var gormLogger logger.Interface

	// 如果禁用日志，直接使用静默模式
	if !c.Config.EnableLogger {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		// 如果启用日志，根据log_level配置
		logLevelMap := map[string]logger.LogLevel{
			"silent": logger.Silent,
			"info":   logger.Info,
			"warn":   logger.Warn,
			"error":  logger.Error,
		}
		level, ok := logLevelMap[c.Config.LogLevel]
		if !ok {
			level = logger.Info
		}

		// 根据配置选择日志器类型
		if c.Config.EnableSlog {
			// 使用slog整合的日志器
			gormLogger = &slogLogger{level: level}
		} else {
			// 使用标准GORM日志器
			gormLogger = logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Duration(c.Config.SlowThresholdMs) * time.Millisecond,
					LogLevel:                  level,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			)
		}
	}

	// 配置 GORM
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Config.TablePrefix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: c.Config.DisableForeignKeyConstraintWhenMigrating,
		Logger:                                   gormLogger,
	}

	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return nil, err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Second)

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		slog.Error("failed to ping database", "err", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// 自动迁移模型（如果启用）
	if c.Config.AutoMigrate {
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			slog.Error("failed to auto migrate", "err", err)
			return nil, fmt.Errorf("failed to auto migrate: %w", err)
		}
		slog.Info("auto migrate success")
	}

	slog.Info("gorm database connected successfully",
		"driver", c.Driver,
		"max_open_conns", c.MaxOpenConns,
		"max_idle_conns", c.MaxIdleConns,
		"conn_max_lifetime", c.ConnMaxLifetime,
		"conn_max_idle_time", c.ConnMaxIdleTime)

	return db, nil
}
