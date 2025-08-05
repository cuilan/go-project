package gin

import (
	"errors"
	"go-project/internal/module"
	"log/slog"

	"github.com/spf13/viper"
)

// GinHttpConfig 是 gin 的配置
type GinHttpConfig struct {
	Mode           string `mapstructure:"mode"`
	Port           int    `mapstructure:"port"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
	EnableTLS      bool   `mapstructure:"enable_tls"`
	TLSCertFile    string `mapstructure:"tls_cert_file"`
	TLSKeyFile     string `mapstructure:"tls_key_file"`
}

// ginHttpModule 是 gin 模块的实现
type ginHttpModule struct {
	server *Server
}

func NewModule() module.Module {
	return &ginHttpModule{}
}

// Name 与配置文件中的 'gin' key 保持一致
func (m *ginHttpModule) Name() string {
	return "gin"
}

// Init 初始化 gin 模块
func (m *ginHttpModule) Init(v *viper.Viper) error {
	if !v.IsSet(m.Name()) {
		slog.Warn("gin config not found, skipping gin module.")
		return nil
	}

	var cfg GinHttpConfig
	if err := v.Sub(m.Name()).Unmarshal(&cfg); err != nil {
		return err
	}

	// 设置默认值
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = 10
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = 10
	}
	if cfg.MaxHeaderBytes == 0 {
		cfg.MaxHeaderBytes = 1024
	}

	if cfg.EnableTLS {
		if cfg.TLSCertFile == "" || cfg.TLSKeyFile == "" {
			slog.Error("TLS is enabled, but tls_cert_file or tls_key_file is not provided")
			return errors.New("TLS is enabled, but tls_cert_file or tls_key_file is not provided")
		}
	}

	server, err := NewServer(&cfg)
	if err != nil {
		slog.Error("gin module new server failed", "err", err)
		return err
	}
	m.server = server
	m.server.Start()

	slog.Info("gin module initialized successfully")
	return nil
}

func (m *ginHttpModule) Close() error {
	if m.server != nil {
		m.server.Shutdown()
	}
	return nil
}
