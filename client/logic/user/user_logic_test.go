package user

import (
	"context"
	upb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"google.golang.org/grpc/metadata"
	"log"
	"testing"
)

type Result struct {
	newUsrId uint32
	balance  int64
}


func TestUserLogic(t *testing.T) {
	result, err := testOption()
	if err != nil {
		t.Log(err.Code())
	}
	t.Log(result)
}

func testOption() (Result, protocol.ServerError) {
	resp := Result{}

	cli, err := NewUserLogicClient()
	if err != nil {
		log.Println("testOption newCli failed: ", err)
	}

	ctx := context.Background()

	// 2. 登录
	loginResp, err := cli.Login(ctx, &upb.LoginReq{
		Account:  "test01",
		Password: "123456",
		})
	if err != nil {
		log.Println("resultFromLogin failed: ", err)
		return resp, protocol.ToServerError(err)
	}


	// 3. 获取 token
	token := loginResp.Token
	//log.Println("token:", token)
	//log.Println("time.new:", time.Now().Unix())

	// 4. token和userInfo 写入ctx
	const grpcToken = "resp_token"
	//const grpcUserInfo = "resp_userInfo"
	//
	//JsonUserInfo, err := json.Marshal(&loginResp.UserInfo)
	//if err != nil {
	//	log.Println("JsonUserInfo failed", err)
	//}
	//log.Println("userInfo:", loginResp.UserInfo)
	//log.Println("JsonUserInfo:", string(JsonUserInfo))

	//StrUserInfo := string(JsonUserInfo)
	ctx = metadata.AppendToOutgoingContext(ctx, grpcToken, token)

	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		log.Println("md[resp_token]:", md["resp_token"][0])
	}else {
		log.Println("FromOutgoingContext failed")
	}

	//err = grpc.SetHeader(ctx, metadata.Pairs(
	//	grpcToken, token,
	//	))
	//
	//if err != nil {
	//	log.Println("setHeader failed:", err)
	//}


	// 5. 充值 （token）
	_, err = cli.Recharge(ctx, &upb.RechargeReq{
		UserId: 1,
		Delta:  1,
		})
	if err != nil {
		log.Println("resultFromRecharge failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return Result{
		balance:  loginResp.Balance,
	}, nil
}


func TestClient_Recharge(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Fatal(err)
	}

	mdToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwLCJpYXQiOjE1ODExMzY5MzAsImlzcyI6InRlc3QwMSIsInN1YiI6InRlc3QifQ.g17JMP06uyWpXBSUjyJbCrQSpYtLHMCPb0_Ve01ojmU"

	ctx := context.Background()

	const grpcToken = "resp_token"

	ctx =  metadata.AppendToOutgoingContext(ctx, grpcToken, mdToken)

	resp, err := cli.Recharge(ctx, &upb.RechargeReq{
		UserId:               1,
		Delta:                1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestClient_Register(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.Register(context.Background(), &upb.RegisterReq{
		Account:              "test02",
		Password:             "123456",
		Nickname:             "test2",
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
		Account:              "test01",
		Password:             "123456",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	t.Log(resp.Token)

}

