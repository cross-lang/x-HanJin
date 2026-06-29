// Package mysql 提供 MySQL 数据库的连接和初始化功能，
// 基于 GORM 实现 ORM 操作。
package mysql

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"x-HanJin/internal/config"
	"x-HanJin/pkg/log"
)

// MySQLdb 全局 MySQL 数据库实例
var MySQLdb *gorm.DB

// NewClient 创建并返回 MySQL 数据库连接。
// 参数 dbName 为目标数据库名称。
func NewClient(dbName string) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.MySQLUser, config.Cfg.MySQLPassword,
		config.Cfg.MySQLHost, config.Cfg.MySQLPort,
		dbName,
	)
	log.Info("正在连接 MySQL 数据库",
		zap.String("host", config.Cfg.MySQLHost),
		zap.Int("port", config.Cfg.MySQLPort),
		zap.String("database", dbName),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("连接 MySQL 数据库失败", zap.Error(err))
		return nil
	}
	log.Info("MySQL 数据库连接成功")
	return db
}

// Init 初始化默认 MySQL 数据库连接
func Init() {
	MySQLdb = NewClient(config.Cfg.MySQLDefaultDBName)
}
