package main

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"

	"goodgoodstudy.com/go-grpc/client/logic/user"
	upb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
	"goodgoodstudy.com/go-grpc/pkg/procotol/encode"

	"google.golang.org/grpc/metadata"
)

type Info struct {
	account string
	userId uint32
}

func main()  {
	/*
	测试, 应该是有一个/很多个进程/协程, 一直在向logic发送 Recharge的请求,
	然后时不时把wallet server停掉, 再启动, 再停掉.

	最后把测试的client进程关掉, 把wallet server打开, 看看数据库的数据对不对, 看看排行榜数据对不对

	用client_test可能不够, 你要自己写个main函数, 搞个for循环一直向logic发请求, 这种测试的小程序就可以放在cmd/ 目录下
	*/


	// get a new client by newClient and send/return chan
	// check err from api. return chan when occurred err
	// keepConn notice newClient when get errSig return chan
	// newClient return clientChan
	// loop check clientChan and get client in main when the chan full

	// need a loop to check sig before get a client from chan
	var wg sync.WaitGroup

	client, err := user.NewUserLogicClient()
	wg.Add(1)
	for {
		go func() {
			if err != nil {
				log.Println("NewUserLogicClient failed:", err)
				return
			}
			userInfo := generateInfo()
			ctx, err := login(client, userInfo.account)
			if err != nil {
				log.Println("login failed:", err)
				return
			}
			err = recharge(ctx, client, userInfo.userId, userInfo.account)
			if err != nil {
				log.Println("recharge failed", err)
				return
			}
			wg.Done()
			wg.Wait()
		}()
		log.Println(runtime.NumGoroutine())
	}
}

func login(cli *user.Client, account string) (ctx context.Context, err error) {
	ctx = context.Background()
	LoginResp, err := cli.Login(ctx, &upb.LoginReq{
		Account:              account,
		Password:             "123456",
	})
	if err != nil {
		log.Println("login failed:", err)
		return ctx, err
	}
	const grpcToken = "resp_token"
	token := LoginResp.Token
	ctx = metadata.AppendToOutgoingContext(ctx, grpcToken, token)
	return
}


func recharge(ctx context.Context, cli *user.Client, userId uint32, account string) (err error) {
	amount := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100)
	_, err = cli.Recharge(ctx, &upb.RechargeReq{
		UserId:  userId,
		Delta:   amount,
		Account: account,
	})
	if err != nil {
		log.Println("Recharge failed:", err)
	}else {
		log.Println("recharge success:", amount)
	}
	return
}


func generateInfo() *Info {
	userIdInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000)
	userIdStr := strconv.FormatInt(userIdInt, 10)

	var account string
	if userIdInt < 10 {
		account = "test0" + userIdStr
	} else {
		account = "test" + userIdStr
	}

	return &Info{
		account: account,
		userId:  uint32(userIdInt),
	}
}


func generateOrderId(userId uint32) string {
	return encode.GenerateMd5(strconv.FormatUint(uint64(userId), 10))
}

