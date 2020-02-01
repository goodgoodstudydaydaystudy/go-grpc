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
	resp, se := cli.Recharge(ctx, &pb.RechargeReq{UserId: 7, Count:8888})
	if se != nil {
		t.Log(se.Code())
	}
	t.Log(resp)
}
