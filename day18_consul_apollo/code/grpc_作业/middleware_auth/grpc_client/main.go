package main

import (
	"context"
	"grpc_client/pb"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// consul
type consul struct {
	consul *api.Client
}

// Token token认证
type Token struct {
	Value string
}

const headerAuthorize string = "authorization"

func doRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "阿平"

	// 创建metadata和context(也可以传递一个ctx).
	md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	ctx = metadata.NewOutgoingContext(ctx, md)

	// 使用带有metadata的context执行RPC调用
	var header, trailer metadata.MD

	res, err := c.SayHello(ctx, &pb.Request{Name: name}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("c.SayHello failed, err:%v\n", err)
	}

	if t, ok := header["timestamp"]; ok {
		log.Printf("timestamp from header:\n")
		for i, e := range t {
			log.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}
	if l, ok := header["location"]; ok {
		log.Printf("location from header:\n")
		for i, e := range l {
			log.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("location expected but doesn't exist in header")
	}
	log.Printf("got reply:%v\n", res.Reply)
}

// middleware auth 方式1:调用auth包
// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{headerAuthorize: t.Value}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return false
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

// ListService 服务发现
func (c *consul) ListService() (map[string]*api.AgentService, error) {
	return c.consul.Agent().ServicesWithFilter("Service==`hello`")
}

// FindTarget 返回一个服务地址
func FindTarget(mc map[string]*api.AgentService) string {
	for _, consulService := range mc {
		addr := consulService.Address
		port := consulService.Port
		if addr != "" && port != 0 {
			return addr + ":" + strconv.Itoa(port)
		}
	}
	return ""
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

	// consul 服务发现
	// 1.连接consul
	addr := viper.GetString("consul.addr")
	consul, err := NewConsul(addr)
	if err != nil {
		log.Fatalf("NewConsul failed, err: %#v", err)
	}

	// 2.服务发现
	mc, err := consul.ListService()
	if err != nil {
		log.Fatalf("ListService failed, err: %#v", err)
	}

	target := FindTarget(mc) // 获取一个服务地址

	/*
		// 从输入的证书文件中为客户端构造TLS凭证
		creds, err := credentials.NewClientTLSFromFile("../tls/server.pem", "go-grpc-example")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
	*/

	// 使用方式1需要解开
	// 构建Token
	token := Token{
		Value: "bearer grpc.auth.token",
	}

	conn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&token), // 方式1:auth包的方法
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// rpc调用客户端
	c := pb.NewHelloClient(conn)

	// 发送普通的RPC调用(一来一回)
	doRPC(c)
}
