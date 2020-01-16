package main

import (
	"goodgoodstudy.com/go-grpc/pkg/DB"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/Account"
	"goodgoodstudy.com/go-grpc/pkg/pb/Pay"
	"goodgoodstudy.com/go-grpc/server"
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
	Pay.RegisterControlServer(s, &server.ControlServer{})
	rpb.RegisterAccountServer(s, &server.Server{})

	// 注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
	err = DB.InitDB()
	if err != nil {
		log.Println("InitDB failed: ", err)
	}
}
