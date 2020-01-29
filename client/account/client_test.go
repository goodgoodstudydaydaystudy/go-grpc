package client

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"testing"
)


func TestClient_Register(t *testing.T) {
	cli, err := NewAccountClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	resp, err := cli.Register(ctx, &pb.RegisterReq{Account:"test01", Password:"123456", Name:"testName", Gender:"F"})
	if err != nil {
		t.Fatal(err)
	}
	//resp, se := cli.Login(ctx, "test01", "123456") // 这里se命名为err的话将不能调用err.Code(); 因为上面err已经被推断为普通error类型了
	//if se != nil {
	//	t.Fatal(se.Code())
	//}
	t.Log(resp)
}
