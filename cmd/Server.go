package main

import (
	"fmt"
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
	"goodgoodstudy.com/go-grpc/pkg/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	Address = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("listen succee")

	// 创建grpc服务器
	s := grpc.NewServer()

	// 注册ControlServer
	pb.RegisterControlServer(s, &server.ControlServer{})
	// 注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
	fmt.Printf("creat grpc server succee")
}
