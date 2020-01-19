package client

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"goodgoodstudy.com/go-grpc/pkg/procotol"
	md "goodgoodstudy.com/go-grpc/pkg/utils"
)

const portRegistered = ":50051"

type Client struct {
	Conn *grpc.ClientConn // TODO 说100次变量名不要乱用大写.
	Cli  pb.AccountClient
}

type Info struct {
	Password string
	Account  string
}

// 注册功能
func NewAccountClient() (*Client, error) {
	conn, err := grpc.Dial(portRegistered,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(client.StatusCodeUnaryInterceptor), // 拦截器
	)
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
func (c *Client) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, protocol.ServerError) {
	req.Password = md.Encryption(req.Password)
	resp, err := c.Cli.Register(ctx, req)
	if err != nil {
		log.Println("cli.Registered failed: ", err)
	}
	return resp, protocol.ToServerError(err)
}

// 登录信息
func (c *Client) Login(ctx context.Context, account string, password string) (*pb.LoginResp, protocol.ServerError) {
	md5Password := md.Encryption(password)
	req := &pb.LoginReq{Account: account, Password: md5Password}
	resp, err := c.Cli.Login(ctx, req)
	if err != nil {
		log.Println("c.cli.LogIn failed:, ", err)
	}
	return resp, protocol.ToServerError(err)
}
