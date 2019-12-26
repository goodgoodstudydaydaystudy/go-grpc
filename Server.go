package main

import (
		"fmt"
		"google.golang.org/grpc/reflection"
		"log"
		"net"
		"context"
		"google.golang.org/grpc"
		pb "gRPC/pb"
)

const (
	Address = "127.0.0.1:50051"
)

type SaveServer struct {

}

func (s *SaveServer) Pay(ctx context.Context, in *pb.ConsumeReq) (*pb.ConsumeResp, error){
	return &pb.ConsumeResp{Message:"消费成功"}, nil
}

//var SaveService = SaveServer{}

func main()  {
	lis, err := net.Listen("TCP", Address)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建grpc服务器
	s := grpc.NewServer()
	pb.RegisterSaveServer(s, &SaveServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to server: %v", err)
	}
}

