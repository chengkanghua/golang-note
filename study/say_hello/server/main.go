package main

import (
	"context"
	"hello_server/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServer
}

// func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
// 	return &pb.Response{
// 		Reply: "hello" + in.Name,
// 	}, nil
// }

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Reply: "hello " + in.Name,
	}, nil
}

func (s *server) ServerStreamHello(in *pb.Request, stream pb.Hello_ServerStreamHelloServer) error {

	name := in.Name

	words := []string{
		"你好",
		"helllo",
		"こんにちは",
		"여보세요",
	}
	for _, word := range words {
		if err := stream.Send(&pb.Response{Reply: word + name}); err != nil {
			log.Printf("stream.Send failed ,err:%v \n", err)
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8974")
	if err != nil {
		log.Fatalf("net.Listem failed , err :%v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	log.Println("server start ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("s.Serve failed,err:%v\n", err)
	}
}
