package main

import (
	"add_server/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedCalServer
}

func (s *server) Do(ctx context.Context, in *proto.Req) (*proto.Res, error) {
	var res int64
	switch in.Op {
	case proto.Op_ADD:
		res = in.X + in.Y
	case proto.Op_SUB:
		res = in.X - in.Y
	default:
		res = 0
	}
	return &proto.Res{Sum: res}, nil
}

func main() {
	// 1 启动tcp服务
	lis, err := net.Listen("tcp", ":8973")
	if err != nil {
		log.Fatalf("net.Listen failed,err:%v\n", err)
	}
	// 2 注册RPC服务
	s := grpc.NewServer()
	proto.RegisterCalServer(s, &server{})
	// 3 启动RPC服务
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("s.Serve failed, err:%v\n", err)
	}
}
