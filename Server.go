package main

import (
	"context"
	//"database/sql"
	"fmt"
	pb "gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	Address = ":50051"
)

type ControlServer struct {

}

func (s *ControlServer) Pay(ctx context.Context, consumeReq *pb.ConsumeReq) (*pb.ConsumeResp, error){
	return &pb.ConsumeResp{OrderId: consumeReq.GetItemId()}, nil	// 返回Resp里的字段？
}

func main()  {
	lis, err := net.Listen("tcp", Address)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("listen succee")

	// 创建grpc服务器
	s := grpc.NewServer()

	// 注册ControlServer
	pb.RegisterControlServer(s, &ControlServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to server: %v", err)
	}
	fmt.Printf("creat grpc server succee")
}

//func ControlMySql()  {
//
//	// connect mysql
//	db, err := sql.Open("mysql", "root:284927463@/order_sql")
//	if err != nil{
//		log.Printf("failel to connect mysql %v\n", err)
//	}
//
//
//}
