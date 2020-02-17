package wallet

import (
	"context"
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"google.golang.org/grpc"
	"log"
)

const portWallet = ":50052"


type Client struct {
	conn  *grpc.ClientConn
	cli	  pb.WalletClient
}


func NewWalletClient() (*Client, error) {
	conn, err := grpc.Dial(portWallet,
		grpc.WithInsecure(),
	grpc.WithUnaryInterceptor(client.StatusCodeUnaryInterceptor))
	if err != nil {
		log.Println("walletClient conn failed: ", err)
	}
	newWalletClient := pb.NewWalletClient(conn)

	return &Client{
		conn: conn,
		cli:  newWalletClient,
	}, nil
}


func (c *Client) Close() error{
	return c.conn.Close()
}


func (c *Client) Recharge(ctx context.Context, account string, uid uint32, delta int64) (*pb.RechargeResp, protocol.ServerError){
	req := &pb.RechargeReq{UserId: uid, Amount: delta, Account:account}
	resp, err := c.cli.Recharge(ctx, req)
	if err != nil {
		log.Println("client Recharge failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}


func (c *Client) GetUserById(ctx context.Context, uid uint32) (*pb.GetUserBalanceResp, protocol.ServerError) {
	req := &pb.GetUserBalanceReq{UserId:uid}
	resp, err := c.cli.GetUserBalance(ctx, req)
	if err != nil {
		log.Println("client GetUserByAccount failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}

func (c *Client) GetTopUser(ctx context.Context, n uint32) (*pb.GetTopUserResp, protocol.ServerError) {
	req := &pb.GetTopUserReq{Top:n}
	resp, err := c.cli.GetTopUser(ctx, req)
	if err != nil {
		log.Println("client GetTopUser failed:", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}

func (c *Client) RecordOrderNoPaid(ctx context.Context, userId uint32, orderId string) (*pb.RecordOrderNoPaidResp, protocol.ServerError) {
	req := &pb.RecordOrderNoPaidReq{UserId: userId, OrderId: orderId}
	resp, err := c.cli.RecordOrderNoPaid(ctx, req)
	if err != nil {
		log.Println("client RecordOrderNoPaid failed:", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}

func (c *Client) Pay(ctx context.Context, orderId string) (*pb.PayResp, protocol.ServerError) {
	req := &pb.PayReq{OrderId: orderId}
	resp, err := c.cli.Pay(ctx, req)
	if err != nil {
		log.Println("client Pay failed:", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}