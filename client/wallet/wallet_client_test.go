package client

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/wallet"
	"testing"
)

func TestWalletClient(t *testing.T) {
	cli, err := NewWalletClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	resp, se := cli.Recharge(ctx, &pb.RechargeReq{UserId: 9, Count:10})
	if se != nil {
		t.Log(se.Code())
	}
	//resp, se := cli.GetUserByAccount(ctx, &pb.GetUserBalanceReq{UserId:10})
	//if se != nil {
	//	t.Log(se.Code())
	//}
	t.Log(resp)
}
