package main

import (
	"context"
	pb "goStudy/go_x_rpc/grpc/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	conent, _ := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conent.Close()
	client := pb.NewSayHelloClient(conent)
	log.Println("client start")
	hello, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "king"})
	//
	log.Println("从服务端拿到的返回结果：", hello.GetMessage())

}
