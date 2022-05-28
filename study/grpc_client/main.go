package main

import (
	"context"
	"flag"
	"grpc_client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

// grpc_client

const (
	defatultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", defatultName, "name to greet")
)

func main() {
	flag.Parse()

	// 连接服务器
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	// rpc调用客户端
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用SayHello方法
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s %v\n", r.GetAnswer(), r.GetTs().AsTime().Format("2006-01-02 15:04:05"))
}
