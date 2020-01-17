package main

import (
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	pb  "goodgoodstudy.com/go-grpc/pkg/pb/pay"
	"goodgoodstudy.com/go-grpc/pkg/server/account"
	"goodgoodstudy.com/go-grpc/pkg/server/pay"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	Address = ":50051"
)

func main() {
	log.Println("listening to:", Address)
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen")

	// 创建grpc服务器
	s := grpc.NewServer()

	// 注册ControlServer
	pb.RegisterControlServer(s, &pay.ControlServer{})
	rpb.RegisterAccountServer(s, &account.Server{})

	// 注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
