package main

import (
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// 记录日志
	var uid int64 = 18967553
	isLogin := true
	name := "杨俊"
	data := []int{1, 2}

	logger.Debug("this is debug level log")

	logger.Info(
		"log info",
		zap.Int64("uid", uid),
		zap.Bool("isLogin", isLogin),
		zap.String("name", name),
		zap.Ints("data", data),
		zap.Any("data", data),
	)

}
