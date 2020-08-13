package grpc

import (
	pb "gin-demo/grpc/protoes"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func client() {

	conn, err := grpc.Dial("127.0.0.1:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewServeRouteClient(conn)

	reqBody1 := &pb.Name{Name: "wang"}
	res1, err := c.Serve1(context.Background(), reqBody1) //就像调用本地函数一样，通过serve1得到返回值
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message from serve: ", res1.Message)

	reqBody2 := &pb.Name{Name: "li"}
	res2, err := c.Serve2(context.Background(), reqBody2) //就像调用本地函数一样，通过serve2得到返回值
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("message from serve: ", res2.Message.Message)
}
