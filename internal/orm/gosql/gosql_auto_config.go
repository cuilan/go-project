package gosql

import (
	"database/sql"
	"go-project/internal/module"
	"go-project/internal/utils"
	"log/slog"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/spf13/viper"
)

// GoSqlConfig 是 gosql 的配置结构体
type GoSqlConfig struct {
	Driver          string `yaml:"driver"`
	Dsn             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time"`
}

// gosqlModule 是 gosql 模块的实现
type gosqlModule struct{}

// NewModule 创建一个新的 gosql 模块实例
func NewModule() module.Module {
	return &gosqlModule{}
}

// Name 返回模块名称，与配置文件中的键 "gosql" 对应
func (m *gosqlModule) Name() string {
	return "gosql"
}

// Init 初始化 database/sql 连接池
func (m *gosqlModule) Init(v *viper.Viper) error {
	if db != nil {
		slog.Info("gosql database already initialized")
		return nil
	}

	var cfg GoSqlConfig
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

	maskedDsn, maskErr := utils.MaskDsn(cfg.Driver, cfg.Dsn)
	if maskErr != nil {
		slog.Error("failed to mask dsn", "err", maskErr)
		return maskErr
	}
	slog.Debug("gosql config loaded", "driver", cfg.Driver, "dsn", maskedDsn)

	// 打开数据库连接
	var err error
	db, err = sql.Open(cfg.Driver, cfg.Dsn)
	if err != nil {
		slog.Error("failed to open database", "driver", cfg.Driver, "err", err)
		return err
	}

	// 配置连接池
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)

	// 测试连接
	if err := db.Ping(); err != nil {
		slog.Error("failed to ping database", "err", err)
		return err
	}

	slog.Info("gosql database connected successfully",
		"driver", cfg.Driver,
		"max_open_conns", cfg.MaxOpenConns,
		"max_idle_conns", cfg.MaxIdleConns)
	return nil
}

// Close 关闭数据库连接
func (m *gosqlModule) Close() error {
	if db != nil {
		if err := db.Close(); err != nil {
			slog.Error("failed to close gosql database", "err", err)
			return err
		}
		slog.Info("gosql database closed successfully")
	}
	return nil
}

// db 全局数据库连接实例
var db *sql.DB

// GetDB 获取数据库连接实例
func GetDB() *sql.DB {
	if db == nil {
		panic("gosql database is not initialized")
	}
	return db
}
