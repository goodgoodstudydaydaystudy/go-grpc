package user

import (
	"context"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"google.golang.org/grpc"
	"log"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
)

const portUserLogic = ":50053"

// 1. client 结构体
type Client struct {
	conn  *grpc.ClientConn
	cli   pb.UserClient
}

// 2. client的new函数返回client的实例
func NewUserLogicClient() (*Client, error) {
	// 2.1 注册服务
	conn, err := grpc.Dial(portUserLogic,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(client.StatusCodeUnaryInterceptor))
	if err != nil {
		log.Println("logic conn failed: ", err)
	}
	// 2.2 获得client的conn
	newUserClient := pb.NewUserClient(conn)

	// 2.3 return
	return &Client{
		conn: conn,
		cli:  newUserClient,
	}, nil
}


func (c *Client) Close() error {
	return c.conn.Close()
}

// account的密码校验(登录)
func (c *Client) CheckoutPassword(ctx context.Context, req *pb.CheckUserPwdReq) (*pb.CheckUserPwdResp, protocol.ServerError) {
	resp, err := c.cli.CheckUserPassword(ctx, req)
	if err != nil {
		log.Println("userClient CheckoutPassword failed: ", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}


// account的register
func (c *Client) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, protocol.ServerError) {
	resp, err := c.cli.Register(ctx, req)
	if err != nil {
		log.Println("userClient Register failed: ", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}