package main

import (
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/"
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
	s := grpc.NewServer(
		grpc.UnaryInterceptor(server.StatusCodeUnaryInterceptor), // 拦截器
		)

	// 注册ConsumeServer
	payServer, err := pay.NewConsumeServer()
	pb.RegisterConsumeServer(s, payServer)

	// 注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
