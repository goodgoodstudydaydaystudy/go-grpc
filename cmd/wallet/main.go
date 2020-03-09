package main

import (
	"log"

	"goodgoodstudy.com/go-grpc/pkg/foundation/base"
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet"

	"google.golang.org/grpc"
)

type basePara struct {
	netType 	string
	address 	string
	grpcServer  grpc.ServerOption
}

func main() {
	log.Println("listening to: localhost:50052")
	b := basePara{
		netType:     "tcp",
		address:     "localhost:50052",
		grpcServer:   grpc.UnaryInterceptor(server.StatusCodeUnaryInterceptor),
	}
	cmdBase := base.NewCmdBase(b.grpcServer, b.netType, b.address)

	walletServer, err := wallet.NewWalletServer()
	log.Println("walletServer", walletServer)
	if err != nil {
		log.Println("main wallet.NewWalletServer failed: ", err)
	}

	pb.RegisterWalletServer(cmdBase.GrpcServer, walletServer)
	go func() {
		if err := cmdBase.GrpcServer.Serve(cmdBase.Listen); err != nil {
			log.Println("failed to server: ", err)
		}
	}()

	log.Printf("waiting exit sig")
	<-cmdBase.Signal
	log.Println("bye")
}
