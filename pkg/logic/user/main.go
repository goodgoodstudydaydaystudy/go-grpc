package main

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server/grpcAuth"
	logicSvr "goodgoodstudy.com/go-grpc/pkg/logic/user/server"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
)

const portUserLogic = ":50053"

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func main() {
	// 和普通server 的main函数一样
	log.Println("listen to", portUserLogic)
	lis, err := net.Listen("tcp", portUserLogic)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen")

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpcAuth.UnaryServerInterceptor(
				grpcAuth.NewAuthFuncBuilder().WithFullMethodException("/user.User/Login").BuildJWT()),
			server.StatusCodeUnaryInterceptor,
		),
	)
	//s := grpc.NewServer(
	//	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//		server.StatusCodeUnaryInterceptor,
	//		server1.LogicReqUnaryInterceptor,
	//	),
	//	))

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
