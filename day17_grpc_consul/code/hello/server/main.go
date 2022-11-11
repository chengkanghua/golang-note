package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"hello_server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedHelloServer
}

// var _ pb.HelloServer = &server{}
var _ pb.HelloServer = (*server)(nil)

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	// 从客户端读取metadata.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
	}
	if t, ok := md["timestamp"]; ok {
		fmt.Printf("timestamp from metadata:\n")
		// []string
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	}
	return &pb.Response{
		Reply: "hello " + in.Name,
	}, nil
}

func (s *server) ServerStreamHello(in *pb.Request, stream pb.Hello_ServerStreamHelloServer) error {
	// 对传进来的人打招呼
	name := in.Name
	// 四国语言打招呼
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"여보세요",
	}

	for _, word := range words {
		// 通过stream写入打招呼的内容
		if err := stream.Send(&pb.Response{Reply: word + name}); err != nil {
			log.Printf("stream.Send failed, err:%v\n", err)
			return err
		}
	}
	return nil
}

// ClientStreamHello 客户端发送流式数据
func (s *server) ClientStreamHello(stream pb.Hello_ClientStreamHelloServer) error {
	// 接收流式发来的请求数据
	var reply string = "你好啊 "
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// 读完了要给客户端返回一个响应
			return stream.SendAndClose(&pb.Response{Reply: reply})
		}
		if err != nil { // 读取请求数据遇到其他错误
			return err
		}
		reply += res.Name
	}
}

// BudiStreamHello 双向流式RPC
func (s *server) BudiStreamHello(stream pb.Hello_BudiStreamHelloServer) error {
	// 服务端 收一条 回复一条
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Printf("stream.Recv failed, err:%v\n", err)
			return err
		}
		// 处理收到的数据拿到要返回的数据
		reply := magic(res.Name)
		// 返回响应
		if err := stream.Send(&pb.Response{Reply: reply}); err != nil {
			log.Printf("stream.Send failed, err:%v\n", err)
			return err
		}
	}
}

// magic 一段价值连城的“人工智能”代码
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

func main() {
	lis, err := net.Listen("tcp", ":8974")
	if err != nil {
		log.Fatalf("net.Listen failed, err:%v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	// 程序启动之前 要注册服务到consul
	register2Consul()

	log.Println("server start...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("s.Serve failed, err:%v\n", err)
	}

}
