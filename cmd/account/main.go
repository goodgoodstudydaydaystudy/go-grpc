package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	"goodgoodstudy.com/go-grpc/pkg/server/account"
)



func main() {
	log.Println("listening to : localhost:50051")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen")

	// 创建gRPC服务器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(server.StatusCodeUnaryInterceptor), // 拦截器
	)

	// 注册ControlServer
	accountServer, err := account.NewAccountServer() // 通过server结构体私有化禁止外部随意访问，指定接口访问内部变量
	rpb.RegisterAccountServer(s, accountServer)
	if err != nil {
		log.Println(err)
	}

	exit := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-exit
		fmt.Printf("server exit")
		done <-true
	}()

	log.Printf("waiting exit sig")
	<-done

	// 注册反射服务
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server: %v", err)
		}
	}()
}
