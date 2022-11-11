package main

import (
	"context"
	"errors"
	"grpc_server/pb"
	"log"
	"net"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const TokenExpireDuration = time.Hour * 2

// server grpc server
type server struct {
	pb.UnimplementedHelloServer
}

// consul
type consul struct {
	consul *api.Client
}

type MyCustomClaims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte("夏天夏天悄悄过去")

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

	return &pb.Response{Reply: "Hello " + in.Name}, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// valid 校验认证信息.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}

	// HTTP header 
	// authorization: Bearer saasdasda.adsadasdasda.daqwfggod

	// TrimPrefix。去掉字符串首 `Bearer ` 字符
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	mc, err := ParseToken(token)
	log.Printf("uid: %s", mc.UID)
	if err != nil || mc.UID != "uid" {
		return false
	}
	return true
}

// 拦截器
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "缺少metadata")
	}
	// 对请求的metadata中携带的认证信息做校验
	if !valid(md["authorization"]) {
		return nil, status.Error(codes.Unauthenticated, "认证失败")
	}
	// 继续执行后续的处理
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error %v\n", err)
	}
	return m, err
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

// RegisterService 将服务注册到consul
func (c *consul) RegisterService() error {
	/* check := &api.AgentServiceCheck{
		GRPC: "192.168.1.33:8999",
		Timeout: "10s",
		Interval: "10s",
		DeregisterCriticalServiceAfter: "20s",
	} */
	address := viper.GetString("grpc_service.address")
	port := viper.GetInt("grpc_service.port")
	srv := &api.AgentServiceRegistration{
		ID:      "hello-127.0.0.1-8999",
		Name:    "hello",
		Tags:    []string{"GuangZhou", "hello", "joil"},
		Address: address,
		Port:    port,
		// Check: check,
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

func main() {
	if err := initViper("./conf/config.yaml"); err != nil {
		log.Fatalf("initViper failed, err: %#v", err)
	}
	viper.WatchConfig()

	// 监听端口
	port := viper.GetString("grpc_service.port")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Listen faild err:%#v", err)
	}

	// 1.连接consul
	addr := viper.GetString("consul.addr")
	c, err := NewConsul(addr)
	if err != nil {
		log.Fatalf("Consul api.NewClient failed, err:%#v\n", err)
	}

	// 2.把服务注册到consul
	err = c.RegisterService()
	if err != nil {
		log.Fatalf("RegisterService failed, err: %#v", err)
	}

	// 创建gRPC服务器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	// 在gRPC服务端注册服务
	pb.RegisterHelloServer(s, &server{})
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		c.DeregisterService("hello-127.0.0.1-8999") // 从consul中注销服务
		log.Fatalf("gRPC Serve failed err:%#v", err)
	}
}
