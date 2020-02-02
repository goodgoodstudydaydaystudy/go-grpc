package main

import (
	"context"
	ap "goodgoodstudy.com/go-grpc/client/account"
	pp "goodgoodstudy.com/go-grpc/client/pay"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/pay"
	md "goodgoodstudy.com/go-grpc/pkg/utils"
	"log"
)


func main() {

	// 初始化客户端
	consumeClient, err := pp.NewConsumeClient()
	if err != nil {
		log.Println("new client:", err)
	}

	accountClient, err := ap.NewAccountClient()
	if err != nil {
		log.Println("new client:", err)
	}

	ctx := context.Background()

	// 付款接口
	Consume(ctx, consumeClient)

	// 注册接口
	Register(ctx, accountClient)

	// 登录接口
	Login(ctx, accountClient)

	_ = consumeClient.Close()	// 不是执行完main就close了？ 为啥不用defer？
	_ = accountClient.Close()	// 执行完close
}

// 付款接口
func Consume(ctx context.Context, consumeClient *pp.ConsumeClient)  {
	consumeResp, err := consumeClient.Pay(ctx, &pb.ConsumeReq{ItemId: 1, ItemNum: 2, UserId:3, Description:"aaaa"})
	if err != nil {
		log.Println("consume failed,", err)
		return
	}
	log.Println(consumeResp)
}

// 注册接口
func Register(ctx context.Context, accountClient *ap.Client) (message string) {
	md5Account := md.Encryption("7777777")
	md5Password := md.Encryption("6666666")

	registerResp, err := accountClient.Register(ctx, &rpb.RegisterReq{Account:md5Account, Password:md5Password})
	if err != nil {
		log.Println("accountClient.Register failed: ", err)
	}
	log.Println(registerResp)
	return
}

// 登录接口
func Login(ctx context.Context, accountClient *ap.Client)  {
	loginResp, err := accountClient.Login(ctx, "account77", "password8888")
	if err != nil {
		log.Println("accountClient.Login failed: ", err)
	}
	log.Println(loginResp)
}

