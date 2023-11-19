package main

import (
	"context"
	pb "goStudy/go_x_rpc/grpc/server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type MyServer struct {
	pb.UnimplementedSayHelloServer
}

func (*MyServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResp, error) {

	log.Println("客户端的请求：", req.Name)
	// 提供出去的 api
	return &pb.HelloResp{
		Message: "hello " + req.Name,
		A:       1,
	}, nil

}

func main() {

	// 启动服务端
	listen, _ := net.Listen("tcp", ":9090")

	server := grpc.NewServer()
	pb.RegisterSayHelloServer(server, &MyServer{})
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
	log.Println("server start")

}
