package logger

// ConsoleConfig 对应配置文件中 "log.console" 的部分
type ConsoleConfig struct {
	EnableConsole bool   `mapstructure:"enable_console"` // 是否开启控制台输出
	Level         string `mapstructure:"level"`          // 控制台日志级别
	AddSource     bool   `mapstructure:"add_source"`     // 是否在日志中记录源码位置
	ConsoleFormat string `mapstructure:"console_format"` // 控制台输出格式: "text" 或 "json"
}

// FileConfig 对应配置文件中 "log.file" 的部分
type FileConfig struct {
	EnableFile     bool   `mapstructure:"enable_file"`      // 是否开启文件输出
	Level          string `mapstructure:"level"`            // 文件日志级别
	Path           string `mapstructure:"path"`             // 日志文件路径
	Filename       string `mapstructure:"filename"`         // 日志文件名
	AddSource      bool   `mapstructure:"add_source"`       // 是否在日志中记录源码位置
	FileFormat     string `mapstructure:"file_format"`      // 文件输出格式: "text" 或 "json"
	FileMaxSize    int    `mapstructure:"file_max_size"`    // 文件最大体积 (MB)
	FileMaxBackups int    `mapstructure:"file_max_backups"` // 最多保留的旧日志文件数量
	FileMaxAge     int    `mapstructure:"file_max_age"`     // 旧日志文件保留天数
	FileCompress   bool   `mapstructure:"file_compress"`    // 是否压缩旧的日志文件
}

// LoggerConfig 对应配置文件中的 "log" 根部分
type LoggerConfig struct {
	Console ConsoleConfig `mapstructure:"console"`
	File    FileConfig    `mapstructure:"file"`
}

// 空实现，用于手动配置日志模块，禁止自动配置
