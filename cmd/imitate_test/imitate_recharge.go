package main

import (
	"context"
	"log"
	"math/rand"
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
	// TODO
	/*
	测试, 应该是有一个/很多个进程/协程, 一直在向logic发送 Recharge的请求,
	然后时不时把wallet server停掉, 再启动, 再停掉.

	最后把测试的client进程关掉, 把wallet server打开, 看看数据库的数据对不对, 看看排行榜数据对不对

	用client_test可能不够, 你要自己写个main函数, 搞个for循环一直向logic发请求, 这种测试的小程序就可以放在cmd/ 目录下
	 */

	// get a new client by newClient and send/return chan
	// keepConn notice newClient when get errSig return chan
	// need a loop to check sig before get a client from chan


	var client *user.Client
	var mut sync.Mutex
	var wg  sync.WaitGroup
	var clientChan  = make(chan *user.Client)
	var errOccurred = make(chan bool)


	client, err := user.NewUserLogicClient()
	if err != nil {
		log.Println("NewUserLogicClient failed:", err)
	}
	// need a loop to check sig before get a client from chan
	wg.Add(1)
	go func() {
		log.Println("go")
		for {
			select {
			case <-errOccurred:		// errOccurred not done. is nil
				client =<-clientChan
			default:
				// nothing
			}
			time.Sleep(time.Second)
			wg.Done()
		}
	}()
	wg.Wait()

	for {
		go func() {
			wg.Add(1)
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			userInfo := generateInfo()

			ctx, err := login(client, userInfo.account)
			if err != nil {
				log.Println("login failed")
				clientChan, errOccurred = keepConn(err)
			}else {
				log.Printf("%v login success", userInfo.userId)
			}

			err = recharge(ctx, client, userInfo.userId, userInfo.account)
			if err != nil {
				log.Printf("%v recharge failed:", userInfo.userId)
				clientChan, errOccurred = keepConn(err)
			} else {
				log.Printf("%v recharge success", userInfo.userId)

			}
			wg.Done()
		}()
		wg.Wait()
	}
}

// the second <return chan *user.Client>
// new client
// channel <-client
// return chan
func newLogicClientChan() chan *user.Client {
	clientChan := make(chan *user.Client)
	client, err := user.NewUserLogicClient()
	if err != nil {
		log.Println("NewUserLogicClient failed:", err)
	}
	log.Printf("get client")
	clientChan <-client
	return clientChan
}

// the first <check errSig>
// check errSig and send notice newLogicClient().
// to new client. return chan
func keepConn(err error) (chan *user.Client, chan bool) {
	var errSig      = make(chan error, 1)
	var clientChan  = make(chan *user.Client)
	var errOccurred = make(chan bool)

	errSig <-err
	select {
		case <-errSig:
			clientChan = newLogicClientChan()
			errOccurred <-true
			return clientChan, errOccurred
		default:
			// nothing
		}
		time.Sleep(time.Second)
		log.Printf("repeat conn wait 1 second")

	return nil, nil
}



func login(cli *user.Client, account string) (ctx context.Context, err error) {
	ctx = context.Background()

	LoginResp, err := cli.Login(ctx, &upb.LoginReq{
		Account:              account,
		Password:             "123456",
	})
	if err != nil {
		log.Println("login failed:", err)
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
	return nil
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

