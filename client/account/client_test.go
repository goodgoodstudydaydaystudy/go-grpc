package client

import (
	"context"
	"testing"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/account"
)

func TestClient_Register(t *testing.T) {
	cli, err := NewAccountClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	resp, err := cli.Register(ctx, &pb.RegisReq{
		Account:  "test01",
		Password: "123456",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}
