# 配置管理

## 📋 概述

本项目采用基于 Viper 的配置管理系统，支持多种配置格式（YAML、JSON、TOML等），支持环境变量覆盖、配置文件热重载等特性。

## 🏗️ 配置架构

### 配置层次结构

```
配置文件加载顺序：
1. 基础配置文件 (app.yaml)
2. 环境配置文件 (app-{profile}.yaml)
3. 环境变量覆盖
4. 命令行参数覆盖
```

### 配置目录结构

```
configs/
├── app.yaml                    # 基础配置文件
├── app-dev.yaml               # 开发环境配置
├── app-prod.yaml              # 生产环境配置
├── app-test.yaml              # 测试环境配置
└── examples/                  # 配置示例
    ├── gorm.yaml             # GORM 配置示例
    ├── gosql.yaml            # database/sql 配置示例
    ├── log.yaml              # 日志配置示例
    └── redis.yaml            # Redis 配置示例
```

## 📝 配置文件详解

### 1. 基础配置文件 (app.yaml)

```yaml
# 应用基础配置
app:
  name: "demo"
  profile: "dev"  # 激活的环境配置

# 日志配置
log:
  console:
    enable_console: true
    level: "debug"
    add_source: true
    console_format: "text"
  file:
    enable_file: false
    level: "info"
    path: "logs"
    filename: "default.log"
    add_source: true
    file_format: "json"
    file_max_size: 100
    file_max_backups: 3
    file_max_age: 7
    file_compress: false

# Redis 配置
redis:
  addr: "127.0.0.1:6379"
  pool_size: 10

# 数据库配置
gorm:
  driver: "mysql"
  dsn: "user:pass@tcp(localhost:3306)/dbname"
  max_open_conns: 20
  max_idle_conns: 10
  conn_max_lifetime: 3600
  conn_max_idle_time: 3600
  config:
    auto_migrate: false
    table_prefix: "t_"
    disable_foreign_key_constraint_when_migrating: true
    enable_logger: true
    enable_slog: true
    log_level: "info"
    slow_threshold_ms: 200

# HTTP 服务配置
gin:
  mode: "debug"
  port: 8080
  read_timeout: 10
  write_timeout: 10
  max_header_bytes: 1024
```

### 2. 环境配置文件

#### 开发环境 (app-dev.yaml)

```yaml
# 开发环境特定配置
log:
  console:
    level: "debug"
  file:
    enable_file: false

redis:
  addr: "localhost:6379"
  pool_size: 5

gorm:
  driver: "mysql"
  dsn: "root:123456@tcp(localhost:3306)/dev_db"
  config:
    enable_logger: true
    log_level: "info"
    auto_migrate: true
```

#### 生产环境 (app-prod.yaml)

```yaml
# 生产环境特定配置
log:
  console:
    level: "info"
  file:
    enable_file: true
    level: "warn"

redis:
  addr: "redis.prod.com:6379"
  pool_size: 20

gorm:
  driver: "mysql"
  dsn: "prod_user:prod_pass@tcp(prod_db.com:3306)/prod_db"
  config:
    enable_logger: false
    auto_migrate: false
```

## 🔧 配置项详解

### 应用配置 (app)

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `name` | string | "demo" | 应用名称 |
| `profile` | string | "dev" | 激活的环境配置 |

### 日志配置 (log)

#### 控制台输出 (console)

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `enable_console` | bool | true | 是否启用控制台输出 |
| `level` | string | "debug" | 日志级别: debug, info, warn, error |
| `add_source` | bool | true | 是否添加源码位置 |
| `console_format` | string | "text" | 输出格式: text, json |

#### 文件输出 (file)

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `enable_file` | bool | false | 是否启用文件输出 |
| `level` | string | "info" | 日志级别 |
| `path` | string | "logs" | 日志文件路径 |
| `filename` | string | "default.log" | 日志文件名 |
| `add_source` | bool | true | 是否添加源码位置 |
| `file_format` | string | "json" | 文件格式: text, json |
| `file_max_size` | int | 100 | 单个文件最大大小(MB) |
| `file_max_backups` | int | 3 | 最大备份文件数 |
| `file_max_age` | int | 7 | 文件保留天数 |
| `file_compress` | bool | false | 是否压缩备份文件 |

