package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger 把日志初始化操作包装到一个函数中
func InitLogger() {
	// 1. encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"                          // 时间的key,默认是ts
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 日期格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // level 大写
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	// 2. writesyncer
	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fileWS := zapcore.AddSync(file)
	consoleWS := zapcore.AddSync(os.Stdout)
	// 思考：如何将日志输出到多个目的地
	// 既把日志输出到文件又把日志输出到终端
	// 3. logLevel
	// zapcore.InfoLevel.Enabled()
	s := "info" // 假设从配置文件读取了一个日志级别;作业
	level, err := zapcore.ParseLevel(s)
	if err != nil {
		level = zapcore.InfoLevel
	}
	// 创建zapcore
	// core := zapcore.NewCore(encoder, fileWS, zapcore.DebugLevel)
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWS, consoleWS), level)
	// 利用core生成logger 赋值给全局的Logger
	Logger = zap.New(core, zap.AddCaller())
}

// InitLogger 把日志初始化操作包装到一个函数中
func InitLogger2() {
	// 1. encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"                          // 时间的key,默认是ts
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 日期格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // level 大写
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	// 2. writesyncer
	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fileWS := zapcore.AddSync(file)
	consoleWS := zapcore.AddSync(os.Stdout)
	// 思考：如何将日志输出到多个目的地
	// 既把日志输出到文件又把日志输出到终端
	// 3. logLevel
	// zapcore.InfoLevel.Enabled()
	s := "info" // 假设从配置文件读取了一个日志级别;作业
	level, err := zapcore.ParseLevel(s)
	if err != nil {
		level = zapcore.InfoLevel
	}
	// 创建zapcore
	// core := zapcore.NewCore(encoder, fileWS, zapcore.DebugLevel)
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWS, consoleWS), level)
	// 利用core生成logger
	logger := zap.New(core, zap.AddCaller())

	// 替换zap全局的logger
	zap.ReplaceGlobals(logger) // 将zap包中默认的logger替换为我们自己的logger
	zap.L().Info("这是一条日志")     // 通过 zap.L() 调用我们的logger
}

// InitLogger 把日志初始化操作包装到一个函数中
func InitLogger3() *zap.Logger {
	// 1. encoder
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"                          // 时间的key,默认是ts
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 日期格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // level 大写
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	// 2. writesyncer
	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fileWS := zapcore.AddSync(file)
	consoleWS := zapcore.AddSync(os.Stdout)
	// 思考：如何将日志输出到多个目的地
	// 既把日志输出到文件又把日志输出到终端
	// 3. logLevel
	// zapcore.InfoLevel.Enabled()
	s := "info" // 假设从配置文件读取了一个日志级别;作业
	level, err := zapcore.ParseLevel(s)
	if err != nil {
		level = zapcore.InfoLevel
	}
	// 创建zapcore
	// core := zapcore.NewCore(encoder, fileWS, zapcore.DebugLevel)
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWS, consoleWS), level)
	// 利用core生成logger
	return zap.New(core, zap.AddCaller())
}
