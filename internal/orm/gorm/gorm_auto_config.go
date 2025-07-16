package gorm

import (
	"fmt"
	"go-project/internal/orm/models"
	"go-project/internal/orm/repository"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// GormConfig - GORM 配置结构体
type GormConfig struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
	Config struct {
		TablePrefix                              string `yaml:"table_prefix"`
		DisableForeignKeyConstraintWhenMigrating bool   `yaml:"disable_foreign_key_constraint_when_migrating"`
		LogLevel                                 string `yaml:"log_level"`
		SlowThresholdMs                          int    `yaml:"slow_threshold_ms"`
	} `yaml:"config"`
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
		return nil, fmt.Errorf("unsupported database driver: %s", c.Driver)
	}

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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(c.Config.SlowThresholdMs) * time.Millisecond,
			LogLevel:                  level,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Config.TablePrefix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: c.Config.DisableForeignKeyConstraintWhenMigrating,
		Logger:                                   newLogger,
	})

	if err != nil {
		return nil, err
	}

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ProvideGormUserRepository - 提供 GORM 用户仓储的 wire provider
func ProvideGormUserRepository(db *gorm.DB) repository.IUserRepository {
	return NewGormUserRepository(db)
}
