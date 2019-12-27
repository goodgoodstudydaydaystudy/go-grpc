package main

import (
	"context"
	pb "gRPC/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	port = ":50051"
)

func main()  {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Printf("Connect succee\n")
	defer conn.Close()

	// 初始化客户端
	client := pb.NewControlClient(conn)

	ConsumeResp, err := client.Pay(context.Background(), &pb.ConsumeReq{
		ItemId: 1,		// 怎么在键盘输入呢	TODO
	})
	if err != nil {
		log.Fatalf("Consume not succee %v", err)
	}
	log.Printf("Pay succee %v", ConsumeResp)
}
