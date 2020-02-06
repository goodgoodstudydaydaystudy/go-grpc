package account

import (
	"context"
	_ "database/sql"
	"log"

	"google.golang.org/grpc"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	"goodgoodstudy.com/go-grpc/pkg/procotol"
)

const portRegistered = ":50051"

type Client struct {
	conn *grpc.ClientConn
	cli  pb.AccountClient
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
func (c *Client) AddUser(ctx context.Context, acc string, pwd string, nickname string, gender pb.Gender) (*pb.AddUserResp, protocol.ServerError) {
	req := &pb.AddUserReq{
		Account:  acc,
		Password: pwd,
		Name:     nickname,
		Gender:   gender,
	}
	resp, err := c.cli.AddUser(ctx, req)
	if err != nil {
		log.Println("client Registered failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}

// 登录信息
func (c *Client) CheckPwd(ctx context.Context, acc, pwd string) (*pb.CheckPwdResp, protocol.ServerError) {
	req := &pb.CheckPwdReq{Account: acc, Password: pwd}
	resp, err := c.cli.CheckPwd(ctx, req)
	if err != nil {
		log.Println("client LogIn failed:, ", err)
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
