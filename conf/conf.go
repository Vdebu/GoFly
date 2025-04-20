package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 配置基础的读取信息
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")
	// 读取文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load config error: %s", err.Error()))
	}
	// 测试输出相关配置信息
	fmt.Println(viper.Get("server.port"))
}
