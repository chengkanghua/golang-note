package dao

import (
	"app_demo/logger"

	"go.uber.org/zap"
)

func f1() {
	logger.Logger.Info("这是一条日志")
	zap.L().Info("程序启动,InitLogger() success")
}
