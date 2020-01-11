package client

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	port = ":50051"
)

type ConsumeClient struct {
	stub pb.ControlClient
}

func NewConsumeClient() (*ConsumeClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	defer conn.Close()

	client := &ConsumeClient{}

	client.stub = pb.NewControlClient(conn)
	return client, nil
}

// 发送 购买订单 数据
func (c *ConsumeClient) Pay(req *pb.ConsumeReq) (resp *pb.ConsumeResp, err error) {
	PayRep, err := c.stub.Pay(context.Background(), req)
	if err != nil {
		log.Println("stub.Pay failed:", err)
	}
	return PayRep, nil
}