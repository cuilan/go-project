package conf

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/beego/beego/v2/core/logs"
	"github.com/spf13/viper"
)

const (
	ConfigType  = "yaml"      // 配置文件格式
	DefaultPath = "./configs" // 默认配置路径
	DefaultName = "app"       // 默认配置名称
)

var (
	AppName  string
	Type     string
	Profile  string
	HttpPort string
	Mode     string

	LogEnable2File bool
	LogPath        string
	LogLevel       int
	LogMaxDays     int

	// ========================================

	// your custom config
)

func Unmarshal(confPath string) {
	UnmarshalProfile(confPath, "")
}

func UnmarshalProfile(confPath, profile string) {
	viper.SetConfigName(DefaultName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("read config failed: %v\n", err)
		panic(err)
	}

	// 全局配置
	AppName = viper.GetString("app.name")
	Type = viper.GetString("app.type")

	// 根据环境配置切换
	if profile == "" {
		Profile = viper.GetString("app.profile")
	} else {
		Profile = profile
	}

	if Profile != "" {
		env := DefaultName + "-" + Profile
		viper.SetConfigName(env)
		viper.SetConfigType(ConfigType)
		if confPath == "" {
			viper.AddConfigPath(DefaultPath)
		} else {
			viper.AddConfigPath(confPath)
		}
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("read config failed: %v\n", err)
			panic(err)
		}
	}

	// 以下均为环境配置
	HttpPort = viper.GetString("app.httpPort")
	Mode = viper.GetString("app.mode")
	LogEnable2File = viper.GetBool("log.enable2file")
	LogPath = viper.GetString("log.path")
	LogLevel = viper.GetInt("log.level")
	LogMaxDays = viper.GetInt("log.maxdays")

	setLoggerConfig()

	logs.Info("======== Application config ========>")
	logs.Info("app name: %s", AppName)
	logs.Info("app type: %s", Type)
	logs.Info("app profile: %s", Profile)
	logs.Info("http port: %v", HttpPort)
	logs.Info("app run mode: %s", Mode)
	logs.Info("-------------------------------------")
	logs.Info("log enable to file: %v", LogEnable2File)
	logs.Info("log path: %s", LogPath)
	logs.Info("log level: %d", LogLevel)
	logs.Info("log max days: %d", LogMaxDays)
	logs.Info("-------------------------------------")

	unmarshalByType()
	logs.Info("<======= Application config =========")
}

// unmarshalByType 根据类型解析配置文件
func unmarshalByType() {

}

func setLoggerConfig() {
	logConf := make(map[string]interface{})
	logConf["filename"] = path.Join(LogPath, AppName+".log")
	logConf["level"] = LogLevel     // 日志保存的时候的级别，默认是 Trace 级别
	logConf["daily"] = true         // 是否按照每天 logrotate，默认是 true
	logConf["maxdays"] = LogMaxDays // 文件最多保存多少天，默认保存 7 天

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	_ = logs.SetLogger(logs.AdapterConsole)
	if LogEnable2File {
		_ = logs.SetLogger(logs.AdapterFile, string(confStr))
	}
	// 开启输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}
