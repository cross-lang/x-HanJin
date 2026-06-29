// Package databases 提供数据库初始化的统一入口，
// 按顺序初始化所有配置中启用的数据库连接。
package databases

import (
	"x-HanJin/internal/databases/kaiwudb"
	"x-HanJin/internal/databases/mysql"
	"x-HanJin/internal/databases/postgresql"
	"x-HanJin/internal/databases/redis"
	"x-HanJin/internal/databases/tdengine"

	"x-HanJin/pkg/log"
	"go.uber.org/zap"
)

// Init 初始化所有数据库连接
func Init() {
	log.Info("开始初始化数据库...")

	mysql.Init()
	postgresql.Init()
	redis.Init()
	tdengine.Init()
	kaiwudb.Init()

	log.Info("数据库初始化完成", zap.Int("databases", 5))
}
