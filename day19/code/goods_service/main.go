package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"goods_service/config"
	"goods_service/dao/mysql"
	"goods_service/handler"
	"goods_service/logger"
	"goods_service/proto"
	"goods_service/registry"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // !!!!
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 可以把启动时的一些列操作放到 bootstap/init 包

	var cfn string
	// 0.从命令行获取可能的conf路径
	// goods_service -conf="./conf/config_qa.yaml"
	// goods_service -conf="./conf/config_online.yaml"
	flag.StringVar(&cfn, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()
	// 1. 加载配置文件
	err := config.Init(cfn)
	if err != nil {
		panic(err) // 程序启动时加载配置文件失败直接退出
	}
	// 2. 加载日志
	err = logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	if err != nil {
		panic(err) // 程序启动时初始化日志模块失败直接退出
	}
	// 3. 初始化MySQL
	err = mysql.Init(config.Conf.MySQLConfig)
	if err != nil {
		panic(err) // 程序启动时初始化MySQL失败直接退出
	}
	// 4. 初始化Consul
	err = registry.Init(config.Conf.ConsulConfig.Addr)
	if err != nil {
		panic(err) // 程序启动时初始化注册中心失败直接退出
	}

	// 监听端口
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Conf.IP, config.Conf.Port))
	if err != nil {
		panic(err)
	}
	// 创建gRPC服务
	s := grpc.NewServer()
	// 商品服务注册RPC服务
	proto.RegisterGoodsServer(s, &handler.GoodsSrv{})
	// 启动gRPC服务
	go func() {
		err = s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
	// 注册服务到consul
	registry.Reg.RegisterService(config.Conf.Name, config.Conf.IP, config.Conf.Port, nil)

	zap.L().Info("service start...")

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext( // RPC客户端
		context.Background(),
		fmt.Sprintf("%s:%d", config.Conf.IP, config.Conf.Port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = proto.RegisterGoodsHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: gwmux,
	}

	zap.L().Info("Serving gRPC-Gateway on http://0.0.0.0:8091")
	go func() {
		err := gwServer.ListenAndServe()
		if err != nil {
			log.Printf("gwServer.ListenAndServe failed, err: %v", err)
			return
		}
	}()

	// 服务退出时要注销服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit // 正常会hang在此处
	// 退出时注销服务
	serviceId := fmt.Sprintf("%s-%s-%d", config.Conf.Name, config.Conf.IP, config.Conf.Port)
	registry.Reg.Deregister(serviceId)
}
