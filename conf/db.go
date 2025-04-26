package conf

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
	"vdebu.gofly.net/model"
)

func InitDB() (*gorm.DB, error) {
	// 定义日志输出级别
	logMode := logger.Info
	// 根据实际环境进行变更
	if !viper.GetBool("mode.develop") {
		// 在生产环境下就设置成error
		logMode = logger.Error
	}
	// 初始化数据库连接池
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 自定义生成表名的前缀
			TablePrefix:   "gofly_", // 设置表名前缀
			SingularTable: true,     // 使用单数表名 gofly_users -> gofly_user
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	// 获取默认数据库连接池并进行配置
	sqlDB, _ := db.DB()
	// 进行相应的修改
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	sqlDB.SetConnMaxIdleTime(time.Hour)
	// 将结构体自动迁移到数据库的表里
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	// 返回数据库链接
	return db, nil
}
