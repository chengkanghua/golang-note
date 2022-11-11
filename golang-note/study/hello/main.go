package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	// register2Consul()
	// 先查询一下服务注册中心，找到hello服务对应的地址
	// 1. 连上consul
	cfg := api.DefaultConfig()
	cfg.Address = "10.211.55.6:8500"
	consul, err := api.NewClient(cfg)
	if err != nil {
		log.Fatalf(" api.NewClient failed, err:%v\n", err)
	}
	// 2. 查询 hello 服务的所有可用地址
	m, err := consul.Agent().ServicesWithFilter("Service==`hello`")
	if err != nil {
		log.Fatalf(" api.NewClient failed, err:%v\n", err)
	}
	fmt.Printf("%#v\n", m)
	var addr string
	for k, v := range m {
		fmt.Printf("%v:%v ---\n", k, v)
		addr = fmt.Sprintf("%v:%v \n", v.Address, v.Port)
		fmt.Println(addr)
		if len(addr) > 0 {
			break
		}
	}

}