### Redis 配置 (redis)

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `network` | string | "tcp" | 网络类型 |
| `addr` | string | "127.0.0.1:6379" | Redis 地址 |
| `client_name` | string | "" | 客户端名称 |
| `password` | string | "" | 密码 |
| `db` | int | 0 | 数据库编号 |
| `pool_size` | int | 10 | 连接池大小 |
| `min_idle_conns` | int | 0 | 最小空闲连接数 |
| `max_idle_conns` | int | 10 | 最大空闲连接数 |
| `conn_max_idle_time` | duration | 30m | 连接最大空闲时间 |
| `conn_max_lifetime` | duration | 0 | 连接最大生命周期 |
| `pool_timeout` | duration | 30s | 连接池超时时间 |
| `read_timeout` | duration | 3s | 读取超时时间 |
| `write_timeout` | duration | 3s | 写入超时时间 |

### GORM 配置 (gorm)

#### 基础配置

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `driver` | string | "mysql" | 数据库驱动: mysql, postgres, sqlite |
| `dsn` | string | "" | 数据源名称 |
| `max_open_conns` | int | 20 | 最大打开连接数 |
| `max_idle_conns` | int | 10 | 最大空闲连接数 |
| `conn_max_lifetime` | int | 3600 | 连接最大生命周期(秒) |
| `conn_max_idle_time` | int | 3600 | 连接最大空闲时间(秒) |

#### GORM 特定配置 (config)

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `auto_migrate` | bool | false | 是否自动迁移表结构 |
| `table_prefix` | string | "" | 全局表前缀 |
| `disable_foreign_key_constraint_when_migrating` | bool | true | 禁用外键约束 |
| `enable_logger` | bool | true | 是否启用GORM日志 |
| `enable_slog` | bool | true | 是否整合到slog日志 |
| `log_level` | string | "info" | 日志级别: silent, info, warn, error |
| `slow_threshold_ms` | int | 200 | 慢查询阈值(毫秒) |

### database/sql 配置 (gosql)

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `driver` | string | "mysql" | 数据库驱动 |
| `dsn` | string | "" | 数据源名称 |
| `max_open_conns` | int | 20 | 最大打开连接数 |
| `max_idle_conns` | int | 10 | 最大空闲连接数 |
| `conn_max_lifetime` | int | 3600 | 连接最大生命周期(秒) |
| `conn_max_idle_time` | int | 3600 | 连接最大空闲时间(秒) |

### HTTP 服务配置 (gin)

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `mode` | string | "debug" | 运行模式: debug, release, test |
| `port` | int | 8080 | 服务端口 |
| `read_timeout` | int | 10 | 读取超时时间(秒) |
| `write_timeout` | int | 10 | 写入超时时间(秒) |
| `max_header_bytes` | int | 1024 | 最大请求头大小(字节) |

## 🔄 配置加载机制

### 1. 配置加载流程

```go
// 1. 加载基础配置文件
v := viper.New()
v.SetConfigName("app")
v.SetConfigType("yaml")
v.AddConfigPath("./configs")

// 2. 读取基础配置
if err := v.ReadInConfig(); err != nil {
    panic(fmt.Errorf("read base config file failed: %w", err))
}

// 3. 获取环境配置
profile := v.GetString("app.profile")

// 4. 加载环境配置文件
if profile != "" {
    v.SetConfigName("app-" + profile)
    if err := v.MergeInConfig(); err != nil {
        panic(fmt.Errorf("merge profile config file failed: %w", err))
    }
}
```

### 2. 环境变量支持

支持通过环境变量覆盖配置：

