package wallet

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"google.golang.org/grpc"
	"log"
)

const portWallet = ":50051"

type walletClient struct {
	conn  *grpc.ClientConn
	cli	  pb.WalletClient
}


func NewWalletClient() (*walletClient, error) {
	conn, err := grpc.Dial(portWallet,
		grpc.WithInsecure(),
	)	// TODO 拦截器
	if err != nil {
		log.Println("walletClient conn failed: ", err)
	}
	newWalletClient := pb.NewWalletClient(conn)

	return &walletClient{
		conn: conn,
		cli:  newWalletClient,
	}, nil
}


func (c *walletClient) Close() error{
	return c.conn.Close()
}



func (c *walletClient) Recharge(ctx context.Context, req *pb.RechargeReq) (*pb.RechargeResp, protocol.ServerError){
	resp, err := c.cli.Recharge(ctx, req)
	if err != nil {
		log.Println("client Recharge failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}


func (c *walletClient) GetUserByAccount(ctx context.Context, req *pb.GetUserBalanceReq) (*pb.GetUserBalanceResp, protocol.ServerError) {
	resp, err := c.cli.GetUserBalance(ctx, req)
	if err != nil {
		log.Println("client GetUserByAccount failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return resp, nil
}