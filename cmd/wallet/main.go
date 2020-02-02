package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet"
)

const portWallet = ":50051"

func main() {
	log.Println("listening to: ", portWallet)
	lis, err := net.Listen("tcp", portWallet)
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
	if err := s.Serve(lis); err != nil {
		log.Println("failed to server: ", err)
	}
}
