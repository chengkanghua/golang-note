package main

import (
	"context"
	"fmt"
	"grpc_server/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// grpc server端

// server 定义服务
type server struct {
	pb.UnimplementedGreeterServer // 之前的pb版本没有这个
}

// SayHello 定义方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 业务逻辑
	var answer string
	if in.Name == "杨俊" {
		answer = "今天翘课了，小wb！"
	} else {
		answer = fmt.Sprintf("好好上课，加油！%s是最棒的！", in.Name)
	}
	return &pb.HelloReply{
		Answer: answer,
		Ts:     timestamppb.Now(),
	}, nil
}

func main() {
	// 向gRPC注册我们的服务
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务

	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