```bash
# 设置应用名称
export APP_NAME="my-app"

# 设置数据库连接
export GORM_DSN="user:pass@tcp(localhost:3306)/dbname"

# 设置Redis地址
export REDIS_ADDR="localhost:6379"
```

### 3. 命令行参数支持

```bash
# 指定配置文件目录
./bin/client --config-dir ./configs

# 指定环境配置
./bin/client --config-dir ./configs --profile prod

# 显示版本信息
./bin/client --version
```

## 🛡️ 配置安全

### 1. 敏感信息脱敏

```go
// 脱敏显示DSN
maskedDsn, maskErr := utils.MaskDsn(cfg.Driver, cfg.Dsn)
if maskErr != nil {
    slog.Error("failed to mask dsn", "err", maskErr)
    return maskErr
}
slog.Debug("config loaded", "driver", cfg.Driver, "dsn", maskedDsn)
```

### 2. 环境变量支持

```yaml
# 支持环境变量覆盖
gorm:
  dsn: "${GORM_DSN}"  # 从环境变量读取
  config:
    password: "${DB_PASSWORD}"  # 从环境变量读取密码
```

### 3. 配置文件权限

```bash
# 设置配置文件权限
chmod 600 configs/app-prod.yaml
chown app:app configs/app-prod.yaml
```

## 🔍 配置验证

### 1. 配置结构体验证

```go
type GormConfig struct {
    Driver string `mapstructure:"driver" validate:"required"`
    Dsn    string `mapstructure:"dsn" validate:"required"`
    Config struct {
        EnableLogger bool   `mapstructure:"enable_logger"`
        LogLevel     string `mapstructure:"log_level" validate:"oneof=silent info warn error"`
    } `mapstructure:"config"`
}
```

### 2. 运行时验证

```go
func (m *gormModule) Init(v *viper.Viper) error {
    var cfg GormConfig
    if err := v.Sub(m.Name()).Unmarshal(&cfg); err != nil {
        return err
    }
    
    // 验证配置
    if cfg.Driver == "" {
        return errors.New("database driver is required")
    }
    if cfg.Dsn == "" {
        return errors.New("database dsn is required")
    }
    
    return nil
}
```

## 📊 配置监控

### 1. 配置变更监控

```go
// 监听配置文件变更
v.WatchConfig()
v.OnConfigChange(func(e fsnotify.Event) {
    slog.Info("config file changed", "file", e.Name)
    // 重新加载配置
})
```

### 2. 配置热重载

```go
// 支持配置热重载
func reloadConfig() {
    // 重新加载配置
    // 重启相关模块
}
```

## 🧪 配置测试

### 1. 配置验证测试

```go
func TestGormConfig(t *testing.T) {
    config := `
gorm:
  driver: "mysql"
  dsn: "user:pass@tcp(localhost:3306)/test"
  config:
    enable_logger: true
    log_level: "info"
`
    
    v := viper.New()
    v.SetConfigType("yaml")
    v.ReadConfig([]byte(config))
    
    var cfg GormConfig
    err := v.Sub("gorm").Unmarshal(&cfg)
    assert.NoError(t, err)
    assert.Equal(t, "mysql", cfg.Driver)
}
```

### 2. 配置集成测试

```go
func TestConfigIntegration(t *testing.T) {
    // 测试完整配置加载流程
    // 测试环境变量覆盖
    // 测试配置文件合并
}
```

## 📚 最佳实践

### 1. 配置文件组织

- 按环境分离配置文件
- 使用有意义的配置项名称
- 提供配置示例和文档

### 2. 配置安全

- 敏感信息使用环境变量
- 配置文件权限控制
- 日志中脱敏显示

### 3. 配置验证

- 使用结构体标签验证
- 运行时配置验证
- 提供配置测试

### 4. 配置监控

- 配置文件变更监控
- 配置热重载支持
- 配置变更日志记录

---

通过这套配置管理系统，项目可以灵活地适应不同的部署环境和运行需求，同时保证了配置的安全性和可维护性。 