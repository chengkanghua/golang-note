package main

import (
	"context"
	"log"
	"time"

	"client_with_lb/pb"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // 匿名导入
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type consul struct {
	consul *api.Client
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

func doRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "阿平"

	res, err := c.SayHello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("c.SayHello failed, err:%v\n", err)
	}

	log.Printf("got reply:%v\n", res.Reply)
}
func main() {
	// 海棠代码的原流程
	// 1. 连接上consul
	// 2. 查询 hello 服务下所有的ip
	// 3. 自己从ip列表中找到一个server去连接 （手动选）
	// 4. 发送RPC请求

	// 1. 连接上consul
	// 2. 根据负载均衡配置自动选服务（自动选）
	// 3. 发起RPC请求
	conn, err := grpc.Dial(
		// "127.0.0.1:8999",
		// consul服务
		"consul://192.168.1.11:8500/hello?healthy=true&wait=14s",
		// 指定round_robin策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	// rpc调用客户端
	c := pb.NewHelloClient(conn)

	// 发送普通的RPC调用(一来一回)
	for i := 0; i < 10; i++ {
		doRPC(c)
	}
}
