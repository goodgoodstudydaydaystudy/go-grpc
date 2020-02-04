package user

import (
	"context"
	"log"
	"testing"

	upb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
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

func testOption() (*Result, protocol.ServerError) {
	resp := &Result{}

	logicCli, err := NewUserLogicClient()
	if err != nil {
		log.Println("testOption newCli failed: ", err)
	}

	ctx := context.Background()

	//// 注册
	//resultFromRegister, err := logicCli.Register(ctx, &upb.RegisterReq{
	//	Account:  "test04",
	//	Password: "777777",
	//	Nickname: "test4",
	//	Gender:   2,
	//	})
	//if err != nil {
	//	log.Println("resultFromRegister failed: ", err)
	//	return resp, protocol.ToServerError(err)
	//}

	// 2. 登录
	resultFromLogin, err :=  logicCli.Login(ctx, &upb.LoginReq{
		Account:  "test04",
		Password: "777777",
		})
	if err != nil {
		log.Println("resultFromLogin failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	// 充值
	_, err = logicCli.Recharge(ctx, &upb.RechargeReq{
		UserId: 13,
		Delta:  11,
		})
	if err != nil {
		log.Println("resultFromRecharge failed: ", err)
		return resp, protocol.ToServerError(err)
	}
	return &Result{
		//newUsrId: resultFromRegister.UserId,
		balance:  resultFromLogin.Balance,
	}, nil
}

func TestClient_Recharge(t *testing.T) {
	cli, err := NewUserLogicClient()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.Login(context.Background(), &upb.LoginReq{
		Account:              "test04",
		Password:             "777777",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

