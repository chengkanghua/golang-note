package main

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zap demo

// 要用zap来记录日志，就需要生成Logger对象

var logger *zap.Logger

// Info 有的公司或有的同学不直接使用logger对象，而是在其外部包装一层
func Info(msg string) {
	logger.Info(msg, zap.Int64("timestamp", time.Now().Unix()))
}

// initLogger 把日志初始化操作包装到一个函数中
func initLogger() {
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
	// 文件和终端都写日志
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWS, consoleWS), level)
	// 利用core生成logger
	logger = zap.New(core, zap.AddCaller())

	// 替换zap全局的logger
	// zap.ReplaceGlobals(logger)
	// zap.L().Info("这是一条日志")
}

// 思考题
// 有的公司会要求把warn和err级别的日志单独记录到一个日志文件中，比如 app.wf.log
// 或者单独把err日志再记录到一个名为 app.err.log（err）中
// 所有日志还记录到 app.log（info、warning、err全量日志）

func main() {

	// initLg()
	// lg.Info("这是一条日志")

	// 获取Logger对象
	// zap.NewExample()
	// zap.NewDevelopment()
	// 生产环境的对象
	// logger, err := zap.NewProduction()  // 默认会在日志中记录调用信息
	// if err != nil {
	// 	panic(err)
	// }

	// 生成定制化的zap日志对象
	// 三部分
	// 1. encoder
	// zapcore.NewConsoleEncoder()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"                          // 时间的key,默认是ts
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 日期格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // level 大写
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	// 2. writesyncer
	// file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// fileWS := zapcore.AddSync(file)
	// consoleWS := zapcore.AddSync(os.Stdout)
	// 思考：如何将日志输出到多个目的地
	// 既把日志输出到文件又把日志输出到终端
	// 3. logLevel
	// zapcore.InfoLevel.Enabled()
	// s := "info" // 假设从配置文件读取了一个日志级别;作业
	// level, err := zapcore.ParseLevel(s) // 从字符串中解析
	// if err != nil {
	// 	level = zapcore.InfoLevel
	// }

	// 创建zapcore
	// core := zapcore.NewCore(encoder, fileWS, zapcore.DebugLevel)
	// core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWS, consoleWS), level)
	// // 利用core生成logger
	// logger = zap.New(core, zap.AddCaller())
	// // logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// 当我们要输出的日志 目的地不同且级别也不同，就需要新建Core
	// 往app.log里面写
	// logFile1, _ := os.Create("./app.log") // io.Writer
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./app.log",
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	core1 := zapcore.NewCore(encoder, zapcore.AddSync(lumberJackLogger), zapcore.InfoLevel)
	// 往app.err.log里面写
	logFile2, _ := os.Create("./app.err.log")
	core2 := zapcore.NewCore(encoder, zapcore.AddSync(logFile2), zapcore.ErrorLevel)
	// 将core1和core2合并成一个新的newCore
	newCore := zapcore.NewTee(core1, core2)
	// 使用新的newCore去创建logger实例
	logger = zap.New(newCore, zap.AddCaller())

	// 记录日志
	var uid int64 = 18967553
	isLogin := true
	name := "杨俊"
	data := []int{1, 2}

	logger.Debug("这是一条debug级别的日志")

	// 默认是输出JSON格式，日志会输出到 标准输出（终端）
	logger.Info(
		"日志信息",
		zap.Int64("uid", uid),
		zap.Bool("isLogin", isLogin),
		zap.String("name", name),
		// zap.Any("data", data),  // 万能的Any
		zap.Ints("data", data),
	)
	logger.Warn("这是一条warn级别的日志")
	// 1. app.log 里面要有这条err日志
	// 2. app.err.log 里面也要有这条err日志
	logger.Error("这是一条err日志", zap.Int64("uid", uid))

	// 使用包装后的logger记录日志
	Info("包装后的日志信息")

	// sLogger := logger.Sugar()
	// // 跟其他日志框架差不多的用法
	// sLogger.Info("sugar logger记录日志", uid, isLogin, name)
	// sLogger.Warnf("sugar logger记录日志, uid:%v isLogin:%v name:%v", uid, isLogin, name)
}
