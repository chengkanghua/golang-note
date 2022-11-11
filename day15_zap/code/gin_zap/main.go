package main

import (
	"fmt"
	"os"

	"gin_zap/dao/mysql"
	"gin_zap/logger"
	"gin_zap/router"
	"gin_zap/setting"

	"go.uber.org/zap"
)

func main() {
	// bootstrap
	// 1. 加载配置
	if len(os.Args) < 2 {
		panic("程序执行时必须通过命令行指定配置文件")
	}
	err := setting.Init(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(setting.Conf.Name)

	// 2. 初始化日志模块
	err = logger.Init()
	if err != nil {
		panic(err)
	}
	defer zap.L().Sync()
	// defer logger.Logger.Sync()

	// 3. 数据库初始化
	err = mysql.Init()
	if err != nil {
		zap.L().Error("mysql.Init failed", zap.Error(err))
	}

	// 4. 路由初始化
	r := router.Setup()

	// 5. 程序启动
	r.Run()
}
