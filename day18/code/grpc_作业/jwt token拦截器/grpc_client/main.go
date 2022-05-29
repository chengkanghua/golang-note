package main

import (
	"context"
	"grpc_client/pb"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const TokenExpireDuration = time.Hour * 2

// consul
type consul struct {
	consul *api.Client
}

type MyCustomClaims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte("夏天夏天悄悄过去")

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

// GenToken 生成JWT
func GenToken(uid string) (string, error) {
	// 创建一个我们自己的声明
	claims := MyCustomClaims{
		uid, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    "grpcJWT",                                  // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySigningKey)
}

// unaryInterceptor 一元拦截器
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	token, err := GenToken("uid") // uid可以通过(在服务端写死了)
	if err != nil {
		return err
	}
	md := metadata.Pairs("authorization", token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	err = invoker(ctx, method, req, reply, cc, opts...) // 此时才真的调用程序。类似gin的next()
	return err
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

	conn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryInterceptor),
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
