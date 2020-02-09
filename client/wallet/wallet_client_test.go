package wallet

import (
	"context"
	"testing"
)

func TestWalletClient(t *testing.T) {
	cli, err := NewWalletClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	resp, se := cli.Recharge(ctx, 9, 10)
	if se != nil {
		t.Log(se.Code())
	}
	//resp, se := cli.GetUserByAccount(ctx, &pb.GetUserBalanceReq{UserId:10})
	//if se != nil {
	//	t.Log(se.Code())
	//}
	t.Log(resp)
}
