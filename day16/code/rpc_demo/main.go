package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// rpc server

// ServiceA 定义一个服务A
type ServiceA struct{}

// Args 定义一个Args参数类型
type Args struct {
	X, Y int
}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	fmt.Println("有人调用我啦！！！")
	// reply：  0x78622000(内存地址)
	// *reply： 0（值）
	*reply = args.X + args.Y  // 30是一个具体的值
	return nil
}

func main() {
	service := new(ServiceA) // new函数返回对应类型的指针

	rpc.Register(service) // 注册RPC服务
	rpc.HandleHTTP()      // 基于HTTP协议

	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)

	// 基于HTTP协议启动了一个RPC Server

}
