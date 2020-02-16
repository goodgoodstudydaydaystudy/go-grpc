package user

import (
	"context"
	upb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"testing"
	"time"
)


func TestClient_Register(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.Register(context.Background(), &upb.RegisterReq{
		Account:              "test11",
		Password:             "123456",
		Nickname:             "test11",
		Gender:               2,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestClient_Login(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Fatal(err)
	}

	resp, err :=cli.Login(context.Background(), &upb.LoginReq{
		Account:              "test03",
		Password:             "123456",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	t.Log(resp.Token)

}

func TestClient_Recharge(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	const grpctoken  = "resp_token"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2luZm8iOnsidXNlcl9pZCI6MiwiYWNjb3VudCI6InRlc3QwMiIsIm5pY2tuYW1lIjoidGVzdDIiLCJnZW5kZXIiOjJ9LCJleHAiOjE1ODE0OTMyNDcsImlhdCI6MTU4MTQ5MjY0Nywic3ViIjoidGVzdCJ9.NFC6vWSqp__c6PC7-8FnP45OAROUbwTeC6dO56kcVck"
	ctx = metadata.AppendToOutgoingContext(ctx, grpctoken, token)

	resp, err := cli.Recharge(ctx, &upb.RechargeReq{
		UserId:               2,
		Delta:                2,
	})

	t.Log(resp)
}

func TestClient_LoginAndRecharge(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	LoginResp, err := cli.Login(ctx, &upb.LoginReq{
		Account:              "test11",
		Password:             "123456",
	})
	if err != nil {
		t.Log("Login failed:", err)
	}

	const grpcToken = "resp_token"
	token := LoginResp.Token
	t.Log("test token:", token)
	ctx = metadata.AppendToOutgoingContext(ctx, grpcToken, token)

	amount := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000)
	t.Log("ready to recharge")
	_, err = cli.Recharge(ctx, &upb.RechargeReq{
		UserId:  11,
		Delta:   amount,
		Account: "test11",
	})
	if err != nil {
		t.Log("Recharge failed:", err)
	}else {
		t.Log("recharge success:", amount)
	}
}

func TestClient_GetTopUser(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Log("test GetTopUser failed:", err)
	}
	ctx := context.Background()

	LoginResp, err := cli.Login(ctx, &upb.LoginReq{
		Account:  "test01",
		Password: "123456",
	})
	if err != nil {
		t.Log("Login failed:", err)
	}

	const grpcToken = "resp_token"
	token := LoginResp.Token
	ctx = metadata.AppendToOutgoingContext(ctx, grpcToken, token)

	resp, err := cli.GetTopUser(ctx, &upb.GetTopUserReq{
		Top: 10,
	})
	if err != nil {
		t.Log("test GetTopUser failed:", err)
	}

	t.Log(resp)
}

func TestClient_OrderNotPay(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Log("test cli failed:", err)
	}
	ctx := context.Background()

	LoginResp, err := cli.Login(ctx, &upb.LoginReq{
		Account:              "test01",
		Password:             "123456",
	})
	if err != nil {
		t.Log("Login failed:", err)
	}

	const grpcToken = "resp_token"
	token := LoginResp.Token
	ctx = metadata.AppendToOutgoingContext(ctx, grpcToken, token)

	orderId, err := cli.OrderNotPay(ctx, &upb.OrderNotPayReq{
		UserId: 1,
	})
	if err != nil {
		t.Log("test OrderNotPay failed:", err)
	}
	t.Log(orderId)
}

func TestClient_Pay(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Log("test cli failed:", err)
	}
	ctx := context.Background()

	LoginResp, err := cli.Login(ctx, &upb.LoginReq{
		Account:              "test01",
		Password:             "123456",
	})
	if err != nil {
		t.Log("Login failed:", err)
	}

	const grpcToken = "resp_token"
	token := LoginResp.Token
	ctx = metadata.AppendToOutgoingContext(ctx, grpcToken, token)

	_, err = cli.Pay(ctx, &upb.PayReq{
		OrderId: "fd5832fc62b6da8bd2b381b4e3bee9c0",
		})
	if err == nil {
		t.Log("pay success")
	}else {
		t.Log("pay failed:", err)
	}
}