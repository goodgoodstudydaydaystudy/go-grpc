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
	return &ConsumeClient{}, nil // TODO
}

// 获取input数据
func (c *ConsumeClient) Pay(req *pb.ConsumeReq) (resp *pb.ConsumeResp, err error) {
	// TODO
	data := &pb.ConsumeReq{ ItemId:req.ItemId,
							ItemNum:req.ItemNum,
							UserId:req.UserId,
							Description:req.Description,
	}
	c.stub.Pay(context.Background(), data)	// TODO 这里有个错误，但不知道原因
	return &pb.ConsumeResp{Message:"success"}, nil		// TODO 但这不是和服务器的返回一样了吗，而且还有orderid…
}
