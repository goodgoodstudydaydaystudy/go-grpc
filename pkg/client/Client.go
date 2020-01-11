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
	conn *grpc.ClientConn // 用于关闭连接等

	cli pb.ControlClient
}

func NewConsumeClient() (*ConsumeClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	controlClient := pb.NewControlClient(conn)
	return &ConsumeClient{
		conn: conn,
		cli:  controlClient,
	}, nil
}

func (c *ConsumeClient) Close() error {
	return c.conn.Close()
}

// 发送 购买订单 数据
func (c *ConsumeClient) Pay(ctx context.Context, req *pb.ConsumeReq) (resp *pb.ConsumeResp, err error) {
	resp, err = c.cli.Pay(ctx, req) // 注意这里没有用:=而已=, 因为函数声明里面返回值已经有名称了, 等于声明了变量
	if err != nil {
		log.Println("Pay failed:", err)
	}
	return // 注意这里没有写返回哪个变量, 因为函数声明的返回值里面给了变量名字, 默认就是返回那些变量
}
