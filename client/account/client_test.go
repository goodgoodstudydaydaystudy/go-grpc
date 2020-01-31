package client

import (
	"context"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"testing"
)

type testType struct {
	*pb.RegisterResp
	*pb.LoginResp

}

/*
func TestClient_Register(t *testing.T) {
	cli, err := NewAccountClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	resp, err := cli.Register(ctx, &pb.RegisterReq{Account:"test01", Password:"123456", Name:"testName", Gender:1})
	if err != nil {
		t.Fatal(err)
	}
	resp, se := cli.Login(ctx, &pb.LoginReq{Account:"test01", Password:"123456"}) // 这里se命名为err的话将不能调用err.Code(); 因为上面err已经被推断为普通error类型了
	if se != nil {
		t.Fatal(se.Code())
	}

	resp, se := cli.Query(ctx, &pb.QueryUserReq{Account:"test01"})
	if se != nil {
		t.Fatal(se.Code())
	}
	t.Log(resp)
}
*/


func TestClient_Register(t *testing.T) {
	resp, err := testAllFeature(2)
	if err != nil {
		t.Fatal(err.Code())
	}
	t.Log(resp)
}

func testAllFeature(n int) (interface{}, protocol.ServerError) {
	cli, err := NewAccountClient()
	if err != nil {
	}
	ctx := context.Background()
	switch n {
	case 0:
		resp, err := cli.Register(ctx, &pb.RegisterReq{Account:"test01", Password:"123456", Name:"testName", Gender:1})
		return resp, err
	case 1:
		resp, se := cli.Login(ctx, &pb.LoginReq{Account:"test01", Password:"123456"}) // 这里se命名为err的话将不能调用err.Code(); 因为上面err已经被推断为普通error类型了
		return resp, se
	case 2:
		resp, se := cli.Query(ctx, &pb.QueryUserReq{Account:"test01"})
		return resp, se
	case 3:
		resp, se := cli.Query(ctx, &pb.QueryUserReq{UserId:9})
		return resp, se
	default:
		return nil, nil
	}
}