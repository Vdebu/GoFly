package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	// 1.日志格式
	// 2.输出位置
	// 3.输出级别信息
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	// 创建多个输出目标防止在终端中看不见日志信息
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), os.Stdout), logMode)
	// 输出使用的日志工具
	return zap.New(core,
		zap.AddCaller(), // 添加调用者信息
		// zap.AddCallerSkip(1) 跳过包装函数输出主要goroutine信息
	).Sugar()
}

// 获取日志格式
func getEncoder() zapcore.Encoder {
	// 获取默认的config
	encoderConfig := zap.NewProductionEncoderConfig()
	// 进行格式上的修改(小写time)
	encoderConfig.TimeKey = "time"
	// 大写形式输出阶级(注意这里不是Color不然识别不出来会乱码)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 自定义时间初始化形式
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		// 2006-01-02 15:04:05
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	// 将输出格式设置为JSON返回
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getWriteSyncer() zapcore.WriteSyncer {
	// 告诉到底写到哪里去
	// 写到当前项目log文件夹的文件中

	// 存储路径分隔符
	stSeparator := string(filepath.Separator)
	// 获取当前工作目录
	stRootDir, _ := os.Getwd()
	// 结合生成输出位置
	// /log/时间.txt
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".txt"
	fmt.Println(stLogFilePath)

	// 使用三方包对输出的日志文件进行规范化(实现了io.Writer可以直接给zap使用)
	lumberjackSyncer := &lumberjack.Logger{
		// 使用viper获取配置信息
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"), // megabytes
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("MaxAge"), //days
		Compress:   false,                  // disabled by default
	}
	// 在分割器的基础上添加分割的输出目标
	return zapcore.AddSync(lumberjackSyncer)
}
