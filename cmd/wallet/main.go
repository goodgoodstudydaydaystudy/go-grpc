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
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet"
)

func main() {
	log.Println("listening to: localhost:50052")
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("wallet server listen failed")
	}

	log.Printf("listen")

	s := grpc.NewServer(
		grpc.UnaryInterceptor(server.StatusCodeUnaryInterceptor),
	)

	walletServer, err := wallet.NewWalletServer()
	if err != nil {
		log.Println("main wallet.NewWalletServer failed: ", err)
	}

	pb.RegisterWalletServer(s, walletServer)

	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Println("failed to server: ", err)
		}
	}()

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

}
