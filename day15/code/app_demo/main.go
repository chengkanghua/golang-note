package main

import (
	"app_demo/logger"

	"go.uber.org/zap"
)

var lg *zap.Logger

func main() {
	// 程序启动之后
	// 1. 加载配置
	// 2. 加载日志模块
	logger.InitLogger()
	lg.Info("xxx")
	logger.Logger.Info("程序启动,InitLogger() success")

	// 第2种方法
	zap.L().Info("程序启动,InitLogger() success")

	// 第三种方法,将初始化后的logger赋值给全局的lg
	lg = logger.InitLogger3()
	lg.Info("程序启动,InitLogger() success")
}
