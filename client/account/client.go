package account

import (
	"context"
	_ "database/sql"
	"log"

	"google.golang.org/grpc"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	"goodgoodstudy.com/go-grpc/pkg/procotol"
	md "goodgoodstudy.com/go-grpc/pkg/utils"
)

const portRegistered = ":50051"

type Client struct {
	conn    *grpc.ClientConn
	cli     pb.AccountClient
	message string
}

// 注册服务功能
func NewAccountClient() (*Client, error) {
	conn, err := grpc.Dial(portRegistered,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(client.StatusCodeUnaryInterceptor), // 拦截器
	)
	if err != nil {
		log.Println("connecting failed")
	}
	newAccountClient := pb.NewAccountClient(conn)

	return &Client{
		conn: conn,
		cli:  newAccountClient,
	}, nil
}

// 关闭连接
func (c *Client) Close() error {
	return c.conn.Close()
}

// 发送注册信息
func (c *Client) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, protocol.ServerError) {
	req.Password = md.Encryption(req.Password)
	resp, err := c.cli.Register(ctx, req)
	if err != nil {
		log.Println("cli.Registered failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}

// 登录信息
func (c *Client) Login(ctx context.Context, acc, pwd string) (*pb.LoginResp, protocol.ServerError) {
	req := &pb.LoginReq{Account: acc, Password: pwd}
	resp, err := c.cli.Login(ctx, req)
	if err != nil {
		log.Println("cli.LogIn failed:, ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}

// 查询user by account
func (c *Client) GetUserByAccount(ctx context.Context, acc string) (*pb.GetUserByAccountResp, protocol.ServerError) {
	req := &pb.GetUserByAccountReq{Account: acc}
	resp, err := c.cli.GetUserByAccount(ctx, req)
	if err != nil {
		log.Println("client GetUserByAccount: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}

// 查询user by userId
func (c *Client) GetUserByUserId(ctx context.Context, uid uint32) (*pb.GetUserByIdResp, protocol.ServerError) {
	req := &pb.GetUserByIdReq{UserId: uid}
	resp, err := c.cli.GetUserById(ctx, req)
	if err != nil {
		log.Println("client GetUserByUserId failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}
