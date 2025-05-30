package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"vdebu.gofly.net/conf"
)

// 全局日志组件
var (
	Logger *zap.SugaredLogger
	DB     *gorm.DB // gorm维护了连接池不需要关闭操作直接放全局用
	RDB    *conf.RedisClient
)
