package main

import (
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"goodgoodstudy.com/go-grpc/pkg/server/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	PORT_LOGIN = ":50051"
)

func main() {
	log.Println("listening to:", PORT_LOGIN)
	lis, err := net.Listen("tcp", PORT_LOGIN)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen")

	// 创建grpc服务器
	s := grpc.NewServer()

	// 注册ControlServer
	rpb.RegisterAccountServer(s, &account.Server{})

	// 注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
