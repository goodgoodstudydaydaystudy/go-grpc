package client

import (
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	port = ":50051"
)

type ConsumeClient struct {
}

func NewConsumeClient() (*ConsumeClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	defer conn.Close()
	client := pb.NewControlClient(conn)
	// 这个client应该在cmd/client调用才是
	// 怎么让cmd/client调用……

	return &ConsumeClient{}, nil // 这个函数要负责什么做什么 TODO
}

func (c *ConsumeClient) Pay(req *pb.ConsumeReq) (resp *pb.ConsumeResp, err error) {
	// TODO
	// 这里的Pay方法应该写什么
	// 怎么联系服务器哦
	// 想有一个葫芦，可以画一下瓢，啊哈哈哈哈哈。
	return
}
