package client

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/Register"
	"google.golang.org/grpc"
	"log"
)


const port_regis = ":50051"

type AccountClient struct {
	conn *grpc.ClientConn
	cli  pb.NewconstomerClient
}

// 注册功能
func NewAccountClient() (*AccountClient, error) {
	conn, err := grpc.Dial(port_regis, grpc.WithInsecure())
	if err != nil {
		log.Println("connecting failed")
	}
	NewAccountClient := pb.NewNewconstomerClient(conn)
	return &AccountClient{
		conn: conn,
		cli:  NewAccountClient,
	}, nil
}

// 关闭连接
func (c *AccountClient) RegClose() error {
	return c.conn.Close()
}

// 发送注册信息
func (c *AccountClient) Regis(ctx context.Context, regireq *pb.RegisReq) (regiresp *pb.RegisResp, err error){
	regiresp, err = c.cli.Registered(ctx, regireq)
	if err != nil {
		log.Println("cli.Registered failed: ", err)
	}
	return
}

// 登录信息
func (c *AccountClient) Login(ctx context.Context, req *pb.LogReq) (logresp *pb.LogResp, err error) {
	logresp, err = c.cli.LogIn(ctx, req)
	if err != nil {
		log.Println("c.cli.LogIn:, ", err)
	}
	return
}