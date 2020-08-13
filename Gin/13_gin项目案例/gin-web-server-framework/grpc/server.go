package grpc

import (
	pb "gin-demo/grpc/protoes"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//通过一个结构体，实现proto中定义的所有服务
type Hello struct{}

func (h Hello) Serve1(ctx context.Context, in *pb.Name) (*pb.Msg1, error) {
	log.Println("serve 1 works: get name: ", in.Name)
	resp := &pb.Msg1{Message: "this is serve 1"}
	return resp, nil
}

func (h Hello) Serve2(ctx context.Context, in *pb.Name) (*pb.Msg2, error) {
	log.Println("serve 2 works, get name: ", in.Name)
	resp := &pb.Msg2{
		Message: &pb.Msg1{Message: "this is serve 2"},
	}
	return resp, nil
}

func server() {

	// Address gRPC服务地址
	listen, err := net.Listen("tcp", "127.0.0.1:3001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// 与http的注册路由类似，此处将所有服务注册到grpc服务器上，
	pb.RegisterServeRouteServer(s, Hello{})

	log.Println("grpc serve running")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
