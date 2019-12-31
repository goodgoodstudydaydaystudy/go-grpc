package client

import (
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
)

const (
	port = ":50051"
)

type ConsumeClient struct {
}

func NewConsumeClient() (*ConsumeClient, error) {
	return &ConsumeClient{}, nil // TODO
}

func (c *ConsumeClient) Pay(req *pb.ConsumeReq) (resp *pb.ConsumeResp, err error) {
	// TODO
	return
}
