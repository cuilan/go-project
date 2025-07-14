package conf

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Profile string `mapstructure:"profile"`
}

var App AppConfig
