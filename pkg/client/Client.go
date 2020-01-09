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

// 获取input数据
func (c *ConsumeClient) Pay(req *pb.ConsumeReq) (resp *pb.ConsumeResp, err error) {
	//req.ItemId = pb.ConsumeReq{}.ItemId
	//req.ItemNum = pb.ConsumeReq{}.ItemNum
	//req.UserId = pb.ConsumeReq{}.UserId
	//req.Description = pb.ConsumeReq{}.Description

	c.stub.Pay(context.Background(), &pb.ConsumeReq{ItemId: req.ItemId,
													ItemNum: req.ItemNum,
													UserId: req.UserId,
													Description: req.Description,
	})
	return , nil		// TODO 但这不是和服务器的返回一样了吗，而且还有orderid…
}
