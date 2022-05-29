package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"grpc_server/pb"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	// 使用方式1时需要解开
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// server grpc server
type server struct {
	pb.UnimplementedHelloServer
}

// consul
type consul struct {
	consul *api.Client
}

// TokenInfo 用户信息
type TokenInfo struct {
	ID    string
	Roles []string
}

// var _ pb.HelloServer = &server{}
var _ pb.HelloServer = (*server)(nil)

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	// 从客户端读取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
	}
	if t, ok := md["timestamp"]; ok {
		log.Printf("timestamp from metadata:\n")

		for i, e := range t {
			log.Printf(" %d. %s\n", i, e)
		}
	}

	// 创建和发送header
	header := metadata.New(map[string]string{"location": "BeiJing", "timestamp": time.Now().Format("2006-01-02 15:04:05")})
	grpc.SetHeader(ctx, header)

	log.Printf("request received: %v, say hello...\n", in)
	// 返回数据的时候告诉对方来自哪个server(通过port区分不同的server)
	reply := fmt.Sprintf("hello %s from %d", in.Name, port)
	return &pb.Response{Reply: reply}, nil
}

// 方式2:调用auth包
// AuthInterceptor 认证拦截器，对以authorization为头部，形式为`bearer token`的Token进行验证
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}
	//使用context.WithValue添加了值后，可以用Value(key)方法获取值
	newCtx := context.WithValue(ctx, tokenInfo.ID, tokenInfo)
	//log.Println(newCtx.Value(tokenInfo.ID))
	return newCtx, nil
}

// 方式1：拦截器
func UnaryInterceptorfunc(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 对所有来自客户端的RPC请求都会做如下check,包括健康检查
	// 需要忽略掉健康检查那部分请求
	fmt.Println("--->", info.FullMethod)
	if info.FullMethod != "/grpc.health.v1.Health/Check" {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "缺少metadata")
		}
		// 对请求的metadata中携带的认证信息做校验
		if !valid(md["authorization"]) {
			return nil, status.Error(codes.Unauthenticated, "认证失败")
		}
	}
	// 继续执行后续的处理
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error %v\n", err)
	}
	return m, err
}

// parseToken 解析token，并进行验证
func parseToken(token string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	if token == "grpc.auth.token" {
		tokenInfo.ID = "1"
		tokenInfo.Roles = []string{"admin"}
		return tokenInfo, nil
	}
	return tokenInfo, errors.New("Token无效: bearer " + token)
}

// userClaimFromToken 从token中获取用户唯一标识
func userClaimFromToken(tokenInfo TokenInfo) string {
	return tokenInfo.ID
}

// NewConsul 连接至consul服务返回一个consul对象
func NewConsul(addr string) (*consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consul{c}, nil
}

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// RegisterService 将服务注册到consul
func (c *consul) RegisterService(port int) error {
	ip, err := GetOutboundIP()
	if err != nil {
		return err
	}
	fmt.Println("-->", ip)
	address := viper.GetString("grpc_service.address")
	// port := viper.GetInt("grpc_service.port")
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ip, port), // 这里一定要写外网地址
		Timeout:                        "10s",                          // 超时
		Interval:                       "10s",                          // 每隔10秒做一次健康检查
		DeregisterCriticalServiceAfter: "20s",                          // 注销不健康的服务
	}

	// 在向consul注册服务的时候告诉consul
	// 你每隔10s请求一下我的健康检查服务吧
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("hello-%s-%d", address, port),
		Name:    "hello",
		Tags:    []string{"GuangZhou", "hello", "joil"},
		Address: address,
		Port:    port,
		Check:   check,
	}
	return c.consul.Agent().ServiceRegister(srv)
}

// DeregisterService 将服务从consul中注销
func (c *consul) DeregisterService(consulID string) error {
	return c.consul.Agent().ServiceDeregister(consulID)
}

// initViper 初始化Viper
func initViper(path string) error {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

var port int

func main() {
	if err := initViper("./conf/config.yaml"); err != nil {
		log.Fatalf("initViper failed, err: %#v", err)
	}
	viper.WatchConfig()

	flag.IntVar(&port, "port", 8999, "端口号")
	flag.Parse()
	// 监听端口
	// port := viper.GetString("grpc_service.port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Listen faild err:%#v", err)
	}

	// 1.连接consul
	addr := viper.GetString("consul.addr")
	c, err := NewConsul(addr)
	if err != nil {
		log.Fatalf("Consul api.NewClient failed, err:%#v\n", err)
	}
	// 2.创建gRPC服务器
	s := grpc.NewServer(
	// 方法1：
	// grpc.UnaryInterceptor(UnaryInterceptorfunc),
	// 方法2：
	// grpc.UnaryInterceptor(
	// 	grpc_middleware.ChainUnaryServer(
	// 		grpc_auth.UnaryServerInterceptor(AuthInterceptor),
	// 	)),
	)
	// 在gRPC服务端注册服务
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(s, healthServer) // 注册健康检查的服务

	pb.RegisterHelloServer(s, &server{}) // 注册自己的hello服务
	// 3.启动gRPC服务
	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		err = s.Serve(lis) // for循环hang住。。。
		if err != nil {

			log.Printf("gRPC Serve failed err:%#v", err)
			return
		}
	}()
	fmt.Println("哈哈哈！愚蠢的人类！")
	// 程序启动之后再去注册服务
	// 4.把服务注册到consul
	err = c.RegisterService(port)
	if err != nil {
		log.Printf("RegisterService failed, err: %#v", err)
		return
	}
	// defer func() {
	// 	fmt.Println("嘿嘿嘿2")
	// 	c.DeregisterService("hello-127.0.0.1-8999") // 从consul中注销服务
	// }()
	// wg.Wait()
	// 程序退出之前去注销服务？
	// 实际生产环境下都是kill -9/Ctrl+C 直接关闭服务
	// 怎么样在代码里去处理这种退出程序的情形

	// 可以在代码里接收操作系统发来的中断信号
	quitChan := make(chan os.Signal)
	// syscall.SIGTERM(kill), syscall.SIGINT(Ctrl+C), syscall.SIGKILL(Kill -9)
	signal.Notify(quitChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-quitChan // hang住
	fmt.Println("嘿嘿嘿")
	c.DeregisterService("hello-127.0.0.1-8999") // 从consul中注销服务
}
