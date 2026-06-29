// Package config 提供应用程序配置管理功能。
// 基于 viper 库实现 YAML 配置文件的加载与解析，
// 支持 MySQL、Redis、Elasticsearch、RabbitMQ 等多种中间件的配置。
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoggerConfig 日志模块配置
type LoggerConfig struct {
	LogDir       string `yaml:"LogDir"`       // 日志文件存储目录
	Level        string `yaml:"Level"`         // 日志级别: debug/info/warn/error
	EnableRemote bool   `yaml:"EnableRemote"`  // 是否启用远程日志推送
	RemoteURL    string `yaml:"RemoteURL"`     // 远程日志服务地址
}

// Config 应用全局配置结构体
type Config struct {
	// Web 服务配置
	WebHost string // Web 服务监听地址
	WebPort int    // Web 服务监听端口

	// MySQL 数据库配置
	MySQLHost          string // MySQL 主机地址
	MySQLPort          int    // MySQL 端口
	MySQLUser          string // MySQL 用户名
	MySQLPassword      string // MySQL 密码
	MySQLDefaultDBName string // MySQL 默认数据库名称

	// RabbitMQ 消息队列配置
	RabbitMQHost             string // RabbitMQ 主机地址
	RabbitMQPort             int    // RabbitMQ 端口
	RabbitMQUser             string // RabbitMQ 用户名
	RabbitMQPassword         string // RabbitMQ 密码
	RabbitMQDefaultQueueName string // RabbitMQ 默认队列名称

	// Redis 缓存配置
	RedisHost      string // Redis 主机地址
	RedisPort      int    // Redis 端口
	RedisDefaultDB int    // Redis 默认数据库编号

	// Elasticsearch 配置
	ESAddress  string // Elasticsearch 服务地址
	ESUser     string // Elasticsearch 用户名
	ESPassword string // Elasticsearch 密码

	// App 应用配置
	AppId  string // 应用 ID
	AppKey string // 应用密钥

	// Logger 日志配置
	Logger LoggerConfig
}

// Cfg 全局配置实例
var Cfg Config

// Init 初始化配置，从 configs/config.yaml 读取配置项。
// 返回 error 以便调用方决定是否终止程序。
func Init() error {
	viper.SetConfigName("config")     // 配置文件名称（不带扩展名）
	viper.SetConfigType("yaml")       // 配置文件类型
	viper.AddConfigPath("./configs")  // 优先查找 configs 目录
	viper.AddConfigPath("./config")   // 兼容旧路径

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 映射配置到结构体
	Cfg = Config{
		// Web 服务
		WebHost: viper.GetString("Web.host"),
		WebPort: viper.GetInt("Web.port"),

		// MySQL
		MySQLHost:          viper.GetString("MySQL.host"),
		MySQLPort:          viper.GetInt("MySQL.port"),
		MySQLUser:          viper.GetString("MySQL.user"),
		MySQLPassword:      viper.GetString("MySQL.password"),
		MySQLDefaultDBName: viper.GetString("MySQL.default_dbname"),

		// RabbitMQ
		RabbitMQHost:             viper.GetString("RabbitMQ.host"),
		RabbitMQPort:             viper.GetInt("RabbitMQ.port"),
		RabbitMQUser:             viper.GetString("RabbitMQ.user"),
		RabbitMQPassword:         viper.GetString("RabbitMQ.password"),
		RabbitMQDefaultQueueName: viper.GetString("RabbitMQ.default_queue_name"),

		// Redis
		RedisHost:      viper.GetString("Redis.host"),
		RedisPort:      viper.GetInt("Redis.port"),
		RedisDefaultDB: viper.GetInt("Redis.default_db"),

		// Elasticsearch
		ESAddress:  viper.GetString("ES.address"),
		ESUser:     viper.GetString("ES.user"),
		ESPassword: viper.GetString("ES.password"),

		// App
		AppId:  viper.GetString("App.app_id"),
		AppKey: viper.GetString("App.app_key"),

		// Logger
		Logger: LoggerConfig{
			LogDir:       viper.GetString("Logger.LogDir"),
			Level:        viper.GetString("Logger.Level"),
			EnableRemote: viper.GetBool("Logger.EnableRemote"),
			RemoteURL:    viper.GetString("Logger.RemoteURL"),
		},
	}

	return nil
}
