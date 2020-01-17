package client

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"google.golang.org/grpc"
	"log"
)


const portRegistered = ":50051"

type Client struct {
	Conn *grpc.ClientConn
	Cli  pb.AccountClient
}

// 注册功能
func NewAccountClient() (*Client, error) {
	conn, err := grpc.Dial(portRegistered, grpc.WithInsecure())
	if err != nil {
		log.Println("connecting failed")
	}
	NewAccountClient := pb.NewAccountClient(conn)
	return &Client{
		Conn: conn,
		Cli:  NewAccountClient,
	}, nil
}

// 关闭连接
func (c *Client) Close() error {
	return c.Conn.Close()
}

// 发送注册信息
func (c *Client) Register(ctx context.Context, Req *pb.RegisReq) (Resp *pb.RegisResp, err error){
	Resp, err = c.Cli.Register(ctx, Req)
	if err != nil {
		log.Println("cli.Registered failed: ", err)
	}
	return
}

// 登录信息
func (c *Client) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp, err = c.Cli.Login(ctx, req)
	if err != nil {
		log.Println("c.cli.LogIn:, ", err)
	}
	return
}