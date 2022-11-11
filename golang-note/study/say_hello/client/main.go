package main

import (
	"context"
	"hello_client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "阿平"

	// md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	// ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := c.SayHello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("c.SayHello failed, err:%v\n", err)
	}
	log.Printf("got reply:%v\n", res.Reply)
}

func main() {
	conn, err := grpc.Dial(":8974", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v\n", err)
	}
	defer conn.Close()
	// 创建rpc client
	client := pb.NewHelloClient(conn)

	// 发送正常的RPC 调用 一来一回
	doRPC(client)
}
