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

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 业务逻辑
	var answer string
	if in.Name == "杨俊" {
		answer = "今天翘课了，小wb"
	} else {
		answer = fmt.Sprintf("好好上课%s", in.Name)
	}
	return &pb.HelloReply{
		Answer: "hello" + answer,
		Ts:     timestamppb.Now(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve :%v", err)
		return
	}
}
