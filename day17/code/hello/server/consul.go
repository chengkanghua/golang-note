package main

// 基于consul实现服务注册

import (
	"log"

	"github.com/hashicorp/consul/api"
)

// register2Consul 注册服务到consul
func register2Consul() {
	// 1. 连上consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("api.NewClient failed, err:%v\n", err)
	}
	// 2. 服务注册()
	srv := &api.AgentServiceRegistration{
		ID:      "hello-127.0.0.1-8974", // 服务名称-ip-端口
		Name:    "hello",
		Tags:    []string{"Beijing-hello", "hello", "q1mi"},
		Address: "127.0.0.1",
		Port:    8974,
	}
	client.Agent().ServiceRegister(srv)
}
