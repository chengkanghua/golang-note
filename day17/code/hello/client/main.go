package main

import (
	"bufio"
	"context"
	"fmt"
	"hello_client/pb"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func doRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "阿平"

	md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := c.SayHello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("c.SayHello failed, err:%v\n", err)
	}
	log.Printf("got reply:%v\n", res.Reply)
}

// doServerStreamRPC 调用服务端流式的RPC
func doServerStreamRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "阿平"
	stream, err := c.ServerStreamHello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("c.SayHello failed, err:%v\n", err)
	}
	// 从stream中依次接收流式返回结果
	for {
		res, err := stream.Recv()
		if err == io.EOF { // io.EOF != nil
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv failed, err:%v\n", err)
		}
		log.Printf("got reply:%v\n", res.Reply) // 把收到的每一次响应打印出来
	}
}

func doClientStreamRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.ClientStreamHello(ctx)
	if err != nil {
		log.Fatalf("c.ClientStreamHello failed, err:%v\n", err)
	}
	// 客户端流式发送请求
	names := []string{
		"根正苗红",
		"三好学生",
		"阿平",
	}
	for _, name := range names {
		stream.Send(&pb.Request{Name: name})
	}
	// 发送结束之后，要告诉服务端并且要开始接收响应
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream.CloseAndRecv() failed, err:%v\n", err)
	}
	// 将响应结果打印出来
	log.Printf("got reply:%v\n", res.Reply)
}

// doBudiStreamRPC 调用双向流式RPC
func doBudiStreamRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	stream, err := c.BudiStreamHello(ctx)
	if err != nil {
		log.Fatalf("c.BudiStreamHello failed, err:%v\n", err)
	}
	waitChan := make(chan struct{})
	// 一边收服务端发来的响应
	go func() { // 开启一个单独的goroutine去接收消息
		defer func() {
			waitChan <- struct{}{}
		}()
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("stream.Recv failed, err:%v\n", err)
				return
			}
			// 将收到的响应打印出来
			log.Printf("AI：%v\n", in.Reply)
		}
	}()
	go func() {
		// 一边还要源源不断的发送请求数据
		// 要发送的数据是用户在终端输入的
		reader := bufio.NewReader(os.Stdin)
		for {
			c, _ := reader.ReadString('\n') // 没有输入hang住
			c = strings.TrimSpace(c)        // 去掉首尾的空格
			if len(c) == 0 {
				continue
			}
			if strings.ToUpper(c) == "QUIT" {
				break
			}
			// 把用户的输入内容发送给服务端
			stream.Send(&pb.Request{Name: c})
		}
		stream.CloseSend() // 关闭发送流
	}()
	<-waitChan
}
func main() {
	// 先查询一下服务注册中心，找到hello服务对应的地址
	// 1. 连上consul
	consul, err := api.NewClient(api.DefaultConfig())
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
		fmt.Printf("%v:%v\n", k, v)
		addr = fmt.Sprintf("%v:%v", v.Address, v.Port)
		fmt.Println(addr)
		if len(addr) > 0 {
			break
		}
	}
	
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed, err:%v\n", err)
	}
	defer conn.Close()
	// 创建rpc client
	client := pb.NewHelloClient(conn)
	log.Printf("client start")

	// 发送普通的RPC调用（一来一回）
	doRPC(client)

	// 调用服务端流式RPC
	// doServerStreamRPC(client)

	// 调用客户端流式RPC
	// doClientStreamRPC(client)

	// 调用双向流式RPC
	// doBudiStreamRPC(client)

}
