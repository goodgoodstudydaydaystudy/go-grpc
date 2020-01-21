package main

import (
	pb "goodgoodstudy.com/go-grpc/pkg/pb/pay"
	"goodgoodstudy.com/go-grpc/pkg/server/pay"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	PayPort = ":50051"
)

func main() {
	log.Println("listening to:", PayPort)
	lis, err := net.Listen("tcp", PayPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen")

	// 创建grpc服务器
	s := grpc.NewServer()

	// 注册ControlServer
	payServer, err := pay.NewConsumeServer()
	pb.RegisterControlServer(s, payServer)

	// 注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
