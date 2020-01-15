package client

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/Account"
	"google.golang.org/grpc"
	"log"
)


const portRegistered = ":50051"

type Client struct {
	conn *grpc.ClientConn
	cli  pb.AccountClient
}

// 注册功能
func NewAccountClient() (*Client, error) {
	conn, err := grpc.Dial(portRegistered, grpc.WithInsecure())
	if err != nil {
		log.Println("connecting failed")
	}
	NewAccountClient := pb.NewAccountClient(conn)
	return &Client{
		conn: conn,
		cli:  NewAccountClient,
	}, nil
}

// 关闭连接
func (c *Client) Close() error {
	return c.conn.Close()
}

// 发送注册信息
func (c *Client) Register(ctx context.Context, regisReq *pb.RegisReq) (regisResp *pb.RegisResp, err error){
	regisResp, err = c.cli.Register(ctx, regisReq)
	if err != nil {
		log.Println("cli.Registered failed: ", err)
	}
	return
}

// 登录信息
func (c *Client) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp, err = c.cli.Login(ctx, req)
	if err != nil {
		log.Println("c.cli.LogIn:, ", err)
	}
	return
}