/**
 * Author: Orange
 * Date: 20-02-04
 */
package main

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server"
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server/grpcAuth"
	logicSvr "goodgoodstudy.com/go-grpc/pkg/logic/auth/server"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/auth"
)

const portAuthLogic = ":50153"

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func main() {
	log.Println("listen to", portAuthLogic)
	lis, err := net.Listen("tcp", portAuthLogic)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening")

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpcAuth.UnaryServerInterceptor(
				grpcAuth.NewAuthFuncBuilder().WithFullMethodException("/user.User/Login").BuildJWT()),
			server.StatusCodeUnaryInterceptor,
		),
	)

	authLogic, err := logicSvr.NewAuthLogic()
	if err != nil {
		log.Println("main wallet.NewWalletServer failed: ", err)
	}

	pb.RegisterAuthServer(s, authLogic)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Println("failed to Serve: ", err)
	}

}
