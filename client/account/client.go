package client

import (
	"context"
	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"

	"google.golang.org/grpc"

	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"goodgoodstudy.com/go-grpc/pkg/procotol"
	md "goodgoodstudy.com/go-grpc/pkg/utils"
)

const portRegistered = ":50051"

type accountClient struct {
	conn *grpc.ClientConn
	cli  pb.AccountClient
	message string
}

// 注册功能
func NewAccountClient() (*accountClient, error) {
	conn, err := grpc.Dial(portRegistered,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(client.StatusCodeUnaryInterceptor), // 拦截器
	)
	if err != nil {
		log.Println("connecting failed")
	}
	newAccountClient := pb.NewAccountClient(conn)

	return &accountClient{
		conn: conn,
		cli:  newAccountClient,
	}, nil
}


// 关闭连接
func (c *accountClient) Close() error {
	return c.conn.Close()
}

// 发送注册信息
func (c *accountClient) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, protocol.ServerError) {
	req.Password = md.Encryption(req.Password)
	resp, err := c.cli.Register(ctx, req)
	if err != nil {
		log.Println("cli.Registered failed: ", err)
		log.Println(resp.GetMessage(), nil)
		return resp, nil
	}
	return resp, protocol.ToServerError(err)
}

// 登录信息
func (c *accountClient) Login(ctx context.Context, account string, password string) (*pb.LoginResp, protocol.ServerError) {
	md5Password := md.Encryption(password)
	req := &pb.LoginReq{Account: account, Password: md5Password}
	resp, err := c.cli.Login(ctx, req)
	if err != nil {
		log.Println("c.cli.LogIn failed:, ", err)
	}
	return resp, protocol.ToServerError(err)
}
