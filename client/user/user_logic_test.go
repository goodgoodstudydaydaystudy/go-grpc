package user

import (
	"context"
	"log"
	"testing"

	"goodgoodstudy.com/go-grpc/client/wallet"
	upb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
	wpb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
)

func TestUserLogic(t *testing.T) {
	result, err := testOption(2)
	if err != nil {
		t.Log(err.Code())

	}
	t.Log(result)
}

func testOption(n int) (interface{}, protocol.ServerError) {
	logicCli, err := NewUserLogicClient()
	if err != nil {
		log.Println("test UserLogicClient error: ", err)
	}
	walletCli, err := wallet.NewWalletClient()
	if err != nil {
		log.Println("test WalletClient error: ", err)
	}

	ctx := context.Background()
	switch n {
	case 1:
		// 1. 注册
		return  logicCli.Register(ctx, &upb.RegisterReq{
			Account:              "test03",
			Password:             "666666",
			Nickname:             "test3",
			Gender:               1,
		})

		// 2. 登录
	case 2:
		return  logicCli.CheckoutPassword(ctx, &upb.CheckUserPwdReq{
			Account:              "test03",
			Password:             "666666",
		})
		// 3. 充值
	case 3:
		return walletCli.Recharge(ctx, &wpb.RechargeReq{
			UserId:               10,
			Count:                88888,
		})
		// 4. 余额
	case 4:
		return walletCli.GetUserById(ctx, 9)
		}
	return nil, nil
}
