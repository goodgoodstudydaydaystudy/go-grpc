package client

import (
	"context"
	"testing"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/pay"
)

func TestConsumeClient_Pay(t *testing.T) {
	cli, err := NewConsumeClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	resp, err := cli.Pay(ctx, &pb.ConsumeReq{UserId:3333, ItemNum:5555, ItemId:88888, Description:"test_pay"})
	if  err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}