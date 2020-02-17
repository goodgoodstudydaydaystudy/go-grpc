package user

import (
	"context"
	"log"

	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
)

const portUserLogic = ":50053"

// 1. client 结构体
type Client struct {
	conn *grpc.ClientConn
	cli  pb.UserClient
}

// 2. client的new函数返回client的实例
func NewUserLogicClient() (*Client, error) {
	// 2.1 注册服务
	conn, err := grpc.Dial(portUserLogic,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				client.StatusCodeUnaryInterceptor,
				),
	))

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
func (c *Client) Login(ctx context.Context, req *pb.LoginReq, opts ...grpc.CallOption) (*pb.LoginResp, protocol.ServerError) {
	resp, err := c.cli.Login(ctx, req, opts...)
	if err != nil {
		log.Println("userClient CheckoutPassword failed: ", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}

// account的register
func (c *Client) Register(ctx context.Context, req *pb.RegisterReq, opts ...grpc.CallOption) (*pb.RegisterResp, protocol.ServerError) {
	resp, err := c.cli.Register(ctx, req, opts...)
	if err != nil {
		log.Println("userClient Register failed: ", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}

func (c *Client) Recharge(ctx context.Context, req *pb.RechargeReq, opts ...grpc.CallOption) (*pb.RechargeResp, protocol.ServerError) {
	resp, err := c.cli.Recharge(ctx, req, opts...)
	if err != nil {
		log.Println("userClient Recharge failed: ", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}

func (c *Client) GetTopUser(ctx context.Context, req *pb.GetTopUserReq, opts ...grpc.CallOption) (*pb.GetTopUserResp, protocol.ServerError) {
	resp, err := c.cli.GetTopUser(ctx, req, opts...)
	if err != nil {
		log.Println("userClient GetTopUser failed:", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}

func (c *Client) RecordOrderNotPay(ctx context.Context, req *pb.RecordOrderNoPaidReq, opts ...grpc.CallOption) (*pb.RecordOrderNoPaidResp, protocol.ServerError) {
	resp, err := c.cli.RecordOrderNoPaid(ctx, req, opts...)
	if err != nil {
		log.Println("userClient OrderNotPay failed:", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}

func (c *Client) Pay(ctx context.Context, req *pb.PayReq, opts ...grpc.CallOption) (*pb.PayResp, protocol.ServerError) {
	resp, err := c.cli.Pay(ctx, req, opts...)
	if err != nil {
		log.Println("userClient Pay failed:", err)
		return nil, protocol.ToServerError(err)
	}
	return resp, nil
}