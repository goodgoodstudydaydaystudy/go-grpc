package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"goodgoodstudy.com/go-grpc/pkg/server/account"
)

const (
	port = ":50051"
)



func main() {
	log.Println("listening to:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen")

	// 创建grpc服务器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(server.StatusCodeUnaryInterceptor), // 拦截器
	)

	// 注册ControlServer
	accountServer, err := account.NewAccountServer()
	rpb.RegisterAccountServer(s, accountServer) // TODO 修改了account_server的自定义返回后，发现会与返回值有冲突

	if err != nil {
		log.Println("account main open db failed")
		return
	}

	// 注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}