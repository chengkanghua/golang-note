package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Args 定义一个Args参数类型
type Args struct {
	X, Y int
}

// 实现RPC 跨程序调用 -> 不在同一个内存空间

func main() {
	// 1.建立HTTP连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 2. 同步调用
	// 2.1 定义参数
	args := &Args{10, 20}
	var reply int
	// 2.2 直接调用server端的方法
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)
}
