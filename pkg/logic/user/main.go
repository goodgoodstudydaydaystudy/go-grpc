package main

import (
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	logicSvr "goodgoodstudy.com/go-grpc/pkg/logic/user/server"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)


func main() {
	// 和普通server 的main函数一样
	log.Println("listen to localhost:50053")
	lis, err := net.Listen("tcp", "localhost:50053")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen")

	s := grpc.NewServer(
		grpc.UnaryInterceptor(server.StatusCodeUnaryInterceptor))

	user, err := logicSvr.NewUserLogic()
	if err != nil {
		log.Println("main wallet.NewWalletServer failed: ", err)
	}

	pb.RegisterUserServer(s, user)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Println("failed to Serve: ", err)
	}

}
