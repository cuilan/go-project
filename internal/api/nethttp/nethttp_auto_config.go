package nethttp

import (
	"errors"
	"go-project/internal/module"
	"log/slog"

	"github.com/spf13/viper"
)

// NetHttpConfig 是 net/http 的配置
type NetHttpConfig struct {
	Port           int    `mapstructure:"port"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
	EnableTLS      bool   `mapstructure:"enable_tls"`
	TLSCertFile    string `mapstructure:"tls_cert_file"`
	TLSKeyFile     string `mapstructure:"tls_key_file"`
}

// nethttpModule 是 net/http 模块的实现
type nethttpModule struct {
	server *Server
}

func NewModule() module.Module {
	return &nethttpModule{}
}

// Name 与配置文件中的 'http' key 保持一致
func (m *nethttpModule) Name() string {
	return "http"
}

func (m *nethttpModule) Init(v *viper.Viper) error {
	if !v.IsSet(m.Name()) {
		slog.Warn("http config not found, skipping nethttp module.")
		return nil
	}

	var cfg NetHttpConfig
	if err := v.Sub(m.Name()).Unmarshal(&cfg); err != nil {
		return err
	}

	// 设置默认值
	if cfg.Port == 0 {
		cfg.Port = 8081
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

	m.server = NewServer(&cfg)
	m.server.Start()

	slog.Info("nethttp module initialized successfully")
	return nil
}

func (m *nethttpModule) Close() error {
	if m.server != nil {
		return m.server.Stop()
	}
	return nil
}
