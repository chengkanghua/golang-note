package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"add_client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 1 建立链接
	conn, err := grpc.Dial(":8973", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Dial failed,err:%v\n", err)
	}
	defer conn.Close()
	// 2 发起调用
	client := proto.NewCalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// res, err := client.Do(ctx, &proto.Req{X: 10, Y: 20, Op: proto.Op_ADD})
	res, err := client.Do(ctx, &proto.Req{X: 10, Y: 20, Op: proto.Op_SUB})
	// res, err := client.Do(ctx, &proto.Req{X: 10, Y: 20})
	if err != nil {
		log.Printf("client.Do failed ,err:%v \n", err)
		return
	}

	//打印结果
	fmt.Println(res.Sum)
}
