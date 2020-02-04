package account

import (
	"context"
	"os"
	"testing"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
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

var cli *Client

func init() {
	var err error
	cli, err = NewAccountClient()
	if err != nil {
		os.Exit(1)
	}
}

func TestClient_Register(t *testing.T) {
	resp, err := cli.AddUsr(context.Background(), "test001", "123456", "testName", pb.Gender_MALE)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestClient_Login(t *testing.T) {
	resp, err := cli.CheckPwd(context.Background(), "test001", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestClient_GetUserByAccount(t *testing.T) {
	resp, err := cli.GetUserByAccount(context.Background(), "test001")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
